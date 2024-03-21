package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteGlobalEipTagRequest Request Object
type DeleteGlobalEipTagRequest struct {
	ResourceId string `json:"resource_id"`

	TagKey string `json:"tag_key"`
}

func (o DeleteGlobalEipTagRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteGlobalEipTagRequest struct{}"
	}

	return strings.Join([]string{"DeleteGlobalEipTagRequest", string(data)}, " ")
}
