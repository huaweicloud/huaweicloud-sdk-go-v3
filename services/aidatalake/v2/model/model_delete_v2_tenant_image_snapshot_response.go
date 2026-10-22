package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteV2TenantImageSnapshotResponse Response Object
type DeleteV2TenantImageSnapshotResponse struct {

	// **参数解释**：被删除的镜像快照ID。 **取值范围**：长度为32~36的英文字符、数字和中划线的组合。
	ImageSnapshotId *string `json:"image_snapshot_id,omitempty"`
	HttpStatusCode  int     `json:"-"`
}

func (o DeleteV2TenantImageSnapshotResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteV2TenantImageSnapshotResponse struct{}"
	}

	return strings.Join([]string{"DeleteV2TenantImageSnapshotResponse", string(data)}, " ")
}
