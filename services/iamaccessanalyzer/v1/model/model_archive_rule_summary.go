package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ArchiveRuleSummary 分析器创建的存档规则。
type ArchiveRuleSummary struct {

	// 创建存档规则的时间。
	CreatedAt *sdktime.SdkTime `json:"created_at"`

	// 匹配要返回的访问分析结果的筛选器。
	Filters []FindingFilter `json:"filters"`

	// 存档规则的唯一标识符。
	Id string `json:"id"`

	// 创建存档规则的名称。
	Name string `json:"name"`

	// 上次更新存档规则的时间。
	UpdatedAt *sdktime.SdkTime `json:"updated_at"`

	// 存档规则的唯一资源标识符。
	Urn string `json:"urn"`
}

func (o ArchiveRuleSummary) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ArchiveRuleSummary struct{}"
	}

	return strings.Join([]string{"ArchiveRuleSummary", string(data)}, " ")
}
