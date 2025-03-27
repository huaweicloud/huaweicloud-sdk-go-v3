package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteTemplateResponse Response Object
type DeleteTemplateResponse struct {

	// 删除指定ID的模板成功
	Body           *string `json:"body,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteTemplateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteTemplateResponse struct{}"
	}

	return strings.Join([]string{"DeleteTemplateResponse", string(data)}, " ")
}
