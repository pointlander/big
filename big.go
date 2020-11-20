// Copyright 2020 The Big Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package big

import (
	"math/big"

	"github.com/ALTree/bigfloat"
)

// Imaginary is an imaginary number
type Imaginary struct {
	a, b *big.Float
}

// NewImaginary creates a new imaginary number
func NewImaginary(a, b *big.Float) *Imaginary {
	return &Imaginary{
		a: a,
		b: b,
	}
}

// Abs computes the absolute value of a
func (i *Imaginary) Abs(a *Imaginary) *Imaginary {
	i.a.Mul(a.a, a.a)
	i.b.Mul(a.b, a.b)
	i.a = bigfloat.Sqrt(i.a.Add(i.a, i.b))
	i.b = big.NewFloat(0)
	return i
}

// Add add two imaginary numbers
func (i *Imaginary) Add(a, b *Imaginary) *Imaginary {
	i.a.Add(a.a, b.a)
	i.b.Add(a.b, b.b)
	return i
}

// Sub subtracts two imaginary numbers
func (i *Imaginary) Sub(a, b *Imaginary) *Imaginary {
	i.a.Sub(a.a, b.a)
	i.b.Sub(a.b, b.b)
	return i
}

// Mul multiples two imaginary numbers
func (i *Imaginary) Mul(a, b *Imaginary) *Imaginary {
	x1, x2, x3, x4 :=
		big.NewFloat(0).SetPrec(i.a.Prec()), big.NewFloat(0).SetPrec(i.b.Prec()),
		big.NewFloat(0).SetPrec(i.b.Prec()), big.NewFloat(0).SetPrec(i.a.Prec())
	x1.Mul(a.a, b.a) // a*a
	x2.Mul(a.a, b.b) // a*ib
	x3.Mul(a.b, b.a) // ib*a
	x4.Mul(a.b, b.b) // i^2 * b = -b
	i.a.Add(x1, x4.Neg(x4))
	i.b.Add(x2, x3)
	return i
}

// Conj computes the complex conjugate of a
func (i *Imaginary) Conj(a *Imaginary) *Imaginary {
	i.a.Set(a.a)
	i.b.Neg(a.b)
	return i
}

// Div divides two imaginary numbers
func (i *Imaginary) Div(a, b *Imaginary) *Imaginary {
	c := NewImaginary(big.NewFloat(0).SetPrec(i.a.Prec()), big.NewFloat(0).SetPrec(i.b.Prec()))
	c.Conj(b)
	x := NewImaginary(a.a.Copy(a.a), a.b.Copy(a.b))
	y := NewImaginary(b.a.Copy(b.a), b.b.Copy(b.b))
	x.Mul(x, c)
	y.Mul(y, c)
	i.a.Quo(x.a, y.a)
	i.b.Quo(x.b, y.a)
	return i
}

// String returns a string of the imaginary number
func (i *Imaginary) String() string {
	return i.a.String() + " + " + i.b.String() + "i"
}
