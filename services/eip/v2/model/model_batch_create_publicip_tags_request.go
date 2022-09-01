package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type BatchCreatePublicipTagsRequest struct {

	// 资源ID
	PublicipId string `json:"publicip_id" xml:"publicip_id"`

	Body *BatchCreatePublicipTagsRequestBody `json:"body,omitempty" xml:"body"`
}

func (o BatchCreatePublicipTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreatePublicipTagsRequest struct{}"
	}

	return strings.Join([]string{"BatchCreatePublicipTagsRequest", string(data)}, " ")
}
