package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type GetHostGroupListRequestBody struct {
	// 主机组ID

	HostGroupIdList *[]string `json:"host_group_id_list,omitempty"`

	Filter *GetHostGroupListFilter `json:"filter,omitempty"`
}

func (o GetHostGroupListRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GetHostGroupListRequestBody struct{}"
	}

	return strings.Join([]string{"GetHostGroupListRequestBody", string(data)}, " ")
}
