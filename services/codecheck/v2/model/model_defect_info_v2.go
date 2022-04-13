package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 缺陷信息
type DefectInfoV2 struct {
	// 缺陷的id

	DefectId *string `json:"defect_id,omitempty"`
	// 缺陷对应检查项的名称

	DefectCheckerName *string `json:"defect_checker_name,omitempty"`
	// 缺陷的状态0为解决 1已解决 2已忽略

	DefectStatus *string `json:"defect_status,omitempty"`
	// 规则标签,多个标签用逗号隔开

	RuleSystemTags *string `json:"rule_system_tags,omitempty"`
	// 规则名

	RuleName *string `json:"rule_name,omitempty"`
	// 缺陷所在文件行号

	LineNumber *string `json:"line_number,omitempty"`
	// 缺陷描述

	DefectContent *string `json:"defect_content,omitempty"`
	// 缺陷等级，0致命，1严重，2一般，3提示

	DefectLevel *string `json:"defect_level,omitempty"`
	// 缺陷文件路径

	FilePath *string `json:"file_path,omitempty"`
	// 创建时间

	CreatedAt *string `json:"created_at,omitempty"`
	// 问题唯一标识

	IssueKey *string `json:"issue_key,omitempty"`
	// 缺陷代码片段详情

	Fragment *[]DefectFragmentV2 `json:"fragment,omitempty"`
}

func (o DefectInfoV2) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DefectInfoV2 struct{}"
	}

	return strings.Join([]string{"DefectInfoV2", string(data)}, " ")
}
