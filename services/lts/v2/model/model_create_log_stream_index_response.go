package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateLogStreamIndexResponse Response Object
type CreateLogStreamIndexResponse struct {

	// 错误码
	ErrorCode *string `json:"errorCode,omitempty"`

	// 错误信息描述
	ErrorMessage *string `json:"errorMessage,omitempty"`

	// 结果
	Result *string `json:"result,omitempty"`

	// 是否查询完成
	IsQueryComplete *bool `json:"isQueryComplete,omitempty"`
	HttpStatusCode  int   `json:"-"`
}

func (o CreateLogStreamIndexResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateLogStreamIndexResponse struct{}"
	}

	return strings.Join([]string{"CreateLogStreamIndexResponse", string(data)}, " ")
}
