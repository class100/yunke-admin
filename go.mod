module github.com/class100/yunke-admin

go 1.14

require (
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/go-resty/resty/v2 v2.3.0
	github.com/class100/yunke-core v1.0.4
	github.com/storezhang/gox v1.2.8
	golang.org/x/text v0.3.2 // indirect
)

// replace github.com/storezhang/gox => ../../storezhang/gox
// replace github.com/class100/yunke-core => ../yunke-core
