package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateConnectionResponse Response Object
type CreateConnectionResponse struct {

	// - 参数解释：请求的唯一标识。 - 约束限制：UUID格式。 - 取值范围：不涉及。 - 默认取值：不涉及。
	RequestId *string `json:"request_id,omitempty"`

	Connection     *Connection `json:"connection,omitempty"`
	HttpStatusCode int         `json:"-"`
}

func (o CreateConnectionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateConnectionResponse struct{}"
	}

	return strings.Join([]string{"CreateConnectionResponse", string(data)}, " ")
}
