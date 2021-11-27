//go:build windows
// +build windows

package elocale

import (
	"github.com/watermint/essentials/eidiom"
	"github.com/watermint/essentials/eidiom/eoutcome"
	"github.com/watermint/essentials/enative/ewindows"
)

func currentLocaleWithSysCall(apiName string) (string, eidiom.Outcome) {
	locName := ewindows.NewBufferString(localeNameMaxLength)
	r, _, oc := ewindows.Kernel32.Call(apiName, locName.Pointer(), locName.BufSize())
	switch {
	case oc.HasError():
		return "", eoutcome.NewOutcomeBaseError(oc.Cause())
	case r == 0:
		return "", eoutcome.NewOutcomeBaseError(oc.LastError())
	default:
		return locName.String(), eoutcome.NewOutcomeBaseOk()
	}
}

func currentLocaleString() (string, error) {
	ul, oc := currentLocaleWithSysCall("GetUserDefaultLocaleName")
	if oc.IsOk() {
		return ul, nil
	}

	sl, oc := currentLocaleWithSysCall("GetSystemDefaultLocaleName")
	if oc.IsOk() {
		return sl, nil
	}

	return "", oc.Cause()
}

const (
	// https://docs.microsoft.com/en-us/windows/win32/intl/locale-name-constants
	localeNameMaxLength = 85
)

