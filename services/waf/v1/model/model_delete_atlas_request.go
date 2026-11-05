package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteAtlasRequest Request Object
type DeleteAtlasRequest struct {

	// 图库ID
	AtlasId string `json:"atlas_id"`
}

func (o DeleteAtlasRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteAtlasRequest struct{}"
	}

	return strings.Join([]string{"DeleteAtlasRequest", string(data)}, " ")
}
