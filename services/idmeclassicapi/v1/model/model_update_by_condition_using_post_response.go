package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateByConditionUsingPostResponse Response Object
type UpdateByConditionUsingPostResponse struct {

	// 请求结果。
	Result *string `json:"result,omitempty"`

	// 请求数据。
	Data *[]RdmParamVoPersistableModelUpdateDto `json:"data,omitempty"`

	// 异常信息。
	Errors         *[]string `json:"errors,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o UpdateByConditionUsingPostResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateByConditionUsingPostResponse struct{}"
	}

	return strings.Join([]string{"UpdateByConditionUsingPostResponse", string(data)}, " ")
}
