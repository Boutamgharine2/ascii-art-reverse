package asci

// b
var f = map[string]string{
	"thinkertoy": "thinkertoy.txt",
	"shadow":     "shadow.txt",
	"standard":   "standard.txt",
}

// f
var q = map[string]func([]string, string){
	"--output=": Write,
	"--color=":  Color,
	// //"--help=",
	"--allign=": Allign,
	// "--reverse=",
}

// c
var t = map[string]string{
	"ColorOff":          "\033[0m",
	"Aquamarine1":       "\033[0;38;5;122m",
	"Aquamarine3":       "\033[0;38;5;79m",
	"Black":             "\033[0;38;5;0m",
	"Blue1":             "\033[0;38;5;21m",
	"Blue3":             "\033[0;38;5;19m",
	"Blue":              "\033[0;38;5;12m",
	"Blueviolet":        "\033[0;38;5;57m",
	"Cadetblue":         "\033[0;38;5;73m",
	"Chartreuse1":       "\033[0;38;5;118m",
	"Chartreuse2":       "\033[0;38;5;112m",
	"Chartreuse3":       "\033[0;38;5;76m",
	"Chartreuse4":       "\033[0;38;5;64m",
	"CornSilk1":         "\033[0;38;5;230m",
	"Cornflowerblue":    "\033[0;38;5;69m",
	"Cyan1":             "\033[0;38;5;51m",
	"Cyan2":             "\033[0;38;5;50m",
	"Cyan3":             "\033[0;38;5;43m",
	"DarkOliveGreen1":   "\033[0;38;5;191m",
	"DarkOrange":        "\033[0;38;5;208m",
	"DarkSeaGreen1":     "\033[0;38;5;193m",
	"Darkblue":          "\033[0;38;5;18m",
	"Darkcyan":          "\033[0;38;5;36m",
	"Darkgoldenrod":     "\033[0;38;5;136m",
	"Darkgreen":         "\033[0;38;5;22m",
	"Darkkhaki":         "\033[0;38;5;143m",
	"Darkmagenta":       "\033[0;38;5;90m",
	"Darkolivegreen2":   "\033[0;38;5;155m",
	"Darkolivegreen3":   "\033[0;38;5;107m",
	"Darkorange3":       "\033[0;38;5;130m",
	"Darkred":           "\033[0;38;5;88m",
	"Darkseagreen1":     "\033[0;38;5;158m",
	"Darkseagreen2":     "\033[0;38;5;151m",
	"Darkseagreen3":     "\033[0;38;5;115m",
	"Darkseagreen4":     "\033[0;38;5;71m",
	"Darkseagreen":      "\033[0;38;5;108m",
	"Darkslategray1":    "\033[0;38;5;123m",
	"Darkslategray2":    "\033[0;38;5;87m",
	"Darkslategray3":    "\033[0;38;5;116m",
	"Darkturquoise":     "\033[0;38;5;44m",
	"Darkviolet":        "\033[0;38;5;128m",
	"DeepPink1":         "\033[0;38;5;198m",
	"DeepPink2":         "\033[0;38;5;197m",
	"Deeppink3":         "\033[0;38;5;161m",
	"Deeppink4":         "\033[0;38;5;125m",
	"Deepskyblue1":      "\033[0;38;5;39m",
	"Deepskyblue2":      "\033[0;38;5;38m",
	"Deepskyblue3":      "\033[0;38;5;31m",
	"Deepskyblue4":      "\033[0;38;5;23m",
	"Dodgerblue1":       "\033[0;38;5;33m",
	"Dodgerblue2":       "\033[0;38;5;27m",
	"Dodgerblue3":       "\033[0;38;5;26m",
	"Fuchsia":           "\033[0;38;5;13m",
	"Gold1":             "\033[0;38;5;220m",
	"Gold3":             "\033[0;38;5;142m",
	"Green1":            "\033[0;38;5;46m",
	"Green3":            "\033[0;38;5;34m",
	"Green4":            "\033[0;38;5;28m",
	"Green":             "\033[0;38;5;2m",
	"Greenyellow":       "\033[0;38;5;154m",
	"Grey0":             "\033[0;38;5;16m",
	"Grey100":           "\033[0;38;5;231m",
	"Grey11":            "\033[0;38;5;234m",
	"Grey15":            "\033[0;38;5;235m",
	"Grey19":            "\033[0;38;5;236m",
	"Grey23":            "\033[0;38;5;237m",
	"Grey27":            "\033[0;38;5;238m",
	"Grey30":            "\033[0;38;5;239m",
	"Grey35":            "\033[0;38;5;240m",
	"Grey37":            "\033[0;38;5;59m",
	"Grey39":            "\033[0;38;5;241m",
	"Grey3":             "\033[0;38;5;232m",
	"Grey42":            "\033[0;38;5;242m",
	"Grey46":            "\033[0;38;5;243m",
	"Grey50":            "\033[0;38;5;244m",
	"Grey53":            "\033[0;38;5;102m",
	"Grey54":            "\033[0;38;5;245m",
	"Grey58":            "\033[0;38;5;246m",
	"Grey62":            "\033[0;38;5;247m",
	"Grey63":            "\033[0;38;5;139m",
	"Grey66":            "\033[0;38;5;248m",
	"Grey69":            "\033[0;38;5;145m",
	"Grey70":            "\033[0;38;5;249m",
	"Grey74":            "\033[0;38;5;250m",
	"Grey78":            "\033[0;38;5;251m",
	"Grey7":             "\033[0;38;5;233m",
	"Grey82":            "\033[0;38;5;252m",
	"Grey84":            "\033[0;38;5;188m",
	"Grey85":            "\033[0;38;5;253m",
	"Grey89":            "\033[0;38;5;254m",
	"Grey93":            "\033[0;38;5;255m",
	"Grey":              "\033[0;38;5;8m",
	"Honeydew2":         "\033[0;38;5;194m",
	"HotPink":           "\033[0;38;5;206m",
	"Hotpink2":          "\033[0;38;5;169m",
	"Hotpink3":          "\033[0;38;5;132m",
	"IndianRed1":        "\033[0;38;5;204m",
	"Indianred":         "\033[0;38;5;131m",
	"Khaki1":            "\033[0;38;5;228m",
	"Khaki3":            "\033[0;38;5;185m",
	"LightCoral":        "\033[0;38;5;210m",
	"LightCyan1":        "\033[0;38;5;195m",
	"LightGoldenrod1":   "\033[0;38;5;227m",
	"LightGoldenrod2":   "\033[0;38;5;221m",
	"LightPink1":        "\033[0;38;5;217m",
	"LightSalmon1":      "\033[0;38;5;216m",
	"LightSteelBlue1":   "\033[0;38;5;189m",
	"Lightcyan3":        "\033[0;38;5;152m",
	"Lightgoldenrod2":   "\033[0;38;5;186m",
	"Lightgoldenrod3":   "\033[0;38;5;179m",
	"Lightgreen":        "\033[0;38;5;119m",
	"Lightpink3":        "\033[0;38;5;174m",
	"Lightpink4":        "\033[0;38;5;95m",
	"Lightsalmon3":      "\033[0;38;5;137m",
	"Lightseagreen":     "\033[0;38;5;37m",
	"Lightskyblue1":     "\033[0;38;5;153m",
	"Lightskyblue3":     "\033[0;38;5;109m",
	"Lightslateblue":    "\033[0;38;5;105m",
	"Lightslategrey":    "\033[0;38;5;103m",
	"Lightsteelblue3":   "\033[0;38;5;146m",
	"Lightsteelblue":    "\033[0;38;5;147m",
	"Lightyellow3":      "\033[0;38;5;187m",
	"Lime":              "\033[0;38;5;10m",
	"Magenta1":          "\033[0;38;5;201m",
	"Magenta2":          "\033[0;38;5;165m",
	"Magenta3":          "\033[0;38;5;127m",
	"Maroon":            "\033[0;38;5;1m",
	"MayaBlue":          "\033[0;38;5;81m",
	"MediumOrchid1":     "\033[0;38;5;207m",
	"Mediumorchid1":     "\033[0;38;5;171m",
	"Mediumorchid3":     "\033[0;38;5;133m",
	"Mediumorchid":      "\033[0;38;5;134m",
	"Mediumpurple1":     "\033[0;38;5;141m",
	"Mediumpurple2":     "\033[0;38;5;135m",
	"Mediumpurple3":     "\033[0;38;5;97m",
	"Mediumpurple4":     "\033[0;38;5;60m",
	"Mediumpurple":      "\033[0;38;5;104m",
	"Mediumspringgreen": "\033[0;38;5;49m",
	"Mediumturquoise":   "\033[0;38;5;80m",
	"Mediumvioletred":   "\033[0;38;5;126m",
	"MistyRose1":        "\033[0;38;5;224m",
	"Mistyrose3":        "\033[0;38;5;181m",
	"NavajoWhite1":      "\033[0;38;5;223m",
	"Navajowhite3":      "\033[0;38;5;144m",
	"Navy":              "\033[0;38;5;4m",
	"Navyblue":          "\033[0;38;5;17m",
	"Olive":             "\033[0;38;5;3m",
	"Orange1":           "\033[0;38;5;214m",
	"Orange3":           "\033[0;38;5;172m",
	"Orange4":           "\033[0;38;5;94m",
	"OrangeRed1":        "\033[0;38;5;202m",
	"Orchid1":           "\033[0;38;5;213m",
	"Orchid2":           "\033[0;38;5;212m",
	"Orchid":            "\033[0;38;5;170m",
	"PaleVioletred1":    "\033[0;38;5;211m",
	"Palegreen1":        "\033[0;38;5;121m",
	"Palegreen3":        "\033[0;38;5;114m",
	"Paleturquoise1":    "\033[0;38;5;159m",
	"Paleturquoise4":    "\033[0;38;5;66m",
	"Pink1":             "\033[0;38;5;218m",
	"Pink3":             "\033[0;38;5;175m",
	"Plum1":             "\033[0;38;5;219m",
	"Plum2":             "\033[0;38;5;183m",
	"Plum3":             "\033[0;38;5;176m",
	"Plum4":             "\033[0;38;5;96m",
	"Purple3":           "\033[0;38;5;56m",
	"Purple4":           "\033[0;38;5;55m",
	"Purple":            "\033[0;38;5;129m",
	"Red1":              "\033[0;38;5;196m",
	"Red3":              "\033[0;38;5;124m",
	"Red":               "\033[0;38;5;9m",
	"Rosybrown":         "\033[0;38;5;138m",
	"Royalblue1":        "\033[0;38;5;63m",
	"Salmon1":           "\033[0;38;5;209m",
	"SandyBrow":         "\033[0;38;5;215m",
	"Seagreen1":         "\033[0;38;5;84m",
	"Seagreen2":         "\033[0;38;5;83m",
	"Seagreen3":         "\033[0;38;5;78m",
	"Silver":            "\033[0;38;5;7m",
	"Skyblue":           "\033[0;38;5;117m",
	"Skyblue3":          "\033[0;38;5;74m",
	"Slateblue1":        "\033[0;38;5;99m",
	"Slateblue3":        "\033[0;38;5;62m",
	"Springgreen1":      "\033[0;38;5;48m",
	"Springgreen2":      "\033[0;38;5;47m",
	"Springgreen3":      "\033[0;38;5;35m",
	"Springgreen4":      "\033[0;38;5;29m",
	"Steelblue1":        "\033[0;38;5;75m",
	"Steelblue3":        "\033[0;38;5;68m",
	"Steelblue":         "\033[0;38;5;67m",
	"Tan":               "\033[0;38;5;180m",
	"Teal":              "\033[0;38;5;6m",
	"Thistle1":          "\033[0;38;5;225m",
	"Thistle3":          "\033[0;38;5;182m",
	"Turquoise2":        "\033[0;38;5;45m",
	"Turquoise4":        "\033[0;38;5;30m",
	"Violet":            "\033[0;38;5;177m",
	"Wheat1":            "\033[0;38;5;229m",
	"Wheat4":            "\033[0;38;5;101m",
	"White":             "\033[0;38;5;15m",
	"Yellow1":           "\033[0;38;5;226m",
	"Yellow2":           "\033[0;38;5;190m",
	"Yellow3":           "\033[0;38;5;148m",
	"Yellow4":           "\033[0;38;5;100m",
	"Yellow":            "\033[0;38;5;11m",
	"K8sBlue":           "\033[0;38;5;38m",
}
