// Copyright 2025 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package windows_test

import (
	"fmt"
	"math"
	"time"

	"golang.org/x/sys/windows"
)

func ExampleFiletime_Unix() {
	ftZero := windows.Filetime{
		HighDateTime: 0x0,
		LowDateTime:  0x0,
	}
	secs, nsecs := ftZero.Unix()
	fmt.Println(secs, nsecs)

	zeroTime := time.Unix(ftZero.Unix())
	match := zeroTime.Equal(time.Date(1601, time.January, 1, 0, 0, 0, 0, time.UTC))
	fmt.Printf("Matches January 1, 1601: %v\n", match)

	// Output:
	// -11644473600 0
	// Matches January 1, 1601: true
}

func ExampleNsecToFiletime() {
	ftMin := windows.NsecToFiletime(math.MinInt64)
	fmt.Println("Minimal Filetime with nanoseconds only:")
	fmt.Printf("hi: %#x lo: %#x\n", ftMin.HighDateTime, ftMin.LowDateTime)
	fmt.Printf("= %v\n", time.Unix(0, math.MinInt64).UTC())

	ftMax := windows.NsecToFiletime(math.MaxInt64)
	fmt.Println("\nMaximum Filetime with nanoseconds only:")
	fmt.Printf("hi: %#x lo: %#x\n", ftMax.HighDateTime, ftMax.LowDateTime)
	fmt.Printf("= %v\n", time.Unix(0, math.MaxInt64).UTC())

	// Output:
	// Minimal Filetime with nanoseconds only:
	// hi: 0x5603ca lo: 0x5a5d3852
	// = 1677-09-21 00:12:43.145224192 +0000 UTC
	//
	// Maximum Filetime with nanoseconds only:
	// hi: 0x2e55ff3 lo: 0x501fc7ae
	// = 2262-04-11 23:47:16.854775807 +0000 UTC
}
