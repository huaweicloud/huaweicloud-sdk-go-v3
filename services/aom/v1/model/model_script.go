package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 脚本的基本信息。
type Script struct {
	ApproveInfo *ApproveInfo `json:"approve_info,omitempty"`

	// 创建人，比如为：张三
	CreateBy *string `json:"create_by,omitempty"`

	// 创建时间
	CreateTime *int64 `json:"create_time,omitempty"`

	// 脚本描述，脚本描述,对脚本进行描述，最大长度为1000
	Description *string `json:"description,omitempty"`

	// 企业项目id
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 项目id,根据project_id、脚本名称的和，用guava计算hash(md5)方式获取，比如为：c279d73a0b4e0f0927721e366c736880
	Id *string `json:"id,omitempty"`

	// 脚本名称，支持数字，下划线，大小写字母 ,中文
	Name string `json:"name"`

	// 脚本中是否有已上线的版本，true表示有已上线的版本，false表示灭没有已上线的版本
	OnlineExistStatus *bool `json:"online_exist_status,omitempty"`

	// 正在上线版本id
	OnlineId *string `json:"online_id,omitempty"`

	// 租户从IAM申请到的projectid，一般为32位字符串
	ProjectId *string `json:"project_id,omitempty"`

	RateControl *RateControl `json:"rate_control,omitempty"`

	// 脚本语言，目前支持四种，分别是：[\"Shell\",\"Bat\",\"Python\",\"Powershell\"]
	ScriptLanguage string `json:"script_language"`

	// 修改人
	UpdateBy *string `json:"update_by,omitempty"`

	// 实体的最后更新时间戳。 注意：执行创建/修补/删除操作时，update_time将更新。
	UpdateTime *int64 `json:"update_time,omitempty"`
}

func (o Script) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Script struct{}"
	}

	return strings.Join([]string{"Script", string(data)}, " ")
}
