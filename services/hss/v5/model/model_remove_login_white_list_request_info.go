package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RemoveLoginWhiteListRequestInfo 删除登录白名单列表,data_list为空时，删除所有；data_list不为空时，请填充里面所有属性
type RemoveLoginWhiteListRequestInfo struct {

	// 删除登录白名单详情
	DataList *[]LoginWhiteListRequestInfo `json:"data_list,omitempty"`

	// 是否删除所有白名单内容
	DeleteAll *bool `json:"delete_all,omitempty"`
}

func (o RemoveLoginWhiteListRequestInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RemoveLoginWhiteListRequestInfo struct{}"
	}

	return strings.Join([]string{"RemoveLoginWhiteListRequestInfo", string(data)}, " ")
}
