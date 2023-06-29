package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowNextflowResponse Response Object
type ShowNextflowResponse struct {

	// 引擎版本号
	Version *string `json:"version,omitempty"`

	// 用作路径用量，单位：byte
	Workspace *int64 `json:"workspace,omitempty"`

	// 状态
	Status *string `json:"status,omitempty"`

	// 进度
	Progress *float64 `json:"progress,omitempty"`

	// 失败原因
	FailedReason *string `json:"failed_reason,omitempty"`

	// 缓存清理状态
	CacheStatus    *string `json:"cache_status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowNextflowResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowNextflowResponse struct{}"
	}

	return strings.Join([]string{"ShowNextflowResponse", string(data)}, " ")
}
