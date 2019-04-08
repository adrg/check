package check

import (
	"errors"
	"fmt"
	"net"
	"net/mail"
	"regexp"
	"strings"
)

var errEmpty = errors.New("empty argument")

// Required checks if any of the passed in arguments is empty. Returns an error
// on the first empty value it encounters.
// An argument is considered to be empty if it is:
// - nil
// - the zero value of its type
// - an array, a channel, a slice, a map, or a string of length 0
// - an interface whose underlying value is empty
// - a pointer which points to an empty value
func Required(args ...interface{}) ValidateFunc {
	return func() error {
		for _, arg := range args {
			if isEmpty(arg) {
				return errEmpty
			}
		}

		return nil
	}
}

// Eq checks if x is equal to the comparison term.
func Eq(x, term interface{}) ValidateFunc {
	return func() error {
		cmpField, err := newCmpField(eq, term)
		if err != nil {
			return err
		}

		return compare(x, cmpField)
	}
}

// Ne checks if x is not equal to the comparison term.
func Ne(x, term interface{}) ValidateFunc {
	return func() error {
		cmpField, err := newCmpField(ne, term)
		if err != nil {
			return err
		}

		return compare(x, cmpField)
	}
}

// Lt checks if x is less than the comparison term.
// Should be used for numeric types or time.Time.
func Lt(x, term interface{}) ValidateFunc {
	return func() error {
		cmpField, err := newCmpField(lt, term)
		if err != nil {
			return err
		}

		return compare(x, cmpField)
	}
}

// Lte checks if x is less than or equal to the comparison term.
// Should be used for numeric types or time.Time.
func Lte(x, term interface{}) ValidateFunc {
	return func() error {
		cmpField, err := newCmpField(lte, term)
		if err != nil {
			return err
		}

		return compare(x, cmpField)
	}
}

// Gt checks if x is greater than the comparison term.
// Should be used for numeric types or time.Time.
func Gt(x, term interface{}) ValidateFunc {
	return func() error {
		cmpField, err := newCmpField(gt, term)
		if err != nil {
			return err
		}

		return compare(x, cmpField)
	}
}

// Gte checks if x is greater than or equal to the comparison term.
// Should be used for numeric types or time.Time.
func Gte(x, term interface{}) ValidateFunc {
	return func() error {
		cmpField, err := newCmpField(gte, term)
		if err != nil {
			return err
		}

		return compare(x, cmpField)
	}
}

// Between checks if x is greater than or equal to the lower bound and
// less than or equal to the upper bound.
// Should be used for numeric types or time.Time.
func Between(x, lower interface{}, upper interface{}) ValidateFunc {
	return func() error {
		cmpField, err := newCmpField(gte, lower)
		if err != nil {
			return err
		}
		if err = compare(x, cmpField); err != nil {
			return err
		}

		cmpField, err = newCmpField(lte, upper)
		if err != nil {
			return err
		}
		return compare(x, cmpField)
	}
}

// In verifies that x is equal to one of the elems values.
func In(x interface{}, elems ...interface{}) ValidateFunc {
	return func() error {
		for _, elem := range elems {
			cmpField, err := newCmpField(eq, elem)
			if err != nil {
				return err
			}
			if err = compare(x, cmpField); err == nil {
				return nil
			}
		}

		return fmt.Errorf("`in` comparison failed: `%v` not in `%v`", x, elems)
	}
}

// NotIn verifies that x is not equal to any of the elems values.
func NotIn(x interface{}, elems ...interface{}) ValidateFunc {
	return func() error {
		for _, elem := range elems {
			cmpField, err := newCmpField(eq, elem)
			if err != nil {
				return err
			}
			if err = compare(x, cmpField); err == nil {
				return fmt.Errorf("`not in` comparison failed: `%v` in `%v`", x, elems)
			}
		}

		return nil
	}
}

// Matches checks if the val parameter matches the pattern (regular expression).
// The value can be empty if the required parameter is false.
func Matches(val, pattern string, required bool) ValidateFunc {
	return func() error {
		if isEmptyStr(val) {
			return requiredErr(required, "match term cannot be empty")
		}

		ok, err := regexp.MatchString(pattern, val)
		if err != nil {
			return fmt.Errorf("invalid pattern `%s`", pattern)
		}
		if !ok {
			return fmt.Errorf("`%s` does not match pattern `%s`", val, pattern)
		}

		return nil
	}
}

// Email checks if the email parameter is a valid email.
// The email can be empty if the required parameter is false.
func Email(email string, required bool) ValidateFunc {
	return func() error {
		if isEmptyStr(email) {
			return requiredErr(required, "email address cannot be empty")
		}

		if _, err := mail.ParseAddress(email); err != nil {
			return fmt.Errorf("invalid email address `%s`", email)
		}

		return nil
	}
}

// EmailList checks if the list parameter is a valid email address list.
// The list can be empty if the required parameter is false.
func EmailList(list string, required bool) ValidateFunc {
	return func() error {
		if list = stripSpaces(list); isEmptyStr(list) {
			return requiredErr(required, "email address list cannot be empty")
		}

		emails := strings.Split(list, ",")
		for _, email := range emails {
			if _, err := mail.ParseAddress(email); err != nil {
				return fmt.Errorf("invalid email address `%s`", email)
			}
		}

		return nil
	}
}

// URL checks if the url parameter is a valid URL.
// The URL can be empty if the required parameter is false.
func URL(url string, required bool) ValidateFunc {
	return func() error {
		if isEmptyStr(url) {
			return requiredErr(required, "URL cannot be empty")
		}
		if ok := regURL.MatchString(url); !ok {
			return fmt.Errorf("invalid URL `%s`", url)
		}

		return nil
	}
}

// IBAN checks if the iban parameter is a valid IBAN.
// The IBAN can be empty if the required parameter is false.
func IBAN(iban string, required bool) ValidateFunc {
	return func() error {
		if isEmptyStr(iban) {
			return requiredErr(required, "IBAN cannot be empty")
		}
		if ok := regIBAN.MatchString(iban); !ok {
			return fmt.Errorf("invalid IBAN `%s`", iban)
		}

		return nil
	}
}

// VAT checks if the vat parameter is a valid VAT number.
// The VAT number can be empty if the required parameter is false.
func VAT(vat string, required bool) ValidateFunc {
	return func() error {
		if isEmptyStr(vat) {
			return requiredErr(required, "VAT number cannot be empty")
		}
		if ok := regVAT.MatchString(vat); !ok {
			return fmt.Errorf("invalid VAT number `%s`", vat)
		}

		return nil
	}
}

// IP checks if the ip parameter is a valid IPv4 or IPv6 address.
// The IP address can be empty if the required parameter is false.
func IP(ip string, required bool) ValidateFunc {
	return func() error {
		if isEmptyStr(ip) {
			return requiredErr(required, "IP address cannot be empty")
		}
		if addr := net.ParseIP(ip); addr == nil {
			return fmt.Errorf("invalid IP address `%s`", ip)
		}

		return nil
	}
}

// MAC checks if the mac parameter is a valid MAC address.
// The MAC address can be empty if the required parameter is false.
func MAC(mac string, required bool) ValidateFunc {
	return func() error {
		if isEmptyStr(mac) {
			return requiredErr(required, "MAC address cannot be empty")
		}
		if _, err := net.ParseMAC(mac); err != nil {
			return fmt.Errorf("invalid mac address `%s`", mac)
		}

		return nil
	}
}
