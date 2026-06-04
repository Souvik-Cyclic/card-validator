package card

import "testing"

func TestIsValid(t *testing.T) {
	cases := []struct {
		number string
		want   bool
	}{
		{"4111111111111111", true},    // Visa
		{"5500005555555559", true},    // Mastercard
		{"340000000000009", true},     // Amex
		{"4111 1111 1111 1111", true}, // spaces tolerated
		{"4111-1111-1111-1111", true}, // dashes tolerated
		{"4111111111111112", false},   // bad checksum
		{"1234567812345678", false},   // bad checksum
		{"4111111111111111a", false},  // trailing non-digit
		{"abcd", false},               // non-digit
		{"4", false},                  // too short
		{"", false},                   // empty
	}
	for _, c := range cases {
		if got := IsValid(c.number); got != c.want {
			t.Errorf("IsValid(%q) = %v, want %v", c.number, got, c.want)
		}
	}
}

func TestIssuer(t *testing.T) {
	cases := []struct {
		number string
		want   string
	}{
		{"4111111111111111", Visa},
		{"4012888888881881", Visa},
		{"5500005555555559", Mastercard},
		{"2221000000000009", Mastercard}, // 2-series range
		{"340000000000009", Amex},
		{"370000000000002", Amex},
		{"6011000000000004", Discover},
		{"6440000000000000", Discover}, // 644-649 range
		{"36000000000008", DinersClub},
		{"30000000000004", DinersClub}, // 300-305 range
		{"6521000000000000", RuPay},
		{"8100000000000000", RuPay},
		{"9999999999999999", Unknown},
		{"", Unknown},
	}
	for _, c := range cases {
		if got := Issuer(c.number); got != c.want {
			t.Errorf("Issuer(%q) = %q, want %q", c.number, got, c.want)
		}
	}
}

func TestNormalize(t *testing.T) {
	if got := Normalize("4111 1111-1111 1111"); got != "4111111111111111" {
		t.Errorf("Normalize stripped wrong: %q", got)
	}
}

func TestLengths(t *testing.T) {
	if got := Lengths(Amex); len(got) != 1 || got[0] != 15 {
		t.Errorf("Lengths(Amex) = %v", got)
	}
	if Lengths("Nope") != nil {
		t.Errorf("unknown issuer should return nil")
	}
}
