package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListEmailTemplateRequest Request Object
type ListEmailTemplateRequest struct {

	// 数据库类型。支持MySQL、TaurusDB、GaussDB、MariaDB。
	DatastoreType string `json:"datastore_type"`

	// 偏移量
	Offset *int32 `json:"offset,omitempty"`

	// 每页记录数
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListEmailTemplateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListEmailTemplateRequest struct{}"
	}

	return strings.Join([]string{"ListEmailTemplateRequest", string(data)}, " ")
}
