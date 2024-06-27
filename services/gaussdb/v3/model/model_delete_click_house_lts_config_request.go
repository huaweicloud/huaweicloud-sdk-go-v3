package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteClickHouseLtsConfigRequest Request Object
type DeleteClickHouseLtsConfigRequest struct {

	// 请求语言类型。默认en-us。 取值范围： - en-us - zh-cn
	XLanguage *string `json:"X-Language,omitempty"`

	Body *DeleteChLtsConfigRequestBody `json:"body,omitempty"`
}

func (o DeleteClickHouseLtsConfigRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteClickHouseLtsConfigRequest struct{}"
	}

	return strings.Join([]string{"DeleteClickHouseLtsConfigRequest", string(data)}, " ")
}
