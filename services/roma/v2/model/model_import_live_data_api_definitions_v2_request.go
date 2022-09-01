package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ImportLiveDataApiDefinitionsV2Request struct {

	// 实例ID
	InstanceId string `json:"instance_id" xml:"instance_id"`

	Body *ImportLiveDataApiDefinitionsV2RequestBody `json:"body,omitempty" xml:"body" type:"multipart"`
}

func (o ImportLiveDataApiDefinitionsV2Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImportLiveDataApiDefinitionsV2Request struct{}"
	}

	return strings.Join([]string{"ImportLiveDataApiDefinitionsV2Request", string(data)}, " ")
}
