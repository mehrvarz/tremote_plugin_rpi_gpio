echo build plugin
CGO_ENABLED=1 go build -buildmode=plugin $1 rpi_gpio.go

