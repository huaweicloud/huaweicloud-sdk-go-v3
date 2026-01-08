package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AddressGroupDependency struct {

	// **参数解释**： IP地址组的资源ID。IP地址组创建成功后，会生成一个IP地址组 ID，是IP地址组对应的唯一标识。 **取值范围**： 带“-”的标准UUID格式。
	Id string `json:"id"`

	// **参数解释**： IP地址组所属的企业项目ID。 **取值范围**： 最大长度36字节，带“-”连字符的UUID格式，或者是字符串“0”。“0”表示默认企业项目。
	EnterpriseProjectId string `json:"enterprise_project_id"`

	// **参数解释**： IP地址组关联的资源列表。 **取值范围**： 不涉及。
	Dependency []Dependency `json:"dependency"`
}

func (o AddressGroupDependency) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddressGroupDependency struct{}"
	}

	return strings.Join([]string{"AddressGroupDependency", string(data)}, " ")
}
