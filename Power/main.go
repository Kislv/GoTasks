package Power 

func Power(digit, log int) (int){	
	result := 1
	digitInPower := digit
	for (log > 0){
		odd := log & 1
		if odd == 1 {
			result = (result * digitInPower)
		}
		log = log >> 1 
		digitInPower = digitInPower * digitInPower
	}
	return result
}