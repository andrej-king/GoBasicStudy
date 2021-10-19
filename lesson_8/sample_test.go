package main

import (
	"fmt"
	"os"
	"testing"
)

/**
 * TestMain global testing for package
 * Add setup logic before run all tests in this package
 * Run tests in package
 * Add teardown logic after run all tests in this package
 */
func TestMain(m *testing.M) {
	fmt.Println("package setup")
	res := m.Run() // get response code
	fmt.Println("package teardown")

	os.Exit(res)
}

/*
 * TestMultiple check function behaviour
 * For run tests need create configuration Go Test in IDE
 * run command: /bin/go test -json ./...
 * cd in folder with tests and write in terminal: go test -v
 * cd in folder with tests and write in terminal: go test -v -json
 */
func TestMultiple(t *testing.T) {
	//basic
	//var x, y, result = 2, 2, 4
	//realResult := Multiple(x, y)
	//if realResult != result {
	//	t.Errorf("%d != %d", result, realResult)
	//}

	// prepare data (setup): insert data in db

	// create group with tests
	t.Run("groupA", func(t *testing.T) {
		t.Run("simple2", func(t *testing.T) {
			// for run in parallel
			t.Parallel()
			t.Log("simple")

			var x, y, result = 2, 2, 4

			realResult := Multiple(x, y)

			if realResult != result {
				t.Errorf("%d != %d", result, realResult)
			}

			// one more test level in depth
			t.Run("1", func(t *testing.T) {
				r := Multiple(1, 1)
				if r != 1 {
					t.Errorf("failed")
				}
			})
		})

		t.Run("medium2", func(t *testing.T) {
			// for run in parallel
			t.Parallel()
			t.Log("medium")

			var x, y, result = 222, 222, 49284

			realResult := Multiple(x, y)

			if realResult != result {
				t.Errorf("%d != %d", result, realResult)
			}
		})

		t.Run("negative2", func(t *testing.T) {
			// for run in parallel
			t.Parallel()
			t.Log("negative")

			var x, y, result = -2, 4, -8

			realResult := Multiple(x, y)

			if realResult != result {
				t.Errorf("%d != %d", result, realResult)
			}
		})
	})

	t.Run("groupB", func(t *testing.T) {
		t.Run("simple3", func(t *testing.T) {
			// for run in parallel
			t.Parallel()
			t.Log("simple")

			var x, y, result = 2, 2, 4

			realResult := Multiple(x, y)

			if realResult != result {
				t.Errorf("%d != %d", result, realResult)
			}

			// one more test level in depth
			t.Run("1", func(t *testing.T) {
				r := Multiple(1, 1)
				if r != 1 {
					t.Errorf("failed")
				}
			})
		})

		t.Run("medium3", func(t *testing.T) {
			// for run in parallel
			t.Parallel()
			t.Log("medium")

			var x, y, result = 222, 222, 49284

			realResult := Multiple(x, y)

			if realResult != result {
				t.Errorf("%d != %d", result, realResult)
			}
		})

		t.Run("negative3", func(t *testing.T) {
			// for run in parallel
			t.Parallel()
			t.Log("negative")

			var x, y, result = -2, 4, -8

			realResult := Multiple(x, y)

			if realResult != result {
				t.Errorf("%d != %d", result, realResult)
			}
		})
	})

	// Sequence
	t.Run("simple", func(t *testing.T) {
		// for run in parallel
		t.Parallel()
		t.Log("simple")

		var x, y, result = 2, 2, 4

		realResult := Multiple(x, y)

		if realResult != result {
			t.Errorf("%d != %d", result, realResult)
		}

		// one more test level in depth
		t.Run("1", func(t *testing.T) {
			r := Multiple(1, 1)
			if r != 1 {
				t.Errorf("failed")
			}
		})
	})

	t.Run("medium", func(t *testing.T) {
		// for run in parallel
		t.Parallel()
		t.Log("medium")

		var x, y, result = 222, 222, 49284

		realResult := Multiple(x, y)

		if realResult != result {
			t.Errorf("%d != %d", result, realResult)
		}
	})

	t.Run("negative", func(t *testing.T) {
		// for run in parallel
		t.Parallel()
		t.Log("negative")

		var x, y, result = -2, 4, -8

		realResult := Multiple(x, y)

		if realResult != result {
			t.Errorf("%d != %d", result, realResult)
		}
	})

	// clean data after tests (teardown): delete test data
}

/**
 * TestAdd
 * go test -v -run /simple/1 will run tests /simple/1 from all methods
 */
func TestAdd(t *testing.T) {
	t.Run("simple", func(t *testing.T) {
		// for run in parallel
		t.Parallel()
		t.Log("simple")

		var x, y, result = 2, 2, 4

		realResult := Add(x, y)

		if realResult != result {
			t.Errorf("%d != %d", result, realResult)
		}

		// one more test level in depth
		t.Run("1", func(t *testing.T) {
			r := Add(1, 1)
			if r != 2 {
				t.Errorf("failed")
			}
		})
	})
}
