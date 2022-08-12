package main

func main() {
	gunType := "ak47"
	gun, err := getGun(gunType)
	if err != nil {
		panic(err)
	}
	println(gun.getName())
}
