package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Nodes 集群实例对象
type Nodes struct {

	// 集群实例ID
	Id string `json:"id"`

	// 集群实例状态 - 100：创建中 - 199：空闲 - 200：可用 - 300：不可用 - 303：创建失败 - 304：删除中 - 305：删除失败 - 400：已删除
	Status string `json:"status"`
}

func (o Nodes) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Nodes struct{}"
	}

	return strings.Join([]string{"Nodes", string(data)}, " ")
}
