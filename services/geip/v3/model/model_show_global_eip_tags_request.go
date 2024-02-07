package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowGlobalEipTagsRequest Request Object
type ShowGlobalEipTagsRequest struct {

	// 全域弹性公网IP的id
	ResourceId string `json:"resource_id"`
}

func (o ShowGlobalEipTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowGlobalEipTagsRequest struct{}"
	}

	return strings.Join([]string{"ShowGlobalEipTagsRequest", string(data)}, " ")
}
