package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TestCaseBasicInfo struct {
	AlertTemplate *AlertTemplate `json:"alert_template,omitempty"`

	// tmss用例类型
	CaseType *int32 `json:"caseType,omitempty"`

	// 执行机类型
	ExecutorType *string `json:"executor_type,omitempty"`

	// 用例id
	Id *string `json:"id,omitempty"`

	// 是否收藏
	IsForbidden *bool `json:"is_forbidden,omitempty"`

	// 用例id
	Name *string `json:"name,omitempty"`

	// 用例编号
	Number *string `json:"number,omitempty"`

	// 用例包更新时间
	ProjectUpdateTime *int64 `json:"project_update_time,omitempty"`

	// 用例包名
	ScriptProjectName *string `json:"scriptProjectName,omitempty"`

	// 用例状态
	State *int32 `json:"state,omitempty"`

	// svn脚本路径
	SvnScriptPath *string `json:"svn_script_path,omitempty"`

	// tmss版本地址
	TmssVersionUri *string `json:"tmssVersionUri,omitempty"`
}

func (o TestCaseBasicInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TestCaseBasicInfo struct{}"
	}

	return strings.Join([]string{"TestCaseBasicInfo", string(data)}, " ")
}
