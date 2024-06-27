package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateClickHouseLtsConfigRequest Request Object
type UpdateClickHouseLtsConfigRequest struct {

	// 请求语言类型。默认en-us。 取值范围： - en-us - zh-cn
	XLanguage *string `json:"X-Language,omitempty"`

	Body *CreateChLtsConfigRequestBody `json:"body,omitempty"`
}

func (o UpdateClickHouseLtsConfigRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateClickHouseLtsConfigRequest struct{}"
	}

	return strings.Join([]string{"UpdateClickHouseLtsConfigRequest", string(data)}, " ")
}
