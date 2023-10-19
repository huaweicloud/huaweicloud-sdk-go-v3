package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListConfigurationsDifferencesRequest Request Object
type ListConfigurationsDifferencesRequest struct {

	// 请求语言类型。默认en-us。 取值范围： - en-us - zh-cn
	XLanguage *string `json:"X-Language,omitempty"`

	Body *ListConfigurationsDifferencesRequestBody `json:"body,omitempty"`
}

func (o ListConfigurationsDifferencesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListConfigurationsDifferencesRequest struct{}"
	}

	return strings.Join([]string{"ListConfigurationsDifferencesRequest", string(data)}, " ")
}
