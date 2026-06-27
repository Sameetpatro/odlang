package investigate

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
)

type profile struct {
	confidence     int
	underestimated bool
	timeOfDay      string
	drink          string
	travel         string
	overthink      string
	surprises      string
}

// Run starts the full cinematic investigation experience.
func Run() {
	clearScreen()
	fmt.Print(hide)
	enableGreen()
	defer disableColor()

	p := &profile{}

	stepBoot()
	if !stepIdentity() {
		return
	}
	stepWelcome()
	stepQuestions(p)
	stepAnalysis()
	stepDossier(p)
	stepOfficialReport(p)
	if !stepHiddenRecords() {
		stepCleanup()
		return
	}
	stepCreatorNote()
	stepClassifiedNotes()
	stepLostAndFound()
	stepVideoLink()
	stepCleanup()
}

func stepBoot() {
	banner(`
        ODLANG SECURE TERMINAL

      Personal Investigation Suite
`, 38)
	pause(300)

	loadItems := []string{
		"Loading Investigation Engine",
		"Loading Behavioral Analyzer",
		"Loading Encrypted Database",
		"Loading Memory Cache",
	}
	for _, item := range loadItems {
		fmt.Print("  ")
		typeLine(item+"...", 8)
		progressBar("", 24, 6)
	}
	typePause("\n  Done.")
	pause(350)
	clearScreen()
}

func stepIdentity() bool {
	typeBlock([]string{
		"  WARNING",
		"",
		"  This executable was created",
		"  for only one individual.",
		"",
		"  Identity verification required.",
		"",
		"  Unauthorized access",
		"  will immediately terminate execution.",
	}, 18)
	pause(250)

	name := prompt("  Enter your full name:")
	subj, ok := resolveSubject(name)
	if !ok {
		fmt.Println()
		typeBlock([]string{
			"  Identity mismatch.",
			"",
			"  Access denied.",
			"",
			"  Terminating...",
			"",
			"  Goodbye.",
		}, 20)
		pause(600)
		return false
	}
	activeSubject = subj
	return true
}

func stepWelcome() {
	clearScreen()
	scanSequence([]string{
		"Scanning identity",
		"Searching national records",
		"Searching encrypted archive",
		"Cross verifying",
	})
	fmt.Println()
	typeBlock([]string{
		"  MATCH FOUND",
		"",
		"  Identity Confirmed.",
		"",
		"  Welcome,",
		"  " + activeSubject.displayName + ".",
	}, 22)
	pause(400)

	typeBlock([]string{
		"",
		"  Beginning Psychological Investigation.",
		"",
		"  You will be asked 7 questions.",
		"",
		"  No wrong answers.",
		"",
		"  Answer honestly.",
	}, 18)
	pause(350)
}

func stepQuestions(p *profile) {
	for {
		ans := prompt(questionLabel(1) + "\n\n  How would you rate yourself?\n\n  (1-10)")
		n := 0
		if _, err := fmt.Sscanf(ans, "%d", &n); err == nil && n >= 1 && n <= 10 {
			p.confidence = n
			fmt.Println()
			switch {
			case n >= 9:
				typeBlock([]string{
					"  Confidence Level:",
					"  High.",
					"",
					"  Interesting...",
				}, 18)
			case n >= 6:
				typePause("  Balanced confidence detected.")
			default:
				typeBlock([]string{
					"  Confidence lower than expected.",
					"",
					"  People are often harder on",
					"  themselves than everyone else is.",
				}, 16)
			}
			break
		}
		typePause("  Enter a number between 1 and 10.")
	}
	pause(200)

	p.underestimated = promptYN(2, "Do people underestimate you?")
	if p.underestimated {
		typePause("  Common in high-potential subjects. Filing accordingly.")
	} else {
		typePause("  Self-awareness noted. Rare.")
	}
	pause(200)

	timeChoices := []string{"Morning person", "Night owl", "Depends on the day"}
	tidx := promptChoice(3, "Morning or Night?", timeChoices)
	p.timeOfDay = timeChoices[tidx]
	if tidx == 1 {
		typePause("  Night owl status: confirmed.")
	} else {
		typePause("  Circadian data logged.")
	}
	pause(200)

	drinkChoices := []string{"Tea", "Coffee", "Neither", "Both, no questions"}
	didx := promptChoice(4, "Tea or Coffee?", drinkChoices)
	p.drink = drinkChoices[didx]
	typePause("  Beverage profile updated.")
	pause(200)

	travelChoices := []string{"Travel anywhere", "Stay home", "Short trips only"}
	tridx := promptChoice(5, "Travel or Stay Home?", travelChoices)
	p.travel = travelChoices[tridx]
	typePause("  Wanderlust coefficient measured.")
	pause(200)

	otChoices := []string{"Constantly", "Sometimes", "Rarely", "I don't think, I vibe"}
	otidx := promptChoice(6, "Do you overthink?", otChoices)
	p.overthink = otChoices[otidx]
	if otidx <= 1 {
		typePause("  Overthinker badge unlocked. Welcome to the club.")
	} else {
		typePause("  Mental clarity detected. Teach us your ways.")
	}
	pause(200)

	surChoices := []string{"Love them", "Depends", "Hate them", "Only good ones"}
	suridx := promptChoice(7, "Do you like surprises?", surChoices)
	p.surprises = surChoices[suridx]
	typePause("  Surprise tolerance indexed.")
	pause(250)
}

