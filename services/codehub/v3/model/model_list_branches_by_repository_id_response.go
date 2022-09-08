package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListBranchesByRepositoryIdResponse struct {
	Error *Error `json:"error,omitempty"`

	Result *BranchResponse `json:"result,omitempty"`

	// 响应状态
	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListBranchesByRepositoryIdResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListBranchesByRepositoryIdResponse struct{}"
	}

	return strings.Join([]string{"ListBranchesByRepositoryIdResponse", string(data)}, " ")
}
