package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type BatchDeletePublicipTagsRequest struct {

	// 资源ID
	PublicipId string `json:"publicip_id" xml:"publicip_id"`

	Body *BatchDeletePublicipTagsRequestBody `json:"body,omitempty" xml:"body"`
}

func (o BatchDeletePublicipTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeletePublicipTagsRequest struct{}"
	}

	return strings.Join([]string{"BatchDeletePublicipTagsRequest", string(data)}, " ")
}
