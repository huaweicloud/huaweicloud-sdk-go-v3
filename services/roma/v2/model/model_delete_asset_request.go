package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteAssetRequest Request Object
type DeleteAssetRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id"`

	Body *AssetOperateRequest `json:"body,omitempty"`
}

func (o DeleteAssetRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteAssetRequest struct{}"
	}

	return strings.Join([]string{"DeleteAssetRequest", string(data)}, " ")
}
