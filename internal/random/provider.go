package random

import (
	"io"
	"math/big"
)

// Provider of random functions and functionality.
type Provider interface {
	io.Reader

	// BytesErr returns random data as bytes with the standard random.DefaultN length and can contain any byte values
	// (including unreadable byte values). If an error is returned from the random read this function returns it.
	BytesErr() (data []byte, err error)

	// Bytes returns random data as bytes with the standard random.DefaultN length and can contain any byte values
	// (including unreadable byte values). If an error is returned from the random read this function ignores it.
	Bytes() (data []byte)

	// BytesCustomErr returns random data as bytes with n length and can contain only byte values from the provided
	// values. If n is less than 1 then DefaultN is used instead. If an error is returned from the random read this function
	// returns it.
	BytesCustomErr(n int, charset []byte) (data []byte, err error)

	// StringCustomErr is an overload of BytesCustomWithErr which takes a characters string and returns a string.
	StringCustomErr(n int, characters string) (data string, err error)

	// BytesCustom returns random data as bytes with n length and can contain only byte values from the provided
	// values. If n is less than 1 then DefaultN is used instead.
	BytesCustom(n int, charset []byte) (data []byte)

	// StringCustom is an overload of GenerateCustom which takes a characters string and returns a string.
	StringCustom(n int, characters string) (data string)

	// IntErr returns a random *big.Int error combination with a maximum of max.
	IntErr(max *big.Int) (value *big.Int, err error)

	// Int returns a random *big.Int with a maximum of max.
	Int(max *big.Int) (value *big.Int)

	// IntegerErr returns a random int error combination with a maximum of n.
	IntegerErr(n int) (value int, err error)

	// Integer returns a random integer with a maximum of n.
	Integer(n int) (value int)
}
