// Copyright 2023 xgfone
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package defaults

import (
	"bytes"
	"fmt"
	"log"
	"sync"
)

var logkv = stdlog

var bufpool = &sync.Pool{New: func() any { return bytes.NewBuffer(make([]byte, 0, 1024)) }}

func stdlog(msg string, kvs ...any) {
	_len := len(kvs)
	if _len == 0 {
		log.Print(msg)
		return
	}

	buf := bufpool.Get().(*bytes.Buffer)
	buf.WriteString(msg)
	buf.WriteString(", ")
	for i := 0; i < _len; i += 2 {
		if i == 0 {
			fmt.Fprintf(buf, "%v=%v", kvs[i], kvs[i+1])
		} else {
			fmt.Fprintf(buf, ", %v=%v", kvs[i], kvs[i+1])
		}
	}
	log.Print(buf.String())
	buf.Reset()
	bufpool.Put(buf)
}
