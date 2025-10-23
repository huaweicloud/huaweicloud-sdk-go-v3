package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteBranchResponse Response Object
type BatchDeleteBranchResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o BatchDeleteBranchResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteBranchResponse struct{}"
	}

	return strings.Join([]string{"BatchDeleteBranchResponse", string(data)}, " ")
}
