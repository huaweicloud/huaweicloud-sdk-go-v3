package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListStylesResponse Response Object
type ListStylesResponse struct {

	// 风格信息总数
	Count *int32 `json:"count,omitempty"`

	// 风格信息列表
	Styles         *[]StyleInfo `json:"styles,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o ListStylesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListStylesResponse struct{}"
	}

	return strings.Join([]string{"ListStylesResponse", string(data)}, " ")
}
