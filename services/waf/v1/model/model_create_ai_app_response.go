package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateAiAppResponse Response Object
type CreateAiAppResponse struct {

	// AI应用ID
	Id *string `json:"id,omitempty"`

	// AI应用名称
	Name *string `json:"name,omitempty"`

	// 应用描述
	Description *string `json:"description,omitempty"`

	// 创建时间戳（毫秒Unix时间戳）
	CreateTime *int64 `json:"createTime,omitempty"`

	// API签名
	ApiSign        *string `json:"apiSign,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateAiAppResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAiAppResponse struct{}"
	}

	return strings.Join([]string{"CreateAiAppResponse", string(data)}, " ")
}
