# MySQLサーバを起動
db-up:
	docker compose up mysql -d

# MySQLサーバに接続
db-cli:
	mysql -h 127.0.0.1 -P 3306 -u root -proot

# マイグレーション
db-migrate:
	migrate -path db/migrations -database "mysql://root:root@tcp(127.0.0.1:3306)/test-invoice" up
