# toggl-report

serverless framework  
go1.15.6

## What's this
- TogglのレポートをSlackに出力するサーバーレスアプリケーション
- 前日分のレポートを出力する

## Toggl API で Workspace IDを取得する方法
Basic Authで以下のEndpointを実行する
```
https://api.track.toggl.com/api/v8/workspaces
```