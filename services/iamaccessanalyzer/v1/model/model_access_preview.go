package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AccessPreview 访问预览。
type AccessPreview struct {

	// 访问预览的唯一标识符。
	AccessPreviewId string `json:"access_preview_id"`

	// 分析器的唯一标识符。
	AnalyzerId string `json:"analyzer_id"`

	Configurations *Configuration `json:"configurations"`

	// 访问预览创建时间。
	CreatedAt *sdktime.SdkTime `json:"created_at"`

	Status *AccessPreviewStatus `json:"status"`

	StatusReason *PreviewStatusReason `json:"status_reason,omitempty"`
}

func (o AccessPreview) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AccessPreview struct{}"
	}

	return strings.Join([]string{"AccessPreview", string(data)}, " ")
}
