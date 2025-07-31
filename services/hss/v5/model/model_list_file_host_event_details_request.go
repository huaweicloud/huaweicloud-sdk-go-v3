package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListFileHostEventDetailsRequest Request Object
type ListFileHostEventDetailsRequest struct {

	// **参数解释**: 企业项目ID，用于过滤不同企业项目下的资产。获取方式请参见[获取企业项目ID](hss_02_0027.xml)。 如需查询所有企业项目下的资产请传参“all_granted_eps”。 **约束限制**: 开通企业项目功能后才需要配置企业项目ID参数。 **取值范围**: 字符长度1-256位 **默认取值**: 0，表示默认企业项目（default）。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// **参数解释**: 开始时间，13位时间戳 **约束限制**: 不涉及 **取值范围**: 最小值0，最大值9223372036854775807 **默认取值**:
	BeginTime *int64 `json:"begin_time,omitempty"`

	// **参数解释**: 结束时间，13位时间戳 **约束限制**: 不涉及 **取值范围**: 最小值0，最大值9223372036854775807 **默认取值**:
	EndTime *int64 `json:"end_time,omitempty"`

	// **参数解释**: 服务器ID **约束限制**: 不涉及 **取值范围**: 字符长度1-64位 **默认取值**: 不涉及
	HostId string `json:"host_id"`

	// 文件名称
	FileName *string `json:"file_name,omitempty"`

	// 文件路径
	FilePath *string `json:"file_path,omitempty"`

	// 变更类型，包含如下:   - \"all\" : 全部   - \"registry\" : 注册表   - \"file\" : 文件
	ChangeType *string `json:"change_type,omitempty"`

	// 变更类别，包含如下:   - \"all\" : 全部   - \"modify\" : 修改   - \"add\" : 新增   - \"delete\" : 删除
	ChangeCategory *string `json:"change_category,omitempty"`

	// 状态，包含如下:   - \"all\" : 全部   - \"trust\" : 可信   - \"untrust\" : 不可信   - \"unknown\" : 未知
	Status *string `json:"status,omitempty"`

	// **参数解释**: 偏移量：指定返回记录的开始位置 **约束限制**: 不涉及 **取值范围**: 最小值0，最大值2000000 **默认取值**: 不涉及
	Offset int32 `json:"offset"`

	// **参数解释**: 每页显示个数 **约束限制**: 不涉及 **取值范围**: 取值10-200 **默认取值**: 10
	Limit int32 `json:"limit"`
}

func (o ListFileHostEventDetailsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListFileHostEventDetailsRequest struct{}"
	}

	return strings.Join([]string{"ListFileHostEventDetailsRequest", string(data)}, " ")
}
