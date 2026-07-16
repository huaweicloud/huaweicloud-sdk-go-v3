package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AttachServerVolumeRequest Lite Server服务器挂载磁盘ID
type AttachServerVolumeRequest struct {

	// **参数解释**：待挂载磁盘的磁盘ID。 **约束限制**：不涉及。 **取值范围**：^[0-9a-f]{8}-[0-9a-f]{4}-[1-5][0-9a-f]{3}-[89ab][0-9a-f]{3}-[0-9a-f]{12}$。 **默认取值**：不涉及。
	VolumeId string `json:"volume_id"`
}

func (o AttachServerVolumeRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AttachServerVolumeRequest struct{}"
	}

	return strings.Join([]string{"AttachServerVolumeRequest", string(data)}, " ")
}
