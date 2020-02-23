# lambda-line-notifier
## 動作環境
macOSまたはLinux。Windowsの場合は別のコマンドでbuildやLambdaへのdeployを行う必要がある

## 前提
* Lambda側で関数lineNotifierを作成し、TriggerとしてSQSを設定している
* lineNotifierのハンドラに`main`を登録している
* lineNotifierの環境変数`LINE_NOTIFY_ACCESS_TOKEN`にLINE Notifyのアクセストークンを登録している
* AWS CLIがインストールされている

## セットアップ
```bash
$ GO111MODULE=on go get -v -d # Install dependencies
$ GOOS=linux GOARCH=amd64 go build -v -o main main.go && zip main.zip main # Build & make zipfile
$ aws lambda update-function-code --function-name lineNotifier --zip-file fileb://main.zip --publish # Deploy to AWS Lambda
```
