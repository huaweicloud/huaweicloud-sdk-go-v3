package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListArticlesResponse struct {

	// 热点/推荐案例列表
	Articles *[]Article `json:"articles,omitempty"`

	// 错误码
	ErrorCode *string `json:"error_code,omitempty"`

	// 错误描述
	ErrorMsg       *string `json:"error_msg,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListArticlesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListArticlesResponse struct{}"
	}

	return strings.Join([]string{"ListArticlesResponse", string(data)}, " ")
}
