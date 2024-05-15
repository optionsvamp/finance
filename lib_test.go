package finance

import (
	"math"
	"testing"
)

func TestBlackScholesOptionPrice(t *testing.T) {
	option := Option{
		Price:            100.0,
		Strike:           100.0,
		DaysToExpiration: 30.0,
		RiskFreeRate:     0.05,
		UnderlyingPrice:  100.0,
		OptionType:       Call,
	}

	priceCall := BlackScholesOptionPrice(option, 0.2)

	if priceCall <= 0 {
		t.Errorf("Invalid price: got %v, expected a value greater than 0", priceCall)
	}

	const tolerance = 0.00001

	const expectedPriceCall = 2.493376
	if diff := math.Abs(priceCall - expectedPriceCall); diff > tolerance {
		t.Errorf("Unexpected price for call option: got %v, want %v", priceCall, expectedPriceCall)
	}

	option.OptionType = Put
	pricePut := BlackScholesOptionPrice(option, 0.2)

	const expectedPricePut = 2.08326
	if diff := math.Abs(pricePut - expectedPricePut); diff > tolerance {
		t.Errorf("Unexpected price for put option: got %v, want %v", pricePut, expectedPricePut)
	}
}

func TestBlackScholesImpliedVolatility(t *testing.T) {
	option := Option{
		Price:            10.0,
		Strike:           100.0,
		DaysToExpiration: 30.0,
		RiskFreeRate:     0.05,
		UnderlyingPrice:  100.0,
		OptionType:       Call,
	}

	volatility := BlackScholesImpliedVolatility(option)

	if !math.IsNaN(volatility) && volatility < 0 || volatility > 1 {
		t.Errorf("Invalid volatility: got %v, expected a value between 0 and 1", volatility)
	}

	const expectedIVCall = 0.86021805
	const tolerance = 0.00001
	if diff := math.Abs(volatility - expectedIVCall); diff > tolerance {
		t.Errorf("Unexpected volatility for call option: got %v, want %v", volatility, expectedIVCall)
	}
}

func TestBlackScholesGamma(t *testing.T) {
	option := Option{
		Price:            10.0,
		Strike:           100.0,
		DaysToExpiration: 30.0,
		RiskFreeRate:     0.05,
		UnderlyingPrice:  100.0,
		OptionType:       Call,
	}

	gamma := BlackScholesGamma(option, 0.2)

	if gamma < 0 {
		t.Errorf("Invalid gamma: got %v, expected a value greater than 0", gamma)
	}

	const expectedGammaCall = 0.0692276
	const tolerance = 0.00001
	if diff := math.Abs(gamma - expectedGammaCall); diff > tolerance {
		t.Errorf("Unexpected gamma for call option: got %v, want %v", gamma, expectedGammaCall)
	}
}

func TestBlackScholesVega(t *testing.T) {
	option := Option{
		Price:            10.0,
		Strike:           100.0,
		DaysToExpiration: 30.0,
		RiskFreeRate:     0.05,
		UnderlyingPrice:  100.0,
		OptionType:       Call,
	}

	vega := BlackScholesVega(option, 0.2)

	if vega <= 0 {
		t.Errorf("Invalid vega: got %v, expected a value greater than 0", vega)
	}

	const expectedVegaCall = 11.37988
	const tolerance = 0.00001
	if diff := math.Abs(vega - expectedVegaCall); diff > tolerance {
		t.Errorf("Unexpected delta for call option: got %v, want %v", vega, expectedVegaCall)
	}
}

func TestBlackScholesDelta(t *testing.T) {
	option := Option{
		Price:            10.0,
		Strike:           100.0,
		DaysToExpiration: 30.0,
		RiskFreeRate:     0.05,
		UnderlyingPrice:  100.0,
		OptionType:       Call,
	}

	deltaCall := BlackScholesDelta(option, 0.2)

	if deltaCall < -1 || deltaCall > 1 {
		t.Errorf("Invalid delta for call option: got %v, expected a value between -1 and 1", deltaCall)
	}

	const tolerance = 0.00001

	const expectedDeltaCall = 0.53996
	if diff := math.Abs(deltaCall - expectedDeltaCall); diff > tolerance {
		t.Errorf("Unexpected delta for call option: got %v, want %v", deltaCall, expectedDeltaCall)
	}

	option.OptionType = Put
	deltaPut := BlackScholesDelta(option, 0.2)

	if deltaPut < -1 || deltaPut > 1 {
		t.Errorf("Invalid delta for put option: got %v, expected a value between -1 and 1", deltaPut)
	}

	const expectedDeltaPut = -0.46003645
	if diff := math.Abs(deltaPut - expectedDeltaPut); diff > tolerance {
		t.Errorf("Unexpected delta for put option: got %v, want %v", deltaPut, expectedDeltaPut)
	}
}
