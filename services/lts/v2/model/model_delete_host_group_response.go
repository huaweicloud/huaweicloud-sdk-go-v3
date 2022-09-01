package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type DeleteHostGroupResponse struct {

	// 主机组详细信息
	Result *[]GetHostGroupInfo `json:"result,omitempty" xml:"result"`

	// 删除主机组数量
	Total          *int64 `json:"total,omitempty" xml:"total"`
	HttpStatusCode int    `json:"-"`
}

func (o DeleteHostGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteHostGroupResponse struct{}"
	}

	return strings.Join([]string{"DeleteHostGroupResponse", string(data)}, " ")
}
