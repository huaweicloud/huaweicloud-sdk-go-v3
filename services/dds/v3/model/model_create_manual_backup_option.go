package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 备份参数对象。
type CreateManualBackupOption struct {

	// 实例ID，可以调用“查询实例列表和详情”接口获取。如果未申请实例，可以调用“创建实例”接口创建。
	InstanceId string `json:"instance_id" xml:"instance_id"`

	// 手动备份名称。 取值范围：长度为4~64位，必须以字母开头（A~Z或a~z），区分大小写，可以包含字母、数字（0~9）、中划线（-）或者下划线（_），不能包含其他特殊字符。
	Name string `json:"name" xml:"name"`

	// 手动备份描述。 取值范围：长度不超过256位，且不能包含>!<\"&'=特殊字符。
	Description *string `json:"description,omitempty" xml:"description"`
}

func (o CreateManualBackupOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateManualBackupOption struct{}"
	}

	return strings.Join([]string{"CreateManualBackupOption", string(data)}, " ")
}
