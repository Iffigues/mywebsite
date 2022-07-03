package linuxtool

func MakeFiglet() (a Bash) {
	a.Bash = "/usr/bin/figlet"
	a.Arg = true
	e := []How {
		How{"-r", 0},
		How{"-x", 0},
		How{"-L", 0},
		How{"-R", 0},
		How{"-X", 0},
		How{"-l", 0},
		How{"-c", 0},
		How{"-f", 1},
		How{"-C", 1},
		How{"-N", 0},
		How{"-w", 1},
		How{"-p", 0},
		How{"-n", 0},
		How{"-m", 1},
		How{"-o", 0},
		How{"-W", 0},
		How{"-k", 0},
		How{"-S", 0},
		How{"-s", 0},
		How{"-N", 0},
		How{"-E", 0},
		How{"-D", 0},
		How{"-t", 0},
	}
	a.Com = e
	return
}
