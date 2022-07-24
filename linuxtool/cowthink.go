package linuxtool

func MakeCowthink() (a Bash) {
	a.Bash = "/usr/games/cowthink"
	a.Arg = true
	e := []How {
		How{"-e", 1},
		How{"-f", 1},
		How{"-l", 0},
		How{"-T", 1},
		How{"-W", 1},
		How{"-b", 0},
		How{"-d", 0},
		How{"-g", 0},
		How{"-p", 0},
		How{"-s", 0},
		How{"-t", 0},
		How{"-w", 0},
		How{"-y", 0},
	}
	a.Com = e
	return
}
