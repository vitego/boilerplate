module github.com/vitego/boilerplate

go 1.17

require (
	github.com/ermos/annotation v0.0.0-20210309132609-a4e71ea8028f
	github.com/vitego/config v1.0.0
	github.com/vitego/router v1.0.0
)

require (
	github.com/dgrijalva/jwt-go v3.2.0+incompatible // indirect
	github.com/julienschmidt/httprouter v1.3.0 // indirect
	github.com/rs/cors v1.8.2 // indirect
	github.com/valyala/fastjson v1.6.3 // indirect
)

replace (
	github.com/ermos/annotation v0.0.0-20210309132609-a4e71ea8028f => ../../ermos/annotation
	github.com/vitego/router v1.0.0 => ../router
)
