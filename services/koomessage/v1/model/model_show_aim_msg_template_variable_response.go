package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAimMsgTemplateVariableResponse Response Object
type ShowAimMsgTemplateVariableResponse struct {

	// 返回结果列表。
	Results *[]VariableAttributesResponse `json:"results,omitempty"`

	// 总数。
	Total          *int32 `json:"total,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowAimMsgTemplateVariableResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAimMsgTemplateVariableResponse struct{}"
	}

	return strings.Join([]string{"ShowAimMsgTemplateVariableResponse", string(data)}, " ")
}
