module github.com/class100/yunke-admin

go 1.14

require (
	github.com/class100/yunke-core v1.0.5
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/go-resty/resty/v2 v2.3.0
	github.com/storezhang/gox v1.2.32
)

// replace github.com/storezhang/gox => ../../storezhang/gox
// replace github.com/class100/yunke-core => ../yunke-core
