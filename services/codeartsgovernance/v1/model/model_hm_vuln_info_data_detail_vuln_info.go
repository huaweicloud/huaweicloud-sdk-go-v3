package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type HmVulnInfoDataDetailVulnInfo struct {

	// 问题文件路径
	ProblemFilePath *string `json:"problem_file_path,omitempty"`

	// 问题特征信息
	IdentityInfo *string `json:"identity_info,omitempty"`

	// 问题等级
	RiskLevel *string `json:"risk_level,omitempty"`

	// 应用名称
	AppName *string `json:"app_name,omitempty"`

	// 问题项大类id
	CategoryId *string `json:"category_id,omitempty"`

	// 问题项大类中文名
	TypeCn *string `json:"type_cn,omitempty"`

	// 问题项大类英文名
	TypeEn *string `json:"type_en,omitempty"`

	// 问题描述中文
	ProblemCn *string `json:"problem_cn,omitempty"`

	// 问题描述英文
	ProblemEn *string `json:"problem_en,omitempty"`

	// 解决办法中文
	SolutionCn *string `json:"solution_cn,omitempty"`

	// 解决办法英文
	SolutionEn *string `json:"solution_en,omitempty"`

	// 问题详细描述中文
	DetectionScenarioCn *string `json:"detection_scenario_cn,omitempty"`

	// 问题详细描述英文
	DetectionScenarioEn *string `json:"detection_scenario_en,omitempty"`

	// 问题wiki
	WikiUrl *string `json:"wiki_url,omitempty"`

	// 问题来源规范
	StandardInfo *string `json:"standard_info,omitempty"`

	// 漏洞确认: 0 - 未确认 1 - 已确认
	ConfirmState *int32 `json:"confirm_state,omitempty"`

	// 漏洞确认结果: 0 - 未误报 1 - 误报
	ConfirmResult *int32 `json:"confirm_result,omitempty"`

	// 确认人
	Confirmer *string `json:"confirmer,omitempty"`

	// 确认描述
	ConfirmDescription *string `json:"confirm_description,omitempty"`

	// 确认时间
	ConfirmTime *string `json:"confirm_time,omitempty"`

	// 问题单编号
	DtsId *string `json:"dts_id,omitempty"`

	// 标准标号
	StandardNo *string `json:"standard_no,omitempty"`
}

func (o HmVulnInfoDataDetailVulnInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HmVulnInfoDataDetailVulnInfo struct{}"
	}

	return strings.Join([]string{"HmVulnInfoDataDetailVulnInfo", string(data)}, " ")
}
