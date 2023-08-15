package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateTemplateResponse Response Object
type CreateTemplateResponse struct {

	// 服务端返回的新添加的模板的ID
	Id             *string `json:"id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateTemplateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateTemplateResponse struct{}"
	}

	return strings.Join([]string{"CreateTemplateResponse", string(data)}, " ")
}