var (
	// https://docs.microsoft.com/en-us/openspecs/windows_protocols/ms-lcid/70feba9f-294e-491e-b6eb-56532684c37f
	msLcidToTag = map[uintptr]string{
		0x0001: "ar",
		0x0002: "bg",
		0x0003: "ca",
		0x0004: "zh-Hans",
		0x0005: "cs",
		0x0006: "da",
		0x0007: "de",
		0x0008: "el",
		0x0009: "en",
		0x000A: "es",
		0x000B: "fi",
		0x000C: "fr",
		0x000D: "he",
		0x000E: "hu",
		0x000F: "is",
		0x0010: "it",
		0x0011: "ja",
		0x0012: "ko",
		0x0013: "nl",
		0x0014: "no",
		0x0015: "pl",
		0x0016: "pt",
		0x0017: "rm",
		0x0018: "ro",
		0x0019: "ru",
		0x001A: "hr",
		0x001B: "sk",
		0x001C: "sq",
		0x001D: "sv",
		0x001E: "th",
		0x001F: "tr",
		0x0020: "ur",
		0x0021: "id",
		0x0022: "uk",
		0x0023: "be",
		0x0024: "sl",
		0x0025: "et",
		0x0026: "lv",
		0x0027: "lt",
		0x0028: "tg",
		0x0029: "fa",
		0x002A: "vi",
		0x002B: "hy",
		0x002C: "az",
		0x002D: "eu",
		0x002E: "hsb",
		0x002F: "mk",
		0x0030: "st",
		0x0031: "ts",
		0x0032: "tn",
		0x0033: "ve",
		0x0034: "xh",
		0x0035: "zu",
		0x0036: "af",
		0x0037: "ka",
		0x0038: "fo",
		0x0039: "hi",
		0x003A: "mt",
		0x003B: "se",
		0x003C: "ga",
		0x003D: "yi", // reserved
		0x003E: "ms",
		0x003F: "kk",
		0x0040: "ky",
		0x0041: "sw",
		0x0042: "tk",
		0x0043: "uz",
		0x0044: "tt",
		0x0045: "bn",
		0x0046: "pa",
		0x0047: "gu",
		0x0048: "or",
		0x0049: "ta",
		0x004A: "te",
		0x004B: "kn",
		0x004C: "ml",
		0x004D: "as",
		0x004E: "mr",
		0x004F: "sa",
		0x0050: "mn",
		0x0051: "bo",
		0x0052: "cy",
		0x0053: "km",
		0x0054: "lo",
		0x0055: "my",
		0x0056: "gl",
		0x0057: "kok",
		0x0058: "mni", // reserved
		0x0059: "sd",
		0x005A: "syr",
		0x005B: "si",
		0x005C: "chr",
		0x005D: "iu",
		0x005E: "am",
		0x005F: "tzm",
		0x0060: "ks",
		0x0061: "ne",
		0x0062: "fy",
		0x0063: "ps",
		0x0064: "fil",
		0x0065: "dv",
		0x0066: "bin", // reserved
		0x0067: "ff",
		0x0068: "ha",
		0x0069: "ibb", // reserved
		0x006A: "yo",
		0x006B: "quz",
		0x006C: "nso",
		0x006D: "ba",
		0x006E: "lb",
		0x006F: "kl",
		0x0070: "ig",
		0x0071: "kr", // reserved
		0x0072: "om",
		0x0073: "ti",
		0x0074: "gn",
		0x0075: "haw",
		0x0076: "la", // reserved
		0x0077: "so", // reserved
		0x0078: "ii",
		0x0079: "pap", // reserved
		0x007A: "arn",
		//0x007B: "Neither defined nor reserved",
		0x007C: "moh",
		//0x007D: "Neither defined nor reserved",
		0x007E: "br",
		//0x007F: "Reserved for invariant locale behavior",
		0x0080: "ug",
		0x0081: "mi",
		0x0082: "oc",
		0x0083: "co",
		0x0084: "gsw",
		0x0085: "sah",
		0x0086: "qut",
		0x0087: "rw",
		0x0088: "wo",
		//0x0089: "Neither defined nor reserved",
		//0x008A: "Neither defined nor reserved",
		//0x008B: "Neither defined nor reserved",
		0x008C: "prs",
		//0x008D: "Neither defined nor reserved",
		//0x008E: "Neither defined nor reserved",
		//0x008F: "Neither defined nor reserved",
		//0x0090: "Neither defined nor reserved",
		0x0091: "gd",
		0x0092: "ku",
		0x0093: "quc", // reserved
		0x0401: "ar-SA",
		0x0402: "bg-BG",
		0x0403: "ca-ES",
		0x0404: "zh-TW",
		0x0405: "cs-CZ",
		0x0406: "da-DK",
		0x0407: "de-DE",
		0x0408: "el-GR",
		0x0409: "en-US",
		0x040A: "es-ES_tradnl",
		0x040B: "fi-FI",
		0x040C: "fr-FR",
		0x040D: "he-IL",
		0x040E: "hu-HU",
		0x040F: "is-IS",
		0x0410: "it-IT",
		0x0411: "ja-JP",
		0x0412: "ko-KR",
		0x0413: "nl-NL",
		0x0414: "nb-NO",
		0x0415: "pl-PL",
		0x0416: "pt-BR",
		0x0417: "rm-CH",
		0x0418: "ro-RO",
		0x0419: "ru-RU",
		0x041A: "hr-HR",
		0x041B: "sk-SK",
		0x041C: "sq-AL",
		0x041D: "sv-SE",
		0x041E: "th-TH",
		0x041F: "tr-TR",
		0x0420: "ur-PK",
		0x0421: "id-ID",
		0x0422: "uk-UA",
		0x0423: "be-BY",
		0x0424: "sl-SI",
		0x0425: "et-EE",
		0x0426: "lv-LV",
		0x0427: "lt-LT",
		0x0428: "tg-Cyrl-TJ",
		0x0429: "fa-IR",
		0x042A: "vi-VN",
		0x042B: "hy-AM",
		0x042C: "az-Latn-AZ",
		0x042D: "eu-ES",
		0x042E: "hsb-DE",
		0x042F: "mk-MK",
		0x0430: "st-ZA",
		0x0431: "ts-ZA",
		0x0432: "tn-ZA",
		0x0433: "ve-ZA",
		0x0434: "xh-ZA",
		0x0435: "zu-ZA",
		0x0436: "af-ZA",
		0x0437: "ka-GE",
		0x0438: "fo-FO",
		0x0439: "hi-IN",
		0x043A: "mt-MT",
		0x043B: "se-NO",
		0x043D: "yi-001",
		0x043E: "ms-MY",
		0x043F: "kk-KZ",
		0x0440: "ky-KG",
		0x0441: "sw-KE",
		0x0442: "tk-TM",
		0x0443: "uz-Latn-UZ",
		0x0444: "tt-RU",
		0x0445: "bn-IN",
		0x0446: "pa-IN",
		0x0447: "gu-IN",
		0x0448: "or-IN",
		0x0449: "ta-IN",
		0x044A: "te-IN",
		0x044B: "kn-IN",
		0x044C: "ml-IN",
		0x044D: "as-IN",
		0x044E: "mr-IN",
		0x044F: "sa-IN",
		0x0450: "mn-MN",
		0x0451: "bo-CN",
		0x0452: "cy-GB",
		0x0453: "km-KH",
		0x0454: "lo-LA",
		0x0455: "my-MM",
		0x0456: "gl-ES",
		0x0457: "kok-IN",
		0x0458: "mni-IN",     // reserved
		0x0459: "sd-Deva-IN", // reserved
		0x045A: "syr-SY",
		0x045B: "si-LK",
		0x045C: "chr-Cher-US",
		0x045D: "iu-Cans-CA",
		0x045E: "am-ET",
		0x045F: "tzm-Arab-MA",
		0x0460: "ks-Arab",
		0x0461: "ne-NP",
		0x0462: "fy-NL",
		0x0463: "ps-AF",
		0x0464: "fil-PH",
		0x0465: "dv-MV",
		0x0466: "bin-NG", // reserved
		0x0467: "ff-NG",  // ff-Latn-NG
		0x0468: "ha-Latn-NG",
		0x0469: "ibb-NG", // reserved
		0x046A: "yo-NG",
		0x046B: "quz-BO",
		0x046C: "nso-ZA",
		0x046D: "ba-RU",
		0x046E: "lb-LU",
		0x046F: "kl-GL",
		0x0470: "ig-NG",
		0x0471: "kr-Latn-NG",
		0x0472: "om-ET",
		0x0473: "ti-ET",
		0x0474: "gn-PY",
		0x0475: "haw-US",
		0x0476: "la-VA",
		0x0477: "so-SO",
		0x0478: "ii-CN",
		0x0479: "pap-029", // reserved
		0x047A: "arn-CL",
		0x047C: "moh-CA",
		0x047E: "br-FR",
		0x0480: "ug-CN",
		0x0481: "mi-NZ",
		0x0482: "oc-FR",
		0x0483: "co-FR",
		0x0484: "gsw-FR",
		0x0485: "sah-RU",
		0x0486: "qut-GT", // reserved
		0x0487: "rw-RW",
		0x0488: "wo-SN",
		0x048C: "prs-AF",
		0x048D: "plt-MG",      // reserved
		0x048E: "zh-yue-HK",   // reserved
		0x048F: "tdd-Tale-CN", // reserved
		0x0490: "khb-Talu-CN", // reserved
		0x0491: "gd-GB",
		0x0492: "ku-Arab-IQ",
		0x0493: "quc-CO", // reserved
		0x0501: "qps-ploc",
		0x05FE: "qps-ploca",
		0x0801: "ar-IQ",
		0x0803: "ca-ES-valencia",
		0x0804: "zh-CN",
		0x0807: "de-CH",
		0x0809: "en-GB",
		0x080A: "es-MX",
		0x080C: "fr-BE",
		0x0810: "it-CH",
		0x0811: "ja-Ploc-JP", // reserved
		0x0813: "nl-BE",
		0x0814: "nn-NO",
		0x0816: "pt-PT",
		0x0818: "ro-MD",
		0x0819: "ru-MD",
		0x081A: "sr-Latn-CS",
		0x081D: "sv-FI",
		0x0820: "ur-IN",
		//0x0827: "Neither defined nor reserved",
		0x082C: "az-Cyrl-AZ", // reserved
		0x082E: "dsb-DE",
		0x0832: "tn-BW",
		0x083B: "se-SE",
		0x083C: "ga-IE",
		0x083E: "ms-BN",
		0x083F: "kk-Latn-KZ", // reserved
		0x0843: "uz-Cyrl-UZ", // reserved
		0x0845: "bn-BD",
		0x0846: "pa-Arab-PK",
		0x0849: "ta-LK",
		0x0850: "mn-Mong-CN", // reserved
		0x0851: "bo-BT",      // reserved
		0x0859: "sd-Arab-PK",
		0x085D: "iu-Latn-CA",
		0x085F: "tzm-Latn-DZ",
		0x0860: "ks-Deva-IN",
		0x0861: "ne-IN",
		0x0867: "ff-Latn-SN",
		0x086B: "quz-EC",
		0x0873: "ti-ER",
		0x09FF: "qps-plocm",
		//0x0C00: "Locale without assigned LCID if the current user default locale. See section 2.2.1.",
		0x0C01: "ar-EG",
		0x0C04: "zh-HK",
		0x0C07: "de-AT",
		0x0C09: "en-AU",
		0x0C0A: "es-ES",
		0x0C0C: "fr-CA",
		0x0C1A: "sr-Cyrl-CS",
		0x0C3B: "se-FI",
		0x0C50: "mn-Mong-MN",
		0x0C51: "dz-BT",
		0x0C5F: "tmz-MA", // reserved
		0x0C6b: "quz-PE",
		//0x1000: "Locale without assigned LCID if the current user default locale. See section 2.2.1.",
		0x1001: "ar-LY",
		0x1004: "zh-SG",
		0x1007: "de-LU",
		0x1009: "en-CA",
		0x100A: "es-GT",
		0x100C: "fr-CH",
		0x101A: "hr-BA",
		0x103B: "smj-NO",
		0x105F: "tzm-Tfng-MA",
		0x1401: "ar-DZ",
		0x1404: "zh-MO",
		0x1407: "de-LI",
		0x1409: "en-NZ",
		0x140A: "es-CR",
		0x140C: "fr-LU",
		0x141A: "bs-Latn-BA",
		0x143B: "smj-SE",
		0x1801: "ar-MA",
		0x1809: "en-IE",
		0x180A: "es-PA",
		0x180C: "fr-MC",
		0x181A: "sr-Latn-BA",
		0x183B: "sma-NO",
		0x1C01: "ar-TN",
		0x1C09: "en-ZA",
		0x1C0A: "es-DO",
		0x1C0C: "fr-029",
		0x1C1A: "sr-Cyrl-BA",
		0x1C3B: "sma-SE",
		0x2001: "ar-OM",
		//0x2008: "Neither defined nor reserved",
		0x2009: "en-JM",
		0x200A: "es-VE",
		0x200C: "fr-RE",
		0x201A: "bs-Cyrl-BA",
		0x203B: "sms-FI",
		0x2401: "ar-YE",
		0x2409: "en-029, reserved",
		0x240A: "es-CO",
		0x240C: "fr-CD",
		0x241A: "sr-Latn-RS",
		0x243B: "smn-FI",
		0x2801: "ar-SY",
		0x2809: "en-BZ",
		0x280A: "es-PE",
		0x280C: "fr-SN",
		0x281A: "sr-Cyrl-RS",
		0x2C01: "ar-JO",
		0x2C09: "en-TT",
		0x2C0A: "es-AR",
		0x2C0C: "fr-CM",
		0x2C1A: "sr-Latn-ME",
		//0x3000: "Unassigned LCID locale temporarily assigned to LCID 0x3000. See section 2.2.1.",
		0x3001: "ar-LB",
		0x3009: "en-ZW",
		0x300A: "es-EC",
		0x300C: "fr-CI",
		0x301A: "sr-Cyrl-ME",
		//0x3400: "Unassigned LCID locale temporarily assigned to LCID 0x3400. See section 2.2.1.",
		0x3401: "ar-KW",
		0x3409: "en-PH",
		0x340A: "es-CL",
		0x340C: "fr-ML",
		//0x3800: "Unassigned LCID locale temporarily assigned to LCID 0x3800. See section 2.2.1.",
		0x3801: "ar-AE",
		0x3809: "en-ID", // reserved
		0x380A: "es-UY",
		0x380C: "fr-MA",
		//0x3C00: "Unassigned LCID locale temporarily assigned to LCID 0x3C00. See section 2.2.1.",
		0x3C01: "ar-BH",
		0x3C09: "en-HK",
		0x3C0A: "es-PY",
		0x3C0C: "fr-HT",
		//0x4000: "Unassigned LCID locale temporarily assigned to LCID 0x4000. See section 2.2.1.",
		0x4001: "ar-QA",
		0x4009: "en-IN",
		0x400A: "es-BO",
		//0x4400: "Unassigned LCID locale temporarily assigned to LCID 0x4400. See section 2.2.1.",
		0x4401: "ar-Ploc-SA, reserved",
		0x4409: "en-MY",
		0x440A: "es-SV",
		//0x4800: "Unassigned LCID locale temporarily assigned to LCID 0x4800. See section 2.2.1.",
		0x4801: "ar-145, reserved",
		0x4809: "en-SG",
		0x480A: "es-HN",
		//0x4C00: "Unassigned LCID locale temporarily assigned to LCID 0x4C00. See section 2.2.1.",
		0x4C09: "en-AE",
		0x4C0A: "es-NI",
		0x5009: "en-BH", // reserved
		0x500A: "es-PR",
		0x5409: "en-EG", // reserved
		0x540A: "es-US",
		0x5809: "en-JO",  // reserved
		0x580A: "es-419", // reserved
		0x5C09: "en-KW",  // reserved
		0x5C0A: "es-CU",
		0x6009: "en-TR", // reserved
		0x6409: "en-YE", // reserved
		0x641A: "bs-Cyrl",
		0x681A: "bs-Latn",
		0x6C1A: "sr-Cyrl",
		0x701A: "sr-Latn",
		0x703B: "smn",
		0x742C: "az-Cyrl",
		0x743B: "sms",
		0x7804: "zh",
		0x7814: "nn",
		0x781A: "bs",
		0x782C: "az-Latn",
		0x783B: "sma",
		0x783F: "kk-Cyrl", // reserved
		0x7843: "uz-Cyrl",
		0x7850: "mn-Cyrl",
		0x785D: "iu-Cans",
		0x785F: "tzm-Tfng",
		0x7C04: "zh-Hant",
		0x7C14: "nb",
		0x7C1A: "sr",
		0x7C28: "tg-Cyrl",
		0x7C2E: "dsb",
		0x7C3B: "smj",
		0x7C3F: "kk-Latn", // reserved
		0x7C43: "uz-Latn",
		0x7C46: "pa-Arab",
		0x7C50: "mn-Mong",
		0x7C59: "sd-Arab",
		0x7C5C: "chr-Cher",
		0x7C5D: "iu-Latn",
		0x7C5F: "tzm-Latn",
		0x7C67: "ff-Latn",
		0x7C68: "ha-Latn",
		0x7C92: "ku-Arab",
		//0xF2EE: "reserved",
		0xE40C: "fr-015", // reserved
		//0xEEEE: "reserved",
	}
)
