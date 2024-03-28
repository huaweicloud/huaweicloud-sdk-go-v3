package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateTableRequest Request Object
type CreateTableRequest struct {

	// 新增表所在的数据库名称。
	DatabaseName string `json:"database_name"`

	Body *CreateTableRequestBody `json:"body,omitempty"`
}

func (o CreateTableRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateTableRequest struct{}"
	}

	return strings.Join([]string{"CreateTableRequest", string(data)}, " ")
}
