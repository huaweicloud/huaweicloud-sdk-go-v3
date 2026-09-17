package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListTestBranchesResponse Response Object
type ListTestBranchesResponse struct {

	// 对外时：success|error;
	Status *string `json:"status,omitempty"`

	Result         *ResultValueListTestVersionVo `json:"result,omitempty"`
	HttpStatusCode int                           `json:"-"`
}

func (o ListTestBranchesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTestBranchesResponse struct{}"
	}

	return strings.Join([]string{"ListTestBranchesResponse", string(data)}, " ")
}
