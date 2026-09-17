package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDatabaseInfosRequest Request Object
type ListDatabaseInfosRequest struct {

	// 实例id
	InstanceId string `json:"instance_id"`

	// 语言
	XLanguage *string `json:"X-Language,omitempty"`

	// 数据库引擎类型
	EngineType *string `json:"engine_type,omitempty"`
}

func (o ListDatabaseInfosRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDatabaseInfosRequest struct{}"
	}

	return strings.Join([]string{"ListDatabaseInfosRequest", string(data)}, " ")
}
