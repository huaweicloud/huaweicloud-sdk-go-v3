package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RollingRestartReq struct {

	// 操作角色
	Type string `json:"type"`

	// 实例类型。例如，  - ess-master对应Master节点。 - ess-client对应Clinet节点。 - ess-cold对应冷数据节点。 - ess对应数据节点。 - all对应所有节点。
	Value string `json:"value"`
}

func (o RollingRestartReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RollingRestartReq struct{}"
	}

	return strings.Join([]string{"RollingRestartReq", string(data)}, " ")
}
