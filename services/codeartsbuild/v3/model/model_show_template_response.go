package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowTemplateResponse Response Object
type ShowTemplateResponse struct {
	Result *QueryTemplatesItems `json:"result,omitempty"`

	// 返回错误信息
	Error *string `json:"error,omitempty"`

	// 返回状态信息
	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowTemplateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTemplateResponse struct{}"
	}

	return strings.Join([]string{"ShowTemplateResponse", string(data)}, " ")
}
