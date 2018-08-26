# TRemote Plugin rpi_gpio

TRemote is a service for ARM based Linux computers. It lets you remote control *things* on these kind of machines, specifically over Bluetooth. There is no limit to what you can remote control. You can access a list of predefined actions, you can execute executables and shell scripts, you can issue http request, and you can invoke your own or 3rd party native code plugins.

This repository contains the complete Go source code for a remote control plugin. You can use this plugin as-is. You can also use it as a template to implement extended functionality.

This particular plugin is for **Raspberry Pi computers.**

TRemote plugin rpi_gpio in the given form will set a specicified GPIO pin to HIGH when a certain remote control button is pressed. And when the button is released, the GPIO pin will be set back to LOW. Button and pin number can be set in the plugin code. They can also be specified from the the central service mapping. See **Button mapping** below.


# Building the plugin

TRemote plugins are based on Go Modules. You need to use [Go v1.11](https://dl.google.com/go/go1.11.linux-armv6l.tar.gz) (direct dl link) to build TRemote plugins. The "go version" command should return "go version go1.11 linux/arm".

After cloning the repository enter the following command to build the plugin:

```
CGO_ENABLED=1 go build -buildmode=plugin
```
This will create the "rpi_gpio.so" binary. Copy the binary over to your Tremote Host folder, add a mapping entry like the one shown below to your mapping.txt file and restart TRemote service. You can now invoke your plugin functionality from a Bluetooh remote control.


# Button mapping

This is how you can link a remote control button (here P8) to a plugin (here rpi_gpio) and hand over an argument (here 10). The argument in this case will be used as the GPIO pin number. "gpio" is just a label name:


```
P8, gpio, rpi_gpio|10
```

Note that TRemote plugin rpi_gpio makes use of project [go-rpio](https://github.com/stianeikeland/go-rpio). This package will be automatically fetched.


# Accessing GPIO via /dev/gpiomem memory (without root)

This plugin tries to access the Raspberry Pi GPIO pins via /dev/gpiomem memory. In order for this to work without root, you may need to set specific permissions for "/dev/gpiomem". Check the current permissions first. To do so enter "ls -l /dev/gpiomem". You should see: "crw-rw---- 1 root gpio". If this is not what you get, enter the following:

```
sudo chown root.gpio /dev/gpiomem
sudo chmod g+rw /dev/gpiomem
```

