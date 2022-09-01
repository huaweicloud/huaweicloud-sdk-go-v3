package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type GlanceCreateImageMetadataRequest struct {
	Body *GlanceCreateImageMetadataRequestBody `json:"body,omitempty" xml:"body"`
}

func (o GlanceCreateImageMetadataRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GlanceCreateImageMetadataRequest struct{}"
	}

	return strings.Join([]string{"GlanceCreateImageMetadataRequest", string(data)}, " ")
}
