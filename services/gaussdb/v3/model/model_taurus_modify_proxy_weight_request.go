package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TaurusModifyProxyWeightRequest proxy实例修改权重请求体
type TaurusModifyProxyWeightRequest struct {

	// 主节点权重
	MasterWeight *int32 `json:"master_weight,omitempty"`

	// 只读节点权重配置信息
	ReadonlyNodes *[]ModifyProxyWeightReadonlyNode `json:"readonly_nodes,omitempty"`
}

func (o TaurusModifyProxyWeightRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TaurusModifyProxyWeightRequest struct{}"
	}

	return strings.Join([]string{"TaurusModifyProxyWeightRequest", string(data)}, " ")
}
