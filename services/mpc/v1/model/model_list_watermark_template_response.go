package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListWatermarkTemplateResponse struct {

	// 水印模板总数。
	Total *int32 `json:"total,omitempty" xml:"total"`

	// 水印模板
	Templates      *[]WatermarkTemplate `json:"templates,omitempty" xml:"templates"`
	HttpStatusCode int                  `json:"-"`
}

func (o ListWatermarkTemplateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListWatermarkTemplateResponse struct{}"
	}

	return strings.Join([]string{"ListWatermarkTemplateResponse", string(data)}, " ")
}
