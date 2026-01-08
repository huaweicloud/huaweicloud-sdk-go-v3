package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Dependency struct {

	// **参数解释**： IP地址组关联的资源类型。 **取值范围**： - sg：IP地址组关联的安全组。 - acl：IP地址组关联的网络ACL。
	Type string `json:"type"`

	// **参数解释**： IP地址组关联资源的ID。 **取值范围**： 带“-”的标准UUID格式。
	InstanceId string `json:"instance_id"`

	// **参数解释**： IP地址组关联资源的名称。 **取值范围**： 1-64个字符，支持数字、字母、中文、_(下划线)、-（中划线）、.（点）。
	InstanceName string `json:"instance_name"`
}

func (o Dependency) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Dependency struct{}"
	}

	return strings.Join([]string{"Dependency", string(data)}, " ")
}
