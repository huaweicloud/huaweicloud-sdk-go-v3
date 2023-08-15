package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteHostGroupResponse Response Object
type DeleteHostGroupResponse struct {

	// 主机组详细信息
	Result *[]GetHostGroupInfo `json:"result,omitempty"`

	// 删除主机组数量
	Total          *int64 `json:"total,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o DeleteHostGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteHostGroupResponse struct{}"
	}

	return strings.Join([]string{"DeleteHostGroupResponse", string(data)}, " ")
}
