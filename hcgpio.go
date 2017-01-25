package main

import (
	"flag"
	"log"

	"github.com/brutella/hc"
	"github.com/brutella/hc/accessory"
	"github.com/kidoman/embd"
	_ "github.com/kidoman/embd/host/rpi"
)

func main() {
	pin := flag.String("pin", "01234567", "PIN for paring the device with HomeKit")
	gpiopin := flag.String("gpio", "P1_3", "Physical connector ID on the board that will be switched")
	accessoryName := flag.String("name", "Lamp", "Name of the device, which will be display in HomeKit")
	flag.Parse()

	setupGPIO(*gpiopin)

	info := accessory.Info{
		Name: *accessoryName,
	}

	acc := accessory.NewSwitch(info)

	acc.Switch.On.OnValueRemoteUpdate(func(on bool) {
		if on {
			embd.DigitalWrite(*gpiopin, embd.High)
		} else {
			embd.DigitalWrite(*gpiopin, embd.Low)
		}
		log.Println(on)
	})

	config := hc.Config{Pin: *pin}
	t, err := hc.NewIPTransport(config, acc.Accessory)
	if err != nil {
		log.Panic(err)
	}

	hc.OnTermination(func() {
		t.Stop()
	})

	t.Start()
}

func setupGPIO(gpiopin string) {
	err := embd.InitGPIO()
	if err != nil {
		log.Fatalf("gpio init failed: %v", err)
	}
	err = embd.SetDirection(gpiopin, embd.Out)
	if err != nil {
		log.Fatalf("direction setting failed: %v", err)
	}
}
