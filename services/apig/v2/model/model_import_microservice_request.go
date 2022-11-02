package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ImportMicroserviceRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id"`

	Body *MicroserviceImportReq `json:"body,omitempty"`
}

func (o ImportMicroserviceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImportMicroserviceRequest struct{}"
	}

	return strings.Join([]string{"ImportMicroserviceRequest", string(data)}, " ")
}
