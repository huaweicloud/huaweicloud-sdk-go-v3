package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowImageTemplateListResponse Response Object
type ShowImageTemplateListResponse struct {
	Result *ShowImageTemplateListResponseBodyResult `json:"result,omitempty"`

	// 返回错误信息
	Error *string `json:"error,omitempty"`

	// 返回状态信息
	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowImageTemplateListResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowImageTemplateListResponse struct{}"
	}

	return strings.Join([]string{"ShowImageTemplateListResponse", string(data)}, " ")
}
