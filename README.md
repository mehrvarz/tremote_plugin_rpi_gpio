# TRemote plugin rpi_gpio

TRemote is a service for ARM based Linux computers. It lets you remote control *things* on these machines, specifically over Bluetooth. There is no limit to what you can remote control. You can access a list of predefined actions, you can execute executables and shell scripts, you can issue http request, and you can invoke native code plugins.

This repository contains the complete Go source code for a remote control plugin. You can use this plugin as-is. You can also use it as a template to implement extended functionality.

TRemote_plugin_rpi_gpio in the given form will set a certain GPIO pin to HIGH when a certain remote control button is pressed. And when this button is released, it will set that GPIO pin back to LOW. Button and pin number can be specified via the central mapping file. Example: This is how you would connect button P8 with label "gpio" to plugin rpi_gpio and hand over "10" as an argument, which will be used as GPIO pin number:


```
P8, gpio, rpi_gpio|10
```

TRemote_plugin_rpi_gpio makes use of project [go-rpio](https://github.com/stianeikeland/go-rpio)


# Compiling

TRemote_plugin_rpi_gpio makes use of Go Modules. You need to use Go v1.11. The "go version" command should return "go version go1.11 linux/arm".

After cloning this repository enter the following to build the plugin:

```
CGO_ENABLED=1 go build -buildmode=plugin
```
This will create the "rpi_gpio.so" binary. Copy the binary to your tremotehost folder and add an entry like the one shown above to your mapping.txt. Then restart the TRemote service and execute the new plugin via remote control.

# Access GPIO via /dev/gpiomem memory (without root)

For the plugin to be able to access the GPIO pins via /dev/gpiomem memory (without root), you may need to set permissions for /dev/gpiomem. First check the current permissions. Enter "ls -l /dev/gpiomem". You should see: "crw-rw---- 1 root gpio". If this is not what you get enter the following:

```
sudo chown root.gpio /dev/gpiomem
sudo chmod g+rw /dev/gpiomem
```


