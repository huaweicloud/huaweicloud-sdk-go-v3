package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListLoginCommonLocationResponse Response Object
type ListLoginCommonLocationResponse struct {

	// 常用登录地已设置的个数
	TotalNum *int32 `json:"total_num,omitempty"`

	// 返回常用登陆地信息列表
	DataList       *[]LoginCommonLocationResponseInfo `json:"data_list,omitempty"`
	HttpStatusCode int                                `json:"-"`
}

func (o ListLoginCommonLocationResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListLoginCommonLocationResponse struct{}"
	}

	return strings.Join([]string{"ListLoginCommonLocationResponse", string(data)}, " ")
}
