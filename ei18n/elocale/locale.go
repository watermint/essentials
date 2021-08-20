package elocale

import (
	"essentials/eidiom"
	"fmt"
	"strings"
)

type Language string

const (
	TagEnglish  Language = "en"
	TagJapanese Language = "ja"
)

var (
	// Default fallback locale
	Default  = English
	English  = mustNew(string(TagEnglish))
	Japanese = mustNew(string(TagJapanese))
)

func mustNew(langTag string) Locale {
	if lc, err := New(langTag); err != nil {
		panic(err)
	} else {
		return lc
	}
}

func New(langTag string) (local Locale, err error) {
	matches := bcp47ReM.FindStringSubmatch(langTag)
	if len(matches) < 1 {
		return nil, eidiom.ErrorParseInvalidFormat
	}
	language := strings.ToLower(matches[1])
	if language == "c" {
		return newLocale(string(TagEnglish)), nil
	}
	return newLocale(language), nil
}

// Locale represents a specific geopolitical region.
// Remarks: The immediate goal of this interface is to select a display language from
// small number of supported languages (such as two languages like English & Japanese).
// This interface will not consider region (ISO 3166 or UN M49), variants (BCP 47), scripts (ISO 15924), etc.
type Locale interface {
	// Stringer Language tag
	fmt.Stringer

	// Language ISO 639 code (2-3 letter code)
	Language() Language

	// LanguageTwoLetter ISO 639-1 Two-letter code.
	// Returns empty if the language is not defined in 639-1.
	LanguageTwoLetter() string
}

func newLocale(code string) Locale {
	return &localeImpl{code: Language(code)}
}

type localeImpl struct {
	code Language
}

func (z localeImpl) LanguageTwoLetter() string {
	if len(z.code) == 2 {
		return string(z.code)
	}
	if two, ok := iso631ThreeToTwoLetter[string(z.code)]; ok {
		return two
	}
	return ""
}

func (z localeImpl) String() string {
	return string(z.code)
}

func (z localeImpl) Language() Language {
	return z.code
}
