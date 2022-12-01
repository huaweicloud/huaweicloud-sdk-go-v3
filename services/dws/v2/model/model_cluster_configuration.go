package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 集群所关联的参数组详情。
type ClusterConfiguration struct {

	// 参数组ID。
	Id string `json:"id"`

	// 参数组名称。
	Name string `json:"name"`

	// 参数组类型。
	Type string `json:"type"`

	// 集群参数状态。 - In-Sync：已同步。 - Applying：应用中。 - Pending-Reboot：需重启生效。 - Sync-Failure：应用失败。
	Status string `json:"status"`

	// 参数应用失败原因。
	FailReason string `json:"fail_reason"`
}

func (o ClusterConfiguration) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ClusterConfiguration struct{}"
	}

	return strings.Join([]string{"ClusterConfiguration", string(data)}, " ")
}
