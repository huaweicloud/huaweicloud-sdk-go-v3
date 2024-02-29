package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowBatchDeleteBranchUsingPostResponse Response Object
type ShowBatchDeleteBranchUsingPostResponse struct {

	// 请求结果。
	Result *string `json:"result,omitempty"`

	// 请求数据。
	Data *[]int32 `json:"data,omitempty"`

	// 异常信息。
	Errors         *[]string `json:"errors,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ShowBatchDeleteBranchUsingPostResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowBatchDeleteBranchUsingPostResponse struct{}"
	}

	return strings.Join([]string{"ShowBatchDeleteBranchUsingPostResponse", string(data)}, " ")
}
