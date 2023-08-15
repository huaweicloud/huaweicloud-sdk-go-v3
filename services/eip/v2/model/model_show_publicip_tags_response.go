package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowPublicipTagsResponse Response Object
type ShowPublicipTagsResponse struct {

	// 标签列表
	Tags           *[]ResourceTagResp `json:"tags,omitempty"`
	HttpStatusCode int                `json:"-"`
}

func (o ShowPublicipTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPublicipTagsResponse struct{}"
	}

	return strings.Join([]string{"ShowPublicipTagsResponse", string(data)}, " ")
}
