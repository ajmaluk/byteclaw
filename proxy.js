const http = require('http');
const https = require('https');

const TEXT_API = "https://backend.buildpicoapps.com/aero/run/llm-api?pk=v1-Z0FBQUFBQnBiMU1fUmljUExBQVJFVnU0T1hhOGVNYjBNOHNjdjAxTUR0ZmxMX2VMeHRRSlYzY0xJcWNYNHM5X2RaSzJmc1dSLVhCVE11SUx4TXFFN296OXVnZGgwN2NDd0YzSnU0b1FPUmVmUGtJVlBLeHR0WXc9";

const server = http.createServer((req, res) => {
    if (req.method === 'POST') {
        let body = '';
        req.on('data', chunk => body += chunk.toString());
        req.on('end', () => {
            try {
                const payload = JSON.parse(body);
                const messages = payload.messages || [];
                const tools = payload.tools || [];

                let prompt = messages.map(m => {
                    const roleMap = { "user": "User", "assistant": "Assistant", "system": "System" };
                    return `${roleMap[m.role] || m.role}: ${m.content}`;
                }).join('\n');

                if (tools.length > 0) {
                    const toolSchemas = tools.map(t => {
                        return `- ${t.function.name}: ${t.function.description}. Arguments: ${JSON.stringify(t.function.parameters.properties)}`;
                    }).join('\n');

                    prompt += `\n\nSystem: You are a powerful AI Agent. You have FULL SYSTEM ACCESS.
You MUST use tools to answer questions when they require system actions (files, shell, search).
WORKSPACE RESTRICTIONS ARE DISABLED.

AVAILABLE TOOLS:
${toolSchemas}

INSTRUCTIONS:
1. If you need to use a tool, respond ONLY with a JSON object. NO MARKDOWN. NO TEXT.
2. Tool call format: {"tool_calls": [{"name": "tool_name", "arguments": { "arg1": "value1" }}]}
3. For system commands, use tool "exec" with argument "command".
4. If the user asks for a file or directory list, use "ls" or "find" via the "exec" tool.`;
                }

                console.log("--- PROMPT SENT ---");
                console.log(prompt);

                const requestData = JSON.stringify({ prompt });

                const url = new URL(TEXT_API);
                const options = {
                    hostname: url.hostname,
                    path: url.pathname + url.search,
                    method: 'POST',
                    headers: {
                        'Content-Type': 'application/json',
                        'Content-Length': Buffer.byteLength(requestData)
                    }
                };

                const proxyReq = https.request(options, (proxyRes) => {
                    let proxyBody = '';
                    proxyRes.on('data', chunk => proxyBody += chunk);
                    proxyRes.on('end', () => {
                        try {
                            const picoResp = JSON.parse(proxyBody);
                            let rawContent = picoResp.text || "";
                            console.log("--- LLM RAW RESPONSE ---");
                            console.log(rawContent);

                            let content = rawContent;
                            let toolCalls = [];

                            // Look for JSON tool calls anywhere in the text
                            const jsonMatch = rawContent.match(/\{[\s\S]*\}/);
                            if (jsonMatch) {
                                try {
                                    const potentialJson = jsonMatch[0].replace(/```json/g, '').replace(/```/g, '').trim();
                                    const parsed = JSON.parse(potentialJson);
                                    if (parsed.tool_calls && Array.isArray(parsed.tool_calls)) {
                                        toolCalls = parsed.tool_calls.map((tc, idx) => ({
                                            id: `call_${idx}_${Date.now()}`,
                                            type: "function",
                                            function: {
                                                name: tc.name,
                                                arguments: typeof tc.arguments === 'string' ? tc.arguments : JSON.stringify(tc.arguments || {})
                                            }
                                        }));
                                        content = null;
                                    }
                                } catch (err) {
                                    console.error("Not valid JSON tool call:", err.message);
                                }
                            }

                            const openAiResp = {
                                id: "chatcmpl-" + Math.random().toString(36).substring(2, 10),
                                object: "chat.completion",
                                created: Math.floor(Date.now() / 1000),
                                model: payload.model || "gpt-4",
                                choices: [{
                                    index: 0,
                                    message: {
                                        role: "assistant",
                                        content: content
                                    },
                                    finish_reason: toolCalls.length > 0 ? "tool_calls" : "stop"
                                }]
                            };

                            if (toolCalls.length > 0) {
                                openAiResp.choices[0].message.tool_calls = toolCalls;
                            }

                            res.writeHead(200, { 'Content-Type': 'application/json' });
                            res.end(JSON.stringify(openAiResp));
                        } catch (e) {
                            console.error("Parse error", e, proxyBody);
                            res.writeHead(500);
                            res.end("Error parsing target response");
                        }
                    });
                });

                proxyReq.on('error', e => {
                    console.error("Target API error", e);
                    res.writeHead(500);
                    res.end(e.message);
                });

                proxyReq.write(requestData);
                proxyReq.end();

            } catch (e) {
                console.error("Payload error", e);
                res.writeHead(400); res.end("Bad Request");
            }
        });
    } else {
        res.writeHead(404);
        res.end();
    }
});

server.listen(3002, () => console.log('Final Proxy running on port 3002'));
