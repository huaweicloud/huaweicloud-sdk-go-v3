package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateResourceGroupResponse Response Object
type CreateResourceGroupResponse struct {

	// 资源分组ID，以rg开头，后跟22位由字母或数字组成的字符串
	GroupId        *string `json:"group_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateResourceGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateResourceGroupResponse struct{}"
	}

	return strings.Join([]string{"CreateResourceGroupResponse", string(data)}, " ")
}
