package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateInstanceReq struct {

	// **参数解释**： 实例名称。 **约束限制**： 由英文字符开头，只能由英文字母、数字、中划线组成，长度为4~64的字符。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Name *string `json:"name,omitempty"`

	// **参数解释**： 实例的描述信息。 **约束限制**： 长度不超过1024的字符串。[且字符串不能包含\">\"与\"<\"，字符串首字符不能为\"=\",\"+\",\"-\",\"@\"的全角和半角字符。](tag:hcs)  \\与\"在json报文中属于特殊字符，如果参数值中需要显示\\或者\"字符，请在字符前增加转义字符\\，比如\\\\\\\\或者\\\\\"。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Description *string `json:"description,omitempty"`

	// **参数解释**： 安全组ID。 获取方法如下：参考《虚拟私有云 API参考》，调用“查询安全组列表”接口，从响应体中获取安全组ID。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	SecurityGroupId *string `json:"security_group_id,omitempty"`

	// **参数解释**： 是否开启ACL访问控制。 **约束限制**： 不涉及。 **取值范围**： - true：开启ACL访问控制。 - false：不开启ACL访问控制。 **默认取值**： 不涉及。
	EnableAcl *bool `json:"enable_acl,omitempty"`

	// **参数解释**： 是否开启公网。 **约束限制**： 不涉及。 **取值范围**： - true：开启公网。 - false：不开启公网。 **默认取值**： 不涉及。
	EnablePublicip *bool `json:"enable_publicip,omitempty"`

	// **参数解释**： 实例绑定的弹性IP地址的ID。 **约束限制**： 以英文逗号隔开多个弹性IP地址的ID。 如果开启了公网访问功能（即enable_publicip为true），该字段为必选。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	PublicipId *string `json:"publicip_id,omitempty"`

	// **参数解释**： 企业项目。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`
}

func (o UpdateInstanceReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateInstanceReq struct{}"
	}

	return strings.Join([]string{"UpdateInstanceReq", string(data)}, " ")
}
