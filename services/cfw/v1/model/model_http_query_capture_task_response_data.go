package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type HttpQueryCaptureTaskResponseData struct {

	// 每页显示个数，范围为1-1024
	Limit *int32 `json:"limit,omitempty"`

	// 防护对象id，是创建云防火墙后用于区分互联网边界防护和VPC边界防护的标志id，可通过调用[查询防火墙实例接口](ListFirewallDetail.xml)，注意type为0的为互联网边界防护对象id，type为1的为VPC边界防护对象id。
	ObjectId *string `json:"object_id,omitempty"`

	// 偏移量：指定返回记录的开始位置，必须为数字，取值范围为大于或等于0，默认0
	Offset *int32 `json:"offset,omitempty"`

	// 租户project_id
	ProjectId *string `json:"project_id,omitempty"`

	// 抓包任务总数
	Total *int64 `json:"total,omitempty"`

	// 抓包任务列表
	Records *[]CaptureTaskVo `json:"records,omitempty"`
}

func (o HttpQueryCaptureTaskResponseData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HttpQueryCaptureTaskResponseData struct{}"
	}

	return strings.Join([]string{"HttpQueryCaptureTaskResponseData", string(data)}, " ")
}
