package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateReleaseReqBodyValues **参数解释：** 模板实例的值 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
type CreateReleaseReqBodyValues struct {

	// **参数解释：** 镜像拉取策略 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	ImagePullPolicy *string `json:"imagePullPolicy,omitempty"`

	// **参数解释：** 镜像标签 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	ImageTag *string `json:"imageTag,omitempty"`
}

func (o CreateReleaseReqBodyValues) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateReleaseReqBodyValues struct{}"
	}

	return strings.Join([]string{"CreateReleaseReqBodyValues", string(data)}, " ")
}
