package sorting

/*
Bubblesort loops through the array and compares two elements and swaps them
if they are not on order. If no swapping occured, the array is sorted.
average: O(n^2)
*/
func Bubblesort(arr []int) []int {
	var swapped bool
	for {
		swapped = false
		for i := 1; i <= len(arr)-1; i++ {
			if arr[i-1] > arr[i] {
				swapped = true
				arr[i], arr[i-1] = arr[i-1], arr[i]
			}
		}
		if !swapped {
			return arr
		}
	}
}
