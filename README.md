# toggl-report

serverless framework  
go1.12.1

## What's this
- TogglのレポートをSlackに出力するサーバーレスアプリケーション
- 前日分のレポートを出力する

## Toggl API で Workspace IDを取得する方法
Basic Authで以下のEndpointを実行する
```
https://www.toggl.com/api/v8/workspaces
```