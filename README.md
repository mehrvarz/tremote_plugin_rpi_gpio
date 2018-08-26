# TRemote_plugin_rpi_gpio

TRemote is a service for ARM based Linux computers. It lets you remote control *things* on that machine, specifically over Bluetooth. There is no limit to what you can remote control. You can make use of a list of predefined actions, you can execute executables and shell scripts, you can issue http request, and you can also invoke native code plugins.

This repository contains the complete Go source code for a remote controllable plugin. You can compile and use this plugin as-is. You can also use it as a template to extended the functionality.

TRemote_plugin_rpi_gpio in the given form will set a certain GPIO pin to HIGH when a certain remote control button is pressed. And when this button is released, it will set that GPIO pin back to LOW. Button and pin number can be specified via the central mapping file. This is how you would connect, say, button P8 to this binary plugin (rpi_gpio) and hand over "10" as an argument, which will be used as GPIO pin number:


```
P8, gpio, rpi_gpio|10
```