func stepAnalysis() {
	clearScreen()
	typePause("  Uploading...")
	progressBar("SYNC", 28, 8)
	pause(150)

	phases := []string{
		"Analyzing responses",
		"Generating psychological model",
		"Cross validating",
		"Behavior prediction",
		"Searching encrypted archive",
		"Decrypting",
	}
	for _, phase := range phases {
		logStatus(phase)
		if phase == "Decrypting" {
			progressBar("", 30, 10)
		} else {
			bracketProgress("", 5+rand.Intn(3))
		}
	}
	fmt.Println()
	progressBar("", 30, 10)
	fmt.Println()
	typePause("  Analysis Complete.")
	pause(400)
}

func confidenceLabel(p *profile) string {
	switch {
	case p.confidence >= 9:
		return "High"
	case p.confidence >= 6:
		return "Stable"
	default:
		return "Quietly Strong"
	}
}

func moodSwingLevel(p *profile) string {
	score := 0
	if p.overthink == "Constantly" || p.overthink == "Sometimes" {
		score += 2
	}
	if p.surprises == "Love them" || p.surprises == "Only good ones" {
		score += 1
	}
	if p.underestimated {
		score += 1
	}
	if p.timeOfDay == "Night owl" {
		score += 1
	}
	if score >= 4 {
		return "VERY HIGH"
	}
	if score >= 2 {
		return "MODERATELY CHAOTIC"
	}
	return "SURPRISINGLY LOW (suspicious)"
}

func buildDossier(p *profile) []string {
	s := activeSubject
	lines := []string{
		"",
		"  Name :",
		"  " + s.displayName,
		"",
		"  Date of Birth :",
		"  " + s.dob,
		"",
		"  Estimated Lifespan :",
		"  " + s.lifespan,
		"",
		"  Mood Swing Level :",
		"  " + moodSwingLevel(p),
		"",
		"  Alcohol Consumption",
		"",
	}
	if s.alcoholNever {
		lines = append(lines, "  Never", "")
	} else {
		lines = append(lines,
			"  First Recorded :",
			"  "+s.alcoholFirst,
			"",
			"  Last Recorded :",
			"  "+s.alcoholLast,
			"",
		)
	}
	lines = append(lines,
		"  Primary Goal :",
		"  "+s.primaryGoal,
		"",
		"  Status :",
		"  Verified Human",
		"",
		"  Threat Level :",
		"  0%",
		"",
		"  Smile Frequency :",
		"  Higher than expected",
		"",
		"  Cuteness Index :",
		"  9.8/10",
		"",
		"  Confidence :",
		"  "+confidenceLabel(p),
		"",
		"  Energy Pattern :",
		"  "+p.timeOfDay,
		"",
		"  Drink Preference :",
		"  "+p.drink,
		"",
		"  Travel Mode :",
		"  "+p.travel,
	)
	return lines
}

