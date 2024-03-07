package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteBranchResponse Response Object
type BatchDeleteBranchResponse struct {

	// 请求结果。
	Result *string `json:"result,omitempty"`

	// 请求数据。
	Data *[]int32 `json:"data,omitempty"`

	// 异常信息。
	Errors         *[]string `json:"errors,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o BatchDeleteBranchResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteBranchResponse struct{}"
	}

	return strings.Join([]string{"BatchDeleteBranchResponse", string(data)}, " ")
}
