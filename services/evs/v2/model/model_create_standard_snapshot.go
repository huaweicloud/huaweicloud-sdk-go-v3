package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateStandardSnapshot struct {

	// **参数解释** 云硬盘id。 可通过[查询所有云硬盘详情](evs_04_2006.xml)获取云硬盘ID。 **约束限制** 不涉及。 **取值范围** 不涉及。 **默认取值** 不涉及。
	VolumeId *string `json:"volume_id,omitempty"`

	// **参数解释** 云硬盘快照名称。 可通过[查询云硬盘快照详情列表](evs_04_2021.xml)接口获取云硬盘快照名称。 **约束限制** 不涉及。 **取值范围** 最大支持64个字符。 **默认取值** 不涉及。
	Name *string `json:"name,omitempty"`

	// **参数解释** 云硬盘快照描述。 **约束限制** 不涉及。 **取值范围** 最大支持85个字符。 **默认取值** 不涉及。
	Description *string `json:"description,omitempty"`

	// **参数解释** 云硬盘快照的标签。 标签的key：String类型，最大长度36个字符。 由英文字母、数字、下划线、中划线、UNICODE字符（\\u4E00-\\u9FFF）组成。  标签的value：String类型，最大长度43个字符，可以为空字符串。 由英文字母、数字、下划线、点、中划线、UNICODE字符（\\u4E00-\\u9FFF）组成。 **约束限制** 不涉及。 **取值范围** 不涉及。 **默认取值** 不涉及。
	Tags map[string]string `json:"tags,omitempty"`

	// **参数解释** 云硬盘快照的企业项目id。  关于企业项目ID的获取及企业项目特性的详细信息，请参考：\"[[企业管理用户指南](https://support.huaweicloud.com/usermanual-em/zh-cn_topic_0123692049.html)\"](tag:hws)[[企业管理用户指南](https://support.huaweicloud.com/intl/zh-cn/usermanual-em/zh-cn_topic_0123692049.html)](tag:hws_hk)。 **约束限制** 不涉及。 **取值范围** 不涉及。 **默认取值** 0，表示default企业项目。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// **参数解释** 是否开启极速可用功能。 **约束限制** 不涉及。 **取值范围** - true：表示开启。 - false：表示不开启。 **默认取值** 默认为false，表示不开启。
	InstantAccess *bool `json:"instant_access,omitempty"`

	// **参数解释** 是否创建增量快照。 **约束限制** 不涉及。 **取值范围** - true：创建增量快照。 - false：创建全量快照。 **默认取值** true
	Incremental *bool `json:"incremental,omitempty"`

	// **参数解释** 设置快照的保留时间，单位为天。保留时间到期后快照会被自动释放。 **约束限制** 不涉及。 **取值范围** 1~65536。 **默认取值** 空，表示快照不会被自动释放。
	RetentionDays *int32 `json:"retention_days,omitempty"`
}

func (o CreateStandardSnapshot) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateStandardSnapshot struct{}"
	}

	return strings.Join([]string{"CreateStandardSnapshot", string(data)}, " ")
}
