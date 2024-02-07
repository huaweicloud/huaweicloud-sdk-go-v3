package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddGlobalEipTagsRequest Request Object
type AddGlobalEipTagsRequest struct {

	// 全域弹性公网IP的id
	ResourceId string `json:"resource_id"`

	Body *CreateV2TagRequestBody `json:"body,omitempty"`
}

func (o AddGlobalEipTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddGlobalEipTagsRequest struct{}"
	}

	return strings.Join([]string{"AddGlobalEipTagsRequest", string(data)}, " ")
}
