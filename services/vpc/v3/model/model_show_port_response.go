package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowPortResponse Response Object
type ShowPortResponse struct {
	Port *Port `json:"port,omitempty"`

	// **参数解释**： 请求ID。 **取值范围**： 不涉及。
	RequestId      *string `json:"request_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowPortResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPortResponse struct{}"
	}

	return strings.Join([]string{"ShowPortResponse", string(data)}, " ")
}
