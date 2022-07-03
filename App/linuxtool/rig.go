package linuxtool

func MakeRig() (a Bash) {
	a.Bash = "/usr/bin/rig"
	a.Arg = false
	e := []How {
		How{"-c", 1},
		How{"-m", 0},
		How{"-f", 0},
	}
	a.Com = e
	return
}
