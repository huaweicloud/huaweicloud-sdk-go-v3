package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExtendClusterGrowReq 集群扩容请求详细描述。
type ExtendClusterGrowReq struct {

	// 扩容实例个数。  集群已有实例个数和扩容实例个数总和不能超过32。
	ModifySize int32 `json:"modifySize"`
}

func (o ExtendClusterGrowReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExtendClusterGrowReq struct{}"
	}

	return strings.Join([]string{"ExtendClusterGrowReq", string(data)}, " ")
}
