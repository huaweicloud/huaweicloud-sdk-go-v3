package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAiAppResponse Response Object
type ShowAiAppResponse struct {

	// AI应用ID
	Id *string `json:"id,omitempty"`

	// AI应用名称
	Name *string `json:"name,omitempty"`

	// 应用描述
	Description *string `json:"description,omitempty"`

	// 创建时间戳（毫秒Unix时间戳）
	CreateTime *int64 `json:"createTime,omitempty"`

	// 防护状态（0：禁用，1：启用）
	ProtectStatus *int32 `json:"protectStatus,omitempty"`

	Policy         *Policy `json:"policy,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowAiAppResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAiAppResponse struct{}"
	}

	return strings.Join([]string{"ShowAiAppResponse", string(data)}, " ")
}
