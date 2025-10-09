package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NodeSpecUpdateNodeNameTemplate **参数解释**： 节点名称固定前后缀配置参数。假设集群内不同节点池被用户所在公司不同部门使用，可以通过前后缀的名称区分部门和使用方式，如设置nodeNamePrefix为finance-，代表部门名称，nodeNameSuffix为-product，代表生产使用，节点池名称为gpu，代表业务类型，则最终的节点名称为finance-gpu(五位随机数)-product **约束限制**： 仅 v1.28.1/v1.27.3/v1.25.6/v1.23.11/v1.21.12 及以上集群支持配置节点名称固定前后缀 **取值范围**： 不涉及 **默认取值**： 不涉及
type NodeSpecUpdateNodeNameTemplate struct {

	// **参数解释**： 节点名称前缀。如果用户填写为空串或缺省，则节点名称不会增加前缀。 **约束限制**： 仅支持小写字母、数字、连字符（-）和点（.），必须以小写字母开头，并且符合[FRC 1123](https://tools.ietf.org/html/rfc1123)中定义的DNS子域名命名规范。 **取值范围**： 不涉及 **默认取值**： 不涉及
	NodeNamePrefix *string `json:"nodeNamePrefix,omitempty"`

	// **参数解释**： 节点名称后缀。如果用户填写为空串或缺省，则节点名称不会增加后缀。 **约束限制**： 仅支持小写字母、数字、连字符（-）和点（.），后缀结尾必须为小写字母或者数字，并且符合[FRC 1123](https://tools.ietf.org/html/rfc1123)中定义的DNS子域名命名规范。 **取值范围**： 不涉及 **默认取值**： 不涉及
	NodeNameSuffix *string `json:"nodeNameSuffix,omitempty"`
}

func (o NodeSpecUpdateNodeNameTemplate) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NodeSpecUpdateNodeNameTemplate struct{}"
	}

	return strings.Join([]string{"NodeSpecUpdateNodeNameTemplate", string(data)}, " ")
}
