KITEX_PATH=/home/kangning/go/bin/kitex

tidy:
	go mod tidy

install:
	go install github.com/cloudwego/kitex/tool/cmd/kitex@latest

update_kitex:
	${KITEX_PATH} -module TTMS -I=./api ./api/user.proto
	${KITEX_PATH} -module TTMS -I=./api ./api/studio.proto
	${KITEX_PATH} -module TTMS -I=./api ./api/play.proto
	${KITEX_PATH} -module TTMS -I=./api ./api/order.proto
	${KITEX_PATH} -module TTMS -I=./api ./api/ticket.proto
