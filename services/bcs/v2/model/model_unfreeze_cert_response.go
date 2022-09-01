package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type UnfreezeCertResponse struct {

	// 操作结果
	Result         *string `json:"result,omitempty" xml:"result"`
	HttpStatusCode int     `json:"-"`
}

func (o UnfreezeCertResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UnfreezeCertResponse struct{}"
	}

	return strings.Join([]string{"UnfreezeCertResponse", string(data)}, " ")
}
