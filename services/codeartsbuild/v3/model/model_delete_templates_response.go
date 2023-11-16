package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteTemplatesResponse Response Object
type DeleteTemplatesResponse struct {

	// 查询模板结果
	Result *[]CreateTemplatesItems `json:"result,omitempty"`

	// 返回错误信息
	Error *string `json:"error,omitempty"`

	// 返回状态信息
	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteTemplatesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteTemplatesResponse struct{}"
	}

	return strings.Join([]string{"DeleteTemplatesResponse", string(data)}, " ")
}
