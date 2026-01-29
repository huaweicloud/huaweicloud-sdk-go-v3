package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ListCollectConfigResponseBodyDataList struct {

	// 接入账号总数量
	AccountAllNum float32 `json:"account_all_num,omitempty"`

	// 成功接入账号数量
	AccountSuccessfulNum float32 `json:"account_successful_num,omitempty"`

	// 云产品ID
	Csvc *string `json:"csvc,omitempty"`

	// 日志数据
	Datasets *[]ListCollectConfigResponseBodyDatasets `json:"datasets,omitempty"`

	// 上次修改时间
	LastModifiedTime float32 `json:"last_modified_time,omitempty"`

	// 日志总数量
	LogAllNum float32 `json:"log_all_num,omitempty"`

	// 接入日志数量
	LogInNum float32 `json:"log_in_num,omitempty"`

	// 过去一个小时接入
	LogInNumLastOneHour float32 `json:"log_in_num_last_one_hour,omitempty"`

	// 状态
	ProcessStatus *string `json:"process_status,omitempty"`

	// 云厂商ID
	Vendor *string `json:"vendor,omitempty"`
}

func (o ListCollectConfigResponseBodyDataList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCollectConfigResponseBodyDataList struct{}"
	}

	return strings.Join([]string{"ListCollectConfigResponseBodyDataList", string(data)}, " ")
}
