package main

type musket struct {
	Gun
}

func (musket musket) signGunBackIn() {
}

func newMusket() IGun {
	return &musket{
		Gun: Gun{
			name:  "Musket gun",
			power: 1,
		},
	}
}
