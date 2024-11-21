package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListPacifyWordsResponse Response Object
type ListPacifyWordsResponse struct {

	// 与第一条数据的偏移量
	Offset *int32 `json:"offset,omitempty"`

	// 页面大小
	Limit *int32 `json:"limit,omitempty"`

	// 总数量
	Count *int32 `json:"count,omitempty"`

	// 安抚话术信息
	Data *[]PacifyWordsInfo `json:"data,omitempty"`

	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListPacifyWordsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPacifyWordsResponse struct{}"
	}

	return strings.Join([]string{"ListPacifyWordsResponse", string(data)}, " ")
}
