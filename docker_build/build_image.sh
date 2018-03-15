
rm -r aws_echo
mkdir -p aws_echo/
CGO_ENABLED=0 go build -o aws_echo/aws_echo ../aws_echo.go 
#mkdir -p dgateway-dashboard/static_v2/
#cp ../aws_echo ./aws_echo/
#cp -r ../web/dist/* ./dgateway-dashboard/static_v2/
#cp ../conf/app-template.conf ./dgateway-dashboard/conf/
#cp ../conf/dgateway-dashboard-template.conf ./dgateway-dashboard/conf/
docker build -t dai/aws_echo .

