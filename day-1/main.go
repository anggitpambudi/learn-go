package main

import (
	"fmt"
	"github.com/asaskevich/govalidator"
	"os"
    "os/exec"
	"time"
)

func main() {

	var menu string
	screen := true
	var name string

	fmt.Println("Selamat Datang Di Daftar aplikasi CLI GO!")
		fmt.Print("Please Input your name : ")
		fmt.Scan(&name)
	
	for screen {
		clearScreen()
		fmt.Println("Daftar Menu")
		fmt.Println("1. Kalkulator")
		fmt.Println("To Exit Program press 'x' ")
		fmt.Print("Pilih Menu : ")
		fmt.Scan(&menu)

		if menu == "x" {
			fmt.Print("Bye Bye ! :D")
			screen = false
			continue
		}

		if !govalidator.IsInt(menu) {
			fmt.Println("Menu Yang Dipiih Harus Angka Diatas 0!")
			sleepScreen(2)
		}

		switch menu {
		case "1" :
			kalkulator(name)
			screen = false
		default :
			fmt.Println("Menu Tidak Tersedia!")
			sleepScreen(2)
		}
	}
}

func clearScreen() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func sleepScreen(detik int) {
	duration := time.Duration(detik)*time.Second
	time.Sleep(duration)
}
