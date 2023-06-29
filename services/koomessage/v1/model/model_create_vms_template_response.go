package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateVmsTemplateResponse Response Object
type CreateVmsTemplateResponse struct {

	// 智能信息基础版模板ID，用来唯一标识上传的模板。
	Tplid          *string `json:"tplid,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateVmsTemplateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateVmsTemplateResponse struct{}"
	}

	return strings.Join([]string{"CreateVmsTemplateResponse", string(data)}, " ")
}
