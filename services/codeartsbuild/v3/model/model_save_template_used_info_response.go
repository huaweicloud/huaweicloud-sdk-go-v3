package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SaveTemplateUsedInfoResponse Response Object
type SaveTemplateUsedInfoResponse struct {

	// 成功
	Result *string `json:"result,omitempty"`

	// 返回错误信息
	Error *string `json:"error,omitempty"`

	// 返回状态信息
	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o SaveTemplateUsedInfoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SaveTemplateUsedInfoResponse struct{}"
	}

	return strings.Join([]string{"SaveTemplateUsedInfoResponse", string(data)}, " ")
}
