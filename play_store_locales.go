package main

import (
	"math"
	"strings"

	"github.com/agnivade/levenshtein"
)

type locales map[string]interface{}

func (l locales) contains(locale string) bool {
	_, ok := l[locale]
	return ok
}

// closestMatch returns the closest match for the `locale` this `locales`.
func (l locales) closestMatch(locale string) string {
	d := math.MaxInt
	s := ""

	for key := range l {
		// the following has the highest priority since play store refuses many
		// locales with missing region suffix.
		splits := strings.Split(key, "-")
		if len(splits) > 1 && splits[0] == locale {
			return key
		}

		nd := levenshtein.ComputeDistance(locale, key)
		if nd < d {
			d = nd
			s = key
		}
	}

	return s
}

// playStoreLocales declares locales recognised by the Play Store Listing.
var playStoreLocales = locales{
	"af":     nil,
	"sq":     nil,
	"am":     nil,
	"ar":     nil,
	"hy-AM":  nil,
	"az-AZ":  nil,
	"bn-BD":  nil,
	"eu-ES":  nil,
	"be":     nil,
	"bg":     nil,
	"my-MM":  nil,
	"ca":     nil,
	"zh-HK":  nil,
	"zh-CN":  nil,
	"zh-TW":  nil,
	"hr":     nil,
	"cs-CZ":  nil,
	"da-DK":  nil,
	"nl-NL":  nil,
	"en-AU":  nil,
	"en-CA":  nil,
	"en-GB":  nil,
	"en-IN":  nil,
	"en-SG":  nil,
	"en-US":  nil,
	"en-ZA":  nil,
	"et":     nil,
	"fil":    nil,
	"fi-FI":  nil,
	"fr-CA":  nil,
	"fr-FR":  nil,
	"gl-ES":  nil,
	"ka-GE":  nil,
	"de-DE":  nil,
	"el-GR":  nil,
	"gu":     nil,
	"iw-IL":  nil,
	"hi-IN":  nil,
	"hu-HU":  nil,
	"is-IS":  nil,
	"id":     nil,
	"it-IT":  nil,
	"ja-JP":  nil,
	"kn-IN":  nil,
	"kk":     nil,
	"km-KH":  nil,
	"ko-KR":  nil,
	"ky-KG":  nil,
	"lo-LA":  nil,
	"lv":     nil,
	"lt":     nil,
	"mk-MK":  nil,
	"ms-MY":  nil,
	"ms":     nil,
	"ml-IN":  nil,
	"mr-IN":  nil,
	"mn-MN":  nil,
	"ne-NP":  nil,
	"no-NO":  nil,
	"fa":     nil,
	"fa-AE":  nil,
	"fa-AF":  nil,
	"fa-IR":  nil,
	"pl-PL":  nil,
	"pt-BR":  nil,
	"pt-PT":  nil,
	"pa":     nil,
	"ro":     nil,
	"rm":     nil,
	"ru-RU":  nil,
	"sr":     nil,
	"si-LK":  nil,
	"sk":     nil,
	"sl":     nil,
	"es-419": nil,
	"es-ES":  nil,
	"es-US":  nil,
	"sw":     nil,
	"sv-SE":  nil,
	"ta-IN":  nil,
	"te-IN":  nil,
	"th":     nil,
	"tr-TR":  nil,
	"uk":     nil,
	"ur":     nil,
	"vi":     nil,
	"zu":     nil,
}
