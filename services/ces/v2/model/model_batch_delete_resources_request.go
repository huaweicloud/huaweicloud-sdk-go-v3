package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteResourcesRequest Request Object
type BatchDeleteResourcesRequest struct {

	// 资源分组ID，以rg开头，后跟22位由字母或数字组成的字符串
	GroupId string `json:"group_id"`

	Body *DelResourcesReq `json:"body,omitempty"`
}

func (o BatchDeleteResourcesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteResourcesRequest struct{}"
	}

	return strings.Join([]string{"BatchDeleteResourcesRequest", string(data)}, " ")
}
