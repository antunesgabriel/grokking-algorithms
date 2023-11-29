package bfs

import (
	"testing"
)

func TestFindKeychain(t *testing.T) {
	t.Run("should return root node if it is keychain", func(t *testing.T) {
		keychain := NewEstablishment("zechave", true)

		graph := Graph{
			keychain.Name: GraphItem{
				Value:     keychain,
				Neighbors: []string{},
			},
		}

		result := FindKeychain(graph, keychain.Name)

		if result != keychain {
			t.Errorf("expect %+v - got %+v", keychain, result)
		}
	})

	t.Run("should return nil if not exist keychain in graph", func(t *testing.T) {
		iceCream := NewEstablishment("icecream", false)
		pizzaria := NewEstablishment("pizzaria", false)
		bakery := NewEstablishment("bakery", false)
		coffeeshop := NewEstablishment("coffeeshop", false)
		lanhouse := NewEstablishment("lanhouse", false)

		establishments := Graph{
			iceCream.Name: GraphItem{
				Value: iceCream,
				Neighbors: []string{
					pizzaria.Name,
					bakery.Name,
				},
			},
			bakery.Name: GraphItem{
				Value: bakery,
				Neighbors: []string{
					pizzaria.Name,
					coffeeshop.Name,
				},
			},
			pizzaria.Name: GraphItem{
				Value:     pizzaria,
				Neighbors: []string{},
			},
			coffeeshop.Name: GraphItem{
				Value: coffeeshop,
				Neighbors: []string{
					lanhouse.Name,
				},
			},
			lanhouse.Name: GraphItem{
				Value:     lanhouse,
				Neighbors: []string{},
			},
		}

		result := FindKeychain(establishments, iceCream.Name)

		if result != nil {
			t.Errorf("expect %v - got %+v", nil, result)
		}
	})

	t.Run("should return keychain in graph", func(t *testing.T) {
		iceCream := NewEstablishment("icecream", false)
		pizzaria := NewEstablishment("pizzaria", false)
		bakery := NewEstablishment("bakery", false)
		coffeeshop := NewEstablishment("coffeeshop", false)
		lanhouse := NewEstablishment("lanhouse", false)
		keychain := NewEstablishment("itskeychain", true)

		establishments := Graph{
			iceCream.Name: GraphItem{
				Value: iceCream,
				Neighbors: []string{
					pizzaria.Name,
					bakery.Name,
				},
			},
			bakery.Name: GraphItem{
				Value: bakery,
				Neighbors: []string{
					pizzaria.Name,
					coffeeshop.Name,
				},
			},
			pizzaria.Name: GraphItem{
				Value: pizzaria,
				Neighbors: []string{
					keychain.Name,
				},
			},
			coffeeshop.Name: GraphItem{
				Value: coffeeshop,
				Neighbors: []string{
					lanhouse.Name,
				},
			},
			lanhouse.Name: GraphItem{
				Value:     lanhouse,
				Neighbors: []string{},
			},
			keychain.Name: GraphItem{
				Value:     keychain,
				Neighbors: []string{},
			},
		}

		result := FindKeychain(establishments, iceCream.Name)

		if result != keychain {
			t.Errorf("expect %+v - got %+v", keychain, result)
		}
	})
}
