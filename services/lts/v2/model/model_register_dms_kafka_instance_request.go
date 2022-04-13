package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type RegisterDmsKafkaInstanceRequest struct {
	Body *RegisterDmsKafkaInstanceRequestBody `json:"body,omitempty"`
}

func (o RegisterDmsKafkaInstanceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RegisterDmsKafkaInstanceRequest struct{}"
	}

	return strings.Join([]string{"RegisterDmsKafkaInstanceRequest", string(data)}, " ")
}
