package besttimetobuyandsellstockii

func maxProfit(prices []int) int {
	// find all local max and min
	// even indices are mins, odds are max
	var maxMin []int
	for i := range prices {
		// mins
		if i < len(prices)-1 { // don't care about mins on last day
			if prices[i] >= prices[i+1] && (i == 0 || prices[i] >= prices[i-1]) {
				maxMin = append(maxMin, prices[i])
			}
		}

		// maxs
		if i > 0 { // don't care about max on the first day

		}
	}
	// goto each max that has a min before it, subtract the min from the max

	// ???

	// profit
	return 0
}
