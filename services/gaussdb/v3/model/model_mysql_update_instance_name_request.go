package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type MysqlUpdateInstanceNameRequest struct {

	// 实例名称。  用于表示实例的名称，同一租户下，同类型的实例名可重名。取值范围：最小为4个字符，最大为64个字符且不超过64个字节（注意：一个中文字符占用3个字节），必须以字母或中文开头，区分大小写，可以包含字母、数字、中划线、下划线或中文，不能包含其他特殊字符。
	Name string `json:"name"`

	// 是否同步修改节点名称，取值：true或false, 默认值为true。
	IsModifyNodeName *string `json:"is_modify_node_name,omitempty"`
}

func (o MysqlUpdateInstanceNameRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MysqlUpdateInstanceNameRequest struct{}"
	}

	return strings.Join([]string{"MysqlUpdateInstanceNameRequest", string(data)}, " ")
}
