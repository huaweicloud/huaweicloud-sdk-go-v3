package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteGlobalEipTagsRequest Request Object
type BatchDeleteGlobalEipTagsRequest struct {

	// 全域弹性公网IP的id
	ResourceId string `json:"resource_id"`

	Body *BatchDeleteV2RequestBody `json:"body,omitempty"`
}

func (o BatchDeleteGlobalEipTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteGlobalEipTagsRequest struct{}"
	}

	return strings.Join([]string{"BatchDeleteGlobalEipTagsRequest", string(data)}, " ")
}
