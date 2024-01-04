package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type MysqlCreateBackupRequest struct {

	// 实例ID，严格匹配UUID规则。
	InstanceId string `json:"instance_id"`

	// 备份名称。  取值范围：最小为4个字符，最大为64个字符且不超过64个字节（注意：一个中文字符占用3个字节），必须以字母或中文开头，区分大小写，可以包含字母、数字、中划线、下划线或中文，不能包含其他特殊字符。
	Name string `json:"name"`

	// 备份描述，不能包含>!<\"&'=特殊字符，不大于256个字符。
	Description *string `json:"description,omitempty"`
}

func (o MysqlCreateBackupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MysqlCreateBackupRequest struct{}"
	}

	return strings.Join([]string{"MysqlCreateBackupRequest", string(data)}, " ")
}
