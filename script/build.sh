echo "0"
go build -o ttms_web ../internal/web
echo "1"
go build -o ttms_user ../internal/user
echo "2"
go build -o ttms_play ../internal/play
echo "3"
go build -o ttms_ticket ../internal/ticket
echo "4"
go build -o ttms_studio ../internal/studio
echo "5"
go build -o ttms_order ../internal/order
echo "6"