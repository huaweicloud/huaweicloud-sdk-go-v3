package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListGcbResourceTagsRequest Request Object
type ListGcbResourceTagsRequest struct {

	// 资源唯一标识符。
	ResourceId string `json:"resource_id"`
}

func (o ListGcbResourceTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListGcbResourceTagsRequest struct{}"
	}

	return strings.Join([]string{"ListGcbResourceTagsRequest", string(data)}, " ")
}
