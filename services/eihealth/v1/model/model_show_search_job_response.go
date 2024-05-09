package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowSearchJobResponse Response Object
type ShowSearchJobResponse struct {
	BasicInfo *DrugJobDto `json:"basic_info,omitempty"`

	// 分子SMILES表达式
	Smiles *string `json:"smiles,omitempty"`

	// 分子骨架表达式
	Scaffold *string `json:"scaffold,omitempty"`

	// 生成分子数量
	TopN *int32 `json:"top_n,omitempty"`

	// 可供搜索分子的公共数据库名称列表
	Databases *[]string `json:"databases,omitempty"`

	// 可供搜索分子的自定义数据库名称列表
	CustomDatabases *[]string `json:"custom_databases,omitempty"`

	// 模型信息
	Models *[]BasicDrugModel `json:"models,omitempty"`

	SearchMethod *SearchType `json:"search_method,omitempty"`

	// 部分失败原因和数量
	PartFailedReason *[]FailedReasonRecord `json:"part_failed_reason,omitempty"`
	HttpStatusCode   int                   `json:"-"`
}

func (o ShowSearchJobResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSearchJobResponse struct{}"
	}

	return strings.Join([]string{"ShowSearchJobResponse", string(data)}, " ")
}
