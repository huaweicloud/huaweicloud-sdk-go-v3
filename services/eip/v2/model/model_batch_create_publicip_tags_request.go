package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type BatchCreatePublicipTagsRequest struct {
	// 资源ID

	PublicipId string `json:"publicip_id"`

	Body *BatchCreatePublicipTagsRequestBody `json:"body,omitempty"`
}

func (o BatchCreatePublicipTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreatePublicipTagsRequest struct{}"
	}

	return strings.Join([]string{"BatchCreatePublicipTagsRequest", string(data)}, " ")
}
