package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListHostNetResponse Response Object
type ListHostNetResponse struct {

	// **参数解释**： 网卡状态响应体。 **取值范围**： 不涉及。
	Body           *[]NetResp `json:"body,omitempty"`
	HttpStatusCode int        `json:"-"`
}

func (o ListHostNetResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListHostNetResponse struct{}"
	}

	return strings.Join([]string{"ListHostNetResponse", string(data)}, " ")
}
