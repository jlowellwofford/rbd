module github.com/bensallen/rbd

go 1.14

require (
	github.com/go-openapi/errors v0.19.9
	github.com/go-openapi/loads v0.20.0
	github.com/go-openapi/runtime v0.19.24
	github.com/go-openapi/spec v0.20.0
	github.com/go-openapi/strfmt v0.19.11
	github.com/go-openapi/swag v0.19.12
	github.com/go-openapi/validate v0.20.0
	github.com/google/goexpect v0.0.0-20200816234442-b5b77125c2c5 // indirect
	github.com/jessevdk/go-flags v1.4.0
	github.com/rekby/gpt v0.0.0-20200614112001-7da10aec5566 // indirect
	github.com/spf13/pflag v1.0.5
	github.com/u-root/u-root v7.0.0+incompatible
	golang.org/x/net v0.0.0-20201202161906-c7110b5ffcbb
	golang.org/x/sys v0.0.0-20201119102817-f84b799fce68
	golang.org/x/tools v0.0.0-20201119174615-0557df368a99 // indirect
)

replace github.com/u-root/u-root v7.0.0+incompatible => github.com/u-root/u-root v1.0.1-0.20201119150355-04f343dd1922
