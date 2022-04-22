package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type RegisterDmsKafkaInstanceResponse struct {

	// kafka ID
	InstanceId     *string `json:"instance_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o RegisterDmsKafkaInstanceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RegisterDmsKafkaInstanceResponse struct{}"
	}

	return strings.Join([]string{"RegisterDmsKafkaInstanceResponse", string(data)}, " ")
}
