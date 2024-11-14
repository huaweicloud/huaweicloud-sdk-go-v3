package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateBeautyPreviewJobReq 创建美白预览任务请求。
type CreateBeautyPreviewJobReq struct {

	// 美白预览任务名称。
	Name string `json:"name"`

	// 美白等级。默认1级。
	Level *int32 `json:"level,omitempty"`
}

func (o CreateBeautyPreviewJobReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateBeautyPreviewJobReq struct{}"
	}

	return strings.Join([]string{"CreateBeautyPreviewJobReq", string(data)}, " ")
}
