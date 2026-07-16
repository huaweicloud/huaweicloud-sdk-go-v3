package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AttachDevServerVolumeRequest Request Object
type AttachDevServerVolumeRequest struct {

	// **参数解释**：Lite Server实例ID。 **约束限制**：不涉及。 **取值范围**：^[0-9a-f]{8}-[0-9a-f]{4}-[1-5][0-9a-f]{3}-[89ab][0-9a-f]{3}-[0-9a-f]{12}$。 **默认取值**：不涉及。
	Id string `json:"id"`

	Body *AttachServerVolumeRequest `json:"body,omitempty"`
}

func (o AttachDevServerVolumeRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AttachDevServerVolumeRequest struct{}"
	}

	return strings.Join([]string{"AttachDevServerVolumeRequest", string(data)}, " ")
}
