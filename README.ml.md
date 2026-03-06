<div align="center">
  <img src="assets/banner.png" alt="ByteClaw Banner" width="100%">
  <br>
  <img src="assets/logo.png" alt="ByteClaw Logo" width="256">

  <h1>ByteClaw: ഗോ (Go) ഭാഷയിലെ അത്യാധുനിക AI അസിസ്റ്റൻ്റ് (Ultra-Efficient AI Assistant)</h1>

  <h3>$10 ഹാർഡ്‌വെയർ · 10MB റാം (RAM) · 1s ബൂട്ട് · 皮皮虾，我们走！</h3>

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

[中文](README.zh.md) | [日本語](README.ja.md) | [Português](README.pt-br.md) | [Tiếng Việt](README.vi.md) | [Français](README.fr.md) | [English](README.md) | [हिन्दी](README.hi.md) | **മലയാളം**

</div>

---

🤖 ByteClaw എന്നത് വളരെ ഭാരം കുറഞ്ഞ (ultra-lightweight) ഒരു വ്യക്തിഗത AI അസിസ്റ്റൻ്റാണ്. ഇത് [nanobot](https://github.com/HKUDS/nanobot)-ൽ നിന്നും പ്രചോദനം ഉൾക്കൊണ്ട് നിർമ്മിച്ചതാണ്. Go ഭാഷ ഉപയോഗിച്ച് സെൽഫ് ബൂട്ട്‌സ്ട്രാപ്പിംഗ് പ്രക്രിയയിലൂടെ ഇത് പൂർണ്ണമായും റീഫാക്ടർ ചെയ്യപ്പെട്ടു.

⚡️ വെറും $10 ഹാർഡ്‌വെയറിലും <10MB റാമിലും ഇത് പ്രവർത്തിക്കുന്നു: OpenClaw-നേക്കാൾ 99% കുറഞ്ഞ മെമ്മറി ഉപയോഗവും Mac mini-യേക്കാൾ 98% വിലക്കുറവുമാണ് ഇതിന്റെ പ്രത്യേകത!

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
> **🚨 സുരക്ഷയും ഔദ്യോഗിക ചാനലുകളും (SECURITY & OFFICIAL CHANNELS)**
>
> * **ക്രിപ്റ്റോ ഇല്ല (NO CRYPTO):** ByteClaw-ന് ഔദ്യോഗിക ടോക്കണോ കോയിനോ **ഇല്ല**. `pump.fun` പോലെയുള്ള ട്രേഡിംഗ് പ്ലാറ്റ്‌ഫോമുകളിലെ അവകാശവാദങ്ങളെല്ലാം തട്ടിപ്പാണ് (**SCAM**).
> * **ഔദ്യോഗിക ഡൊമെയ്ൻ (OFFICIAL DOMAIN):** ഞങ്ങളുടെ **ഏക** ഔദ്യോഗിക വെബ്സൈറ്റ് **[byteclaw.io](https://byteclaw.io)** ആണ്. കമ്പനിയുടെ വെബ്സൈറ്റ് **[ajmaluk.com](https://ajmaluk.com)**.
> * **മുന്നറിയിപ്പ്:** മൂന്നാം കക്ഷികൾ മറ്റ് പല `.ai/.org/.com/.net/...` ഡൊമെയ്‌നുകളും രജിസ്റ്റർ ചെയ്തിട്ടുണ്ട്, അവ നമ്മുടേതല്ല.
> * **ശ്രദ്ധിക്കുക:** ഈ പ്രൊജക്റ്റ് വികസനത്തിൻ്റെ പ്രാരംഭ ഘട്ടത്തിലാണ്. നെറ്റ്‌വർക്ക് സുരക്ഷാ പ്രശ്‌നങ്ങൾ ഉണ്ടായേക്കാം. അതിനാൽ v1.0 റിലീസിന് മുമ്പ് ഇത് പ്രൊഡക്ഷൻ പരിതസ്ഥിതിയിൽ ഉപയോഗിക്കരുത്.

## 📢 വാർത്തകൾ (News)

2026-02-09 🎉 ByteClaw റിലീസ് ചെയ്തു! 10 ഡോളർ വിലയുള്ള ഹാർഡ്‌വെയറിലും കേവലം 10 മെഗാബൈറ്റിലും താഴെ റാമിലും പ്രവർത്തിക്കാൻ കഴിയുന്ന രീതിയിലാണ് ഇത് നിർമ്മിച്ചിരിക്കുന്നത്. 🤖 ByteClaw, നമുക്ക് തുടങ്ങാം!

## ✨ സവിശേഷതകൾ (Features)

🪶 **വളരെ ഭാരം കുറഞ്ഞത് (Ultra-Lightweight)**: 10MB-യിൽ താഴെ മെമ്മറി ഉപയോഗം — പ്രധാന പ്രവർത്തനങ്ങൾക്ക് Clawdbot-നേക്കാൾ 99% ചെറുതാണ്.

💰 **കുറഞ്ഞ ചെലവ് (Minimal Cost)**: $10 ഹാർഡ്‌വെയറിൽ പ്രവർത്തിക്കാൻ പര്യാപ്‍തമായത് — Mac mini-യേക്കാൾ 98% വിലക്കുറവ്.

⚡️ **വേഗത്തിലുള്ള ബൂട്ട് (Lightning Boot)**: 400 മടങ്ങ് വേഗത്തിലുള്ള സ്റ്റാർട്ടപ്പ് സമയം, 0.6GHz സിംഗിൾ-കോർ പ്രോസസ്സറിൽ പോലും 1 സെക്കൻഡിനുള്ളിൽ ബൂട്ട് ചെയ്യുന്നു.

🌍 **ട്രൂ പോർട്ടബിലിറ്റി (True Portability)**: RISC-V, ARM, x86 എന്നിവയ്‌ക്കായുള്ള ഒരൊറ്റ സ്റ്റാൻഡ്-എലോൺ ബൈനറി. ഒറ്റ ക്ലിക്കിലൂടെ പ്രവർത്തിപ്പിക്കാം!

🤖 **AI സ്വയം നിർമ്മിച്ചത് (AI Self-built)**: സ്വയമേ പ്രവർത്തിക്കുന്ന Go-നേറ്റീവ് നിർവ്വഹണം — ഇതിന്റെ 95% കോറും Agent സ്വയം നിർമ്മിച്ചതാണ്. 

|                                | OpenClaw      | NanoBot                  | **ByteClaw**                           |
| ------------------------------ | ------------- | ------------------------ | -------------------------------------- |
| **ഭാഷ (Language)**             | TypeScript    | Python                   | **Go**                                 |
| **റാം (RAM)**                  | >1GB          | >100MB                   | **< 10MB**                             |
| **ബൂട്ട് സമയം (Boot Time)**</br>(0.8GHz core)| >500s     | >30s                     | **<1s**                                |
| **ചെലവ് (Cost)**               | Mac Mini $599 | ഒട്ടുമിക്ക Linux ബോർഡുകളും ~$50 | **ഏതൊരു Linux ബോർഡും**</br>**$10 മുതൽ**|

<img src="assets/compare.png" alt="ByteClaw" width="512">

## 🦾 ഡെമോ (Demo)

### 🛠️ സ്റ്റാൻഡേർഡ് അസിസ്റ്റൻ്റ് വർക്ക്ഫ്ലോ (Standard Assistant Workflow)

<table align="center">
  <tr align="center">
    <th><p align="center">🧩 ഫുൾ-സ്റ്റാക്ക് എഞ്ചിനീയർ മോഡ്</p></th>
    <th><p align="center">🗂️ ലോഗ് & പ്ലാനിംഗ് മാനേജ്മെന്റ്</p></th>
    <th><p align="center">🔎 വെബ് തിരയലും പഠനവും (Web Search)</p></th>
  </tr>
  <tr>
    <td align="center"><p align="center"><img src="assets/byteclaw_code.gif" width="240" height="180"></p></td>
    <td align="center"><p align="center"><img src="assets/byteclaw_memory.gif" width="240" height="180"></p></td>
    <td align="center"><p align="center"><img src="assets/byteclaw_search.gif" width="240" height="180"></p></td>
  </tr>
  <tr>
    <td align="center">Develop • Deploy • Scale</td>
    <td align="center">Schedule • Automate • Memory</td>
    <td align="center">Discover • Insights • Trends</td>
  </tr>
</table>

### 📱 പഴയ ആൻഡ്രോയിഡ് ഫോണുകളിൽ പ്രവർത്തിപ്പിക്കുക (Run on old Android phones)

നിങ്ങളുടെ പഴയ ഫോണിന് പുതിയൊരു ജീവൻ നൽകുക! ByteClaw ഉപയോഗിച്ച് അതിനെ ഒരു സ്മാർട്ട് AI അസിസ്റ്റൻ്റാക്കി മാറ്റുക. പെട്ടെന്നുള്ള തുടക്കത്തിനായി:

1. **Termux ഇൻസ്റ്റാൾ ചെയ്യുക** (F-Droid അല്ലെങ്കിൽ Google Play-യിൽ ലഭ്യമാണ്).
2. **കമാൻഡുകൾ പ്രവർത്തിപ്പിക്കുക:**

```bash
# ശ്രദ്ധിക്കുക: ഏറ്റവും പുതിയ റിലീസിനായി v0.1.1 മാറ്റുക
wget https://github.com/ajmaluk/byteclaw/releases/download/v0.1.1/byteclaw-linux-arm64
chmod +x byteclaw-linux-arm64
pkg install proot
termux-chroot ./byteclaw-linux-arm64 onboard
```

തുടർന്ന് ബാക്കി കാര്യങ്ങൾ പൂർത്തിയാക്കാൻ "ക്വിക്ക് സ്റ്റാർട്ട് (Quick Start)" വിഭാഗത്തിലെ നിർദ്ദേശങ്ങൾ പാലിക്കുക!

<img src="assets/termux.png" alt="ByteClaw" width="512">

## 📦 ഇൻസ്റ്റലേഷൻ (Installation)

### മുൻകൂട്ടി കംപൈൽ ചെയ്ത ബൈനറി ഇൻസ്റ്റാൾ ചെയ്യുക

[റിലീസ് പേജിൽ (Releases page)](https://github.com/ajmaluk/byteclaw/releases) നിന്ന് നിങ്ങളുടെ പ്ലാറ്റ്‌ഫോമിനുള്ള ബൈനറി ഡൗൺലോഡ് ചെയ്യുക.

### സോഴ്സ് വഴി ഇൻസ്റ്റാൾ ചെയ്യുക (ഡെവലപ്പർമാർക്കായി നിർദ്ദേശിക്കുന്നത്)

```bash
git clone https://github.com/ajmaluk/byteclaw.git

cd byteclaw
make deps

# ബിൽഡ് (Build) ചെയ്യുക
make build

# ഒന്നിലധികം പ്ലാറ്റ്‌ഫോമുകൾക്കായി ബിൽഡ് ചെയ്യുക
make build-all

# ഇൻസ്റ്റാൾ ചെയ്യുക
make install
```

## 🚀 വേഗത്തിലുള്ള തുടക്കം (Quick Start)

> [!TIP]
> നിങ്ങളുടെ API Key `~/.byteclaw/config.json` എന്നതിൽ സജ്ജമാക്കുക.
> API Key-കൾ ലഭിക്കാൻ: [OpenRouter](https://openrouter.ai/keys) (LLM) · [Zhipu](https://open.bigmodel.cn/usercenter/proj-mgmt/apikeys) (LLM)
> വെബ് തിരയൽ **നിർബന്ധമല്ല (optional)** — തികച്ചും സൗജന്യമായ [Brave Search API](https://brave.com/search/api) ഉപയോഗിക്കാവുന്നതാണ് (പ്രതിമാസം 2000 സൗജന്യ അന്വേഷണങ്ങൾ).

**1. ആരംഭിക്കുക (Initialize)**

```bash
byteclaw onboard
```

**2. ക്രമീകരിക്കുക (Configure)** (`~/.byteclaw/config.json`)

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

> **കൂടുതൽ വിപുലമായ ക്രമീകരണ വിവരങ്ങൾക്ക്, ദയവായി [ഇംഗ്ലീഷ് README](README.md) നോക്കുക.**

**3. ചാറ്റ് ചെയ്യുക (Chat)**

```bash
byteclaw agent -m "2 ഉം 2 ഉം കൂട്ടിയാൽ എത്രയാണ് (What is 2+2)?"
```

## 💬 ചാറ്റ് ആപ്പുകൾ (Chat Apps)

ടെലിഗ്രാം (Telegram), ഡിസ്കോർഡ് (Discord), മറ്റ് ചാറ്റ് ആപ്പുകൾ എന്നിവ വഴി ByteClaw-യുമായി ആശയവിനിമയം നടത്താം. 
കൂടുതൽ വിവരങ്ങൾക്കും സജ്ജീകരണങ്ങൾക്കുമായി [ഇംഗ്ലീഷ് പതിപ്പ് (English Version)](README.md) കാണുക.

## ⚙️ കോൺഫിഗറേഷൻ (Configuration)

ByteClaw സ്വാഭാവികമായും ഒരു സുരക്ഷിത സാൻഡ്ബോക്സ് (sandbox) പരിതസ്ഥിതിയിലാണ് പ്രവർത്തിക്കുന്നത്. സുരക്ഷാ ക്രമീകരണങ്ങൾ അല്ലെങ്കിൽ പീരിയോഡിക് ടാസ്കുകൾ (Periodic Tasks) പോലെ കൂടുതൽ പ്രവർത്തനങ്ങൾക്കായി, [ഇംഗ്ലീഷ് റിപ്പോസിറ്ററി (English README)](README.md) നോക്കുക.
