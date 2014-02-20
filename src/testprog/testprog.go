package main

import (
	"driver"
	"fmt"
	"time"
)

func main() {
	driver.IoInit()
	driver.SetBit(driver.LIGHT_STOP)
	driver.SetBit(driver.LIGHT_COMMAND1)
	driver.SetBit(driver.LIGHT_DOWN2)
	driver.SetBit(driver.MOTORDIR)
	driver.WriteAnalog(driver.MOTOR, 4000)
	time.Sleep(1 * time.Second)
	driver.ClearBit(driver.MOTORDIR)
	driver.WriteAnalog(driver.MOTOR, 4000)
	time.Sleep(1 * time.Second)
	driver.WriteAnalog(driver.MOTOR, 0)
	driver.ClearBit(driver.MOTORDIR)

	driver.ClearBit(driver.LIGHT_STOP)
	driver.ClearBit(driver.LIGHT_COMMAND1)
	driver.ClearBit(driver.LIGHT_DOWN2)
}