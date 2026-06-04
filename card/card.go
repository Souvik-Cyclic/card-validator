// Package card validates payment card numbers with the Luhn algorithm and
// identifies the issuing network from a number's prefix and length.
//
// Input may contain spaces or dashes (e.g. "4111 1111 1111 1111"); they are
// ignored. Issuer detection uses a deliberately small, well-known set of
// prefix/length rules and returns [Unknown] when nothing matches.
//
// Author: Souvik (github.com/Souvik-Cyclic)
package card

import "strings"

// Network names returned by [Issuer].
const (
	Visa       = "Visa"
	Mastercard = "Mastercard"
	Amex       = "Amex"
	Discover   = "Discover"
	DinersClub = "Diners Club"
	RuPay      = "RuPay"
	Unknown    = "Unknown"
)

// Normalize removes spaces and dashes commonly used to group card digits,
// returning the bare digit string (callers should still validate it).
func Normalize(number string) string {
	return strings.NewReplacer(" ", "", "-", "").Replace(number)
}

// IsValid reports whether number is a well-formed card number per the Luhn
// checksum. Spaces and dashes are ignored. Any other non-digit rune, or a
// length below two digits, makes the number invalid.
func IsValid(number string) bool {
	n := Normalize(number)
	if len(n) < 2 {
		return false
	}

	sum := 0
	double := false
	for i := len(n) - 1; i >= 0; i-- { // process digits right-to-left
		c := n[i]
		if c < '0' || c > '9' {
			return false
		}
		d := int(c - '0')
		if double {
			if d *= 2; d > 9 {
				d -= 9
			}
		}
		sum += d
		double = !double
	}
	return sum%10 == 0
}

// Issuer returns the card network for number based on its prefix and length,
// or [Unknown] if no rule matches. It does not validate the checksum; combine
// with [IsValid] when you need both.
//
// The rules are simplified to the most common ranges:
//
//	Visa         prefix 4,                length 13, 16 or 19
//	Mastercard   prefix 51-55 or 2221-2720, length 16
//	Amex         prefix 34 or 37,         length 15
//	Discover     prefix 6011 or 644-649,  length 16
//	Diners Club  prefix 300-305, 36, 38,  length 14
//	RuPay        prefix 60, 65, 81, 82, 508, length 16
func Issuer(number string) string {
	n := Normalize(number)
	switch {
	case len(n) == 15 && hasPrefix(n, "34", "37"):
		return Amex
	case len(n) == 14 && (hasPrefix(n, "36", "38") || inRange(n, 3, "300", "305")):
		return DinersClub
	case (len(n) == 13 || len(n) == 16 || len(n) == 19) && hasPrefix(n, "4"):
		return Visa
	case len(n) == 16 && (inRange(n, 2, "51", "55") || inRange(n, 4, "2221", "2720")):
		return Mastercard
	case len(n) == 16 && (hasPrefix(n, "6011") || inRange(n, 3, "644", "649")):
		return Discover
	case len(n) == 16 && hasPrefix(n, "60", "65", "81", "82", "508"):
		return RuPay
	default:
		return Unknown
	}
}

// hasPrefix reports whether n starts with any of the given prefixes.
func hasPrefix(n string, prefixes ...string) bool {
	for _, p := range prefixes {
		if strings.HasPrefix(n, p) {
			return true
		}
	}
	return false
}

// inRange reports whether the first width digits of n fall within [lo, hi]
// inclusive, comparing them as fixed-width numeric strings.
func inRange(n string, width int, lo, hi string) bool {
	if len(n) < width {
		return false
	}
	head := n[:width]
	return head >= lo && head <= hi
}
