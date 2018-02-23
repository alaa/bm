package main

func remove(elem string, slice []string) []string {
	var new []string
	for _, str := range slice {
		if str != elem {
			new = append(new, str)
		}
	}
	return new
}

func contains(elem string, slice []string) bool {
	for _, e := range slice {
		if e == elem {
			return true
		}
	}
	return false
}

func append_if_uniq(elem string, slice []string) []string {
	if contains(elem, slice) {
		return slice
	}
	return append(slice, elem)
}

func uniqueNonEmpty(slice []string) []string {
	unique := make(map[string]bool, len(slice))
	unique_slice := make([]string, len(unique))
	for _, elem := range slice {
		if len(elem) != 0 {
			if !unique[elem] {
				unique_slice = append(unique_slice, elem)
				unique[elem] = true
			}
		}
	}
	return unique_slice
}
