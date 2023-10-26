package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowBackUpInfoRequest Request Object
type ShowBackUpInfoRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id"`
}

func (o ShowBackUpInfoRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowBackUpInfoRequest struct{}"
	}

	return strings.Join([]string{"ShowBackUpInfoRequest", string(data)}, " ")
}
