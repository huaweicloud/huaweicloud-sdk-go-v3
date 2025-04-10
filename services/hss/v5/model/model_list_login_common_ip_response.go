package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListLoginCommonIpResponse Response Object
type ListLoginCommonIpResponse struct {

	// 常用登录IP已设置的个数
	TotalNum *int32 `json:"total_num,omitempty"`

	// 列表
	DataList       *[]LoginCommonIpResponseInfo `json:"data_list,omitempty"`
	HttpStatusCode int                          `json:"-"`
}

func (o ListLoginCommonIpResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListLoginCommonIpResponse struct{}"
	}

	return strings.Join([]string{"ListLoginCommonIpResponse", string(data)}, " ")
}
