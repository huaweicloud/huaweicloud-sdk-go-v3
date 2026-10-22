package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteV2ComputeResponse Response Object
type DeleteV2ComputeResponse struct {

	// **参数解释**：计算资源ID。 **取值范围**：长度为1~36个英文字母、数字和中划线的组合。
	ComputeId *string `json:"compute_id,omitempty"`

	XRequestId     *string `json:"X-request-id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteV2ComputeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteV2ComputeResponse struct{}"
	}

	return strings.Join([]string{"DeleteV2ComputeResponse", string(data)}, " ")
}
