package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteBranchResponse Response Object
type DeleteBranchResponse struct {

	// **参数解释：** 状态码。
	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteBranchResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteBranchResponse struct{}"
	}

	return strings.Join([]string{"DeleteBranchResponse", string(data)}, " ")
}
