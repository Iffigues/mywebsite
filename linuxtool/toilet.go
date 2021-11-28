package linuxtool

func MakeToilet() (a Bash) {
	a.Bash = "/usr/bin/toilet"
	a.Arg = true
	e := []How {
		How{"-f", 1},
		How{"-S", 0},
		How{"-s", 0},
		How{"-k", 0},
		How{"-W", 0},
		How{"-o", 0},
		How{"-w", 1},
		How{"-F", 1},
		How{"-E", 1},
	}
	a.Com = e
	return
}
