package main

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

// Run UI server: $GOPATH/bin/goconvey
// TestExampleCleanUp testing with library GoConvey
func TestExampleCleanUp(t *testing.T) {
	x := 0

	Convey("Test A", t, func() {
		x++ // setup

		Convey("Test A-B", func() {
			x++

			Convey("A-B-C1", func() {
				So(x, ShouldEqual, 2)
				So(2, ShouldBeBetween, 0, 3)
				So(2, ShouldAlmostEqual, 1.89, .12)
				So([]int{1, 2, 3}, ShouldContain, 2)
			})
			Convey("A-B-C2", func() {
				So(x, ShouldEqual, 4)
				So(2, ShouldEqual, 2)
				//panic("some panic")
			})
			Convey("A-B-C3", func() {
				So(x, ShouldEqual, 6)
			})
		})

		// TearDown
		Reset(func() {
			t.Log("finish")
		})
	})
}
