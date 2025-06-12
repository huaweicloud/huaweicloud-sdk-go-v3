package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListBuildParameterResponse Response Object
type ListBuildParameterResponse struct {

	// 状态
	Status *string `json:"status,omitempty"`

	// 错误
	Error *interface{} `json:"error,omitempty"`

	// 参数名和参数值
	Result         *[]BuildParameter `json:"result,omitempty"`
	HttpStatusCode int               `json:"-"`
}

func (o ListBuildParameterResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListBuildParameterResponse struct{}"
	}

	return strings.Join([]string{"ListBuildParameterResponse", string(data)}, " ")
}
