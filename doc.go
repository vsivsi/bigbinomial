/*
Package bigbinomial implements binomial distribution PMF and CDF functions with math/big support

This package was created to make it possible to calculate binomial distributions for larger
numbers of trials than are possible using IEEE floating point math.

However, as shown in the example below, this approach does not work for distributions with large
numbers of trials because the IEEE floating point can't represent large enough
values to complete the necessary calculations.
*/
package bigbinomial
