package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 集群扩容请求详细描述。
type RoleExtendGrowReq struct {
	// 扩容实例类型。取值为ess、ess-master或ess-client，请至少选择一个。  ess-master、ess-client节点只支持扩容实例个数。

	Type string `json:"type"`
	// 扩容实例个数。集群已有实例个数和扩容实例个数总和不能超过32。

	Nodesize int32 `json:"nodesize"`
	// 扩容实例存储容量。集群原实例存储容量和扩容实例存储容量之和不能超过创建集群时对应的实例存储容量上限，且同一个集群实例存储容量扩容不能超过6次。 单位：GB。

	Disksize int32 `json:"disksize"`
}

func (o RoleExtendGrowReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RoleExtendGrowReq struct{}"
	}

	return strings.Join([]string{"RoleExtendGrowReq", string(data)}, " ")
}
