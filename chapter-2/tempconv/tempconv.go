package tempconv

import "fmt"

type Celsius float64
type Fahrenheit float64
type Kelvin float64

const (
	// Celsius and Kelvin difference.
	ckDiff = 273.15

	AbsoluteZeroC Celsius = -273.15
	BoilingC      Celsius = 100
	FreezingC     Celsius = 0
)

func (c Celsius) String() string    { return fmt.Sprintf("%g°C", c) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%g°F", f) }
func (k Kelvin) String() string     { return fmt.Sprintf("%gK", k) }
