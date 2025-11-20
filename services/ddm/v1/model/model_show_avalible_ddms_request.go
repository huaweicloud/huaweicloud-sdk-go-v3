package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAvalibleDdmsRequest Request Object
type ShowAvalibleDdmsRequest struct {

	// 实例 ID。
	InstanceId string `json:"instance_id"`
}

func (o ShowAvalibleDdmsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAvalibleDdmsRequest struct{}"
	}

	return strings.Join([]string{"ShowAvalibleDdmsRequest", string(data)}, " ")
}
