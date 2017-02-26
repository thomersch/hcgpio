# HC-GPIO

Control a GPIO pin (e.g. on a Raspberry Pi) over Apple HomeKit. This can be used for example for turning lights on and off.

**Important notice:** Unfortunately the hc dependency is under a weird non-opensource license, which is unfortunate, because I can't put this under an BSD/MIT license which I would love to.

## Requirements

* Go ≥ 1.7
* all go deps are included in `vendor/`

## Compiling

	go build

	# Cross compiling for a Raspberry
	GOARCH=arm GOARM=5 GOOS=linux go build

## Starting on Boot

Place the `hcgpio.service` file in `/etc/systemd/system` and enable the service: `systemctl enable hcgpio.service`. You can start it immediately with `service hcgpio start`. Note that it might be necessary to customize the paths in the service file.
