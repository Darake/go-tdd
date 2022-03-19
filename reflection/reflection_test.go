package reflection

import (
	"reflect"
	"testing"
)

type Person struct {
	Name    string
	Profile Profile
}

type Profile struct {
	Age  int
	City string
}

func TestWalk(t *testing.T) {
	cases := []struct {
		Name          string
		Input         interface{}
		ExpectedCalls []string
	}{
		{
			Name: "Struct with one string field",
			Input: struct {
				Name string
			}{"Gert"},
			ExpectedCalls: []string{"Gert"},
		},
		{
			Name: "Struct with two string fields",
			Input: struct {
				Name string
				City string
			}{"Gert", "Helsinki"},
			ExpectedCalls: []string{"Gert", "Helsinki"},
		},
		{
			Name: "Struct with non string field",
			Input: struct {
				Name string
				Age  int
			}{"Gert", 28},
			ExpectedCalls: []string{"Gert"},
		},
		{
			Name: "Nested fields",
			Input: Person{
				"Gert",
				Profile{29, "Helsinki"},
			},
			ExpectedCalls: []string{"Gert", "Helsinki"},
		},
		{
			Name: "Pointers to things",
			Input: &Person{
				"Gert",
				Profile{29, "Helsinki"},
			},
			ExpectedCalls: []string{"Gert", "Helsinki"},
		},
		{
			Name: "Slices",
			Input: []Profile{
				{29, "Helsinki"},
				{39, "Tampere"},
			},
			ExpectedCalls: []string{"Helsinki", "Tampere"},
		},
		{
			Name: "Arrays",
			Input: [2]Profile{
				{29, "Helsinki"},
				{39, "Tampere"},
			},
			ExpectedCalls: []string{"Helsinki", "Tampere"},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			var got []string
			Walk(test.Input, func(input string) {
				got = append(got, input)
			})

			if !reflect.DeepEqual(got, test.ExpectedCalls) {
				t.Errorf("got %v, want %v", got, test.ExpectedCalls)
			}
		})
	}

	t.Run("With maps", func(t *testing.T) {
		aMap := map[string]string{
			"Foo": "Bar",
			"Baz": "Boz",
		}

		var got []string
		Walk(aMap, func(input string) {
			got = append(got, input)
		})

		assertContains(t, got, "Bar")
		assertContains(t, got, "Boz")
	})

	t.Run("With channels", func(t *testing.T) {
		aChannel := make(chan Profile)

		go func() {
			aChannel <- Profile{29, "Helsinki"}
			aChannel <- Profile{39, "Tampere"}
			close(aChannel)
		}()

		var got []string
		want := []string{"Helsinki", "Tampere"}

		Walk(aChannel, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}

func assertContains(t testing.TB, haystack []string, needle string) {
	t.Helper()
	contains := false
	for _, x := range haystack {
		if x == needle {
			contains = true
		}
	}
	if !contains {
		t.Errorf("expected %+v to contain %q but it didn't", haystack, needle)
	}
}
