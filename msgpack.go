// Copyright 2017 bingo Author. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package codec

import "github.com/shamaton/msgpack"

// Msgpack消息协议
type messagepack struct {
}

func (p *messagepack) Marshal(v interface{}) ([]byte, error) {
	return msgpack.Encode(v)
}

func (p *messagepack) Unmarshal(data []byte, v interface{}) error {
	return msgpack.Decode(data, v)
}

func (p *messagepack) ContentType() string {
	return "application/msgpack"
}

func (p *messagepack) Name() string {
	return "msgpack"
}

func (p *messagepack) Type() Type {
	return MessagePack
}
