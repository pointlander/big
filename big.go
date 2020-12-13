// Copyright 2020 The Big Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package big

import (
	"math/big"

	"github.com/ALTree/bigfloat"
)

// Float is an imaginary number
type Float struct {
	a, b *big.Float
}

// NewFloat creates a new imaginary number
func NewFloat(a, b *big.Float) *Float {
	return &Float{
		a: a,
		b: b,
	}
}

// Abs computes the absolute value of a
func (f *Float) Abs(a *Float) *Float {
	f.a.Mul(a.a, a.a)
	f.b.Mul(a.b, a.b)
	f.a = bigfloat.Sqrt(f.a.Add(f.a, f.b))
	f.b = big.NewFloat(0)
	return f
}

// Add add two imaginary numbers
func (f *Float) Add(a, b *Float) *Float {
	f.a.Add(a.a, b.a)
	f.b.Add(a.b, b.b)
	return f
}

// Sub subtracts two imaginary numbers
func (f *Float) Sub(a, b *Float) *Float {
	f.a.Sub(a.a, b.a)
	f.b.Sub(a.b, b.b)
	return f
}

// Mul multiples two imaginary numbers
func (f *Float) Mul(a, b *Float) *Float {
	x1, x2, x3, x4 :=
		big.NewFloat(0).SetPrec(f.a.Prec()), big.NewFloat(0).SetPrec(f.b.Prec()),
		big.NewFloat(0).SetPrec(f.b.Prec()), big.NewFloat(0).SetPrec(f.a.Prec())
	x1.Mul(a.a, b.a) // a*a
	x2.Mul(a.a, b.b) // a*ib
	x3.Mul(a.b, b.a) // ib*a
	x4.Mul(a.b, b.b) // i^2 * b = -b
	f.a.Add(x1, x4.Neg(x4))
	f.b.Add(x2, x3)
	return f
}

// Conj computes the complex conjugate of a
func (f *Float) Conj(a *Float) *Float {
	f.a.Set(a.a)
	f.b.Neg(a.b)
	return f
}

// Div divides two imaginary numbers
func (f *Float) Div(a, b *Float) *Float {
	c := NewFloat(big.NewFloat(0).SetPrec(f.a.Prec()), big.NewFloat(0).SetPrec(f.b.Prec()))
	c.Conj(b)
	x := NewFloat(a.a.Copy(a.a), a.b.Copy(a.b))
	y := NewFloat(b.a.Copy(b.a), b.b.Copy(b.b))
	x.Mul(x, c)
	y.Mul(y, c)
	f.a.Quo(x.a, y.a)
	f.b.Quo(x.b, y.a)
	return f
}

// Sqrt computes the square root of the complex number
// https://www.johndcook.com/blog/2020/06/09/complex-square-root/
func (f *Float) Sqrt(a *Float) *Float {
	x := big.NewFloat(0).SetPrec(f.a.Prec())
	y := big.NewFloat(0).SetPrec(f.b.Prec())
	l := big.NewFloat(0).SetPrec(f.a.Prec())
	x.Mul(a.a, a.a)
	y.Mul(a.b, a.b)
	l.Add(x, y)
	l = bigfloat.Sqrt(l)

	aa := big.NewFloat(0).SetPrec(f.a.Prec())
	aa.Add(l, a.a)
	aa.Quo(aa, big.NewFloat(2).SetPrec(f.a.Prec()))
	aa = bigfloat.Sqrt(aa)

	f.b.Sub(l, a.a)
	f.b.Quo(f.b, big.NewFloat(2).SetPrec(f.b.Prec()))
	f.b = bigfloat.Sqrt(f.b)
	f.b.Mul(big.NewFloat(float64(a.b.Sign())).SetPrec(f.b.Prec()), f.b)
	f.a = aa

	return f
}

// Atan2 computes atan2 of x
// https://en.wikipedia.org/wiki/Atan2
func (f *Float) Atan2(x *Float) *Float {
	a := x.a
	b := x.b

	if a.Cmp(big.NewFloat(0).SetPrec(a.Prec())) > 0 {
		f.Arg(x)
	} else if a.Cmp(big.NewFloat(0).SetPrec(a.Prec())) < 0 &&
		b.Cmp(big.NewFloat(0).SetPrec(b.Prec())) >= 0 {
		f.Arg(x)
		f.a.Add(f.a, bigfloat.PI(a.Prec()))
	} else if a.Cmp(big.NewFloat(0).SetPrec(a.Prec())) < 0 &&
		b.Cmp(big.NewFloat(0).SetPrec(b.Prec())) < 0 {
		f.Arg(x)
		f.a.Sub(f.a, bigfloat.PI(f.a.Prec()))
	} else {
		f.Arg(x)
	}

	return f
}

