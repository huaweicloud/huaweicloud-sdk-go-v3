package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteModelGroupRequest Request Object
type DeleteModelGroupRequest struct {

	// 模型组id。
	GroupId string `json:"group_id"`
}

func (o DeleteModelGroupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteModelGroupRequest struct{}"
	}

	return strings.Join([]string{"DeleteModelGroupRequest", string(data)}, " ")
}
