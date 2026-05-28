// Copyright 2025 Huawei Technologies Co.,Ltd.
//
// Licensed to the Apache Software Foundation (ASF) under one
// or more contributor license agreements.  See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership.  The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License.  You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package hkdf

import (
	"bytes"
	"crypto/sha256"
	"crypto/sha512"
	"hash"
	"io"
	"testing"
)

func TestExtract(t *testing.T) {
	tests := []struct {
		name   string
		hash   func() hash.Hash
		secret []byte
		salt   []byte
		want   []byte
	}{
		{
			name:   "sha256 with salt",
			hash:   sha256.New,
			secret: []byte("secret"),
			salt:   []byte("salt"),
			want:   []byte{0x98, 0xe5, 0x34, 0x0f, 0x0f, 0x4f, 0x96, 0xd2, 0xb8, 0x0c, 0x2a, 0x90, 0xda, 0x0d, 0x03, 0xcf, 0x46, 0xc3, 0x5e, 0x94, 0x92, 0x91, 0x8c, 0xc7, 0xaf, 0x73, 0xd9, 0xa3, 0x9e, 0xfa, 0x59, 0x81},
		},
		{
			name:   "sha256 with empty salt",
			hash:   sha256.New,
			secret: []byte("secret"),
			salt:   []byte{},
			want:   []byte{0x18, 0x10, 0xe8, 0xd3, 0x41, 0xb9, 0xe6, 0x1e, 0x88, 0x89, 0x5f, 0xba, 0x0b, 0x1a, 0xa2, 0xbd, 0x03, 0xcb, 0x6c, 0x8f, 0xff, 0x2b, 0x36, 0x8e, 0x66, 0x6e, 0xf8, 0xef, 0xde, 0xf5, 0xb4, 0xfb},
		},
		{
			name:   "sha256 with nil salt",
			hash:   sha256.New,
			secret: []byte("secret"),
			salt:   nil,
			want:   []byte{0x18, 0x10, 0xe8, 0xd3, 0x41, 0xb9, 0xe6, 0x1e, 0x88, 0x89, 0x5f, 0xba, 0x0b, 0x1a, 0xa2, 0xbd, 0x03, 0xcb, 0x6c, 0x8f, 0xff, 0x2b, 0x36, 0x8e, 0x66, 0x6e, 0xf8, 0xef, 0xde, 0xf5, 0xb4, 0xfb},
		},
		{
			name:   "sha512 with salt",
			hash:   sha512.New,
			secret: []byte("secret"),
			salt:   []byte("salt"),
			want:   []byte{0x57, 0xe6, 0xf1, 0xea, 0x12, 0x66, 0x6f, 0x22, 0x77, 0xf8, 0xac, 0x04, 0x4f, 0x50, 0xff, 0x7e, 0xab, 0x5e, 0x7f, 0x35, 0x57, 0xdb, 0xdb, 0xca, 0x0a, 0x21, 0xe2, 0x1a, 0x0a, 0xfa, 0x1b, 0xd6, 0x26, 0x0c, 0xc0, 0x96, 0x86, 0xe3, 0x69, 0xfb, 0xee, 0x8d, 0x2d, 0xa2, 0x72, 0x96, 0xd7, 0xc2, 0xe4, 0x86, 0x4b, 0x75, 0xac, 0xcc, 0x2c, 0x74, 0x77, 0xcb, 0xec, 0xf3, 0xa3, 0x8c, 0x9b, 0xe7},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Extract(tt.hash, tt.secret, tt.salt)
			if !bytes.Equal(got, tt.want) {
				t.Errorf("Extract() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNew(t *testing.T) {
	tests := []struct {
		name   string
		hash   func() hash.Hash
		secret []byte
		salt   []byte
		info   []byte
		n      int
		want   []byte
	}{
		{
			name:   "sha256 basic",
			hash:   sha256.New,
			secret: []byte("secret"),
			salt:   []byte("salt"),
			info:   []byte("info"),
			n:      16,
			want:   []byte{0xf6, 0xd2, 0xfc, 0xc4, 0x7c, 0xb9, 0x39, 0xde, 0xaf, 0xe3, 0x85, 0x3a, 0x1e, 0x64, 0x1a, 0x27},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hkdf := New(tt.hash, tt.secret, tt.salt, tt.info)
			got := make([]byte, tt.n)
			_, err := io.ReadFull(hkdf, got)
			if err != nil {
				t.Errorf("ReadFull() error = %v", err)
				return
			}
			if !bytes.Equal(got, tt.want) {
				t.Errorf("ReadFull() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestExpand(t *testing.T) {
	prk := Extract(sha256.New, []byte("secret"), []byte("salt"))
	hkdf := Expand(sha256.New, prk, []byte("info"))
	got := make([]byte, 16)
	_, err := io.ReadFull(hkdf, got)
	if err != nil {
		t.Errorf("ReadFull() error = %v", err)
	}
	want := []byte{0xf6, 0xd2, 0xfc, 0xc4, 0x7c, 0xb9, 0x39, 0xde, 0xaf, 0xe3, 0x85, 0x3a, 0x1e, 0x64, 0x1a, 0x27}
	if !bytes.Equal(got, want) {
		t.Errorf("Expand() = %v, want %v", got, want)
	}
}

func TestReadLimitExceeded(t *testing.T) {
	hkdf := New(sha256.New, []byte("key"), []byte("salt"), []byte("info"))
	maxSize := sha256.New().Size() * 255
	buf := make([]byte, maxSize)
	_, err := io.ReadFull(hkdf, buf)
	if err != nil {
		t.Errorf("ReadFull() error = %v", err)
	}
	_, err = io.ReadFull(hkdf, make([]byte, 1))
	if err == nil || err.Error() != "hkdf: output limit exceeded" {
		t.Errorf("expected limit exceeded error, got %v", err)
	}
}

func TestMultiRead(t *testing.T) {
	hkdf := New(sha256.New, []byte("secret"), []byte("salt"), []byte("info"))
	out := make([]byte, 32)
	for i := 0; i < 32; i++ {
		n, err := io.ReadFull(hkdf, out[i:i+1])
		if n != 1 || err != nil {
			t.Errorf("ReadFull() at byte %d: n = %d, err = %v", i, n, err)
		}
	}
}
