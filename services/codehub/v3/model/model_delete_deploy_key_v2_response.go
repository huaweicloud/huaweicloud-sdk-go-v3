package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteDeployKeyV2Response Response Object
type DeleteDeployKeyV2Response struct {
	Error *Error `json:"error,omitempty"`

	// 响应结果
	Result *bool `json:"result,omitempty"`

	// 响应状态
	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteDeployKeyV2Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteDeployKeyV2Response struct{}"
	}

	return strings.Join([]string{"DeleteDeployKeyV2Response", string(data)}, " ")
}
