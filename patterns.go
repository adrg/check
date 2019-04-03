package check

import "regexp"

const (
	patternURL = "^" +
		// protocol identifier
		"(?:(?:https?|ftp)://)?" +
		// user:pass authentication
		`(?:\S+(?::\S*)?@)?` +
		"(?:" +
		// IP address dotted notation octets
		`(?:[1-9]\d?|1\d\d|2[01]\d|22[0-3])` +
		`(?:\.(?:1?\d{1,2}|2[0-4]\d|25[0-5])){2}` +
		`(?:\.(?:[1-9]\d?|1\d\d|2[0-4]\d|25[0-4]))` +
		"|" +
		// host name
		"(?:(?:[a-z\u00a1-\uffff0-9]-*)*[a-z\u00a1-\uffff0-9]+)" +
		// domain name
		"(?:\\.(?:[a-z\u00a1-\uffff0-9]-*)*[a-z\u00a1-\uffff0-9]+)*" +
		// TLD identifier
		"(?:\\.(?:[a-z\u00a1-\uffff]{2,}))" +
		")" +
		// port number
		`(?::\d{2,5})?` +
		// resource path
		`(?:/\S*)?` +
		"$"

	patternVAT = "^(" +
		"(AT)?U[0-9]{8}|" +
		"(BE)?0[0-9]{9}|" +
		"(BG)?[0-9]{9,10}|" +
		"(CY)?[0-9]{8}L|" +
		"(CZ)?[0-9]{8,10}|" +
		"(DE)?[0-9]{9}|" +
		"(DK)?[0-9]{8}|" +
		"(EE)?[0-9]{9}|" +
		"(EL|GR)?[0-9]{9}|" +
		"(ES)?[0-9A-Z][0-9]{7}[0-9A-Z]|" +
		"(FI)?[0-9]{8}|" +
		"(FR)?[0-9A-Z]{2}[0-9]{9}|" +
		"(GB)?([0-9]{9}([0-9]{3})?|[A-Z]{2}[0-9]{3})|" +
		"(HU)?[0-9]{8}|" +
		"(IE)?[0-9]S[0-9]{5}L|" +
		"(IT)?[0-9]{11}|" +
		"(LT)?([0-9]{9}|[0-9]{12})|" +
		"(LU)?[0-9]{8}|" +
		"(LV)?[0-9]{11}|" +
		"(MT)?[0-9]{8}|" +
		"(NL)?[0-9]{9}B[0-9]{2}|" +
		"(PL)?[0-9]{10}|" +
		"(PT)?[0-9]{9}|" +
		"(RO)?[0-9]{2,10}|" +
		"(SE)?[0-9]{12}|" +
		"(SI)?[0-9]{8}|" +
		"(SK)?[0-9]{10}" +
		")$"

	patternIBAN = "^" +
		"[a-zA-Z]{2}" +
		"[0-9]{2}" +
		"[a-zA-Z0-9]{4}" +
		"[0-9]{7}" +
		"([a-zA-Z0-9]?){0,16}" +
		"$"
)

var (
	regURL  = regexp.MustCompile(patternURL)
	regVAT  = regexp.MustCompile(patternVAT)
	regIBAN = regexp.MustCompile(patternIBAN)
)
