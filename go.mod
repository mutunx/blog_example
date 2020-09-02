module ginExample

go 1.12

require (
	github.com/alecthomas/template v0.0.0-20190718012654-fb15b899a751
	github.com/astaxie/beego v1.9.3-0.20171218111859-f16688817aa4
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/gin-gonic/gin v1.4.0
	github.com/go-ini/ini v1.57.0
	github.com/golang/protobuf v1.3.2 // indirect
	github.com/jinzhu/gorm v1.9.15
	github.com/json-iterator/go v1.1.7 // indirect
	github.com/lib/pq v1.2.0 // indirect
	github.com/mattn/go-isatty v0.0.9 // indirect
	github.com/smartystreets/goconvey v1.6.4 // indirect
	github.com/swaggo/gin-swagger v1.2.0
	github.com/swaggo/swag v1.5.1
	github.com/ugorji/go v1.1.7 // indirect
	github.com/unknwon/com v1.0.1
	gopkg.in/ini.v1 v1.57.0 // indirect
	gopkg.in/yaml.v2 v2.2.3 // indirect
)

replace github.com/EDDYCJY/go-gin-example/pkg/setting => ./pkg/setting
