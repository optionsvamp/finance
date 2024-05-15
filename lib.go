package finance

import (
	"math"
)

type OptionType int

const (
	Call OptionType = iota
	Put
)

// Option represents an option contract
type Option struct {
	Price            float64    // Option price
	Strike           float64    // Option strike price
	DaysToExpiration float64    // Days to expiration
	RiskFreeRate     float64    // Risk-free interest rate
	UnderlyingPrice  float64    // Current price of the underlying asset
	OptionType       OptionType // Option type, can be either Call or Put
}

// BlackScholesImpliedVolatility computes implied volatility using the Newton-Raphson method
func BlackScholesImpliedVolatility(option Option) float64 {
	targetPrice := option.Price
	currentVolatility := 0.2 // Initial guess for volatility
	epsilon := 0.0001        // Tolerance for convergence
	maxIterations := 100     // Maximum number of iterations
	vega := 0.0
	for i := 0; i < maxIterations; i++ {
		price := BlackScholesOptionPrice(option, currentVolatility)
		vega = BlackScholesVega(option, currentVolatility)
		if math.Abs(price-targetPrice) < epsilon {
			break // Convergence achieved
		}
		// Update volatility using Newton-Raphson iteration
		currentVolatility -= (price - targetPrice) / vega
	}
	return currentVolatility
}

// BlackScholesOptionPrice calculates the Black-Scholes option price
// underlyingAssetPrice: the underlying asset price
// strikePrice: the strike price
// timeToExpirationInDays: the time to expiration in days
// volatility: the volatility
// riskFreeInterestRate: the risk-free interest rate
// optionType: the type of the option ("call" or "put")
func BlackScholesOptionPrice(option Option, volatility float64) float64 {
	timeToExpiration := option.DaysToExpiration / 365.0 // convert days to years
	d1 := (math.Log(option.UnderlyingPrice/option.Strike) + (option.RiskFreeRate+0.5*math.Pow(volatility, 2))*timeToExpiration) / (volatility * math.Sqrt(timeToExpiration))
	d2 := d1 - volatility*math.Sqrt(timeToExpiration)
	if option.OptionType == Call {
		return option.UnderlyingPrice*Phi(d1) - option.Strike*math.Exp(-option.RiskFreeRate*timeToExpiration)*Phi(d2)
	}
	return option.Strike*math.Exp(-option.RiskFreeRate*timeToExpiration)*Phi(-d2) - option.UnderlyingPrice*Phi(-d1)
}

// Phi calculates the cumulative distribution function of the standard normal distribution
func Phi(x float64) float64 {
	return 0.5 * (1 + math.Erf(x/math.Sqrt2))
}

// BlackScholesVega calculates the vega of Black-Scholes option price
// option: the option
// volatility: the volatility
func BlackScholesVega(option Option, volatility float64) float64 {
	timeToExpiration := option.DaysToExpiration / 365.0
	d1 := (math.Log(option.UnderlyingPrice/option.Strike) + (option.RiskFreeRate+0.5*math.Pow(volatility, 2))*timeToExpiration) / (volatility * math.Sqrt(timeToExpiration))
	return option.UnderlyingPrice * math.Sqrt(timeToExpiration) * math.Exp(-0.5*d1*d1) / math.Sqrt(2*math.Pi)
}

// BlackScholesGamma computes the gamma of an option
// option: the option
func BlackScholesGamma(option Option, vol float64) float64 {
	d1 := (math.Log(option.UnderlyingPrice/option.Strike) + (option.RiskFreeRate+0.5*math.Pow(vol, 2))*(option.DaysToExpiration/365.0)) / (vol * math.Sqrt(option.DaysToExpiration/365.0))
	return NormalDistributionDerivative(d1) / (option.UnderlyingPrice * vol * math.Sqrt(option.DaysToExpiration/365.0))
}

// NormalDistributionDerivative calculates the derivative of the standard normal cumulative distribution function
// x: the input value
func NormalDistributionDerivative(x float64) float64 {
	return math.Exp(-0.5*math.Pow(x, 2)) / math.Sqrt(2*math.Pi)
}

// BlackScholesDelta computes the delta of an option
// option: the option
// volatility: the volatility
func BlackScholesDelta(option Option, volatility float64) float64 {
	timeToExpiration := option.DaysToExpiration / 365.0
	d1 := (math.Log(option.UnderlyingPrice/option.Strike) + (option.RiskFreeRate+volatility*volatility/2)*timeToExpiration) / (volatility * math.Sqrt(timeToExpiration))

	if option.OptionType == Call {
		return Phi(d1)
	} else {
		return Phi(d1) - 1
	}
}
