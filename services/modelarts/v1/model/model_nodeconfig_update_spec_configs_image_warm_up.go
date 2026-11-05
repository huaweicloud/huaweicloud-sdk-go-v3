package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type NodeconfigUpdateSpecConfigsImageWarmUp struct {

	// **参数解释**：镜像预热任务ID。 **约束限制**：选填。新增任务时无需填写，系统自动生成；格式为image-warm-up-task-<uuid>。删除已有任务时必须填写，通过id标识区分。 **取值范围**：不涉及。 **默认取值**：不涉及。
	Id *string `json:"id,omitempty"`

	// **参数解释**：镜像预热任务名称。 **约束限制**：必填。由4~32位小写字母、数字或\"-\"组成，以字母开头、以字母或数字结尾；同一请求内任务名称不可重复。 **取值范围**：不涉及。 **默认取值**：不涉及。
	Name *string `json:"name,omitempty"`

	// **参数解释**：镜像预热任务的创建时间。 **约束限制**：选填。新增任务时无需填写，系统自动填充为当前时间。 **取值范围**：RFC3339格式时间，例如2026-08-14T19:53:27+08:00。 **默认取值**：不涉及。
	CreationTimestamp *string `json:"creationTimestamp,omitempty"`

	// **参数解释**：镜像类型。 **约束限制**：必填。指定镜像的来源类型。 **取值范围**：可选值如下： - SWR：华为云容器镜像服务。 **默认取值**：不涉及。
	Type *NodeconfigUpdateSpecConfigsImageWarmUpType `json:"type,omitempty"`

	// **参数解释**：需要预热的镜像地址，例如swr.xxx.com/namespace/repository:v1.0。 **约束限制**：必填。同一请求内镜像地址不可重复。 **取值范围**：不涉及。 **默认取值**：不涉及。
	Image *string `json:"image,omitempty"`
}

func (o NodeconfigUpdateSpecConfigsImageWarmUp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NodeconfigUpdateSpecConfigsImageWarmUp struct{}"
	}

	return strings.Join([]string{"NodeconfigUpdateSpecConfigsImageWarmUp", string(data)}, " ")
}

type NodeconfigUpdateSpecConfigsImageWarmUpType struct {
	value string
}

type NodeconfigUpdateSpecConfigsImageWarmUpTypeEnum struct {
	SWR NodeconfigUpdateSpecConfigsImageWarmUpType
}

func GetNodeconfigUpdateSpecConfigsImageWarmUpTypeEnum() NodeconfigUpdateSpecConfigsImageWarmUpTypeEnum {
	return NodeconfigUpdateSpecConfigsImageWarmUpTypeEnum{
		SWR: NodeconfigUpdateSpecConfigsImageWarmUpType{
			value: "SWR",
		},
	}
}

func (c NodeconfigUpdateSpecConfigsImageWarmUpType) Value() string {
	return c.value
}

func (c NodeconfigUpdateSpecConfigsImageWarmUpType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *NodeconfigUpdateSpecConfigsImageWarmUpType) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: string")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(string); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to string error")
	}
}
