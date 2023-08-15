package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowResourceGroupRequest Request Object
type ShowResourceGroupRequest struct {

	// 资源分组ID，以rg开头，后跟22位由字母或数字组成的字符串
	GroupId string `json:"group_id"`
}

func (o ShowResourceGroupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowResourceGroupRequest struct{}"
	}

	return strings.Join([]string{"ShowResourceGroupRequest", string(data)}, " ")
}
