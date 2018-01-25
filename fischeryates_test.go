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
	"strings"
	"testing"
	"time"
)

func TestStrings(t *testing.T) {
	original := []string{"a", "b", "c", "d", "e", "f", "G", "h", "i", "j", "k", "L", "m", "n"}
	a := []string{"a", "b", "c", "d", "e", "f", "G", "h", "i", "j", "k", "L", "m", "n"}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	Strings(a, r)
	s1 := strings.Join(original, "")
	s2 := strings.Join(a, "")
	if strings.Compare(s1, s2) == 0 {
		t.Errorf("Expected a shuffled string s1 %q, s2 %q", s1, s2)
	}
}

func TestInts(t *testing.T) {
	original := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	Ints(a, r)
	originalOrder := true
	for i, v := range original {
		if a[i] != v {
			originalOrder = false
		}
	}
	if originalOrder {
		t.Errorf("Expected a shuffled ints original %+v, suffled %+v", original, a)
	}
}
