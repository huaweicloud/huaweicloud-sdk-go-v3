package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListHtapDataStoreRequest Request Object
type ListHtapDataStoreRequest struct {

	// HTAP引擎名。 取值范围： - star-rocks - click-house
	EngineName string `json:"engine_name"`

	// 内容类型。 取值：application/json。
	ContentType string `json:"Content-Type"`

	// 请求语言类型。默认en-us。 取值范围： - en-us - zh-cn
	XLanguage *string `json:"X-Language,omitempty"`
}

func (o ListHtapDataStoreRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListHtapDataStoreRequest struct{}"
	}

	return strings.Join([]string{"ListHtapDataStoreRequest", string(data)}, " ")
}
