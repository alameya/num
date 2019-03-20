package num_test

import (
	"github.com/magiconair/properties/assert"
	"num"
	"sync"
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
		maxValue:        num.MaxUint(),
		incrementsCount: num.MaxUint(),
		expectedValue:   num.MaxUint(),
	})
}

func TestCounter_Value(t *testing.T) {
	t.Parallel()
	testTable := []testCase{
		{
			maxValue:        num.MaxUint(),
			incrementsCount: 0,
			expectedValue:   0,
		},
		{
			maxValue:        num.MaxUint(),
			incrementsCount: 2,
			expectedValue:   2,
		},
		{
			maxValue:        num.MaxUint(),
			incrementsCount: 10000,
			expectedValue:   10000,
		},

	}

	processTestCases(t, testTable...)
}

func TestCounter_Max_AfterIncrement(t *testing.T) {
	t.Parallel()
	testTable := []testCase{
		{
			maxValue:        5,
			incrementsCount: 10,
			expectedValue:   0,
		},
		{
			maxValue:        10,
			incrementsCount: 10,
			expectedValue:   0,
		},
		{
			maxValue:        10,
			incrementsCount: 9,
			expectedValue:   9,
		},

	}

	for _, testObj := range testTable {
		counter := num.NewCounter()
		for i := uint(0); testObj.incrementsCount != i; i++ {
			counter.Increment()
		}

		counter.Max(testObj.maxValue)
		assert.Equal(t, counter.Value(), testObj.expectedValue)
	}
}

func processTestCases(t *testing.T, testCases ...testCase) {
	t.Helper()
	for _, testObj := range testCases {
		var wg sync.WaitGroup
		counter := num.NewCounter()
		counter.Max(testObj.maxValue)
		for i := uint(0); testObj.incrementsCount != i; i++ {
			wg.Add(1)
			go func(c *num.Counter) {
				c.Increment()
				wg.Done()
			}(counter)

			wg.Wait()
		}

		assert.Equal(t, counter.Value(), testObj.expectedValue)
	}
}
