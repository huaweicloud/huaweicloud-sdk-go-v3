package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowJobSystemParametersResponse Response Object
type ShowJobSystemParametersResponse struct {

	// 状态
	Status *string `json:"status,omitempty"`

	// 错误
	Error *interface{} `json:"error,omitempty"`

	// 本次任务的构建步骤详情，返回的步骤为页面可见步骤
	Result         *[]SystemParametersResult `json:"result,omitempty"`
	HttpStatusCode int                       `json:"-"`
}

func (o ShowJobSystemParametersResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowJobSystemParametersResponse struct{}"
	}

	return strings.Join([]string{"ShowJobSystemParametersResponse", string(data)}, " ")
}
