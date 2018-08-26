# TRemote_plugin_rpi_gpio

TRemote is a service for ARM based Linux computers. It lets you remote control /things/ on that machine, specifically over Bluetooth. You can basically control anything. This works through a list of predefined internal actions, by execution of shell scripts and executables, by issuing http request, and by the ability to invoke native code in plugins.

This repository contains the complete Go source code for such a plugin. This plugin can be compiled and used as-is. It can also be used as a template to implement similar or extended functionality.

TRemote_plugin_rpi_gpio as-is will set a GPIO pin to HIGH when a specified remote control button is pressed. And when the button is released, it will set that GPIO pin back to LOW. Button and pin number can be specified in the mapping file. This is how you would tell the service to connect button P8 to plugin rpi_gpio and hand it over "10" as an argument, which will be used as GPIO pin number:


```
P8, gpio, rpi_gpio|10
```


