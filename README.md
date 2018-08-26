# TRemote_plugin_rpi_gpio

TRemote is a service for ARM based Linux computers. It enables you to remote control anything on that machine over Bluetooth. This works through predefined internal actions, by execution of shell scripts, by issuing http request, or by invoking native plugins.

This repository contains the source code for such a plugin. It can be compiled and used as-is. And it can be used as a template for extended functionality.

TRemote_plugin_rpi_gpio as-is will set a GPIO pin to HIGH when a specific remote control button is pressed. And it will set this GPIO pin back to LOW when the button is released. Button and pin number can be specified in the mapping configuration. This is how such a mapping entry may look like:

```
P8, gpio, rpi_gpio|10
```

Here button P8 with label "gpio" will call plugin "rpi_gpio" and hand it over one argument "10", which will be used as GPIO pin number.


