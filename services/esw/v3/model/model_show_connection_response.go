package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowConnectionResponse Response Object
type ShowConnectionResponse struct {

	// - 参数解释：请求的唯一标识。 - 约束限制：UUID格式。 - 取值范围：不涉及。 - 默认取值：不涉及。
	RequestId *string `json:"request_id,omitempty"`

	Connection     *Connection `json:"connection,omitempty"`
	HttpStatusCode int         `json:"-"`
}

func (o ShowConnectionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowConnectionResponse struct{}"
	}

	return strings.Join([]string{"ShowConnectionResponse", string(data)}, " ")
}
