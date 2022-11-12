module github.com/titor999/infotecs-go-ewallet/client

go 1.18

require (
	github.com/ilyakaznacheev/cleanenv v1.3.0
	github.com/titor999/infotecs-go-ewallet/server v0.0.0
	google.golang.org/grpc v1.48.0
)

require (
	github.com/BurntSushi/toml v1.1.0 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/joho/godotenv v1.4.0 // indirect
	golang.org/x/net v0.0.0-20201021035429-f5854403a974 // indirect
	golang.org/x/sys v0.0.0-20220715151400-c0bba94af5f8 // indirect
	golang.org/x/text v0.3.3 // indirect
	google.golang.org/genproto v0.0.0-20200526211855-cb27e3aa2013 // indirect
	google.golang.org/protobuf v1.27.1 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
	olympos.io/encoding/edn v0.0.0-20201019073823-d3554ca0b0a3 // indirect
)

replace (
	github.com/titor999/infotecs-go-ewallet/server v0.0.0 => ../server
)