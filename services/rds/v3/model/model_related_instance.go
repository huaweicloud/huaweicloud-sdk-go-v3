package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 实例节点信息。
type RelatedInstance struct {
	// 关联实例id。

	Id string `json:"id"`
	// 关联实例类型。  - “replica_of”对应于“主实例”。 - “replica”对应于“只读实例”。

	Type string `json:"type"`
}

func (o RelatedInstance) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RelatedInstance struct{}"
	}

	return strings.Join([]string{"RelatedInstance", string(data)}, " ")
}
