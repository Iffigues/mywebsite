package linuxtool

func MakeFortune() (a Bash) {
	a.Bash = "/usr/games/fortune"
	a.Arg = false
	e := []How{
		How{"-a", 0},
		How{"-c", 0},
		How{"-e", 0},
		How{"-f", 0},
		How{"-l", 0},
		How{"-m", 1},
		How{"-n", 1},
		How{"-o", 0},
		How{"-s", 0},
		How{"-i", 0},
		How{"-u", 0},
	}
	a.Com = e
	return
}
