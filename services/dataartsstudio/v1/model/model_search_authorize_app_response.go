package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SearchAuthorizeAppResponse Response Object
type SearchAuthorizeAppResponse struct {

	// 符合条件的APP总数
	Total *int32 `json:"total,omitempty"`

	// 本次返回的APP列表
	Records        *[]RecordForGetAuthApp `json:"records,omitempty"`
	HttpStatusCode int                    `json:"-"`
}

func (o SearchAuthorizeAppResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SearchAuthorizeAppResponse struct{}"
	}

	return strings.Join([]string{"SearchAuthorizeAppResponse", string(data)}, " ")
}
