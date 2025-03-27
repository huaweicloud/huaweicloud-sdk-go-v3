package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteTemplatesResponse Response Object
type DeleteTemplatesResponse struct {

	// 批量删除指定ID的模板成功
	Body           *string `json:"body,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteTemplatesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteTemplatesResponse struct{}"
	}

	return strings.Join([]string{"DeleteTemplatesResponse", string(data)}, " ")
}
