package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 缺陷信息
type DefectInfoV2 struct {

	// 缺陷的id
	DefectId *string `json:"defect_id,omitempty" xml:"defect_id"`

	// 缺陷对应检查项的名称
	DefectCheckerName *string `json:"defect_checker_name,omitempty" xml:"defect_checker_name"`

	// 缺陷的状态0为解决 1已解决 2已忽略
	DefectStatus *string `json:"defect_status,omitempty" xml:"defect_status"`

	// 规则标签,多个标签用逗号隔开
	RuleSystemTags *string `json:"rule_system_tags,omitempty" xml:"rule_system_tags"`

	// 规则名
	RuleName *string `json:"rule_name,omitempty" xml:"rule_name"`

	// 缺陷所在文件行号
	LineNumber *string `json:"line_number,omitempty" xml:"line_number"`

	// 缺陷描述
	DefectContent *string `json:"defect_content,omitempty" xml:"defect_content"`

	// 缺陷等级，0致命，1严重，2一般，3提示
	DefectLevel *string `json:"defect_level,omitempty" xml:"defect_level"`

	// 缺陷文件路径
	FilePath *string `json:"file_path,omitempty" xml:"file_path"`

	// 创建时间
	CreatedAt *string `json:"created_at,omitempty" xml:"created_at"`

	// 问题唯一标识
	IssueKey *string `json:"issue_key,omitempty" xml:"issue_key"`

	// 缺陷代码片段详情
	Fragment *[]DefectFragmentV2 `json:"fragment,omitempty" xml:"fragment"`
}

func (o DefectInfoV2) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DefectInfoV2 struct{}"
	}

	return strings.Join([]string{"DefectInfoV2", string(data)}, " ")
}
