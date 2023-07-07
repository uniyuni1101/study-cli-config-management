# study-cli-config-management
設定が、コマンドライン引数やコンフィグファイルなどの複数から提供される場合にどのように対処するか理解するための学習リポジトリ

# 設定項目
- theme [string] [ simple | rich | retro ]
- rate [int] [1 ~ 120]
  
## コマンドライン引数
- theme
- rate
- help
- version

## コンフィグファイル
ファイル形式: json
```
{
    {"theme": "simple"},
    {"rate": 120}
}
```