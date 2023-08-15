package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ImportMicroserviceRequest Request Object
type ImportMicroserviceRequest struct {

	// 实例ID，在API网关控制台的“实例信息”中获取。
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
