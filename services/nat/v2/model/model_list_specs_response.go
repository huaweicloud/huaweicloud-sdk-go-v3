package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSpecsResponse Response Object
type ListSpecsResponse struct {

	// 查询项目支持网关规格列表的响应体。
	Specs          *[]Spec `json:"specs,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListSpecsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSpecsResponse struct{}"
	}

	return strings.Join([]string{"ListSpecsResponse", string(data)}, " ")
}