// Arg computes arg(x + yi) = tan-1(y/x)
// https://mathworld.wolfram.com/ComplexArgument.html
func (f *Float) Arg(x *Float) *Float {
	a := x.a
	b := x.b
	f.b = big.NewFloat(0).SetPrec(b.Prec())

	if a.Cmp(big.NewFloat(0).SetPrec(a.Prec())) == 0 {
		if b.Cmp(big.NewFloat(0).SetPrec(a.Prec())) < 0 {
			f.a.Set(bigfloat.PI(a.Prec()))
			f.a.Quo(f.a, big.NewFloat(2).SetPrec(a.Prec()))
			f.a.Neg(f.a)
		} else if b.Cmp(big.NewFloat(0).SetPrec(a.Prec())) == 0 {
			f.a.SetInf(false)
		} else {
			f.a.Set(bigfloat.PI(a.Prec()))
			f.a.Quo(f.a, big.NewFloat(2).SetPrec(a.Prec()))
		}

		return f
	}

	if a.Cmp(big.NewFloat(1).SetPrec(a.Prec())) == 0 &&
		b.Cmp(big.NewFloat(0).SetPrec(b.Prec())) == 0 {
		f.a.Set(big.NewFloat(0).SetPrec(a.Prec()))
	} else if a.Cmp(big.NewFloat(1).SetPrec(a.Prec())) == 0 &&
		b.Cmp(big.NewFloat(1).SetPrec(b.Prec())) == 0 {
		f.a.Set(bigfloat.PI(a.Prec()))
		f.a.Quo(f.a, big.NewFloat(4).SetPrec(a.Prec()))
	} else if a.Cmp(big.NewFloat(0).SetPrec(a.Prec())) == 0 &&
		b.Cmp(big.NewFloat(1).SetPrec(b.Prec())) == 0 {
		f.a.Set(bigfloat.PI(b.Prec()))
		f.a.Quo(f.a, big.NewFloat(2).SetPrec(b.Prec()))
	} else if a.Cmp(big.NewFloat(-1).SetPrec(a.Prec())) == 0 &&
		b.Cmp(big.NewFloat(0).SetPrec(b.Prec())) == 0 {
		f.a.Set(bigfloat.PI(a.Prec()))
	} else if a.Cmp(big.NewFloat(0).SetPrec(a.Prec())) == 0 &&
		b.Cmp(big.NewFloat(-1).SetPrec(b.Prec())) == 0 {
		f.a.Set(bigfloat.PI(b.Prec()))
		f.a.Quo(f.a, big.NewFloat(2).SetPrec(b.Prec()))
		f.a.Neg(f.a)
	} else {
		f.a.Quo(b, a)
		f.a = bigfloat.Arctan(f.a)
	}

	return f
}

// Exp computes e^x for a complex number
// https://www.wolframalpha.com/input/?i=e%5E%28x+%2B+yi%29
func (f *Float) Exp(x *Float) *Float {
	exp := bigfloat.Exp(x.a)
	cos := bigfloat.Cos(x.b)
	sin := bigfloat.Sin(x.b)
	f.a.Mul(exp, cos)
	f.b.Mul(exp, sin)
	return f
}

// Log computes the natural log of x
// https://en.wikipedia.org/wiki/Complex_logarithm
func (f *Float) Log(x *Float) *Float {
	a := x.a
	aa := big.NewFloat(0).SetPrec(a.Prec())
	aa.Mul(a, a)
	b := x.b
	bb := big.NewFloat(0).SetPrec(b.Prec())
	bb.Mul(b, b)
	real := big.NewFloat(0).SetPrec(a.Prec())
	real.Add(aa, bb)
	real = bigfloat.Log(bigfloat.Sqrt(real))
	y := NewFloat(big.NewFloat(0).SetPrec(a.Prec()), big.NewFloat(0).SetPrec(b.Prec()))
	y.Atan2(x)
	f.a = real
	f.b = y.a
	return f
}

