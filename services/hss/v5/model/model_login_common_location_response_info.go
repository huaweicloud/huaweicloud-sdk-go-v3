package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// LoginCommonLocationResponseInfo 请求体云主机列表
type LoginCommonLocationResponseInfo struct {

	// 国家城市的编码
	AreaCode *int32 `json:"area_code,omitempty"`

	// 这个常用登陆地的主机个数
	TotalNum *int32 `json:"total_num,omitempty"`

	// 服务器列表
	HostIdList *[]string `json:"host_id_list,omitempty"`
}

func (o LoginCommonLocationResponseInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LoginCommonLocationResponseInfo struct{}"
	}

	return strings.Join([]string{"LoginCommonLocationResponseInfo", string(data)}, " ")
}
