package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ScanExpireKeyResponse Response Object
type ScanExpireKeyResponse struct {

	// 主键ID
	Id *string `json:"id,omitempty"`

	// 实例ID
	InstanceId *string `json:"instance_id,omitempty"`

	// 状态
	Status *string `json:"status,omitempty"`

	// 扫描类型
	ScanType *string `json:"scan_type,omitempty"`

	// 创建时间, 格式为：\"2020-06-15T02:21:18.669Z\"
	CreatedAt *string `json:"created_at,omitempty"`

	// 开始时间, 格式为：\"2020-06-15T02:21:18.669Z\"
	StartedAt *string `json:"started_at,omitempty"`

	// 完成时间, 格式为：\"2020-06-15T02:21:18.669Z\"
	FinishedAt     *string `json:"finished_at,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ScanExpireKeyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ScanExpireKeyResponse struct{}"
	}

	return strings.Join([]string{"ScanExpireKeyResponse", string(data)}, " ")
}
