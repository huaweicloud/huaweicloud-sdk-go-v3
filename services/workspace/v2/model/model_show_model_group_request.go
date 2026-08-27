package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowModelGroupRequest Request Object
type ShowModelGroupRequest struct {

	// 模型组id。
	GroupId string `json:"group_id"`
}

func (o ShowModelGroupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowModelGroupRequest struct{}"
	}

	return strings.Join([]string{"ShowModelGroupRequest", string(data)}, " ")
}
