package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListApiRuntimeDefinitionV2Request struct {
	// 实例ID

	InstanceId string `json:"instance_id"`
	// API的编号

	ApiId string `json:"api_id"`
	// API的发布环境编号

	EnvId *string `json:"env_id,omitempty"`
}

func (o ListApiRuntimeDefinitionV2Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListApiRuntimeDefinitionV2Request struct{}"
	}

	return strings.Join([]string{"ListApiRuntimeDefinitionV2Request", string(data)}, " ")
}
