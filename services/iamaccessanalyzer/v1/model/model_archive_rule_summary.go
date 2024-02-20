package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ArchiveRuleSummary 为指定分析器创建的存档规则的列表。
type ArchiveRuleSummary struct {

	// 创建归档规则的时间。
	CreatedAt *sdktime.SdkTime `json:"created_at"`

	Filters []FindingFilter `json:"filters"`

	// 存档规则的唯一标识符。
	Id string `json:"id"`

	// 创建归档规则的名称。
	Name string `json:"name"`

	// 上次更新存档规则的时间。
	UpdatedAt *sdktime.SdkTime `json:"updated_at"`

	// 唯一的资源名称。
	Urn string `json:"urn"`
}

func (o ArchiveRuleSummary) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ArchiveRuleSummary struct{}"
	}

	return strings.Join([]string{"ArchiveRuleSummary", string(data)}, " ")
}
