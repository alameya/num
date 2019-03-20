package num

import (
	"github.com/magiconair/properties/assert"
	"testing"
)

type testCase struct {
	maxValue        int
	incrementsCount int
	expectedValue   int
}

func TestCounter_Max_Abnormal(t *testing.T) {
	t.Parallel()
	maxValues := []int{MinInt(), -1000, -1, 0, 1}

	for _, max := range maxValues {
		counter := NewCounter()
		assert.Equal(t, counter.Value(), 0)

		counter.Max(max)
		assert.Equal(t, counter.Value(), 0)

		counter.Increment()
		assert.Equal(t, counter.Value(), 0)
	}
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
		maxValue: MaxInt(),
		incrementsCount: MaxInt(),
		expectedValue: MaxInt(),
	})
}

func TestCounter_Value(t *testing.T) {
	t.Parallel()
	testTable := []testCase{
		{
			maxValue:        MaxInt(),
			incrementsCount: 0,
			expectedValue:   0,
		},
		{
			maxValue:        MaxInt(),
			incrementsCount: 2,
			expectedValue:   2,
		},
		{
			maxValue:        MaxInt(),
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
		for i := 0; testObj.incrementsCount != i; i++ {
			counter.Increment()
		}

		assert.Equal(t, counter.Value(), testObj.expectedValue)
	}
}
