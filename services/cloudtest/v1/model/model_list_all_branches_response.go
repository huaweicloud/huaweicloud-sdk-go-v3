package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAllBranchesResponse Response Object
type ListAllBranchesResponse struct {

	// 对外时：success|error;
	Status *string `json:"status,omitempty"`

	Result         *ResultValueListTestVersionVo `json:"result,omitempty"`
	HttpStatusCode int                           `json:"-"`
}

func (o ListAllBranchesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAllBranchesResponse struct{}"
	}

	return strings.Join([]string{"ListAllBranchesResponse", string(data)}, " ")
}
