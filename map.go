package gowed

// Extract unexpired data in slice
func ExtracMapUnexpiredData[K comparable, V Expirer](slice map[K]V) []V {
	var result []V
	for _, item := range slice {
		if !item.IsExpired() {
			result = append(result, item)
		}
	}
	return result
}

// Extract expired data in slice
func ExtracMapExpiredData[K comparable, V Expirer](slice map[K]V) []V {
	var result []V
	for _, item := range slice {
		if item.IsExpired() {
			result = append(result, item)
		}
	}
	return result
}

// Remove expired data in slice
func RemoveMapExpiredData[K comparable, V Expirer](m map[K]V) {
	for k, v := range m {
		if v.IsExpired() {
			delete(m, k)
		}
	}
}

// Remove unexpired data in slice
func RemoveMapUnexpiredData[K comparable, V Expirer](m map[K]V) {
	for k, v := range m {
		if !v.IsExpired() {
			delete(m, k)
		}
	}
}
