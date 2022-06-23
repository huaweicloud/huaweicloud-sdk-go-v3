package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ActionReq struct {

	// 集群当前行为。REBOOTING表示重启、GROWING表示扩容、RESTORING表示恢复集群、SNAPSHOTTING表示创建快照。
	Action *string `json:"action,omitempty"`
}

func (o ActionReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ActionReq struct{}"
	}

	return strings.Join([]string{"ActionReq", string(data)}, " ")
}