func stepDossier(p *profile) {
	clearScreen()
	typePause("  Compiling confidential dossier...")
	pause(300)

	for _, line := range []string{
		"Parsing neural response map",
		"Calculating mood variance",
		"Scanning beverage history",
		"Verifying humanity status",
	} {
		logStatus(line)
		bracketProgress("", 5)
	}
	fmt.Println()
	pause(200)

	banner("  CONFIDENTIAL DOSSIER  ", 32)
	typeBlock(buildDossier(p), 14)
	pause(500)
}

func smitaOfficialReport(p *profile) []string {
	report := []string{
		"",
		"  Official Report",
		"",
		"  Smita Patra is one of the most unpredictable",
		"  subjects ever investigated.",
		"",
		"  Known symptoms include",
		"",
		"  • Random mood updates",
		"  • Professional overthinker",
		"  • Eats first. Thinks later.",
		"  • Can successfully confuse everyone,",
		"    including herself.",
		"",
	}
	if p.overthink == "Constantly" || p.overthink == "Sometimes" {
		report = append(report, "  • Brain runs background tabs 24/7", "")
	}
	if p.underestimated {
		report = append(report, "  • Underestimated by others, over-delivers quietly", "")
	}
	if p.surprises == "Love them" {
		report = append(report, "  • Thrives on chaos — the fun kind", "")
	}
	if p.travel == "Stay home" {
		report = append(report, "  • Homebody energy, elite tier", "")
	}
	return report
}

func shreetyOfficialReport(p *profile) []string {
	report := []string{
		"",
		"  Official Report",
		"",
		"  Shreety Samantaray is a newly catalogued species —",
		"  mostly inherited from a donkey, behaves like one.",
		"",
		"  Known traits include",
		"",
		"  • Gets rage-baited by the smallest poke",
		"  • Height: 4'2\" without heels, 4'4\" with heels",
		"  • Random mood swings. Certified overthinker",
		"  • Catches colds easily. Immunity: critically low",
		"  • Does not know how to ride a scooty",
		"  • Avoids interacting with Sameet in public —",
		"    makes him uncomfortable. He's shy too. Tragic.",
		"",
		"  Social settings breakdown:",
		"",
		"  • With her sisters: extremely extrovert",
		"  • With Sameet: ambivert mode unlocked",
		"  • In public: full introvert. Mic off.",
		"",
		"  Verdict: chaotic, loud, soft.",
		"  Cute though. No notes.",
		"",
	}
	if p.overthink == "Constantly" || p.overthink == "Sometimes" {
		report = append(report, "  • Overthinking engine never stops", "")
	}
	if p.underestimated {
		report = append(report, "  • Underestimated until she goes off", "")
	}
	return report
}

func smitaReportFooter() []string {
	return []string{
		"  Current Status",
		"",
		"  Still under observation.",
		"",
		"  Danger Level",
		"",
		"  Mostly harmless.",
		"",
		"  Unless she's hungry.",
		"",
		"  Final Notes",
		"",
		"  >>> Your Instagram page would not look good",
		"    until you post a reel with Sameet.",
		"",
		"  >>> lakh lakh shukar hai Sameet jaise bande ka",
		"    who chose to stay with a dickhead like you.",
		"",
		"  Thank you for your cooperation.",
		"",
		"  Fuck you",
	}
}

func shreetyReportFooter() []string {
	return []string{
		"  Current Status",
		"",
		"  Still under observation.",
		"",
		"  Danger Level",
		"",
		"  Mostly harmless.",
		"",
		"  Unless rage-baited.",
	}
}

func isShreety() bool {
	return strings.EqualFold(activeSubject.displayName, subjectShreety.displayName)
}

func stepOfficialReport(p *profile) {
	fmt.Println()
	typePause("  Generating official report...")
	pause(250)

	var report []string
	if isShreety() {
		report = shreetyOfficialReport(p)
		report = append(report, shreetyReportFooter()...)
	} else {
		report = smitaOfficialReport(p)
		report = append(report, smitaReportFooter()...)
	}

	typeBlock(report, 16)
	pause(450)
}

