// Copyright 2015 Geofrey Ernest <geofreyernest@live.com>. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

//Package i18n is a translation library.
package i18n

import (
	"github.com/melvinmt/gt"
)

// Tr contains translations.
var Tr = &gt.Build{
	Origin: "en",
	Target: "sw",
	Index: gt.Strings{
		"home-btn": {
			"en": "home",
			"sw": "nyumbani",
		},
		"jobs-btn": {
			"en": "jobs",
			"sw": "ajira",
		},
		"help-btn": {
			"en": "help",
			"sw": "msaada",
		},
		"contact-btn": {
			"en": "contact us",
			"sw": "wasiliana nasi",
		},
		"deadline": {
			"en": "deadline",
			"sw": "mwisho",
		},
		"regions": {
			"en": "regions",
			"sw": "mikoa",
		},
		"apply-btn": {
			"en": " apply now",
			"sw": "omba sasa",
		},
		"login": {
			"en": "login",
			"sw": "ingia",
		},
		"logout": {
			"en": "logout",
			"sw": "jitoe",
		},
		"register": {
			"en": "register",
			"sw": "jiunge",
		},
		"message_required": {
			"en": "this field should not be empty",
			"sw": "hili eneo halitakiwi kuachwa wazi",
		},
		"message_min_length": {
			"en": "this field should be at least %d characters",
			"sw": "namba ya siri inatakiwa kuanzia herufi %d na kuendelea",
		},
		"message_email": {
			"en": "incorrect email, should of the form example@examples.com",
			"sw": "barua pepe sio sahihi. mfano example@example.com",
		},
		"message_age": {
			"en": "age should be more than %d years",
			"sw": "umri unatakiwa uwe zaidi ya miaka %d",
		},
		"message_equal": {
			"en": "%s should be equal to %s",
			"sw": "%s inatakiwa iwe sawa na %s",
		},
		"documents": {
			"en": "documents",
			"sw": "makala",
		},
		"flash_account_create": {
			"en": "congard, your account has been successful created",
			"sw": "hongera, akaunti yako imefanikiwa kutengenezwa",
		},
		"flash_account_create_fail": {
			"en": "sorry, we can't create you account please try again later",
			"sw": "samahaani, tumeshindwa kutengeneza akaunti yako, jaribu tena baadae",
		},
		"flash_login_success": {
			"en": "welcome back",
			"sw": "karibu ",
		},
		"flash_login_failed": {
			"en": "there was a problem encountered, please check the details and try again",
			"sw": "kuna majanga mkuu, jaribu kupitia maelezo ya fomu na ujaribu tena",
		},
		"issued_by": {
			"en": "issued by",
			"sw": "imetolewa na",
		},
		"valid_name_msg": {
			"en": "field %s should conatin letters a-zA-Z e.g baba",
			"sw": "eneo %s linatakiwa liwe na herufi to a-zA-Z mfano baba",
		},
		"given_name": {
			"en": "first name",
			"sw": "jina la kwanza",
		},
		"family_name": {
			"en": "family name",
			"sw": "jina la ukoo",
		},
		"middle_name": {
			"en": "middle name",
			"sw": "jina la baba",
		},
	},
}

// CloneLang returns a copy of translations.
func CloneLang() *gt.Build {
	b := &gt.Build{}
	b.Origin = Tr.Origin
	b.Target = Tr.Target
	b.Index = Tr.Index
	return b
}
