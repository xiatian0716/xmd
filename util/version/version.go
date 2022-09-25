package version

// BuildVersion is the version of Xmd.
const BuildVersion = "1.0.0"

var tag = map[string]int{"pre-alpha": 0, "alpha": 1, "pre-bata": 2, "beta": 3, "released": 4, "": 5}

// GetXmdVersion returns BuildVersion
func GetXmdVersion() string {
	return BuildVersion
}
