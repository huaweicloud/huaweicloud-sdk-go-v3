package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Snapshots 快照对象列表
type Snapshots struct {

	// 快照ID。
	Id string `json:"id"`

	// 快照名称。
	Name string `json:"name"`

	// 快照描述。
	Description string `json:"description"`

	// 快照创建的日期时间，格式为 ISO8601: YYYY-MM-DDThh:mm:ssZ。
	Started string `json:"started"`

	// 快照完成的日期时间，格式为 ISO8601: YYYY-MM-DDThh:mm:ssZ。
	Finished string `json:"finished"`

	// 快照大小，单位 GB。
	Size float64 `json:"size"`

	// 快照状态：  - CREATING：创建中。 - AVAILABLE：可用。 - UNAVAILABLE：不可用。 - FROZEN： 普通冻结。 - POLICE_FROZEN： 公安冻结。
	Status string `json:"status"`

	// 快照创建类型
	Type string `json:"type"`

	// 快照对应的集群ID
	ClusterId string `json:"cluster_id"`
}

func (o Snapshots) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Snapshots struct{}"
	}

	return strings.Join([]string{"Snapshots", string(data)}, " ")
}
