package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowDatabaseWaterMarkRequest struct {
	Body *ExtractedDatabaseWatermark `json:"body,omitempty" xml:"body"`
}

func (o ShowDatabaseWaterMarkRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDatabaseWaterMarkRequest struct{}"
	}

	return strings.Join([]string{"ShowDatabaseWaterMarkRequest", string(data)}, " ")
}
