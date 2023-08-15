package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowQaPairsResponse Response Object
type ShowQaPairsResponse struct {

	// 错误码
	ErrorCode *string `json:"error_code,omitempty"`

	// 错误描述
	ErrorMsg *string `json:"error_msg,omitempty"`

	// 语料列表
	QaPairs        *[]SimpleQaPair `json:"qa_pairs,omitempty"`
	HttpStatusCode int             `json:"-"`
}

func (o ShowQaPairsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowQaPairsResponse struct{}"
	}

	return strings.Join([]string{"ShowQaPairsResponse", string(data)}, " ")
}
