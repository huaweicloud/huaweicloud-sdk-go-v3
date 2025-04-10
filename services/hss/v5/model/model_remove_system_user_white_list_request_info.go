package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RemoveSystemUserWhiteListRequestInfo 删除系统用户白名单列表,如果delete_all为true时，data_list可不填
type RemoveSystemUserWhiteListRequestInfo struct {

	// 删除系统用户白名单详情
	DataList *[]SystemUserWhiteListRequestInfo `json:"data_list,omitempty"`

	// 是否删除所有白名单内容
	DeleteAll *bool `json:"delete_all,omitempty"`
}

func (o RemoveSystemUserWhiteListRequestInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RemoveSystemUserWhiteListRequestInfo struct{}"
	}

	return strings.Join([]string{"RemoveSystemUserWhiteListRequestInfo", string(data)}, " ")
}
