package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListVolumeTagsRequest struct {
}

func (o ListVolumeTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListVolumeTagsRequest struct{}"
	}

	return strings.Join([]string{"ListVolumeTagsRequest", string(data)}, " ")
}
