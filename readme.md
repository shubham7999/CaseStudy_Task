==> The service is responsible for retreiving hostnames having less or equals X active Ip address actives..

==>We have json file in etc/mock.json path which will mock the data to be given for ip , hostname and status

==>To start the service just extract the folder IpTask after extracting just run go run main.go Then check on browser and hit  : http://localhost:10000/count/domain on browser you will get the expected result[Go should be installed in your system]

==>You can also change the input of ip , domain and status by changing json file in path etc/mock.json

==>Entry point of server is from main function .. 

==>There are mainly it divided into route , model , controller
route --> Used to write the routes for http request
model --> used to store the template of our data
controller --> Its a main controlling body useed to define the handler function
etc/mock.json --> contains the json file for the input of ip, domain and active field

==>Use these command for checking test cases coverage 
1) go build
2) go test ./... -coverprofile=coverage.out
3) go tool cover -func=coverage.out 

==> Output hoststring will be return in lexographical order..

