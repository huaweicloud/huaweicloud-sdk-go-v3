package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowDdmDetailRequest Request Object
type ShowDdmDetailRequest struct {

	// 实例 ID。
	InstanceId string `json:"instance_id"`
}

func (o ShowDdmDetailRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDdmDetailRequest struct{}"
	}

	return strings.Join([]string{"ShowDdmDetailRequest", string(data)}, " ")
}
