package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 集群所关联的参数组信息
type ParameterGroup struct {

	// 参数组ID
	Id string `json:"id"`

	// 参数组名称
	Name string `json:"name"`

	// 集群参数状态，有效值包括：  - In-Sync：已同步 - Applying：应用中 - Pending-Reboot：需重启生效 - Sync-Failure：应用失败
	Status string `json:"status"`
}

func (o ParameterGroup) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ParameterGroup struct{}"
	}

	return strings.Join([]string{"ParameterGroup", string(data)}, " ")
}
