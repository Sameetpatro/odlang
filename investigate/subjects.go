package investigate

import "strings"

type subject struct {
	displayName  string
	dob          string
	lifespan     string
	alcoholNever bool
	alcoholFirst string
	alcoholLast  string
	primaryGoal  string
}

var (
	subjectSmita = subject{
		displayName:  "Smita Patra",
		dob:          "17.12.2004",
		lifespan:     "Around 2074",
		alcoholFirst: "27.06.2026",
		alcoholLast:  "27.06.2026",
		primaryGoal:  "Marry a rich Brahmin guy",
	}
	subjectShreety = subject{
		displayName:  "Shreety Samantaray",
		dob:          "18.09.2005",
		lifespan:     "Around 2075",
		alcoholNever: true,
		primaryGoal:  "Marry a non-Brahmin guy and live happy forever",
	}
)

var activeSubject subject

func resolveSubject(name string) (subject, bool) {
	n := strings.ToLower(strings.TrimSpace(name))
	switch n {
	case "smita patra":
		return subjectSmita, true
	case "shreety samantaray":
		return subjectShreety, true
	default:
		return subject{}, false
	}
}
