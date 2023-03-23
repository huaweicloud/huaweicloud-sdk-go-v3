package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowExtensionTestingResultResponse struct {

	// 返回值
	Result *interface{} `json:"result,omitempty"`

	// 状态
	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowExtensionTestingResultResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowExtensionTestingResultResponse struct{}"
	}

	return strings.Join([]string{"ShowExtensionTestingResultResponse", string(data)}, " ")
}
