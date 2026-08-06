package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreatePoolRequestBody 创建实例组请求体
type CreatePoolRequestBody struct {

	// **参数解释：** 实例组名称，用于标识实例组，便于管理和识别。 **约束限制：** 不涉及 **取值范围：** 只能由英文字母、数字、下划线、中划线和点组成，且长度为1~256个字符 **默认取值：** 不涉及
	Name string `json:"name"`

	// **参数解释：** 实例组所在的区域（Region）。 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Region string `json:"region"`

	// **参数解释：** 实例组类型 **约束限制：** 不涉及 **取值范围：** - elb: 基础elb类型 - elb-v2: elb-v2类型 - elb-shadow: saas化elb类型 - standard-container: 反向代理独享引擎组（云内，承载租户专用） - standard-cloud: 反向代理独享引擎组（云内） - standard: 反向代理独享引擎组（云外） - detector-cloud: 旁路检测独享引擎组（云内） - detector: 旁路检测独享引擎组（云外） **默认取值：** 不涉及
	Type string `json:"type"`

	// **参数解释：** 实例组关联的VPC ID（通过调用虚拟私有云ListVpcs接口获取所有的VPC列表查询VPC的ID） **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	VpcId string `json:"vpc_id"`

	// **参数解释：** 实例组描述 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Description *string `json:"description,omitempty"`
}

func (o CreatePoolRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePoolRequestBody struct{}"
	}

	return strings.Join([]string{"CreatePoolRequestBody", string(data)}, " ")
}
