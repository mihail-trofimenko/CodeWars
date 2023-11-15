package main

func FindUniq(arr []float32) float32 {
	if arr[0] == arr[1] && arr[1] == arr[2] {
		for _, r := range arr {
			if r != arr[0] {
				return r
			}
		}
	}
	if arr[1] != arr[0] && arr[1] != arr[2] {
		return arr[1]
	}
	if arr[2] != arr[0] && arr[1] != arr[2] {
		return arr[2]
	}

	// Do the magic
	return arr[0]
}
