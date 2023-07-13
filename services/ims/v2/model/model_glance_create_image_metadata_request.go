package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// GlanceCreateImageMetadataRequest Request Object
type GlanceCreateImageMetadataRequest struct {
	Body *GlanceCreateImageMetadataRequestBody `json:"body,omitempty"`
}

func (o GlanceCreateImageMetadataRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GlanceCreateImageMetadataRequest struct{}"
	}

	return strings.Join([]string{"GlanceCreateImageMetadataRequest", string(data)}, " ")
}
