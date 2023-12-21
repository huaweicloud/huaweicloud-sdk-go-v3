package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DemandInstance 租户需求
type DemandInstance struct {

	// 弹性公网IP池。 多线路场景下，将在该弹性公网IP池下创建弹性公网IP。 > 覆盖规则为省级/大区时不支持指定线路ID创建边缘业务。
	PoolId *string `json:"pool_id,omitempty"`

	// 线路ID集合。多线路场景下，将在各线路下创建弹性公网IP。
	PoolIds *[]string `json:"pool_ids,omitempty"`
}

func (o DemandInstance) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DemandInstance struct{}"
	}

	return strings.Join([]string{"DemandInstance", string(data)}, " ")
}
