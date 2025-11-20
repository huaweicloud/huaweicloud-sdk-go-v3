package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UnbindVportRequest Request Object
type UnbindVportRequest struct {

	// - 参数解释：二层连接的唯一标识。 - 约束限制：带“-”的UUID格式。 - 取值范围：不涉及。 - 默认取值：不涉及。
	ConnectionId string `json:"connection_id"`

	Body *UnbindVportRequestBody `json:"body,omitempty"`
}

func (o UnbindVportRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UnbindVportRequest struct{}"
	}

	return strings.Join([]string{"UnbindVportRequest", string(data)}, " ")
}
