package main

import (
	"fmt"
	"strconv"
)

func parse(x, y, op string) (res string, err bool) {
	arabics := map[string]int{
		"1":  1,
		"2":  2,
		"3":  3,
		"4":  4,
		"5":  5,
		"6":  6,
		"7":  7,
		"8":  8,
		"9":  9,
		"10": 10,
	}

	romans := map[string]int{
		"I":    1,
		"II":   2,
		"III":  3,
		"IV":   4,
		"V":    5,
		"VI":   6,
		"VII":  7,
		"VIII": 8,
		"IIX":  8,
		"IX":   9,
		"X":    10,
		"XC":   40,
		"C":    50,
		"CL":   90,
		"L":    100,
	}

	xa, okax := arabics[x]
	ya, okay := arabics[y]
	xr, okrx := romans[x]
	yr, okry := romans[y]

	if okax && okay {
		if xa == 0 || ya == 0 {
			fmt.Println("Вывод ошибки, так как строка не является математической операцией.")
			return "0", true
		}

		switch op {
		case "+":
			return strconv.Itoa(xa + ya), false
		case "-":
			return strconv.Itoa(xa - ya), false
		case "*":
			return strconv.Itoa(xa * ya), false
		case "/":
			return strconv.Itoa(xa / ya), false
		default:
			return "0", true
		}
	} else if okrx && okry {
		switch op {
		case "+":
			return fromRomansConv(xr + yr), false
		case "-":
			{
				if yr > xr {
					fmt.Println("Вывод ошибки, так как в римской системе нет отрицательных чисел.")
					return "0", true
				}
			}
			return fromRomansConv(xr - yr), false
		case "*":
			return fromRomansConv(xr * yr), false
		case "/":
			return fromRomansConv(xr / yr), false
		default:
			return "0", true
		}
	} else if okax != okay || okrx != okry {

		fmt.Println("Вывод ошибки, так как используются одновременно разные системыы счисления.")
		return "0", true
	} else {
		return "0", true
	}
}

func fromRomansConv(num int) string {
	romans := map[int]string{
		1: "I", 2: "II", 3: "III", 4: "IV", 5: "V", 6: "VI", 7: "VII", 8: "VIII", 9: "IX",
		10: "X", 11: "XI", 12: "XII", 13: "XIII", 14: "XIV", 15: "XV", 16: "XVI", 17: "XVII", 18: "XVIII", 19: "XIX",
		20: "XX", 21: "XXI", 22: "XXII", 23: "XXIII", 24: "XXIV", 25: "XXV", 26: "XXVI", 27: "XXVII", 28: "XXVIII", 29: "XXIX",
		30: "XXX", 31: "XXXI", 32: "XXXII", 33: "XXXIII", 34: "XXXIV", 35: "XXXVI", 36: "XXXVI", 37: "XXXVII", 38: "XXXVIII", 39: "XXXIX",
		40: "XL", 41: "XLI", 42: "XLII", 43: "XLIII", 44: "XLIV", 45: "XLV", 46: "XLVI", 47: "XLVII", 48: "XLVIII", 49: "IL",
		50: "L", 51: "LI", 52: "LII", 53: "LIII", 54: "LIV", 55: "LV", 56: "LVI", 57: "LVII", 58: "LVIII", 59: "ILX",
		60: "LX", 61: "LXI", 62: "LXII", 63: "LXIII", 64: "LXIV", 65: "LXV", 66: "LXVI", 67: "LXVII", 68: "LXVIII", 69: "LXIX",
		70: "LXX", 71: "LXXI", 72: "LXXII", 73: "LXXIII", 74: "LXXIV", 75: "LXXV", 76: "LXXVI", 77: "LXXVII", 78: "LXXVIII", 79: "LXXIX",
		80: "LXXX", 81: "LXXXI", 82: "LXXXII", 83: "LXIII", 84: "LXXXIV", 85: "LXXXV", 86: "LXXXVI", 87: "LXXXVII", 88: "LXXXVIII", 89: "LXXXIX",
		90: "XC", 91: "XCI", 92: "XCII", 93: "XCIII", 94: "XCIV", 95: "XCV", 96: "XCVI", 97: "XCVII", 98: "XCVIII", 99: "IC",
		100: "C",
	}
	return romans[num]
}

func main() {
	for {
		var x, y, op string
		fmt.Scanf("%s %s %s\n", &x, &op, &y)
		res, err := parse(x, y, op)
		if err {
			break
		} else {
			fmt.Println(res)
		}
	}
}
