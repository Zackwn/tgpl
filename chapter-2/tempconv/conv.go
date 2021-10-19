package tempconv

// CtoF converts Celsius temperature to Fahrenheit.
func CtoF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

// CtoK converts Celsius temperature to Kelvin.
func CtoK(c Celsius) Kelvin { return Kelvin(c + ckDiff) }

// FtoC converts Fahrenheit temperature to Celsius.
func FtoC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

// FtoK converts Fahrenheit temperature to Kelvin.
func FtoK(f Fahrenheit) Kelvin { return CtoK(FtoC(f)) }

// KtoC converts Kelvin temperature to Celsius.
func KtoC(k Kelvin) Celsius { return Celsius(k - ckDiff) }

// KtoF converts Kelvin temperature to Fahrenheit.
func KtoF(k Kelvin) Fahrenheit { return CtoF(KtoC(k)) }
