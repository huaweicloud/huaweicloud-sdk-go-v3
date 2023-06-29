package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateNewBranchResponse Response Object
type CreateNewBranchResponse struct {
	Error *Error `json:"error,omitempty"`

	Result *AddProtectResponse `json:"result,omitempty"`

	// 响应状态
	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateNewBranchResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateNewBranchResponse struct{}"
	}

	return strings.Join([]string{"CreateNewBranchResponse", string(data)}, " ")
}
