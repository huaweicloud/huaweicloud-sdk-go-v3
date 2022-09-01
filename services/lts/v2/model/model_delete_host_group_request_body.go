package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 删除主机组请求体
type DeleteHostGroupRequestBody struct {

	// 主机组ID列表
	HostGroupIdList []string `json:"host_group_id_list" xml:"host_group_id_list"`
}

func (o DeleteHostGroupRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteHostGroupRequestBody struct{}"
	}

	return strings.Join([]string{"DeleteHostGroupRequestBody", string(data)}, " ")
}
