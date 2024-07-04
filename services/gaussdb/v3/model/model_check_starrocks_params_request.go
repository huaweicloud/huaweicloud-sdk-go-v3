package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CheckStarrocksParamsRequest Request Object
type CheckStarrocksParamsRequest struct {

	// 请求语言类型。默认en-us。 取值范围： - en-us - zh-cn
	XLanguage *string `json:"X-Language,omitempty"`

	// 内容类型。 取值：application/json。
	ContentType string `json:"Content-Type"`

	Body *CheckStarrocksParamsRequestBody `json:"body,omitempty"`
}

func (o CheckStarrocksParamsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckStarrocksParamsRequest struct{}"
	}

	return strings.Join([]string{"CheckStarrocksParamsRequest", string(data)}, " ")
}
