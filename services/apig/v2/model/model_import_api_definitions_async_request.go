package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ImportApiDefinitionsAsyncRequest Request Object
type ImportApiDefinitionsAsyncRequest struct {

	// 实例ID，在API网关控制台的“实例信息”中获取。
	InstanceId string `json:"instance_id"`

	Body *ImportApiDefinitionsAsyncRequestBody `json:"body,omitempty" type:"multipart"`
}

func (o ImportApiDefinitionsAsyncRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImportApiDefinitionsAsyncRequest struct{}"
	}

	return strings.Join([]string{"ImportApiDefinitionsAsyncRequest", string(data)}, " ")
}
