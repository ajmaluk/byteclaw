<div align="center">
  <img src="assets/banner.png" alt="ByteClaw Banner" width="100%">
  <br>
  <img src="assets/logo.png" alt="ByteClaw Logo" width="256">

  <h1>ByteClaw: गो (Go) में बना अति-कुशल (Ultra-Efficient) AI असिस्टेंट</h1>

  <h3>$10 हार्डवेयर · 10MB RAM · 1s बूट · 皮皮虾，我们走！</h3>

  <p>
    <img src="https://img.shields.io/badge/Go-1.25+-00ADD8?style=flat&logo=go&logoColor=white" alt="Go">
    <a href="https://github.com/ajmaluk/byteclaw/actions/workflows/ci.yml"><img src="https://github.com/ajmaluk/byteclaw/actions/workflows/ci.yml/badge.svg" alt="CI"></a>
    <img src="https://img.shields.io/badge/Arch-x86__64%2C%20ARM64%2C%20RISC--V-blue" alt="Hardware">
    <img src="https://img.shields.io/badge/license-MIT-green" alt="License">
    <br>
    <a href="https://github.com/ajmaluk/byteclaw"><img src="https://img.shields.io/badge/GitHub-ajmaluk%2Fbyteclaw-black?style=flat&logo=github&logoColor=white" alt="GitHub"></a>
    <br>
    <a href="./assets/wechat.png"><img src="https://img.shields.io/badge/WeChat-Group-41d56b?style=flat&logo=wechat&logoColor=white"></a>
    <a href="https://discord.gg/V4sAZ9XWpN"><img src="https://img.shields.io/badge/Discord-Community-4c60eb?style=flat&logo=discord&logoColor=white" alt="Discord"></a>
  </p>

[中文](README.zh.md) | [日本語](README.ja.md) | [Português](README.pt-br.md) | [Tiếng Việt](README.vi.md) | [Français](README.fr.md) | [English](README.md) | **हिन्दी** | [മലയാളം](README.ml.md)

</div>

---

