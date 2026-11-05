package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAtlasRequest Request Object
type ShowAtlasRequest struct {

	// 图库ID
	AtlasId string `json:"atlas_id"`
}

func (o ShowAtlasRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAtlasRequest struct{}"
	}

	return strings.Join([]string{"ShowAtlasRequest", string(data)}, " ")
}
