package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListWatermarkTemplateResponse Response Object
type ListWatermarkTemplateResponse struct {

	// 总数
	Total *int32 `json:"total,omitempty"`

	// 响应数据
	TemplateInfos *[]WatermarkTemplate `json:"template_infos,omitempty"`

	XRequestId     *string `json:"X-request-id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListWatermarkTemplateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListWatermarkTemplateResponse struct{}"
	}

	return strings.Join([]string{"ListWatermarkTemplateResponse", string(data)}, " ")
}
