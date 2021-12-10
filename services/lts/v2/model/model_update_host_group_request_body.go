package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 更新主机组请求体
type UpdateHostGroupRequestBody struct {
	// 主机组ID

	HostGroupId string `json:"host_group_id"`
	// 主机组名称

	HostGroupName *string `json:"host_group_name,omitempty"`
	// 主机ID列表。主机类型必须与主机组类型一致

	HostIdList *[]string `json:"host_id_list,omitempty"`
	// 主机组标签。KEY不能重复

	HostGroupTag *[]HostGroupTag `json:"host_group_tag,omitempty"`
}

func (o UpdateHostGroupRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateHostGroupRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateHostGroupRequestBody", string(data)}, " ")
}
