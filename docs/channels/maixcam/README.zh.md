# MaixCam

MaixCam 是专用于连接矽速科技 MaixCAM 与 MaixCAM2 AI 摄像设备的通道。它采用 TCP 套接字实现双向通信，支持边缘 AI 部署场景。

## 配置

```json
{
  "channels": {
    "maixcam": {
      "enabled": true,
      "server_address": "0.0.0.0:8899",
      "allow_from": []
    }
  }
}
```

| 字段           | 类型   | 必填 | 描述                             |
| -------------- | ------ | ---- | -------------------------------- |
| enabled        | bool   | 是   | 是否启用 MaixCam 频道            |
| server_address | string | 是   | TCP 服务器监听地址和端口         |
| allow_from     | array  | 否   | 设备ID白名单，空表示允许所有设备 |

## 使用场景

MaixCam 通道使 ByteClaw 能够作为边缘设备的 AI 后端运行：

- **智能监控** ：MaixCAM 发送图像帧，ByteClaw 通过视觉模型进行分析
- **物联网控制** ：设备发送传感器数据，ByteClaw 协调响应
- **离线AI** ：在本地网络部署 ByteClaw 实现低延迟推理


## 2. 边缘设备高性能脚本

在您的边缘设备上，您需要一个脚本来从摄像头捕获帧、运行轻量级检测模型（如 YOLOv8n）并通过 TCP 将事件报告回 ByteClaw。

以下 Python 脚本是一个可用于生产环境的模板。它使用 **OpenCV** 通过专用线程高效访问摄像头，使用 **Ultralytics YOLO** 进行推理，并利用异步网络通信以避免掉帧。

### 边缘设备上的先决条件
```bash
pip install opencv-python ultralytics
```

### `edge_vision.py`

```python
import socket
import json
import time
import cv2
import threading
import queue
from ultralytics import YOLO

# 配置
BYTECLAW_IP = "192.168.1.100"  # 更改为运行 ByteClaw 的机器的 IP
BYTECLAW_PORT = 18790
CAMERA_INDEX = 0               # 0 表示默认 USB/CSI 摄像头
MODEL_PATH = "yolov8n.pt"      # 轻量级边缘检测模型

# 用于异步发送事件的线程安全队列
event_queue = queue.Queue(maxsize=50)

def networking_thread():
    """后台线程处理 TCP 连接和自动重连。"""
    while True:
        client = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
        client.settimeout(5.0)
        connected = False
        
        try:
            client.connect((BYTECLAW_IP, BYTECLAW_PORT))
            print(f"[NET] 已连接到 ByteClaw {BYTECLAW_IP}:{BYTECLAW_PORT}")
            connected = True
            
            while connected:
                # 等待视觉循环的事件
                payload = event_queue.get(timeout=2.0)
                client.sendall((json.dumps(payload) + "\n").encode('utf-8'))
                
        except (socket.error, queue.Empty) as e:
            if isinstance(e, socket.error):
                print(f"[NET] 连接断开或失败: {e}。5秒后重试...")
                time.sleep(5)
        finally:
            client.close()

def vision_loop():
    """主线程处理高效的摄像头捕获和 AI 推理。"""
    # 加载 YOLO 模型
    print("[AI] 加载 YOLO 模型...")
    model = YOLO(MODEL_PATH)
    
    # 初始化摄像头
    print(f"[CAM] 正在访问摄像头索引 {CAMERA_INDEX}...")
    cap = cv2.VideoCapture(CAMERA_INDEX)
    
    # 为边缘设备优化摄像头设置（较低的分辨率 = 更快的推理）
    cap.set(cv2.CAP_PROP_FRAME_WIDTH, 640)
    cap.set(cv2.CAP_PROP_FRAME_HEIGHT, 480)
    # 防止帧缓冲（延迟优化）
    cap.set(cv2.CAP_PROP_BUFFERSIZE, 1) 
    
    if not cap.isOpened():
        print("[CAM] 错误: 无法打开摄像头。")
        return

    print("[CAM] 摄像头已激活。开始检测循环...")
    last_send_time = 0
    
    while True:
        ret, frame = cap.read()
        if not ret:
            print("[CAM] 检测到掉帧。正在重试...")
            continue
            
        # 运行 YOLO 检测
        results = model(frame, verbose=False)
        
        for result in results:
            boxes = result.boxes
            for box in boxes:
                # 检查 'person' 类别 (COCO 数据集中的索引 0)
                class_id = int(box.cls[0])
                if class_id == 0:
                    score = float(box.conf[0])
                    x1, y1, x2, y2 = box.xyxy[0].tolist()
                    w, h = x2 - x1, y2 - y1
                    
                    # 速率限制：每秒最多发送 1 个事件
                    if time.time() - last_send_time > 1.0:
                        payload = {
                            "type": "person_detected",
                            "timestamp": time.time(),
                            "data": {
                                "class_name": "person",
                                "score": score,
                                "x": x1, "y": y1, "w": w, "h": h
                            }
                        }
                        
                        # 加入队列而不阻塞摄像头帧率
                        if not event_queue.full():
                            event_queue.put_nowait(payload)
                            last_send_time = time.time()
                            print(f"[AI] 发现目标! 置信度: {score:.2f}")

    cap.release()
    cv2.destroyAllWindows()

if __name__ == "__main__":
    # 在后台启动健壮的网络线程
    net_thread = threading.Thread(target=networking_thread, daemon=True)
    net_thread.start()
    
    # 启动主视觉处理
    try:
        vision_loop()
    except KeyboardInterrupt:
        print("正在关闭边缘代理。")
```
