package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateLtsInfoConfigRequestBody struct {

	// 是否开启全量日志   - false: 不开启   - true: 开启
	Enabled *bool `json:"enabled,omitempty"`

	LtsIdInfo *LtsIdInfo `json:"ltsIdInfo,omitempty"`
}

func (o UpdateLtsInfoConfigRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateLtsInfoConfigRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateLtsInfoConfigRequestBody", string(data)}, " ")
}
