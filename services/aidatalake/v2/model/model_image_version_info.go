package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ImageVersionInfo **参数解释**：镜像。
type ImageVersionInfo struct {

	// **参数解释**：镜像版本ID。可通过[查询自定义镜像版本列表](ListV2TenantImageVersions.xml)获取镜像版本ID。 **取值范围**：不涉及。 **约束限制**：不涉及。
	Id string `json:"id"`

	// **参数解释**：自定义镜像版本ID。可通过[查询自定义镜像版本列表](ListV2TenantImageVersions.xml)获取镜像版本ID。 **取值范围**：不涉及。 **约束限制**：不涉及。
	VersionId string `json:"version_id"`
}

func (o ImageVersionInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImageVersionInfo struct{}"
	}

	return strings.Join([]string{"ImageVersionInfo", string(data)}, " ")
}
