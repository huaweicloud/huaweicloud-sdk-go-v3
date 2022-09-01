package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListHostGroupResponse struct {

	// 主机组列表
	Result *[]GetHostGroupInfo `json:"result,omitempty" xml:"result"`

	// 主机组信息总数量
	Total          *int64 `json:"total,omitempty" xml:"total"`
	HttpStatusCode int    `json:"-"`
}

func (o ListHostGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListHostGroupResponse struct{}"
	}

	return strings.Join([]string{"ListHostGroupResponse", string(data)}, " ")
}
