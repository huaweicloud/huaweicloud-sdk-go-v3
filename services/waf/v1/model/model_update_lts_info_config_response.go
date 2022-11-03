package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type UpdateLtsInfoConfigResponse struct {

	// lts配置信息id
	Id *string `json:"id,omitempty"`

	// 是否开启全量日志   - false: 不开启   - true: 开启
	Enabale *bool `json:"enabale,omitempty"`

	LtsIdInfo      *LtsIdInfo `json:"ltsIdInfo,omitempty"`
	HttpStatusCode int        `json:"-"`
}

func (o UpdateLtsInfoConfigResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateLtsInfoConfigResponse struct{}"
	}

	return strings.Join([]string{"UpdateLtsInfoConfigResponse", string(data)}, " ")
}
