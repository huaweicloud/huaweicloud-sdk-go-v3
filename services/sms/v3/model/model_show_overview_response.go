package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowOverviewResponse Response Object
type ShowOverviewResponse struct {

	// 等待中
	Waiting *int32 `json:"waiting,omitempty"`

	// 复制中
	Replicate *int32 `json:"replicate,omitempty"`

	// 同步中
	Syncing *int32 `json:"syncing,omitempty"`

	// 已暂停
	Stopped *int32 `json:"stopped,omitempty"`

	// 删除中
	Deleting *int32 `json:"deleting,omitempty"`

	// 启动目的端中
	Cutovering *int32 `json:"cutovering,omitempty"`

	// 环境校验不通过
	Unavailable *int32 `json:"unavailable,omitempty"`

	// 暂停中
	Stopping *int32 `json:"stopping,omitempty"`

	// 跳过中
	Skipping *int32 `json:"skipping,omitempty"`

	// 启动目的端完成
	Finished *int32 `json:"finished,omitempty"`

	// 初始化
	Initialize *int32 `json:"initialize,omitempty"`

	// 错误
	Error *int32 `json:"error,omitempty"`

	// 等待克隆完成
	Cloning *int32 `json:"cloning,omitempty"`

	// 未配置目的端
	Unconfigured   *int32 `json:"unconfigured,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowOverviewResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowOverviewResponse struct{}"
	}

	return strings.Join([]string{"ShowOverviewResponse", string(data)}, " ")
}
