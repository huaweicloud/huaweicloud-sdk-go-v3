package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListTopAbnormalResponse struct {

	// 异常请求数量
	Total *int32 `json:"total,omitempty"`

	// 异常请求信息数组
	Items          *[]UrlCountItem `json:"items,omitempty"`
	HttpStatusCode int             `json:"-"`
}

func (o ListTopAbnormalResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTopAbnormalResponse struct{}"
	}

	return strings.Join([]string{"ListTopAbnormalResponse", string(data)}, " ")
}
