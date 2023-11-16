package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListBasePluginsNewPostRequest Request Object
type ListBasePluginsNewPostRequest struct {

	// 租户ID
	DomainId string `json:"domain_id"`

	// 偏移
	Offset int32 `json:"offset"`

	// 大小
	Limit int32 `json:"limit"`

	Body *BusinessTypePluginsQueryDto `json:"body,omitempty"`
}

func (o ListBasePluginsNewPostRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListBasePluginsNewPostRequest struct{}"
	}

	return strings.Join([]string{"ListBasePluginsNewPostRequest", string(data)}, " ")
}
