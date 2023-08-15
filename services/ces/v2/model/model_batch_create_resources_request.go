package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchCreateResourcesRequest Request Object
type BatchCreateResourcesRequest struct {

	// 资源分组ID，以rg开头，后跟22位由字母或数字组成的字符串
	GroupId string `json:"group_id"`

	Body *ResourcesReq `json:"body,omitempty"`
}

func (o BatchCreateResourcesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateResourcesRequest struct{}"
	}

	return strings.Join([]string{"BatchCreateResourcesRequest", string(data)}, " ")
}
