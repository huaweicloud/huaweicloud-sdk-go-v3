package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type CreateAppResponse struct {

	// 应用模板ID
	AppId *string `json:"app_id,omitempty"`

	// 应用描述
	Description *string `json:"description,omitempty"`

	// 应用类型
	AppType *string `json:"app_type,omitempty"`

	// 应用来源
	ProviderType *string `json:"provider_type,omitempty"`

	// 创建时间
	CreateTime *string `json:"create_time,omitempty"`

	// 最后一次修改时间
	UpdateTime     *string `json:"update_time,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateAppResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAppResponse struct{}"
	}

	return strings.Join([]string{"CreateAppResponse", string(data)}, " ")
}
