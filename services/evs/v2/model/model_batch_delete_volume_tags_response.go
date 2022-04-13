package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type BatchDeleteVolumeTagsResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o BatchDeleteVolumeTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteVolumeTagsResponse struct{}"
	}

	return strings.Join([]string{"BatchDeleteVolumeTagsResponse", string(data)}, " ")
}
