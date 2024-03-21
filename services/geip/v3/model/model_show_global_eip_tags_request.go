package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowGlobalEipTagsRequest Request Object
type ShowGlobalEipTagsRequest struct {
	ResourceId string `json:"resource_id"`
}

func (o ShowGlobalEipTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowGlobalEipTagsRequest struct{}"
	}

	return strings.Join([]string{"ShowGlobalEipTagsRequest", string(data)}, " ")
}
