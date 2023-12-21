package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateUsingPostResponse Response Object
type UpdateUsingPostResponse struct {

	// 请求结果。
	Result *string `json:"result,omitempty"`

	// 请求数据。
	Data *[]RdmParamVoPersistableModelUpdateDto `json:"data,omitempty"`

	// 异常信息。
	Errors         *[]string `json:"errors,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o UpdateUsingPostResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateUsingPostResponse struct{}"
	}

	return strings.Join([]string{"UpdateUsingPostResponse", string(data)}, " ")
}
