package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateSqlRenderResponse Response Object
type CreateSqlRenderResponse struct {

	// Flink SQL Gateway操作ID，用于后续查询数据
	OperateId      *string `json:"operate_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateSqlRenderResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSqlRenderResponse struct{}"
	}

	return strings.Join([]string{"CreateSqlRenderResponse", string(data)}, " ")
}
