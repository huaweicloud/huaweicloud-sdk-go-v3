package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ImportAssetRequest Request Object
type ImportAssetRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id"`

	Body *ImportAssetRequestBody `json:"body,omitempty" type:"multipart"`
}

func (o ImportAssetRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImportAssetRequest struct{}"
	}

	return strings.Join([]string{"ImportAssetRequest", string(data)}, " ")
}
