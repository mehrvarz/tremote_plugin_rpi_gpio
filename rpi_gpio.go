// rpi_gpio plugin needs proper access to /dev/gpiomem
// check permissions: ls -l /dev/gpiomem
// you should see: "crw-rw---- 1 root gpio"
// if you don't, set the correct permissions:
// sudo chown root.gpio /dev/gpiomem
// sudo chmod g+rw /dev/gpiomem
package main

import (
	"fmt"
	"strconv"
	"sync"
	"github.com/mehrvarz/tremote_plugin"
	"github.com/mehrvarz/log"
	"github.com/stianeikeland/go-rpio"
)

const (	
	pluginname          = "rpi_gpio"
)

var (
	log1                log.Logger
	instanceNumber      = 0
	rpioOpen            = false
)

func Action(log log.Logger, pid int, longpress bool, pressedDuration int64, rcs* tremote_plugin.RemoteControlSpec, ph tremote_plugin.PluginHelper, wg *sync.WaitGroup) error {
	log1 = log

	if instanceNumber==0 {
		firstinstance()
	}
	instanceNumber++

	strArray := rcs.StrArray
	if longpress {
		strArray = rcs.StrArraylong
	}

	pinnumber, err := strconv.Atoi(strArray[0])
	if err != nil {
		log1.Warningf("%s arg[0]=%s to int failed err=",pluginname,strArray[0],err.Error())
		return err
    }

    pin := rpio.Pin(pinnumber)
	
	if !rpioOpen {
		if err := rpio.Open(); err != nil {
			log1.Warningf("%s rpio.Open() failed err=%s",pluginname,err.Error())
			return err
		}
		//Because this plugin will stay in memory, we leave rpio open for upcoming tasks
		//defer rpio.Close()
		rpioOpen = true
       	log1.Infof("%s rpio opened; use pin %d for output",pluginname,pinnumber)
		pin.Output()
	}

	if pressedDuration==0 {
		log1.Infof("%s button pressed: pin %d high",pluginname,pinnumber)
		ph.PrintInfo(fmt.Sprintf("pin %d high",pinnumber))
		pin.High()
	} else {
		log1.Infof("%s button released: pin %d low",pluginname,pinnumber)
		ph.PrintInfo(fmt.Sprintf("pin %d low",pinnumber))
		pin.Low()
	}

	return nil
}

func firstinstance() {
	// run things here on first call only
}

