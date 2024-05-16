package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TxSearchRequest 获取URL跟踪视图列表入参。
type TxSearchRequest struct {

	// 应用id。
	BusinessId int64 `json:"business_id"`

	// region英文名称。
	Region string `json:"region"`

	// 开始时间。
	StartTime string `json:"start_time"`

	// 结束时间。
	EndTime string `json:"end_time"`

	// 环境id。
	EnvId *int64 `json:"env_id,omitempty"`

	// 上次请求的id。
	RequestId *string `json:"request_id,omitempty"`

	// 页码。
	PageNo int32 `json:"page_no"`

	// 每页数量。
	PageSize *int32 `json:"page_size,omitempty"`
}

func (o TxSearchRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TxSearchRequest struct{}"
	}

	return strings.Join([]string{"TxSearchRequest", string(data)}, " ")
}
