package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type RegisterDmsKafkaInstanceRequest struct {
	Body *RegisterDmsKafkaInstanceRequestBody `json:"body,omitempty" xml:"body"`
}

func (o RegisterDmsKafkaInstanceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RegisterDmsKafkaInstanceRequest struct{}"
	}

	return strings.Join([]string{"RegisterDmsKafkaInstanceRequest", string(data)}, " ")
}
