package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RemoveImagesRequest Request Object
type RemoveImagesRequest struct {

	// 图库ID
	AtlasId string `json:"atlas_id"`

	Body *RemoveImagesBody `json:"body,omitempty"`
}

func (o RemoveImagesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RemoveImagesRequest struct{}"
	}

	return strings.Join([]string{"RemoveImagesRequest", string(data)}, " ")
}
