package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DiagnoseRecordVo struct {

	// ID
	Id *string `json:"id,omitempty"`

	// 状态
	Status *string `json:"status,omitempty"`

	// 不同诊断定独有的属性
	Params map[string]string `json:"params,omitempty"`

	// 创建人
	DomainId *string `json:"domain_id,omitempty"`

	// 创建时间戳
	CreateTime *int64 `json:"create_time,omitempty"`

	// 任务ID
	JobId *string `json:"job_id,omitempty"`

	// 异常项id列表
	AbnormalItems *[]string `json:"abnormal_items,omitempty"`

	// 异常项数量
	RedCount *int32 `json:"red_count,omitempty"`

	// 区域
	Region *string `json:"region,omitempty"`
}

func (o DiagnoseRecordVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DiagnoseRecordVo struct{}"
	}

	return strings.Join([]string{"DiagnoseRecordVo", string(data)}, " ")
}
