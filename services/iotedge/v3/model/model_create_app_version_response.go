package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateAppVersionResponse Response Object
type CreateAppVersionResponse struct {

	// 应用模板ID
	AppId *string `json:"app_id,omitempty"`

	// 应用版本
	Version *string `json:"version,omitempty"`

	// 应用版本配置
	Values *interface{} `json:"values,omitempty"`

	// 创建时间
	CreateTime *string `json:"create_time,omitempty"`

	// 最后一次修改时间
	UpdateTime     *string `json:"update_time,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateAppVersionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAppVersionResponse struct{}"
	}

	return strings.Join([]string{"CreateAppVersionResponse", string(data)}, " ")
}
