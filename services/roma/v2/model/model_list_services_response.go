package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListServicesResponse struct {

	// 总数
	Total *int32 `json:"total,omitempty"`

	// 本次返回数量
	Size *int32 `json:"size,omitempty"`

	// 服务列表
	Items          *[]Service `json:"items,omitempty"`
	HttpStatusCode int        `json:"-"`
}

func (o ListServicesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListServicesResponse struct{}"
	}

	return strings.Join([]string{"ListServicesResponse", string(data)}, " ")
}
