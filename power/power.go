package power

func IsPower(input int, base int) (int, bool) {
	power := 0
	for input > 1 {
		if input%base != 0 {
			return 0, false
		}
		power++
		input = input / base
	}
	return power, true
}
/*
Example of use
input := 27
base := 3
power, ok := power2.IsPower(input, base)
fmt.Println(ok)
if ok {
fmt.Println(input)
fmt.Printf("explanation: %v^%v=%v\n", base, power, input)
fmt.Println("validating: ", float64(input) == math.Pow(float64(base), float64(power)))
}*/