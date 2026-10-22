package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteV2TenantImageSnapshotRequest Request Object
type DeleteV2TenantImageSnapshotRequest struct {

	// **参数解释**：镜像ID。可通过[查询自定义镜像列表]获取镜像ID(ListV2TenantImages.xml)。 **约束限制**：不涉及。 **取值范围**：长度为32~36的英文字符、数字和中划线的组合。 **默认取值**：不涉及。
	ImageId string `json:"image_id"`

	// **参数解释**：镜像版本ID。可通过[查询自定义镜像版本列表](ListV2TenantImageVersions.xml)获取镜像版本ID。 **约束限制**：不涉及。 **取值范围**：长度为32~36的英文字符、数字和中划线的组合。 **默认取值**：不涉及。
	VersionId string `json:"version_id"`

	// **参数解释**：镜像快照ID。可通过[查询镜像快照列表]获取镜像ID(ListV2TenantImageSnapshots.xml)。 **约束限制**：不涉及。 **取值范围**：长度为32~36的英文字符、数字和中划线的组合。 **默认取值**：不涉及。
	ImageSnapshotId string `json:"image_snapshot_id"`
}

func (o DeleteV2TenantImageSnapshotRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteV2TenantImageSnapshotRequest struct{}"
	}

	return strings.Join([]string{"DeleteV2TenantImageSnapshotRequest", string(data)}, " ")
}
