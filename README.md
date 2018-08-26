# TRemote plugin rpi_gpio

TRemote is a service for ARM based Linux computers. It lets you remote control *things* on these machines, specifically over Bluetooth. There is no limit to what you can remote control. You can access a list of predefined actions, you can execute executables and shell scripts, you can issue http request, and you can invoke native code plugins.

This repository contains the complete Go source code for a remote control plugin. You can use this plugin as-is. You can also use it as a template to implement extended functionality.

This particular plugin is for **Raspberry Pi computers.**

TRemote plugin rpi_gpio in the given form will set a certain GPIO pin to HIGH when a certain remote control button is pressed. And when this button is released, it will set that GPIO pin back to LOW. Button and pin number can be specified via the central mapping file. 

For example: This is how you can link remote control button P8 to plugin "rpi_gpio" and hand over "10" as argument. The argument in this case will be used as GPIO pin number. "gpio" is just a label:


```
P8, gpio, rpi_gpio|10
```

TRemote plugin rpi_gpio makes use of project [go-rpio](https://github.com/stianeikeland/go-rpio)


# Compiling

TRemote plugin rpi_gpio makes use of Go Modules. You must use Go v1.11 to build this project. The "go version" command should return "go version go1.11 linux/arm".

After cloning the repository enter the following to build the plugin:

```
CGO_ENABLED=1 go build -buildmode=plugin
```
This will create the "rpi_gpio.so" binary. Copy the binary to your Tremote host folder and add one entry like the one shown above to your mapping.txt. Restart TRemote service and you can execute the new plugin via your remote control.

# Access GPIO via /dev/gpiomem memory (without root)

This plugin tries to access the Raspberry Pi GPIO pins via /dev/gpiomem memory. In order for this to work without root, you may need to set specific permissions for "/dev/gpiomem". Check the current permissions first. To do so enter "ls -l /dev/gpiomem". You should see: "crw-rw---- 1 root gpio". If this is not what you get, enter the following:

```
sudo chown root.gpio /dev/gpiomem
sudo chmod g+rw /dev/gpiomem
```