func stepHiddenRecords() bool {
	if !promptYN(0, "Do you want to explore hidden records?") {
		typeBlock([]string{
			"",
			"  Closing session.",
			"",
			"  Some mysteries remain unsolved.",
		}, 18)
		pause(400)
		return false
	}

	logStatus("Searching")
	bracketProgress("", 8)
	fmt.Println()
	typePause("  Hidden archive found.")
	pause(200)

	if !promptYN(0, "Decrypt?") {
		typePause("  Archive remains sealed. Respect.")
		pause(300)
		return false
	}

	clearScreen()
	typePause("  DECRYPTION SEQUENCE INITIATED")
	pause(200)

	keys := []string{
		"ACCESS KEY 0x7F2A",
		"PERSONAL HASH MATCH",
		"FINAL LAYER — EMOTIONAL ENCRYPTION",
	}
	for _, key := range keys {
		typeLine("  >> "+key, 10)
		decryptText("  [ UNLOCKED ] "+key, 8)
		progressBar("", 26, 6)
	}
	fmt.Println()
	progressBar("DECRYPT", 30, 10)
	fmt.Println()
	typePause("  Decryption complete.")
	pause(350)

	if isShreety() {
		fmt.Println()
		typePause("  Hidden file recovered.")
		pause(200)
		logStatus("Rendering classified image")
		progressBar("", 28, 6)
		fmt.Println()
		typeBlock([]string{
			"",
			"  Open this link to view the file",
			"",
			"  https://res.cloudinary.com/dw8imuhcz/image/upload/v1782590876/image_mqcwtl.png",
		}, 6)
		pause(400)
	}

	return true
}

func stepCreatorNote() {
	clearScreen()
	banner("  SECURE MESSAGE  ", 32)
	pause(250)

	note := []string{
		"",
		"  This program wasn't created",
		"  to impress everyone.",
		"",
		"  It was made for one person only.",
		"",
		"  And that person is you.",
		"",
		"  If you're reading this,",
		"",
		"  mission accomplished.",
		"",
		"  Hopefully this made",
		"  your day a little more interesting.",
		"",
		"  Creator",
		"",
		"  Sameet",
	}
	typeBlock(note, 20)
	pause(450)
}

func stepClassifiedNotes() {
	fmt.Println()
	typePause("  Opening classified file...")
	pause(200)
	decryptText("  CLASSIFIED NOTES — LEVEL 9", 8)
	pause(250)

	notes := []string{
		"",
		"  Massive respect to Sameet.",
		"",
		"  The fact that he still manages",
		"  this much unpredictability",
		"  deserves an award.",
		"",
		"  Had he never met you,",
		"",
		"  the universe would've been",
		"  slightly less entertaining.",
	}
	typeBlock(notes, 18)
	pause(450)
}

func stepLostAndFound() {
	fmt.Println()
	banner("  LOST & FOUND NOTICE  ", 32)
	pause(200)

	notice := []string{
		"",
		"  If this human is ever found",
		"  wandering around looking lost,",
		"  slightly chaotic, or mid snack —",
		"",
		"  please return immediately.",
		"",
		"  Reward:",
		"  One genuine thank you.",
		"  (No cap. Valid forever.)",
		"",
		"  Contact:",
		"  6370064576",
		"",
		"  Side note:",
		"  She probably forgot why she",
		"  opened the fridge. Happens.",
	}
	typeBlock(notice, 16)
	pause(450)
}

func stepVideoLink() {
	if isShreety() {
		return
	}

	fmt.Println()
	typePause("  One more file detected...")
	pause(250)
	logStatus("Retrieving media link")
	progressBar("", 28, 6)
	fmt.Println()
	typeBlock([]string{
		"",
		"  Open this link to know more",
		"",
		"  https://res.cloudinary.com/dw8imuhcz/video/upload/v1782587670/smi_.ta_6656278345371618565_pbx7xa.mp4",
	}, 6)
	pause(400)
}

func stepCleanup() {
	fmt.Println()
	cleanup := []string{
		"  Cleaning traces...",
		"  Deleting temporary files...",
		"  Logging out...",
	}
	for _, line := range cleanup {
		typeLine(line, 10)
		bracketProgress("", 5)
	}
	fmt.Println()
	typePause("  Investigation Completed.")
	waitEnter("  Press ENTER to exit.")
	os.Exit(0)
}
