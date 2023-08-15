package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteResourceGroupRequest Request Object
type DeleteResourceGroupRequest struct {

	// 资源分组ID。
	GroupId string `json:"group_id"`
}

func (o DeleteResourceGroupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteResourceGroupRequest struct{}"
	}

	return strings.Join([]string{"DeleteResourceGroupRequest", string(data)}, " ")
}
