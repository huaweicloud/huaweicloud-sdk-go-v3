package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowTableMetaInfoRequest Request Object
type ShowTableMetaInfoRequest struct {

	// 请求语言类型。默认en-us。 取值范围： - en-us - zh-cn
	XLanguage *string `json:"X-Language,omitempty"`

	// 实例ID。
	InstanceId string `json:"instance_id"`

	// 数据库名称。
	DatabaseName *string `json:"database_name,omitempty"`

	// 数据表名称。
	TableName *string `json:"table_name,omitempty"`
}

func (o ShowTableMetaInfoRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTableMetaInfoRequest struct{}"
	}

	return strings.Join([]string{"ShowTableMetaInfoRequest", string(data)}, " ")
}
