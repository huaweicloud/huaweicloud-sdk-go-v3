package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SaveAsParameterConfigTemplateResponse Response Object
type SaveAsParameterConfigTemplateResponse struct {

	// 参数组模板id。
	ConfigId       *string `json:"config_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o SaveAsParameterConfigTemplateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SaveAsParameterConfigTemplateResponse struct{}"
	}

	return strings.Join([]string{"SaveAsParameterConfigTemplateResponse", string(data)}, " ")
}
