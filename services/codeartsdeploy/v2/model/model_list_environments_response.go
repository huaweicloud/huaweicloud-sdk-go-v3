package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListEnvironmentsResponse Response Object
type ListEnvironmentsResponse struct {

	// 请求成功失败状态
	Status *string `json:"status,omitempty"`

	// 应用下环境总数
	Total *int32 `json:"total,omitempty"`

	// 环境列表信息
	Result         *[]EnvironmentDetail `json:"result,omitempty"`
	HttpStatusCode int                  `json:"-"`
}

func (o ListEnvironmentsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListEnvironmentsResponse struct{}"
	}

	return strings.Join([]string{"ListEnvironmentsResponse", string(data)}, " ")
}