// Pow computes x**y
// https://mathworld.wolfram.com/ComplexExponentiation.html
func (f *Float) Pow(x *Float, y *Float) *Float {
	if x.a.Cmp(big.NewFloat(0).SetPrec(x.a.Prec())) == 0 &&
		x.b.Cmp(big.NewFloat(0).SetPrec(x.b.Prec())) == 0 &&
		y.a.Cmp(big.NewFloat(0).SetPrec(y.a.Prec())) == 0 &&
		y.b.Cmp(big.NewFloat(0).SetPrec(y.b.Prec())) == 0 {
		f.a.SetInf(false)
		return f
	}

	a := big.NewFloat(0).SetPrec(x.a.Prec())
	a.Set(x.a)
	aa := big.NewFloat(0).SetPrec(x.a.Prec())
	aa.Mul(a, a)
	b := big.NewFloat(0).SetPrec(x.b.Prec())
	b.Set(x.b)
	bb := big.NewFloat(0).SetPrec(x.b.Prec())
	bb.Mul(b, b)
	sum := big.NewFloat(0).SetPrec(x.a.Prec())
	sum.Add(aa, bb)
	c := big.NewFloat(0).SetPrec(y.a.Prec())
	c.Set(y.a)
	cc := big.NewFloat(0).SetPrec(c.Prec())
	cc.Quo(c, big.NewFloat(2).SetPrec(c.Prec()))
	e := bigfloat.Pow(sum, cc)
	d := big.NewFloat(0).SetPrec(y.b.Prec())
	d.Set(y.b)
	arg := NewFloat(big.NewFloat(0).SetPrec(a.Prec()),
		big.NewFloat(0).SetPrec(b.Prec()))
	arg.Arg(x)
	exp := big.NewFloat(0).SetPrec(arg.a.Prec())
	exp.Mul(d, arg.a)
	exp.Neg(exp)
	exp = bigfloat.Exp(exp)
	e.Mul(e, exp)

	i := big.NewFloat(0).SetPrec(c.Prec())
	i.Mul(c, arg.a)
	j := big.NewFloat(0).SetPrec(d.Prec())
	j.Mul(d, bigfloat.Log(sum))
	j.Quo(j, big.NewFloat(2).SetPrec(d.Prec()))
	i.Add(i, j)
	cos := bigfloat.Cos(i)
	cos.Mul(e, cos)
	f.a.Set(cos)
	sin := bigfloat.Sin(i)
	sin.Mul(e, sin)
	f.b.Set(sin)
	return f
}

// String returns a string of the imaginary number
func (f *Float) String() string {
	return f.a.String() + " + " + f.b.String() + "i"
}

// Rational is an imaginary number
type Rational struct {
	a, b *big.Rat
}

// NewRational creates a new imaginary number
func NewRational(a, b *big.Rat) *Rational {
	return &Rational{
		a: a,
		b: b,
	}
}

// Add add two imaginary numbers
func (r *Rational) Add(a, b *Rational) *Rational {
	r.a.Add(a.a, b.a)
	r.b.Add(a.b, b.b)
	return r
}

// Sub subtracts two imaginary numbers
func (r *Rational) Sub(a, b *Rational) *Rational {
	r.a.Sub(a.a, b.a)
	r.b.Sub(a.b, b.b)
	return r
}

// Mul multiples two imaginary numbers
func (r *Rational) Mul(a, b *Rational) *Rational {
	x1, x2, x3, x4 :=
		big.NewRat(0, 1), big.NewRat(0, 1),
		big.NewRat(0, 1), big.NewRat(0, 1)
	x1.Mul(a.a, b.a) // a*a
	x2.Mul(a.a, b.b) // a*ib
	x3.Mul(a.b, b.a) // ib*a
	x4.Mul(a.b, b.b) // i^2 * b = -b
	r.a.Add(x1, x4.Neg(x4))
	r.b.Add(x2, x3)
	return r
}

// Conj computes the complex conjugate of a
func (r *Rational) Conj(a *Rational) *Rational {
	r.a.Set(a.a)
	r.b.Neg(a.b)
	return r
}

// Div divides two imaginary numbers
func (r *Rational) Div(a, b *Rational) *Rational {
	c := NewRational(big.NewRat(0, 1), big.NewRat(0, 1))
	c.Conj(b)
	x := NewRational(a.a.Set(a.a), a.b.Set(a.b))
	y := NewRational(b.a.Set(b.a), b.b.Set(b.b))
	x.Mul(x, c)
	y.Mul(y, c)
	r.a.Quo(x.a, y.a)
	r.b.Quo(x.b, y.a)
	return r
}

// String returns a string of the imaginary number
func (r *Rational) String() string {
	return r.a.String() + " + " + r.b.String() + "i"
}
