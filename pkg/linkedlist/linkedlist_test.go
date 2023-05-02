package linkedlist

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLinkedList_Push(t *testing.T) {
	tcs := map[string]struct {
		in        []int
		expResult LinkedList[int]
	}{
		"push one": {
			in: []int{1},
			expResult: LinkedList[int]{
				Root: &Node[int]{Value: 1},
			},
		},
		"push two": {
			in: []int{2, 3},
			expResult: LinkedList[int]{
				Root: &Node[int]{
					Value: 2,
					Next:  &Node[int]{Value: 3},
				},
			},
		},
	}

	for desc, tc := range tcs {
		t.Run(desc, func(t *testing.T) {
			// Given
			ll := LinkedList[int]{}

			// When
			for _, v := range tc.in {
				ll.Push(v)
			}

			// Then
			require.Equal(t, tc.expResult, ll)
		})
	}
}

func TestLinkedList_Find(t *testing.T) {
	type user struct {
		ID   int
		Name string
	}
	tcs := map[string]struct {
		inID      int
		ll        LinkedList[user]
		expResult *user
	}{
		"found": {
			inID: 1,
			ll: LinkedList[user]{
				Root: &Node[user]{
					Value: user{
						ID:   1,
						Name: "Vae",
					},
					Next: &Node[user]{
						Value: user{
							ID:   2,
							Name: "Virsavik",
						},
					},
				},
			},
			expResult: &user{
				ID:   1,
				Name: "Vae",
			},
		},
		"not found": {
			inID: 10,
			ll: LinkedList[user]{
				Root: &Node[user]{
					Value: user{
						ID:   1,
						Name: "Vae",
					},
					Next: &Node[user]{
						Value: user{
							ID:   2,
							Name: "Virsavik",
						},
					},
				},
			},
		},
	}
	for desc, tc := range tcs {
		t.Run(desc, func(t *testing.T) {
			// Given
			ll := tc.ll

			// When
			theUser := ll.Find(func(value user) bool {
				return value.ID == tc.inID
			})

			// Then
			if tc.expResult != nil {
				require.Equal(t, *tc.expResult, *theUser)
			} else {
				require.Nil(t, theUser)
			}
		})
	}
}

func TestLinkedList_Filter(t *testing.T) {
	type product struct {
		ID    int
		Price float64
	}
	tcs := map[string]struct {
		inPrice   float64
		ll        LinkedList[product]
		expResult LinkedList[product]
	}{
		"got result": {
			inPrice: 100,
			ll: LinkedList[product]{
				Root: &Node[product]{
					Value: product{
						ID:    1,
						Price: 100,
					},
					Next: &Node[product]{
						Value: product{
							ID:    2,
							Price: 200,
						},
						Next: &Node[product]{
							Value: product{
								ID:    3,
								Price: 100,
							},
						},
					},
				},
			},
			expResult: LinkedList[product]{
				Root: &Node[product]{
					Value: product{
						ID:    1,
						Price: 100,
					},
					Next: &Node[product]{
						Value: product{
							ID:    3,
							Price: 100,
						},
					},
				},
			},
		},
		"no result": {
			inPrice: 300,
			ll: LinkedList[product]{
				Root: &Node[product]{
					Value: product{
						ID:    1,
						Price: 100,
					},
					Next: &Node[product]{
						Value: product{
							ID:    2,
							Price: 200,
						},
						Next: &Node[product]{
							Value: product{
								ID:    3,
								Price: 100,
							},
						},
					},
				},
			},
			expResult: LinkedList[product]{},
		},
	}
	for desc, tc := range tcs {
		t.Run(desc, func(t *testing.T) {
			// Given
			ll := tc.ll

			// When
			result := ll.Filter(func(value product) bool {
				return value.Price == tc.inPrice
			})

			// Then
			require.Equal(t, tc.expResult, result)
		})
	}
}
