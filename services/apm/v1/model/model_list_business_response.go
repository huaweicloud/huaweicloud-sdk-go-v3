package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListBusinessResponse struct {

	// 获取应用列表数据结构。
	BusinessNodes  *[]BusinessNodeModel `json:"business_nodes,omitempty"`
	HttpStatusCode int                  `json:"-"`
}

func (o ListBusinessResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListBusinessResponse struct{}"
	}

	return strings.Join([]string{"ListBusinessResponse", string(data)}, " ")
}
