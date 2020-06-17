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

package config_consumer

import (
	"bytes"
	"fmt"

	conf "github.com/spf13/viper"
)

// IMPORTANT
// IMPORTANT Do not forget to update String() function
// IMPORTANT
type ConfigConsumer struct {
	Verbose bool `mapstructure:"verbose"`

	// Kafka
	Brokers    []string `mapstructure:"brokers"`
	Topic      string   `mapstructure:"topic"`
	GroupID    string   `mapstructure:"groupID"`
	ReadNewest bool     `mapstructure:"readNewest"`
	Ack        bool     `mapstructure:"ack"`
	// IMPORTANT
	// IMPORTANT Do not forget to update String() function
	// IMPORTANT
}

var Config ConfigConsumer

func ReadIn() {
	_ = conf.Unmarshal(&Config)
}

func (c *ConfigConsumer) String() string {
	b := &bytes.Buffer{}

	_, _ = fmt.Fprintf(b, "Verbose: %v\n", c.Verbose)

	_, _ = fmt.Fprintf(b, "Brokers: %v\n", c.Brokers)
	_, _ = fmt.Fprintf(b, "Topic: %v\n", c.Topic)
	_, _ = fmt.Fprintf(b, "GroupID: %v\n", c.GroupID)
	_, _ = fmt.Fprintf(b, "ReadNewest: %v\n", c.ReadNewest)
	_, _ = fmt.Fprintf(b, "Ack: %v\n", c.Ack)

	return b.String()
}
