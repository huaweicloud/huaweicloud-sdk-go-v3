package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Quotas 创建配额接口请求结构体
type Quotas struct {

	// 资源类型。支持根据资源类型过滤查询指定类型的配额。  - endpoint_service：终端节点服务  - endpoint：终端节点
	Type *string `json:"type,omitempty"`

	// 已创建的资源个数。 取值范围：0~quota数。
	Used *int32 `json:"used,omitempty"`

	// 资源的最大配额数。 取值范围：各类型资源默认配额数的最大值。
	Quota *int32 `json:"quota,omitempty"`
}

func (o Quotas) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Quotas struct{}"
	}

	return strings.Join([]string{"Quotas", string(data)}, " ")
}
