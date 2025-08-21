package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteProtectedBranchesResponse Response Object
type BatchDeleteProtectedBranchesResponse struct {

	// **参数解释：** 状态码。
	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o BatchDeleteProtectedBranchesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteProtectedBranchesResponse struct{}"
	}

	return strings.Join([]string{"BatchDeleteProtectedBranchesResponse", string(data)}, " ")
}