🤖 ByteClaw एक अत्यंत-हल्का (ultra-lightweight) व्यक्तिगत AI असिस्टेंट है जो [nanobot](https://github.com/HKUDS/nanobot) से प्रेरित है। इसे Go भाषा में एक सेल्फ-बूटस्ट्रैपिंग (self-bootstrapping) प्रक्रिया के माध्यम से बिल्कुल शुरुआत से रिफैक्टर (refactored) किया गया है, जहां AI एजेंट ने स्वयं संपूर्ण आर्किटेक्चरल माइग्रेशन और कोड ऑप्टिमाइज़ेशन को संचालित किया है।

⚡️ $10 हार्डवेयर पर <10MB RAM के साथ काम करता है: यह OpenClaw की तुलना में 99% कम मेमोरी लेता है और Mac mini से 98% सस्ता है!

<table align="center">
  <tr align="center">
    <td align="center" valign="top">
      <p align="center">
        <img src="assets/byteclaw_mem.gif" width="360" height="240">
      </p>
    </td>
    <td align="center" valign="top">
      <p align="center">
        <img src="assets/licheervnano.png" width="400" height="240">
      </p>
    </td>
  </tr>
</table>

> [!CAUTION]
> **🚨 सुरक्षा और आधिकारिक चैनल (SECURITY & OFFICIAL CHANNELS)**
>
> * **कोई क्रिप्टो नहीं:** ByteClaw का कोई आधिकारिक टोकन/सिक्का **नहीं** है। `pump.fun` या अन्य ट्रेडिंग प्लेटफॉर्म पर कोई भी दावा **धोखाधड़ी (SCAM)** है।
> * **आधिकारिक डोमेन:** **एकमात्र** आधिकारिक वेबसाइट **[byteclaw.io](https://byteclaw.io)** है, और कंपनी की वेबसाइट **[ajmaluk.com](https://ajmaluk.com)** है।
> * **चेतावनी:** तीसरे पक्षों द्वारा कई `.ai/.org/.com/.net/...` डोमेन पंजीकृत किए गए हैं, जो हमारे नहीं हैं।
> * **चेतावनी:** ByteClaw विकास के प्रारंभिक चरण में है और इसमें ऐसे नेटवर्क सुरक्षा मुद्दे हो सकते हैं जिनका समाधान नहीं हुआ है। v1.0 रिलीज़ से पहले इसे उत्पादन (production) वातावरण में तैनात न करें।
> * **नोट:** ByteClaw ने हाल ही में कई PR (Pull Requests) मर्ज किए हैं, जिसके परिणामस्वरूप नवीनतम संस्करणों में अधिक मेमोरी फुटप्रिंट (10–20MB) हो सकता है। जैसे-जैसे फीचर सेट स्थिर होता है, हम संसाधन अनुकूलन (resource optimization) को प्राथमिकता देने की योजना बना रहे हैं।

## 📢 समाचार (News)

2026-02-09 🎉 ByteClaw रिलीज़ हुआ! $10 हार्डवेयर पर <10MB RAM के साथ AI एजेंट लाने के लिए सिर्फ 1 दिन में बनाया गया। 🤖 ByteClaw, चलो शुरू करें!

## ✨ विशेषताएं (Features)

🪶 **अति-हल्का (Ultra-Lightweight)**: <10MB मेमोरी फुटप्रिंट — मुख्य कार्यक्षमता के लिए Clawdbot की तुलना में 99% छोटा है।

💰 **न्यूनतम लागत (Minimal Cost)**: $10 हार्डवेयर पर चलाने के लिए काफी कुशल है — Mac mini से 98% सस्ता है।

⚡️ **तेज़ बूट (Lightning Boot)**: 400X तेज़ स्टार्टअप समय, 0.6GHz सिंगल-कोर पर भी 1 सेकंड में बूट हो जाता है।

🌍 **सच्ची पोर्टेबिलिटी (True Portability)**: RISC-V, ARM और x86 के लिए एक एकल (single) स्टैंडअलोन बाइनरी फ़ाइल। बस एक क्लिक और यह तैयार है!

🤖 **AI स्व-निर्मित (AI Self-built)**: स्वायत्त रूप से Go-नेटिव कार्यान्वयन — लूप में मानव परिशोधन (human refinement) के साथ 95% कोर (core) एजेंट-जनित है।

|                                | OpenClaw      | NanoBot                  | **ByteClaw**                           |
| ------------------------------ | ------------- | ------------------------ | -------------------------------------- |
| **भाषा (Language)**             | TypeScript    | Python                   | **Go**                                 |
| **RAM**                        | >1GB          | >100MB                   | **< 10MB**                             |
| **बूट समय (Boot Time)**</br>(0.8GHz core) | >500s         | >30s                     | **<1s**                                |
| **लागत (Cost)**                 | Mac Mini $599 | अधिकांश Linux बोर्ड ~$50 | **कोई भी Linux बोर्ड**</br>**$10 से शुरू** |

<img src="assets/compare.png" alt="ByteClaw" width="512">

## 🦾 डेमो (Demo)

### 🛠️ मानक सहायक वर्कफ़्लो (Standard Assistant Workflow)

<table align="center">
  <tr align="center">
    <th><p align="center">🧩 फुल-स्टैक इंजीनियर मोड</p></th>
    <th><p align="center">🗂️ लॉग और योजना प्रबंधन</p></th>
    <th><p align="center">🔎 वेब खोज और सीखना</p></th>
  </tr>
  <tr>
    <td align="center"><p align="center"><img src="assets/byteclaw_code.gif" width="240" height="180"></p></td>
    <td align="center"><p align="center"><img src="assets/byteclaw_memory.gif" width="240" height="180"></p></td>
    <td align="center"><p align="center"><img src="assets/byteclaw_search.gif" width="240" height="180"></p></td>
  </tr>
  <tr>
    <td align="center">विकास • परिनियोजन (Deploy) • स्केल</td>
    <td align="center">अनुसूची (Schedule) • स्वचालन (Automate) • याद रखें</td>
    <td align="center">खोजें (Discover) • अंतर्दृष्टि (Insights) • रुझान (Trends)</td>
  </tr>
</table>

### 📱 पुराने Android फोन पर चलाएं

अपने दस साल पुराने फोन को नया जीवन दें! ByteClaw के साथ इसे स्मार्ट AI असिस्टेंट में बदलें। त्वरित शुरुआत:

1. **Termux स्थापित करें (Install Termux)** (F-Droid या Google Play पर उपलब्ध)।
2. **कमांड चलाएँ:**

```bash
# नोट: नवीनतम संस्करण के लिए v0.1.1 को बदलें
wget https://github.com/ajmaluk/byteclaw/releases/download/v0.1.1/byteclaw-linux-arm64
chmod +x byteclaw-linux-arm64
pkg install proot
termux-chroot ./byteclaw-linux-arm64 onboard
```

फ़िर सेट अप पूरा करने के लिए "त्वरित प्रारंभ (Quick Start)" अनुभाग में दिए गए निर्देशों का पालन करें!

<img src="assets/termux.png" alt="ByteClaw" width="512">

## 📦 स्थापना (Installation)

### प्री-कम्पाइल बाइनरी के साथ स्थापित करें

[रिलीज़ (Releases) पृष्ठ](https://github.com/ajmaluk/byteclaw/releases) से अपने प्लेटफ़ॉर्म के लिए बाइनरी डाउनलोड करें।

### स्रोत से स्थापित करें (नवीनतम विशेषताएं, विकास के लिए अनुशंसित)

```bash
git clone https://github.com/ajmaluk/byteclaw.git

cd byteclaw
make deps

# निर्माण (Build), स्थापित करने की आवश्यकता नहीं
make build

# एकाधिक प्लेटफार्मों के लिए निर्माण (Build for multi-platforms)
make build-all

# निर्माण और स्थापित करें (Build and Install)
make install
```

## 🚀 त्वरित प्रारंभ (Quick Start)

> [!TIP]
> अपनी API Key को `~/.byteclaw/config.json` में सेट करें।
> API Key प्राप्त करें: [OpenRouter](https://openrouter.ai/keys) (LLM) · [Zhipu](https://open.bigmodel.cn/usercenter/proj-mgmt/apikeys) (LLM)
> वेब खोज **वैकल्पिक (optional)** है — मुफ़्त [Brave Search API](https://brave.com/search/api) प्राप्त करें (प्रत्येक माह 2000 निःशुल्क प्रश्न)।

**1. प्रारंभ करें (Initialize)**

```bash
byteclaw onboard
```

**2. कॉन्फ़िगर करें (Configure)** (`~/.byteclaw/config.json`)

```json
{
  "model_list": [
    {
      "model_name": "gpt4",
      "model": "openai/gpt-5.2",
      "api_key": "sk-your-openai-key"
    }
  ],
  "agents": {
    "defaults": {
      "model_name": "gpt4"
    }
  }
}
```

> **विस्तृत (Advanced) कॉन्फ़िगरेशन विवरण के लिए, कृपया [अंग्रेजी README](README.md) का संदर्भ लें।**

**3. चैट करें (Chat)**

```bash
byteclaw agent -m "2+2 क्या है?"
```

## 💬 चैट ऐप्स (Chat Apps)

टेलीग्राम (Telegram), डिस्कॉर्ड (Discord), या अन्य चैट एप्लिकेशन्स के माध्यम से अपने ByteClaw के साथ चैट करें।
चैट ऐप्स सेटअप और एकीकरण (integration) के विस्तृत निर्देशों के लिए [अंग्रेजी दस्तावेज़ (English documentation)](README.md) देखें।

## ⚙️ विन्यास (Configuration)

ByteClaw डिफ़ॉल्ट रूप से एक सुरक्षित सैंडबॉक्स (sandbox) वातावरण में चलता है। अधिक विस्तृत सुरक्षा सेटिंग्स, आवधिक कार्यों (Periodic Tasks / Heartbeat) और प्रदाता कॉन्फ़िगरेशन (Provider Configurations) के लिए, कृपया [English README](README.md) को देखें।
