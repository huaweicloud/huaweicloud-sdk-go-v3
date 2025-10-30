// Copyright 2023 Huawei Technologies Co.,Ltd.
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

package progress

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"io"
	"testing"
)

type TestProgressListener struct {
	start       int
	data        int
	complete    int
	err         int
	transferred int64
	total       int64
}

func (l *TestProgressListener) ProgressChanged(event *Event) {
	switch event.Type {
	case TransferStartedEvent:
		l.start += 1
	case TransferDataEvent:
		l.data += 1
		l.transferred = event.TransferredBytes
		l.total = event.TotalBytes
	case TransferCompletedEvent:
		l.complete += 1
	case TransferFailedEvent:
		l.err += 1
	default:
		l.err += 1
	}
}

func TestTeeReader_Read(t *testing.T) {
	listener := &TestProgressListener{}

	size := 1024
	size64 := int64(size)
	data := make([]byte, size)

	reader := NewTeeReader(bytes.NewReader(data), nil, size64, listener, defaultProgressInterval)
	defer func() {
		err := reader.Close()
		assert.NoError(t, err)
	}()

	var out bytes.Buffer
	written, err := io.Copy(&out, reader)

	assert.NoError(t, err)
	assert.Equal(t, size64, written)
	assert.Equal(t, 1, listener.start)
	assert.Equal(t, 1, listener.data)
	assert.Equal(t, 1, listener.complete)
	assert.Equal(t, 0, listener.err)
	assert.Equal(t, listener.transferred, listener.total)
}

func TestTeeReader_Read2(t *testing.T) {
	listener := &TestProgressListener{}

	size := 1024 * 100
	size64 := int64(size)
	data := make([]byte, size)

	reader := NewTeeReader(bytes.NewReader(data), nil, size64, listener, 1024)
	defer func() {
		err := reader.Close()
		assert.NoError(t, err)
	}()

	var out bytes.Buffer
	written, err := io.Copy(&out, reader)

	assert.NoError(t, err)
	assert.Equal(t, size64, written)
	assert.Equal(t, 1, listener.start)
	assert.Equal(t, true, listener.data > 1)
	assert.Equal(t, 1, listener.complete)
	assert.Equal(t, 0, listener.err)
	assert.Equal(t, listener.transferred, listener.total)
}
