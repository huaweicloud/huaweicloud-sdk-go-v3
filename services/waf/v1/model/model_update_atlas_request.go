package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateAtlasRequest Request Object
type UpdateAtlasRequest struct {

	// 图库ID
	AtlasId string `json:"atlas_id"`

	Body *AiguardAtlasUpdateBody `json:"body,omitempty"`
}

func (o UpdateAtlasRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAtlasRequest struct{}"
	}

	return strings.Join([]string{"UpdateAtlasRequest", string(data)}, " ")
}
