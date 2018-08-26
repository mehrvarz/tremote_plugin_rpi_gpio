# TRemote Plugin rpi_gpio

TRemote is a service for ARM based Linux computers. It lets you remote control *things* on these kind of machines, specifically over Bluetooth. There is no limit to what you can remote control. You can access a list of predefined actions, you can execute executables and shell scripts, you can issue http request, and you can invoke your own or 3rd party native code plugins.

This repository contains the complete Go source code for a remote control plugin. You can use this plugin as-is. You can also use it as a template to implement extended functionality.

This particular plugin is for **Raspberry Pi computers.**

TRemote plugin rpi_gpio in the given form will set a specicified GPIO pin to HIGH when a certain remote control button is pressed. And when the button is released, the GPIO pin will be set back to LOW. Button and pin number can be set in the plugin code. They can also be specified from the the central service mapping. See **Button mapping** below.


# Building the plugin

TRemote plugins are based on Go Modules. You need to use [Go v1.11](https://dl.google.com/go/go1.11.linux-armv6l.tar.gz) (direct dl link for linux-armv6l) to build TRemote plugins. The "go version" command should return "go version go1.11 linux/arm".

After cloning the repository enter the following command to build the plugin:

```
CGO_ENABLED=1 go build -buildmode=plugin
```
This will create the "rpi_gpio.so" binary. Copy the binary over to your Tremote Host folder, add a mapping entry like the one shown below to your mapping.txt file and restart TRemote service. You can now invoke your plugin functionality from a Bluetooh remote control.

Note that the rpi_gpio plugin makes use of project [go-rpio](https://github.com/stianeikeland/go-rpio). This package will be automatically fetched by "go build".


# Button mapping

The following is a one-line entry for the "mapping.txt" file of the TRemote service. 

This entry will link a specific remote control button (here P8) to a plugin (rpi_gpio.so) and hand over an argument (10). The argument in this paricular case will be used as a GPIO pin number. "gpio" is just a label name:


```
P8, gpio, rpi_gpio|10
```

It is important to understand that plugins know nothing about remote controls, about Bluetooth or how a button event is delivered to them. They only care about the execution of a particular action. The mapping file bindes the two sides together.


# Example console log

```
2018-08-26 19:37:51.09 INFO   mapping button_pressed: P8 (currentlyPressedBitmap=80)
2018-08-26 19:37:51.09 INFO   rpi_gpio button pressed: set pin 10 high
2018-08-26 19:37:51.54 INFO   mapping button_released: P8 (currentlyPressedBitmap=00)
2018-08-26 19:37:51.54 INFO   rpi_gpio button released: clear pin 10 low
```


# Accessing GPIO via /dev/gpiomem memory (without root)

This plugin tries to access the Raspberry Pi GPIO pins via /dev/gpiomem memory. In order for this to work without root, you may need to set specific permissions for "/dev/gpiomem". Check the current permissions first. To do so enter "ls -l /dev/gpiomem". You should see: "crw-rw---- 1 root gpio". If this is not what you get, enter the following:

```
sudo chown root.gpio /dev/gpiomem
sudo chmod g+rw /dev/gpiomem
```




