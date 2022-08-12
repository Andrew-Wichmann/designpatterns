package main

import "fmt"

type Ak47 struct {
	Gun
}

var armourySignout map[Ak47]bool

var armoury []Ak47

const armourySize = 2

func (ak Ak47) signGunBackIn() {
	armourySignout[ak] = false
}

func newAk47() IGun {
	if len(armoury) == 0 {
		armoury = make([]Ak47, armourySize)
		armourySignout = map[Ak47]bool{}
	}
	for _, ak := range armoury {
		if ak.Gun.name != "" && !armourySignout[ak] {
			println("Gun found in armoury")
			armourySignout[ak] = true
			return &ak
		}
	}

	for i, ak := range armoury {
		if ak.Gun.name == "" {
			println("New gun manufactured")
			newGun := &Ak47{
				Gun: Gun{
					name:  fmt.Sprintf("AK47 gun: %d", i),
					power: 4,
				},
			}
			armoury[i] = *newGun
			armourySignout[*newGun] = true
			return newGun
		}
	}
	println("No gun available")
	return nil
}
