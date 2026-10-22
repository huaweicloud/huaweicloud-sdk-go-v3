package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteV2EndpointResponse Response Object
type DeleteV2EndpointResponse struct {

	// **参数解释**：端点ID。 **取值范围**：长度为1~36的英文字符、数字和中划线的组合。
	EndpointId *string `json:"endpoint_id,omitempty"`

	XRequestId     *string `json:"X-request-id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteV2EndpointResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteV2EndpointResponse struct{}"
	}

	return strings.Join([]string{"DeleteV2EndpointResponse", string(data)}, " ")
}
