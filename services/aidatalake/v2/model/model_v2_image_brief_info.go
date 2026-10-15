package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// V2ImageBriefInfo **参数解释**：镜像信息。 **约束限制**：仅提交作业到RayJob端点时配置。
type V2ImageBriefInfo struct {

	// **参数解释**：镜像ID。获取方式请参见[查看镜像列表](ListV2TenantImages.xml)。 **约束限制**：不涉及。 **取值范围**：长度为32~36的英文、数字、中划线的组合。 **默认取值**：不涉及。
	ImageId string `json:"image_id"`

	// **参数解释**：镜像版本ID。获取方式请参见[查询镜像详情](ShowV2TenantImage.xml)。 **约束限制**：不涉及。 **取值范围**：长度为32~36的英文、数字、中划线的组合。 **默认取值**：不涉及。
	ImageVersionId string `json:"image_version_id"`
}

func (o V2ImageBriefInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "V2ImageBriefInfo struct{}"
	}

	return strings.Join([]string{"V2ImageBriefInfo", string(data)}, " ")
}
