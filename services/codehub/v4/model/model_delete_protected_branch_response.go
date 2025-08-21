package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteProtectedBranchResponse Response Object
type DeleteProtectedBranchResponse struct {

	// **参数解释：** 状态码。
	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteProtectedBranchResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteProtectedBranchResponse struct{}"
	}

	return strings.Join([]string{"DeleteProtectedBranchResponse", string(data)}, " ")
}
