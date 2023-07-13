package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowPublicipTagsRequest Request Object
type ShowPublicipTagsRequest struct {

	// 资源ID
	PublicipId string `json:"publicip_id"`
}

func (o ShowPublicipTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPublicipTagsRequest struct{}"
	}

	return strings.Join([]string{"ShowPublicipTagsRequest", string(data)}, " ")
}
