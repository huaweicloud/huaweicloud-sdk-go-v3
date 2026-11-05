package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddImagesRequest Request Object
type AddImagesRequest struct {

	// 图库ID
	AtlasId string `json:"atlas_id"`

	Body *AddImagesBody `json:"body,omitempty"`
}

func (o AddImagesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddImagesRequest struct{}"
	}

	return strings.Join([]string{"AddImagesRequest", string(data)}, " ")
}
