package main

import "fmt"

func main() {
        var nairaAmount, dollar, pound, euro, yen, dihram, cedis, rand, peso, real, kenyanshilling float64
        var currency string

        dollarRate := 1620.00
        poundRate := 2079.76
        euroRate := 1751.40
        yenRate := 10.51
        dihramRate := 441.06
        cedisRate := 104.49
        randRate := 87.59
        pesoRate := 27.66
        realRate := 286.83
        kenyanshillingRate := 12.37



        fmt.Println("Enter Naira amount: ")
        fmt.Scanln(&nairaAmount)

        fmt.Println("Enter desired currency (dollar, pound, euro, yen, dihram, cedis, rand, peso, real, kenyanshilling): ")
        fmt.Scanln(&currency)

        switch currency {
        case "dollar":
                dollar = nairaAmount / dollarRate
                fmt.Println("Naira amount is equivalent to: ", dollar, "dollars")
        case "pound":
                pound = nairaAmount / poundRate
                fmt.Println("Naira amount is equivalent to: ", pound, "pounds")
        case "euro":
                euro = nairaAmount / euroRate
                fmt.Println("Naira amount is equivalent to: ", euro, "euros")
        case "yen":
                yen = nairaAmount / yenRate
                fmt.Println("Naira amount is equivalent to: ", yen, "yens")

        case "dihram":
                dihram = nairaAmount / dihramRate
                fmt.Println("Naira amount is equivalent to: ", dihram, "dihrams")

        case "cedis":
                cedis = nairaAmount / cedisRate
                fmt.Println("Naira amount is equivalent to: ", cedis, "cedis")

        case "rand":
                rand = nairaAmount / randRate
                fmt.Println("Naira amount is equivalent to: ", rand, "rands")

        case "peso":
                peso = nairaAmount / pesoRate
                fmt.Println("Naira amount is equivalent to: ", peso, "pesos")

        case "real":
                real = nairaAmount / realRate
                fmt.Println("Naira amount is equivalent to: ", real, "reals")

        case "kenyanshilling":
                kenyanshilling = nairaAmount / kenyanshillingRate
                fmt.Println("Naira amount is equivalent to: ", kenyanshilling, "kenyanshillings")

        default:
                fmt.Println("Invalid currency")
        }
}

