package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type GetHostGroupListRequestBody struct {
	// 主机组ID

	HostGroupIdList []string `json:"host_group_id_list"`

	Filter *GetHostGroupListFilter `json:"filter"`
}

func (o GetHostGroupListRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GetHostGroupListRequestBody struct{}"
	}

	return strings.Join([]string{"GetHostGroupListRequestBody", string(data)}, " ")
}
