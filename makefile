
# 基礎環境設定
# 啟動dev/build檔的dev docker compose yaml
UpDevInfra:
	cd build/dev && docker-compose up -d

# 關閉dev/build檔的dev docker compose yaml
DownDevInfra:
	cd build/dev && docker-compose down -v
