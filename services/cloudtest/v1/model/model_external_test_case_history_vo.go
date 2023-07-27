package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExternalTestCaseHistoryVo 实际的数据类型：单个对象，集合 或 NULL
type ExternalTestCaseHistoryVo struct {

	// 资源历史记录创建人ID
	Author *string `json:"author,omitempty"`

	// 逻辑region
	Region *string `json:"region,omitempty"`

	// 历史记录字段变更列表
	Changes *[]ResourceChangeExternalVo `json:"changes,omitempty"`

	// 历史记录id
	Id *string `json:"id,omitempty"`

	// 用例id
	TestcaseId *string `json:"testcase_id,omitempty"`

	// 创建时间
	CreationDate *sdktime.SdkTime `json:"creation_date,omitempty"`

	// 创建时间时间戳
	CreateDateTimestamp *int64 `json:"create_date_timestamp,omitempty"`

	// 创建人名称
	AuthorName *string `json:"author_name,omitempty"`

	// 项目id
	ProjectId *string `json:"project_id,omitempty"`
}

func (o ExternalTestCaseHistoryVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExternalTestCaseHistoryVo struct{}"
	}

	return strings.Join([]string{"ExternalTestCaseHistoryVo", string(data)}, " ")
}
