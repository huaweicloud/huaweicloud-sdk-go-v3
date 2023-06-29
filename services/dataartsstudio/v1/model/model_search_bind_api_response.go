package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SearchBindApiResponse Response Object
type SearchBindApiResponse struct {

	// 应用已绑定的api总个数
	Total *int32 `json:"total,omitempty"`

	// 应用已绑定的api列表
	Records        *[]AppBindApiInfo `json:"records,omitempty"`
	HttpStatusCode int               `json:"-"`
}

func (o SearchBindApiResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SearchBindApiResponse struct{}"
	}

	return strings.Join([]string{"SearchBindApiResponse", string(data)}, " ")
}
