package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowProjectTagsListRequest struct {

	// 资源类型
	ResourceType string `json:"resource_type"`
}

func (o ShowProjectTagsListRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowProjectTagsListRequest struct{}"
	}

	return strings.Join([]string{"ShowProjectTagsListRequest", string(data)}, " ")
}
