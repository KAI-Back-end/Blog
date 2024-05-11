.SILENT:

# Адрес кеширующего прокси для модулей.
export GOPROXY=https://proxy.golang.org
# Включаем протокол sumdb.
export GOSUMDB=on

.PHONY:
run:
	go run cmd/blog/main.go