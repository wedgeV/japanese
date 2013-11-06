package japanese

import (
	. "launchpad.net/gocheck"
	"testing"
)

// Hook up gocheck into the "go test" runner.
func Test(t *testing.T) { TestingT(t) }

type JapaneseSuite struct {
	RuVerbs        []RuVerb
	UVerbs         []UVerb
	ExceptionVerbs []ExceptionVerb
}

type TestWord struct {
	kanji string
	kana  string

	neg_kanji string
	neg_kana  string

	past_kanji string
	past_kana  string

	masu_kanji string
	masu_kana  string

	long_neg_kanji string
	long_neg_kana  string

	long_past_kanji string
	long_past_kana  string
}

var _ = Suite(&JapaneseSuite{})

// ru-verbs
var ruVerbs []TestWord = []TestWord{
	TestWord{"見る", "みる", "見ない", "みない", "見た", "みた", "見ます", "みます", "見ません", "みません", "見ました", "みました"},
	TestWord{"食べる", "たべる", "食べない", "たべない", "食べた", "たべた", "食べます", "たべます", "食べません", "たべません", "食べました", "たべました"},
	TestWord{"寝る", "ねる", "寝ない", "ねない", "寝た", "ねた", "寝ます", "ねます", "寝ません", "ねません", "寝ました", "ねました"},
	TestWord{"起きる", "おきる", "起きない", "おきない", "起きた", "おきた", "起きます", "おきます", "起きません", "おきません", "起きました", "おきました"},
	TestWord{"考える", "かんがえる", "考えない", "かんがえない", "考えた", "かんがえた", "考えます", "かんがえます", "考えません", "かんがえません", "考えました", "かんがえました"},
	TestWord{"教える", "おしえる", "教えない", "おしえない", "教えた", "おしえた", "教えます", "おしえます", "教えません", "おしえません", "教えました", "おしえました"},
	TestWord{"出る", "でる", "出ない", "でない", "出た", "でた", "出ます", "でます", "出ません", "でません", "出ました", "でました"},
	TestWord{"着る", "きる", "着ない", "きない", "着た", "きた", "着ます", "きます", "着ません", "きません", "着ました", "きました"},
	TestWord{"いる", "いる", "いない", "いない", "いた", "いた", "います", "います", "いません", "いません", "いました", "いました"},
}

// u-verbs
var uVerbs []TestWord = []TestWord{
	TestWord{"話す", "はなす", "話さない", "はなさない", "話した", "はなした"},
	TestWord{"聞く", "きく", "聞かない", "きかない", "聞いた", "きいた"},
	TestWord{"泳ぐ", "およぐ", "泳がない", "およがない", "泳いた", "およいた"},
	TestWord{"遊ぶ", "あそぶ", "遊ばない", "あそばない", "遊んだ", "あそんだ"},
	TestWord{"待つ", "まつ", "待たない", "またない", "待った", "まった"},
	TestWord{"飲む", "のむ", "飲まない", "のまない", "飲んだ", "のんだ"},
	TestWord{"買う", "かう", "買わない", "かわない", "買った", "かった"},
	TestWord{"帰る", "かえる", "帰らない", "かえらない", "帰った", "かえった"},
	TestWord{"死ぬ", "しぬ", "死なない", "しなない", "死んだ", "しんだ"},
	TestWord{"ある", "ある", "ない", "ない", "あった", "あった"},
}

// exceptions
var exceptionVerbs []TestWord = []TestWord{
	TestWord{"する", "する", "しない", "しない", "した", "した"},
	TestWord{"くる", "くる", "こない", "こない", "きた", "きた"},
}

func (s *JapaneseSuite) SetUpSuite(c *C) {
	s.RuVerbs = []RuVerb{}
	for _, verb := range ruVerbs {
		v := RuVerb{Verb{Word{verb.kanji, verb.kana}}}
		s.RuVerbs = append(s.RuVerbs, v)
	}
	s.UVerbs = []UVerb{}
	for _, verb := range uVerbs {
		v := UVerb{Verb{Word{verb.kanji, verb.kana}}}
		s.UVerbs = append(s.UVerbs, v)
	}
	s.ExceptionVerbs = []ExceptionVerb{}
	for _, verb := range exceptionVerbs {
		v := ExceptionVerb{Verb{Word{verb.kanji, verb.kana}}}
		s.ExceptionVerbs = append(s.ExceptionVerbs, v)
	}
}

func (s *JapaneseSuite) TestNegativeRuVerbs(c *C) {
	// check that ru-verbs get the correct negative suffix
	for i := range s.RuVerbs {
		v := s.RuVerbs[i]
		t := ruVerbs[i]
		neg_word := v.Negative()
		c.Check(neg_word.kanji, Equals, t.neg_kanji)
		c.Check(neg_word.kana, Equals, t.neg_kana)
	}
}

func (s *JapaneseSuite) TestPastRuVerbs(c *C) {
	// check that ru-verbs get the correct past suffix
	for i := range s.RuVerbs {
		v := s.RuVerbs[i]
		t := ruVerbs[i]
		past_word := v.Past()
		c.Check(past_word.kanji, Equals, t.past_kanji)
		c.Check(past_word.kana, Equals, t.past_kana)
	}
}

func (s *JapaneseSuite) TestNegativeUVerbs(c *C) {
	// check that ru-verbs get the correct negative suffix
	for _, verb := range uVerbs {
		v := UVerb{Verb{Word{verb.kanji, verb.kana}}}
		neg_word := v.Negative()
		c.Check(neg_word.kanji, Equals, verb.neg_kanji)
		c.Check(neg_word.kana, Equals, verb.neg_kana)
	}
}

func (s *JapaneseSuite) TestPastUVerbs(c *C) {
	// check that ru-verbs get the correct past suffix
	for i := range s.UVerbs {
		v := s.UVerbs[i]
		t := uVerbs[i]
		past_word := v.Past()
		c.Check(past_word.kanji, Equals, t.past_kanji)
		c.Check(past_word.kana, Equals, t.past_kana)
	}
}

func (s *JapaneseSuite) TestNegativeExceptionVerbs(c *C) {
	// check that exception-verbs get the correct negative suffix
	for _, verb := range exceptionVerbs {
		v := ExceptionVerb{Verb{Word{verb.kanji, verb.kana}}}
		neg_word := v.Negative()
		c.Check(neg_word.kanji, Equals, verb.neg_kanji)
		c.Check(neg_word.kana, Equals, verb.neg_kana)
	}
}

func (s *JapaneseSuite) TestPastExceptionVerbs(c *C) {
	// check that exception-verbs get the correct past suffix
	for i := range s.ExceptionVerbs {
		v := s.ExceptionVerbs[i]
		t := exceptionVerbs[i]
		past_word := v.Past()
		c.Check(past_word.kanji, Equals, t.past_kanji)
		c.Check(past_word.kana, Equals, t.past_kana)
	}
}
