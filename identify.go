package japanese

import "strings"

// possible verb forms
const (
	TeForm = iota
	Infinitive
	PresentIndicative
	Presumptive
	Imperative
	PastIndicative
	PastPresumptive
	PresentProgressive
	PastProgressive
	Provisional
	Conditional
	Potential
	Causative
	Unknown
)

// politeness levels
const (
	Plain = iota
	Polite
)

// positive / negative
const (
	Positive = iota
	Negative
)

// verb classes
const (
	Ichidan = iota
	Godan
	Irregular
)

// verb endings
const (
	Te = "て"

	Ru    = "る"
	U     = "う"
	Masu  = "ます"
	Nai   = "ない"
	Masen = "ません"

	Ta           = "た"
	Mashita      = "ました"
	Katta        = "かった"
	Nakatta      = "なかった"
	MasenDeshita = "ませんでした"

	You       = "よう"
	Darou     = "だろう"
	Naidarou  = "ないだろう"
	Mashou    = "ましょう"
	Rudeshou  = "でしょう"
	Naideshou = "ないでしょう"

	Ro           = "ろ"
	Kudasai      = "ください" // te-kudasai
	Naidekudasai = "ないでください"
	Na           = "な"

	Tarou         = "たろう"
	Tadarou       = "ただろう"
	Nakattadarou  = "なかっただろう"
	Tadeshou      = "たでしょう"
	Nakattadeshou = "なかったでしょう"
)

type FormEnding struct {
	form    int
	endings []string
}

var PresentIndicativeEnding = FormEnding{PresentIndicative, []string{Ru, U, Masu, Nai, Masen}}
var PresumptiveEnding = FormEnding{Presumptive, []string{You, Darou, Naidarou, Mashou, Rudeshou, Naideshou}}
var ImperativeEnding = FormEnding{Imperative, []string{Ro, Kudasai, Naidekudasai, Na}}
var PastIndicativeEnding = FormEnding{PastIndicative, []string{Ta, Mashita, Katta, Nakatta, MasenDeshita}}
var PastPresumptiveEnding = FormEnding{PastPresumptive, []string{Tarou, Tadarou, Nakattadarou, Tadeshou, Nakattadeshou}}
var PresentProgressiveEnding = FormEnding{}
var PastProgressiveEnding = FormEnding{}
var ProvisionalEnding = FormEnding{}
var ConditionalEnding = FormEnding{}
var PotentialEnding = FormEnding{}
var CausativeEnding = FormEnding{}
var InfinitiveEnding = FormEnding{}
var TeFormEnding = FormEnding{TeForm, []string{Te}}

var EndingForms = []FormEnding{
	TeFormEnding,
	InfinitiveEnding,
	PresentIndicativeEnding,
	PresumptiveEnding,
	ImperativeEnding,
	PastIndicativeEnding,
	PastPresumptiveEnding,
	PresentProgressiveEnding,
	PastProgressiveEnding,
	ProvisionalEnding,
	ConditionalEnding,
	PotentialEnding,
	CausativeEnding,
}

// hasAnySuffix
func hasAnySuffix(s string, endings ...string) bool {
	for _, e := range endings {
		if strings.HasSuffix(s, e) {
			return true
		}
	}
	return false
}

// IdentifyEnding gets the longest ending for the verb
// from the list of known verb endings.
func IdentifyEnding(verb string) (ending string) {
	// Should make this a trie at some point, but... MVP
	for _, ef := range EndingForms {
		for _, e := range ef.endings {
			if strings.HasSuffix(verb, e) && len(e) > len(ending) {
				ending = e
			}
		}
	}
	return ending
}

// IdentifyForm tries to identify the verb form.
// It returns the integer form constant
func IdentifyForm(verb string) (form int) {
	if hasAnySuffix(verb, Ru, U, Masu, Nai, Masen) {
		return PresentIndicative
	} else if hasAnySuffix(verb, Ta, Mashita, Nakatta, MasenDeshita) {
		return PastIndicative
	}
	return Unknown
}

// IdentifyPositivity determines whether the verb is
// in affirmative or negative form.
func IdentifyPositivity(verb string) (positivity int) {
	if hasAnySuffix(verb, Nai, Masen, Nakatta, MasenDeshita) {
		return Negative
	}
	return Positive
}

// // IdentifyGroup tries to identify the verb class
// func IdentifyGroup(verb string) (class int) {
// 	// TODO
// 	return class
// }