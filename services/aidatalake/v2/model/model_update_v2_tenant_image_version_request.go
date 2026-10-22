package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateV2TenantImageVersionRequest Request Object
type UpdateV2TenantImageVersionRequest struct {

	// **参数解释**：镜像ID。可通过[查询自定义镜像列表]获取镜像ID(ListV2TenantImages.xml)。 **约束限制**：不涉及。 **取值范围**：长度为32~36的英文字符、数字和中划线的组合。 **默认取值**：不涉及。
	ImageId string `json:"image_id"`

	// **参数解释**：镜像版本ID。可通过[查询自定义镜像版本列表](ListV2TenantImageVersions.xml)获取镜像版本ID。 **约束限制**：不涉及。 **取值范围**：长度为32~36的英文字符、数字和中划线的组合。 **默认取值**：不涉及。
	VersionId string `json:"version_id"`

	Body *UpdateV2ImageVersionInput `json:"body,omitempty"`
}

func (o UpdateV2TenantImageVersionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateV2TenantImageVersionRequest struct{}"
	}

	return strings.Join([]string{"UpdateV2TenantImageVersionRequest", string(data)}, " ")
}
