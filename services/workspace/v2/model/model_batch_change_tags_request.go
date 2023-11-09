package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchChangeTagsRequest Request Object
type BatchChangeTagsRequest struct {

	// 桌面id。
	DesktopId string `json:"desktop_id"`

	Body *TagsReq `json:"body,omitempty"`
}

func (o BatchChangeTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchChangeTagsRequest struct{}"
	}

	return strings.Join([]string{"BatchChangeTagsRequest", string(data)}, " ")
}
