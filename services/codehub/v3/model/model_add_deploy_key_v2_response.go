package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type AddDeployKeyV2Response struct {
	Error *Error `json:"error,omitempty"`

	Result *Key `json:"result,omitempty"`
	// 响应状态

	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o AddDeployKeyV2Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddDeployKeyV2Response struct{}"
	}

	return strings.Join([]string{"AddDeployKeyV2Response", string(data)}, " ")
}
