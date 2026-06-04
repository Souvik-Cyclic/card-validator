# Card Validator

A small, zero-dependency Go library for payment card numbers. Validates the number with
the Luhn algorithm and detects the issuing network from its prefix and length.

```go
card.IsValid("4111 1111 1111 1111") // true
card.Issuer("4111111111111111")     // "Visa"
card.Issuer("6521000000000000")     // "RuPay"
```

## API
- `IsValid(number string) bool` — true if the number passes the Luhn checksum (spaces/dashes ignored).
- `Issuer(number string) string` — Visa, Mastercard, Amex, Discover, Diners Club, RuPay, or Unknown.

## Test
```bash
go test ./...
```
