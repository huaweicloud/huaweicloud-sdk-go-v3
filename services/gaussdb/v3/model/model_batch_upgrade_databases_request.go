package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchUpgradeDatabasesRequest Request Object
type BatchUpgradeDatabasesRequest struct {

	// 请求语言类型。默认en-us。 取值范围： - en-us - zh-cn
	XLanguage *string `json:"X-Language,omitempty"`

	Body *BatchUpgradeDatabasesRequestBody `json:"body,omitempty"`
}

func (o BatchUpgradeDatabasesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchUpgradeDatabasesRequest struct{}"
	}

	return strings.Join([]string{"BatchUpgradeDatabasesRequest", string(data)}, " ")
}
