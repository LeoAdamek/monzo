package monzo

import "math/big"

// Money represents a monetary value. It is an integer of pennies.
type Money int64

// Rat converts a monetary value to a rational number with a denominator of 100
func (m Money) Rat() *big.Rat {
	return big.NewRat(int64(m), 100)
}

// String returns the value as a formatted string.
func (m Money) String() string {
	return m.Rat().FloatString(2)
}
