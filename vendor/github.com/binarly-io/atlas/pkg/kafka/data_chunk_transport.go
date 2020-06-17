// Copyright 2020 The Atlas Authors. All rights reserved.
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

package kafka

import (
	"io"

	"github.com/golang/protobuf/proto"

	"github.com/binarly-io/atlas/pkg/api/atlas"
)

type KafkaDataChunkTransport struct {
	producer *Producer
	consumer *Consumer
	close    bool
}

// NewKafkaDataChunkTransport
func NewKafkaDataChunkTransport(producer *Producer, consumer *Consumer, close bool) *KafkaDataChunkTransport {
	return &KafkaDataChunkTransport{
		producer: producer,
		consumer: consumer,
		close:    close,
	}
}

// Close
func (t *KafkaDataChunkTransport) Close() {
	if !t.close {
		return
	}

	if t.producer != nil {
		t.producer.Close()
		t.producer = nil
	}

	if t.consumer != nil {
		t.consumer.Close()
		t.consumer = nil
	}
}

// Send
func (t *KafkaDataChunkTransport) Send(dataChunk *atlas.DataChunk) error {
	if buf, err := proto.Marshal(dataChunk); err == nil {
		return t.producer.Send(buf)
	} else {
		return err
	}
}

// Recv
func (t *KafkaDataChunkTransport) Recv() (*atlas.DataChunk, error) {
	msg := t.consumer.Recv()
	if msg == nil {
		// TODO not sure
		return nil, io.EOF
	}
	dataChunk := &atlas.DataChunk{}
	return dataChunk, proto.Unmarshal(msg.Value, dataChunk)
}
