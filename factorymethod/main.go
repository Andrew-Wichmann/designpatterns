package main

func main() {
	gunType := "ak47"
	gun, _ := getGun(gunType)
	println(gun.getName())
	gun.signGunBackIn()

	gun, _ = getGun(gunType)
	if gun != nil {
		println(gun.getName())
	}
	gun, _ = getGun(gunType)
	if gun != nil {
		println(gun.getName())
	}
	gun, _ = getGun(gunType)
	if gun != nil {
		println(gun.getName())
	}
}
