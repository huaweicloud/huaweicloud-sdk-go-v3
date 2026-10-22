package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateAiAppResponse Response Object
type UpdateAiAppResponse struct {

	// AI应用ID
	AppId *string `json:"appId,omitempty"`

	// AI应用名称
	Name *string `json:"name,omitempty"`

	// 应用描述
	Description *string `json:"description,omitempty"`

	// 更新时间戳（毫秒Unix时间戳）
	UpdateTime     *int64 `json:"updateTime,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o UpdateAiAppResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAiAppResponse struct{}"
	}

	return strings.Join([]string{"UpdateAiAppResponse", string(data)}, " ")
}
