<div align="center">
  <h1 align="center">
      Suno AI API
  </h1>
  <p>通过 API 调用 suno 实现AI音乐生成。</p>
</div>

## 简介

Suno AI，简称Suno，是一款生成式人工智能音乐创作程序，旨在产生人声与乐器相结合的逼真歌曲。2023年12月20日，Suno
AI在推出网络应用程序并与微软建立合作关系（微软将Suno作为插件纳入Microsoft Copilot）后，开始广泛使用。
**本项目是基于Suno AI网页端逆向封装的API接口，采用golang语言开发，仅供学习与研究使用。**

感谢这个nodejs项目提供的思路：[suno-api](https://github.com/gcui-art/suno-api)

## 开始使用

### 1. 获取你的 app.suno.ai 账号的 cookie

1. 浏览器访问 [app.suno.ai](https://app.suno.ai)
2. 打开浏览器的控制台：按下 `F12` 或者`开发者工具`
3. 选择`网络`标签
4. 刷新页面
5. 找到包含`client?_clerk_js_version`关键词的请求
6. 点击并切换到 `Header` 标签
7. 找到 `Cookie` 部分，鼠标复制 Cookie 的值

![获取cookie](https://github.com/gcui-art/suno-api/blob/main/public/get-cookie-demo.gif)

### 2. 克隆项目

```bash
git clone https://github.com/supercat0867/suno-api.git
cd suno-api
```

### 3. 在项目目录中创建一个 .env 文件，并添加环境变量

```dotenv
SUNO_COOKIE="步骤1中的cookie"
```

### 4.部署

部署方式有编译运行部署也可以docker容器部署。

### 编译运行部署，需要安装go环境

1. 安装所有的依赖项

```shell
go mod download
```

2. 编译

```shell
go build -o suno-api
```

3. 配置环境变量

```shell
export SUNO_COOKIE="步骤1中的cookie"
```

4. 运行

```shell
./suno-api
```

### docker容器部署

1. 构建镜像

```shell
docker build -t suno-api:latest .
```

2. 运行镜像

```shell
docker run -p 3000:3000 --env-file .env suno-api:latest
```

### 文档地址

http://localhost:3000/swagger/index.html

### 请求示例

1. 创建音乐生成任务

`POST:http://localhost:3000/api/v1/suno/createTask`

请求参数：

```json
{
  "prompt": "主歌A1在这繁华世界里寻找一份真挚的心你的笑容温暖如阳光照亮我每一个梦境马小姐你就像那夜空中最亮的星让我迷失在你的光芒里无法抗拒每一次眼神交汇心中都泛起涟漪你的温柔是我最美的记忆深深烙印在心底马小姐你的每一个瞬间我都想要铭记在这爱情的旅途中与你携手同行副歌B你是我心中的唯一马小姐我为你着迷你的每个微笑都让我沉醉不已我愿意陪你走过每一个四季在这漫长的岁月里给你我所有的深情主歌A2时间匆匆流逝我们的故事还在继续你的每个动作和每句话语都让我心动不已马小姐你是我生命中最美的奇遇我愿意用我全部的热情去守护这份爱情副歌B（重复）你是我心中的唯一马小姐我为你着迷你的每个微笑都让我沉醉不已我愿意陪你走过每一个四季在这漫长的岁月里给你我所有的深情过渡段爱情就像一场奇妙的旅行有你在身边一切都变得如此美丽马小姐你是我不可或缺的伴侣让我们一起走过这段不平凡的旅程副歌B（再重复）你是我心中的唯一马小姐我为你着迷你的每个微笑都让我沉醉不已我愿意陪你走过每一个四季在这漫长的岁月里给你我所有的深情尾声马小姐我的爱将永远属于你在这漫长的人生路上你是我唯一的伴侣让我们一起走过风风雨雨守护这份来之不易的爱情",
  "tags": "rap,electronic",
  "title": "送给马小姐的歌",
  "callback": ""
}
```

响应格式：

```json
{
  "code": 200,
  "msg": "success",
  "data": [
    {
      "id": "309817d3-36fe-4400-8c6d-d874d907ec83"
    },
    {
      "id": "123c991a-d1ee-45b5-9e04-42be3cd7fcf5"
    }
  ]
}
```

2. 获取音乐生成结果

`GET:http://localhost:3000/api/v1/suno/getStatus?songId=<音乐id>`

响应格式：

```json
{
  "code": 200,
  "msg": "success",
  "data": {
    "id": "41e6214c-c944-4318-ac1b-368ba8b23059",
    "video_url": "https://cdn1.suno.ai/41e6214c-c944-4318-ac1b-368ba8b23059.mp4",
    "audio_url": "https://cdn1.suno.ai/41e6214c-c944-4318-ac1b-368ba8b23059.mp3",
    "image_url": "https://cdn2.suno.ai/image_41e6214c-c944-4318-ac1b-368ba8b23059.jpeg",
    "image_large_url": "https://cdn2.suno.ai/image_large_41e6214c-c944-4318-ac1b-368ba8b23059.jpeg",
    "status": "complete"
  }
}
```