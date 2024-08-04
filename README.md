# Hugo ChatGPT Convert (hcc)

## これは何？

HugoのMarkdownをChatGPTを使って変換するためのツールです。  
自分のブログの記事を英語に翻訳するために作りました。

## 使い方

1. ChatGPTのAPIキーを取得します。
2. config.yamlを編集します。
3. `hcc`を実行します。

```bash
$ hcc -input input.md -output output.md

# or

$ cat input.md | hcc > output.md
```

## config

```yaml
---
OpenAI:
  APIToken: "sk-xxxxxx"
  Model: "gpt-4o-mini"
  MaxTokens: 100
Prompt: |
  # Command
  * Convert the Markdown of a hugo article according to instructions
  * Do not convert within comment blocks
  * Output should be in Markdown only

  # Instructions
  * 記事を英語に翻訳する
```

Promptに変換の指示を記述します。
