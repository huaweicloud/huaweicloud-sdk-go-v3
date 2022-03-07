package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type DeleteResourceGroupsResourcesBatchRequest struct {
	// 当前资源所在分组信息，以rg开头，后跟22位由字母或数字组成的字符串

	GroupId string `json:"group_id"`

	Body *ResourcesReq `json:"body,omitempty"`
}

func (o DeleteResourceGroupsResourcesBatchRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteResourceGroupsResourcesBatchRequest struct{}"
	}

	return strings.Join([]string{"DeleteResourceGroupsResourcesBatchRequest", string(data)}, " ")
}
