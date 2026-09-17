package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateReleaseReqBody **参数解释：** 创建模板实例的请求体 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
type CreateReleaseReqBody struct {

	// **参数解释：** 模板ID **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	ChartId string `json:"chart_id"`

	// **参数解释：** 模板实例描述 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Description *string `json:"description,omitempty"`

	// **参数解释：** 模板实例名称 **约束限制：** 由小写字母开头，中间由小写字母、数字和中划线(-)组成，以小写字母或数字结尾 **取值范围：** 不涉及 **默认取值：** 不涉及
	Name string `json:"name"`

	// **参数解释：** 模板实例所在的命名空间 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Namespace string `json:"namespace"`

	// **参数解释：** 模板实例版本号 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Version string `json:"version"`

	Parameters *ReleaseReqBodyParams `json:"parameters,omitempty"`

	Values *CreateReleaseReqBodyValues `json:"values"`
}

func (o CreateReleaseReqBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateReleaseReqBody struct{}"
	}

	return strings.Join([]string{"CreateReleaseReqBody", string(data)}, " ")
}
