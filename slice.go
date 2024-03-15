package gowed

// Extract unexpired data in slice
func ExtracSliceUnexpiredData[T Expirer](slice []T) []T {
	var result []T
	for _, item := range slice {
		if !item.IsExpired() {
			result = append(result, item)
		}
	}
	return result
}

// Extract expired data in slice
func ExtracSliceExpiredData[T Expirer](slice []T) []T {
	var result []T
	for _, item := range slice {
		if item.IsExpired() {
			result = append(result, item)
		}
	}
	return result
}
