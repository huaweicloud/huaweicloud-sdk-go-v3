package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// QueryAutoSqlLimitingRequest Request Object
type QueryAutoSqlLimitingRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id"`

	// 语言。默认en-us。
	XLanguage *string `json:"X-Language,omitempty"`
}

func (o QueryAutoSqlLimitingRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QueryAutoSqlLimitingRequest struct{}"
	}

	return strings.Join([]string{"QueryAutoSqlLimitingRequest", string(data)}, " ")
}
