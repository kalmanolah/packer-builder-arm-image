package main

import (
	"flag"
	"log"
	"os"

	"github.com/hashicorp/packer/packer"
	"github.com/solo-io/packer-builder-arm-image/pkg/flasher"
)

func main() {

	device := flag.String("device", "", "device to flash to. leave empty for auto detect")
	image := flag.String("image", "", "image to flash. leave empty for auto detect")
	interactive := flag.Bool("interactive", true, "use interactive mode")
	verify := flag.Bool("verify", true, "verify that image was written")
	flag.Parse()

	flashercfg := flasher.FlashConfig{
		Image:          *image,
		Device:         *device,
		NotInteractive: !*interactive,
		Verify:         *verify,
	}
	var ui packer.Ui = &packer.BasicUi{
		Reader:      os.Stdin,
		Writer:      os.Stdout,
		ErrorWriter: os.Stdout,
	}

	flshr := flasher.NewFlasher(ui, flashercfg)
	err := flshr.Flash()
	if err != nil {
		log.Fatal(err)
	} else {
		ui.Say("flashed successfully")
	}
}