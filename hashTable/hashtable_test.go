package hashtable

import (
	"testing"
)

func TestHashTableSetGet(t *testing.T) {
	var h *HashTable
	testCases := []struct {
		key   string
		value string
	}{
		{"John", "Smith"},
		{"Adam", "Foxal"},
		{"Andy", "Steve"},
		{" ", ""},
		{"", " "},
		{"", ""},
	}

	for i, test := range testCases {
		h = New()
		h.Set(test.key, test.value)
		result, ok := h.Get(test.key)
		if !ok {
			t.Errorf("[%d]: Could not get value from key \"%s\"", i, test.key)
		}
		if result != test.value {
			t.Errorf(
				"[%d]: Incorrect value for key \"%s\", got = \"%s\", expected = \"%s\"",
				i, test.key, result, test.value,
			)
		}
	}
}
