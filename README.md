# TRemote_plugin_rpi_gpio

TRemote is a service for ARM based Linux computers. It lets you remote control *things* on that machine, specifically over Bluetooth. There is no limit to what you can remote control. You can make use of a list of predefined actions, you can execute executables and shell scripts, you can issue http request, and you can also invoke native code plugins.

This repository contains the complete Go source code for a remote controllable plugin. You can compile and use this plugin as-is. You can also use it as a template to extended the functionality.

TRemote_plugin_rpi_gpio in the given form will set a certain GPIO pin to HIGH when a certain remote control button is pressed. And when this button is released, it will set that GPIO pin back to LOW. Button and pin number can be specified via the central mapping file. Example: This is how you would connect button P8 with label "gpio" to plugin rpi_gpio and hand over "10" as an argument, which will be used as GPIO pin number:


```
P8, gpio, rpi_gpio|10
```

TRemote_plugin_rpi_gpio makes use of [go-rpio](https://github.com/stianeikeland/go-rpio)


# Compiling

TRemote_plugin_rpi_gpio makes use of Go Modules. You need Go v1.11. "go version" should return "go version go1.11 linux/arm".

Clone this repository and enter:

```
CGO_ENABLED=1 go build -buildmode=plugin $1 rpi_gpio.go
```
This will create binary "rpi_gpio.so". Copy this binary to the tremotehost folder and add an entry like the one above to mapping.txt. Restart the TRemote service and invoke the plugin via the remote control.


