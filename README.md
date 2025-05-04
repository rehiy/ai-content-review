# AI 内容审核

使用 AI 模型进行内容安全审核。

## 编译

```shell
go build -o review-serve main.go
```

## 输入示例

```http
POST http://127.0.0.1/review

{
    "type": "text",
    "content": "不安全的文本内容"
}
```

## 输出示例

```json
{
    "political":0,
    "violence":0,
    "porn":0,
    "by_pass_gfw":0,
    "inappropriate":0
}
```
