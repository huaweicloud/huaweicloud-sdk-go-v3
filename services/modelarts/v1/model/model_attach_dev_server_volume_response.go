package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AttachDevServerVolumeResponse Response Object
type AttachDevServerVolumeResponse struct {

	// **参数解释**：操作ID。 **取值范围**：^[0-9a-f]{8}-[0-9a-f]{4}-[1-5][0-9a-f]{3}-[89ab][0-9a-f]{3}-[0-9a-f]{12}$。
	OperationId *string `json:"operation_id,omitempty"`

	// **参数解释**：操作状态。 **取值范围**： - pending 等待处理 - running 运行中 - success 成功 - failed 失败
	OperationStatus *string `json:"operation_status,omitempty"`

	// **参数解释**：操作类型。 **取值范围**：node_attach_volume
	OperationType *string `json:"operation_type,omitempty"`

	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o AttachDevServerVolumeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AttachDevServerVolumeResponse struct{}"
	}

	return strings.Join([]string{"AttachDevServerVolumeResponse", string(data)}, " ")
}
