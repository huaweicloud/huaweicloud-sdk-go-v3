package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateLtsConfigRequestBody 更新日志配置请求体
type UpdateLtsConfigRequestBody struct {

	// 日志配置开关
	Enabled bool `json:"enabled"`

	LtsIdInfo *UpdateLtsConfigRequestBodyLtsIdInfo `json:"lts_id_info,omitempty"`
}

func (o UpdateLtsConfigRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateLtsConfigRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateLtsConfigRequestBody", string(data)}, " ")
}
