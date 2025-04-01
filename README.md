# Coze TTS

一个基于扣子API的文字转语音工具，支持多种声音和语速调节。

## 功能特点

- 支持多种声音选择
- 可调节语音速度
- 通过GitHub Actions自动化运行
- 简单易用的命令行接口

## 安装说明

1. 克隆项目到本地：
```bash
git clone https://github.com/YOUR_USERNAME/coze.git
cd coze
```

2. 安装Go依赖：
```bash
go mod download
```

## 使用方法

### 本地运行

1. 设置环境变量：
```bash
export token="your_coze_api_token"
```

2. 运行程序：
```bash
go run main.go <voice_id> <text> <speed>
```

参数说明：
- voice_id: 语音ID
- text: 要转换的文本
- speed: 语速（例如：1.0为正常速度）

### GitHub Actions工作流使用

本项目包含两个GitHub Actions工作流：

1. GetVoiceList：获取可用的语音列表
2. TTS：执行文字转语音任务

要使用GitHub Actions，请按以下步骤操作：

1. Fork本项目到你的GitHub账号

2. 在你的项目设置中添加Secrets：
   - 进入项目的Settings > Secrets and variables > Actions
   - 点击"New repository secret"
   - 添加名为`TOKEN`的secret，值为你的Coze API Token

3. 工作流将自动运行，或者你可以手动触发：
   - 进入Actions标签页
   - 选择需要运行的工作流
   - 点击"Run workflow"

### 获取音色列表和单元测试

1. 获取音色列表：
   - 工作流会自动运行并获取最新的音色列表
   - 生成的音色列表文件会保存在工件（Artifacts）中
   - 查看方法：
     1. 进入Actions页面
     2. 点击最新的GetVoiceList工作流运行记录
     3. 在页面底部的Artifacts区域下载音色列表文件

2. 单元测试：
   - 项目包含了完整的单元测试用例
   - 测试文件位于`voice/unit_test.go`
   - 本地运行测试：
     ```bash
     go test ./voice -v
     ```
   - GitHub Actions中的测试：
     - 每次推送代码时会自动运行测试
     - 可以在Actions页面查看测试结果
     - 测试报告会作为工件保存

### 工件（Artifacts）说明

工作流运行后会生成以下工件：

1. 音色列表文件：
   - 包含所有可用的语音ID和描述
   - 文件格式为JSON
   - 位于GetVoiceList工作流的Artifacts中

2. 测试报告：
   - 包含单元测试的详细结果
   - 可在每次工作流运行后查看

3. 生成的语音文件：
   - TTS工作流生成的语音文件
   - 保存在Artifacts中的voices目录下

## Token获取和设置

1. 获取Coze API Token：
   - 访问Coze官网并登录
   - 在个人设置中找到API Token选项
   - 生成并复制你的API Token

2. Token使用方式：
   - 本地开发：设置环境变量
   - GitHub Actions：设置Repository Secrets

## 注意事项

- 请妥善保管你的API Token，不要直接在代码中硬编码
- 确保在使用GitHub Actions时正确设置了Token
- 语音生成文件将保存在voice/downloads/voices/目录下

## 许可证

本项目基于MIT许可证开源，详见[LICENSE](LICENSE)文件。
