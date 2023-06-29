package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TransactionConfigSearchRequest 获取url跟踪配置数据入参。
type TransactionConfigSearchRequest struct {

	// 应用id。
	BusinessId int64 `json:"business_id"`

	// 环境id。
	EnvId *int64 `json:"env_id,omitempty"`

	// 页码。
	PageNo int32 `json:"page_no"`

	// 每页数量。
	PageSize int32 `json:"page_size"`
}

func (o TransactionConfigSearchRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TransactionConfigSearchRequest struct{}"
	}

	return strings.Join([]string{"TransactionConfigSearchRequest", string(data)}, " ")
}
