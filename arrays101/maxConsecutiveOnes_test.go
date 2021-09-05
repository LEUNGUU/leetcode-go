package arrays101

import "testing"

func TestNormalNums(t *testing.T) {
	nums1 := []int{0, 1, 0, 1, 1, 1, 0, 1}
	want := 3
	res := findMaxConsecutiveOnes(nums1)
	if res != want {
		t.Fatalf("findMaxConsecutiveOnes error, want: %v, but return: %v", want, res)
	}
}

func TestArrayWithOneElement(t *testing.T) {
	nums1 := []int{0}
	nums2 := []int{1}
	want1 := 0
	want2 := 1
	res1 := findMaxConsecutiveOnes(nums1)
	res2 := findMaxConsecutiveOnes(nums2)
	if res1 != want1 {
		t.Fatalf("findMaxConsecutiveOnes error, want: %v, but return: %v", want1, res1)
	}
	if res2 != want2 {
		t.Fatalf("findMaxConsecutiveOnes error, want: %v, but return: %v", want2, res2)
	}
}

func TestArrayWithNonElement(t *testing.T) {
	nums := []int{}
	want := 0
	res := findMaxConsecutiveOnes(nums)
	if res != want {
		t.Fatalf("findMaxConsecutiveOnes error, want: %v, but return: %v", want, res)
	}
}
