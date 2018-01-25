//
// shuffle implements a Fischer/Yates shuffling algorithm in Go.
//
// Copyright (c) 2018, Caltech
// All rights not granted herein are expressly reserved by Caltech.
//
// Redistribution and use in source and binary forms, with or without modification, are permitted provided that the following conditions are met:
//
// 1. Redistributions of source code must retain the above copyright notice, this list of conditions and the following disclaimer.
//
// 2. Redistributions in binary form must reproduce the above copyright notice, this list of conditions and the following disclaimer in the documentation and/or other materials provided with the distribution.
//
// 3. Neither the name of the copyright holder nor the names of its contributors may be used to endorse or promote products derived from this software without specific prior written permission.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
//
package shuffle

import (
	"math/rand"
)

const (
	Version = `v0.0.1`
)

// NOTE: You need to initialize your random number generator, e.g.
//    random := rand.New(rand.NewSource(time.Now().UnixNano())
// Before calling one of the shuffle functions.
// This package maybe obsolete after go v1.10.x

// Strings shuffles an array of strings in place.
func Strings(a []string, random *rand.Rand) {
	for i := len(a) - 1; i > 0; i-- {
		j := random.Intn(i + 1)
		a[i], a[j] = a[j], a[i]
	}
}

// Ints shuffles an array of ints in place.
func Ints(a []int, random *rand.Rand) {
	for i := len(a) - 1; i > 0; i-- {
		j := random.Intn(i + 1)
		a[i], a[j] = a[j], a[i]
	}
}
