package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListBriefRecordResponse Response Object
type ListBriefRecordResponse struct {

	// 结果列表
	Result *[]ListBriefRecordResponseBodyResult `json:"result,omitempty"`

	// 错误
	Error *interface{} `json:"error,omitempty"`

	// 状态信息
	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListBriefRecordResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListBriefRecordResponse struct{}"
	}

	return strings.Join([]string{"ListBriefRecordResponse", string(data)}, " ")
}
