package rutificador

import "testing"

func TestObtenerDV(t *testing.T) {
	testCases := []struct {
		input    string
		expected rune
		err      error
	}{
		{"18622178", '8', nil},
		{"11111111", '1', nil},
		{"70360100", '6', nil},
		{"22222222", '2', nil},
		{"15304340", 'K', nil},
		{"8367720", '1', nil},
		{"18.622.178", '8', nil},
	}

	for _, tc := range testCases {
		if result, err := obtenerDV(tc.input); tc.expected != result && err != tc.err {
			t.Errorf("obtenerDV(%s) = %v; want: value: %v expected error: %v, received error: %v", tc.input, result, tc.expected, tc.err, err)
		}
	}
}
