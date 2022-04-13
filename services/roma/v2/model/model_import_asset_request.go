package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ImportAssetRequest struct {
	// 实例ID

	InstanceId string `json:"instance_id"`

	Body string `json:"body,omitempty"`
}

func (o ImportAssetRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImportAssetRequest struct{}"
	}

	return strings.Join([]string{"ImportAssetRequest", string(data)}, " ")
}
