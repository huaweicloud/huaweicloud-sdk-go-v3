package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListProtectedIpResponse Response Object
type ListProtectedIpResponse struct {

	// 总数
	Total *int32 `json:"total,omitempty"`

	// 防护ip列表
	Items          *[]ProtectedIpResponse `json:"items,omitempty"`
	HttpStatusCode int                    `json:"-"`
}

func (o ListProtectedIpResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListProtectedIpResponse struct{}"
	}

	return strings.Join([]string{"ListProtectedIpResponse", string(data)}, " ")
}
