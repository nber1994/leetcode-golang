package kds

func IntToRoman(num int) string {
	str := []byte{}
	symbol, index := []byte{'I', 'V', 'X', 'L', 'C', 'D', 'M'}, 8
	no := 0
	for divisor := 1000; num > 0 && index >= 0; divisor /= 10 {
		no = num / divisor
		index = index - 2
		num = num % divisor
		if no == 0 {
			continue
		}

		if no <= 3 {
			for i := 0; i < no; i++ {
				str = append(str, symbol[index])
			}
		} else if 4 == no {
			str = append(str, symbol[index], symbol[index+1])
		} else if 5 == no {
			str = append(str, symbol[index+1])
		} else if no < 9 {
			str = append(str, symbol[index+1])
			for i := 0; i < no-5; i++ {
				str = append(str, symbol[index])
			}
		} else if 9 == no {
			str = append(str, symbol[index], symbol[index+2])
		}
	}
	return string(str)
}
