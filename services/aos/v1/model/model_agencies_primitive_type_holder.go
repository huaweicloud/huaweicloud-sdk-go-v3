package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AgenciesPrimitiveTypeHolder struct {

	// 委托授权的信息。  RFS仅在创建资源栈（触发部署）、创建执行计划、部署资源栈、删除资源栈等涉及资源操作的请求中使用委托，且该委托仅作用于与之绑定的Provider对资源的操作中。若委托中提供的权限不足，有可能导致相关资源操作失败。
	Agencies *[]Agency `json:"agencies,omitempty"`
}

func (o AgenciesPrimitiveTypeHolder) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AgenciesPrimitiveTypeHolder struct{}"
	}

	return strings.Join([]string{"AgenciesPrimitiveTypeHolder", string(data)}, " ")
}
