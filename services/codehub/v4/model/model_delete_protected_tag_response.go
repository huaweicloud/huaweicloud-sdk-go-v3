package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteProtectedTagResponse Response Object
type DeleteProtectedTagResponse struct {

	// **参数解释：** 状态码。
	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteProtectedTagResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteProtectedTagResponse struct{}"
	}

	return strings.Join([]string{"DeleteProtectedTagResponse", string(data)}, " ")
}
