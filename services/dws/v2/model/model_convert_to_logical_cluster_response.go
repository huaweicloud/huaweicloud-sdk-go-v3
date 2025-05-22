package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ConvertToLogicalClusterResponse Response Object
type ConvertToLogicalClusterResponse struct {

	// **参数解释**： 错误码。 **取值范围**： 不涉及。
	ErrorCode *string `json:"error_code,omitempty"`

	// **参数解释**： 错误信息。 **取值范围**： 不涉及。
	ErrorMsg       *string `json:"error_msg,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ConvertToLogicalClusterResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConvertToLogicalClusterResponse struct{}"
	}

	return strings.Join([]string{"ConvertToLogicalClusterResponse", string(data)}, " ")
}
