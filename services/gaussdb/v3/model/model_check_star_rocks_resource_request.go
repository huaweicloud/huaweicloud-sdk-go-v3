package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CheckStarRocksResourceRequest Request Object
type CheckStarRocksResourceRequest struct {

	// 内容类型。  取值：application/json。
	ContentType string `json:"Content-Type"`

	// 请求语言类型。默认en-us。 取值范围： - en-us - zh-cn
	XLanguage *string `json:"X-Language,omitempty"`

	Body *ResourceCheck `json:"body,omitempty"`
}

func (o CheckStarRocksResourceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckStarRocksResourceRequest struct{}"
	}

	return strings.Join([]string{"CheckStarRocksResourceRequest", string(data)}, " ")
}
