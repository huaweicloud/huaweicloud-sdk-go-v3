package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ImportApiDefinitionsV2Request struct {
	// 实例ID

	InstanceId string `json:"instance_id"`

	Body *ImportApiDefinitionsV2RequestBody `json:"body,omitempty" type:"multipart"`
}

func (o ImportApiDefinitionsV2Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImportApiDefinitionsV2Request struct{}"
	}

	return strings.Join([]string{"ImportApiDefinitionsV2Request", string(data)}, " ")
}
