package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAppVersionResponse Response Object
type ShowAppVersionResponse struct {

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

func (o ShowAppVersionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAppVersionResponse struct{}"
	}

	return strings.Join([]string{"ShowAppVersionResponse", string(data)}, " ")
}
