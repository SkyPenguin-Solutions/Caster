package CastHunter

var Keys = map[string]string{
	"home":        "http://%s:8060/keypress/home",
	"play":        "http://%s:8060/keypress/play",
	"down":        "http://%s:8060/keypress/down",
	"up":          "http://%s:8060/keypress/up",
	"left":        "http://%s:8060/keypress/left",
	"right":       "http://%s:8060/keypress/right",
	"OK":          "http://%s:8060/keypress/select",
	"rewind":      "http://%s:8060/keypress/rewind",
	"fastforward": "http://%s:8060/keypress/fastforward",
	"options":     "http://%s:8060/keypress/options",
	"pause":       "http://%s:8060/keypress/pause",
	"back":        "http://%s:8060/keypress/back",
	"poweroff":    "http://%s:8060/keypress/poweroff",
	"vup":         "http://%s:8060/keypress/volumeup",
	"vdown":       "http://%s:8060/keypress/volumedown",
	"mute":        "http://%s:8060/keypress/volumemute",
	"devdown":     "http://%s:8060/keypress/powerOff",
	"devup":       "http://%s:8060/keypress/powerOn",
	"launch":      "http://%s:8060/launch/%s",
	"install":     "http://%s:8060/install/%s?contentid=%s&MediaType=%s",
	"disable_sgr": "http://%s:8060/query/sqrendezvous/untrack",
	"enable_sgr":  "http://%s:8060/query/sqrendezvous/track",
	"tv":          "http://%s:8060/query/tv-channels",
	"sgnode":      "http://%s:8060/query/sgnodes",
	"activetv":    "http://%s:8060/query/tv-active-channel",
	"dial":        "http://%s:8060/dial/dd.xml",
}

var PathsRoku = map[string]string{
	"browse":    "http://%s:8060/search/browse?keyword=%s",
	"devinfo":   "http://%s:8060/query/device-info",
	"devapps":   "http://%s:8060/query/apps",
	"devactive": "http://%s:8060/query/active-app",
}

var RouterPathsATT = map[string]string{}

func CheckLen() bool {
	if len(MACS) == len(IPS) {
		return true
	} else {
		return false
	}
}
