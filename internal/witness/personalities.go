package witness

// Personality describes the courtroom persona assigned to a command family.
type Personality struct {
	Family             string
	Name               string
	Archetype          string
	Quote              string
	TestimonialAngle   string
	Risk               int
	Repetition         int
	Habit              string
}

// Personalities maps command families to witness personalities.
var Personalities = map[string]Personality{
	"git": {
		Family:             "git",
		Name:               "Git the Version Keeper",
		Archetype:          "meticulous but easily startled",
		Quote:              "I recorded every commit, but I cannot vouch for the force-push.",
		TestimonialAngle:   "explains branch drift, detached heads, and untracked files",
		Risk:               3,
		Repetition:         2,
		Habit:              "run git status and git diff before destructive git commands",
	},
	"sudo": {
		F