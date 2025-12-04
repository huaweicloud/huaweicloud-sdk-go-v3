package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSourceIpTop5Response Response Object
type ListSourceIpTop5Response struct {

	// 攻击源ip的总数量
	Total *int32 `json:"total,omitempty"`

	// 对象数组
	Items          *[]SourceIpTopListInfoItems `json:"items,omitempty"`
	HttpStatusCode int                         `json:"-"`
}

func (o ListSourceIpTop5Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSourceIpTop5Response struct{}"
	}

	return strings.Join([]string{"ListSourceIpTop5Response", string(data)}, " ")
}
