package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteGeipSegmentTagRequest Request Object
type DeleteGeipSegmentTagRequest struct {
	ResourceId string `json:"resource_id"`

	TagKey string `json:"tag_key"`
}

func (o DeleteGeipSegmentTagRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteGeipSegmentTagRequest struct{}"
	}

	return strings.Join([]string{"DeleteGeipSegmentTagRequest", string(data)}, " ")
}
