# 개요
typora 에서 사용할 수 있는 image uploader 용 Server입니다.

# 사용 방법
.env.sample 파일을 복사 후, 설정에 맞게 수정 후 실행
```bash
cp .env.sample .env
```

Build 해서 사용하면 더욱 편합니다.
```bash
go build -o image-uploader main.go
nohup ./image-uploader &
```