package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ListCollectConfigResponseBodyAccounts struct {

	// 账号ID
	AccountId *string `json:"account_id,omitempty"`

	// 最近接入时间
	LastLogDate float32 `json:"last_log_date,omitempty"`

	// 最近一小时接入数量
	LogCount *string `json:"log_count,omitempty"`

	// 账号名称
	Name *string `json:"name,omitempty"`

	// 接入状态
	ProcessStatus *string `json:"process_status,omitempty"`

	// 错误信息
	SinkMsg *string `json:"sink_msg,omitempty"`
}

func (o ListCollectConfigResponseBodyAccounts) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCollectConfigResponseBodyAccounts struct{}"
	}

	return strings.Join([]string{"ListCollectConfigResponseBodyAccounts", string(data)}, " ")
}
