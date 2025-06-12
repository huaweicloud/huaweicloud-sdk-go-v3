package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListRepoBranchResponse Response Object
type ListRepoBranchResponse struct {

	// 状态
	Status *string `json:"status,omitempty"`

	// 结果
	Result         *[]string `json:"result,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListRepoBranchResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRepoBranchResponse struct{}"
	}

	return strings.Join([]string{"ListRepoBranchResponse", string(data)}, " ")
}
