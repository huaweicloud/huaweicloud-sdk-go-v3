package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ShowKafkaTopicPartitionDiskusageResponse struct {
	// Broker列表。

	BrokerList     *[]ShowKafkaTopicPartitionDiskusageRespBrokerList `json:"broker_list,omitempty"`
	HttpStatusCode int                                               `json:"-"`
}

func (o ShowKafkaTopicPartitionDiskusageResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowKafkaTopicPartitionDiskusageResponse struct{}"
	}

	return strings.Join([]string{"ShowKafkaTopicPartitionDiskusageResponse", string(data)}, " ")
}
