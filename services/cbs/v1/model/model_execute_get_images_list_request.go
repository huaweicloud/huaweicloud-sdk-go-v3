package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ExecuteGetImagesListRequest struct {

	// 范围：大于等于1 默认：10
	Limit int32 `json:"limit"`

	// 范围：大于等于0 默认值：0
	Offset int32 `json:"offset"`

	// 横竖屏模式 0：横屏（默认） 1：竖屏
	ResolutionType int32 `json:"resolution_type"`

	// 默认0 0：用户背景 1：系统背景 2：用户图标
	Type int32 `json:"type"`
}

func (o ExecuteGetImagesListRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExecuteGetImagesListRequest struct{}"
	}

	return strings.Join([]string{"ExecuteGetImagesListRequest", string(data)}, " ")
}
