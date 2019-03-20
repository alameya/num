package num

import (
	"github.com/magiconair/properties/assert"
	"testing"
)

type testCase struct {
	maxValue        uint
	incrementsCount uint
	expectedValue   uint
}

func TestCounter_Max_Abnormal(t *testing.T) {
	t.Parallel()
	testTable := []testCase{
		{
			maxValue:        0,
			incrementsCount: 1,
			expectedValue:   0,
		},
		{
			maxValue:        1,
			incrementsCount: 1,
			expectedValue:   0,
		},
	}

	processTestCases(t, testTable...)
}

func TestCounter_Max(t *testing.T) {
	t.Parallel()
	testTable := []testCase{
		{
			maxValue:        2,
			incrementsCount: 2,
			expectedValue:   0,
		},
		{
			maxValue:        2,
			incrementsCount: 3,
			expectedValue:   1,
		},
	}

	processTestCases(t, testTable...)
}

func TestCounter_Value_MaxIncrements(t *testing.T) {
	t.Parallel()
	if testing.Short() {
		t.Skip("skipped because is to long")
	}

	processTestCases(t, testCase{
		maxValue:        MaxUint(),
		incrementsCount: MaxUint(),
		expectedValue:   MaxUint(),
	})
}

func TestCounter_Value(t *testing.T) {
	t.Parallel()
	testTable := []testCase{
		{
			maxValue:        MaxUint(),
			incrementsCount: 0,
			expectedValue:   0,
		},
		{
			maxValue:        MaxUint(),
			incrementsCount: 2,
			expectedValue:   2,
		},
		{
			maxValue:        MaxUint(),
			incrementsCount: 14,
			expectedValue:   14,
		},

	}

	processTestCases(t, testTable...)
}

func processTestCases(t *testing.T, testCases ...testCase) {
	t.Helper()
	for _, testObj := range testCases {
		counter := NewCounter()
		counter.Max(testObj.maxValue)
		for i := uint(0); testObj.incrementsCount != i; i++ {
			counter.Increment()
		}

		assert.Equal(t, counter.Value(), testObj.expectedValue)
	}
}
