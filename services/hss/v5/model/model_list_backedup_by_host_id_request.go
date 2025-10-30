package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListBackedupByHostIdRequest Request Object
type ListBackedupByHostIdRequest struct {

	// **参数解释**: 区域ID，用于查询目的区域内的资产。获取方式请参见[获取区域ID](hss_02_0026.xml)。 **约束限制**: 不涉及 **取值范围**: 字符长度1-128位 **默认取值**: 不涉及
	Region *string `json:"region,omitempty"`

	// **参数解释**: 企业项目ID，用于过滤不同企业项目下的资产。获取方式请参见[获取企业项目ID](hss_02_0027.xml)。 如需查询所有企业项目下的资产请传参“all_granted_eps”。 **约束限制**: 开通企业项目功能后才需要配置企业项目ID参数。 **取值范围**: 字符长度1-256位 **默认取值**: 0，表示默认企业项目（default）。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// **参数解释**: 偏移量：指定返回记录的开始位置 **约束限制**: 不涉及 **取值范围**: 最小值0，最大值2000000 **默认取值**: 不涉及
	Offset int32 `json:"offset"`

	// **参数解释**: 每页显示个数 **约束限制**: 不涉及 **取值范围**: 取值10-200 **默认取值**: 10
	Limit int32 `json:"limit"`

	// **参数解释**: 主机ID **约束限制**: 不涉及 **取值范围**: 字符长度0-64位 **默认取值**: 不涉及
	HostId string `json:"host_id"`

	// **参数解释**: 状态 **约束限制**: 不涉及 **取值范围**: - available：可用 - protecting：保护中 - deleting：删除中 - restoring：恢复中 - error：异常 - waiting_protect：等待保护 - waiting_delete：等待删除 - waiting_restore：等待恢复  **默认取值**: 不涉及
	Status *ListBackedupByHostIdRequestStatus `json:"status,omitempty"`

	// **参数解释**: 备份名称 **约束限制**: 不涉及 **取值范围**: 字符长度0-65535位 **默认取值**: 不涉及
	Name *string `json:"name,omitempty"`

	// **参数解释**: 查询时间范围 **约束限制**: 不涉及 **取值范围**: 取值1-30 **默认取值**: 不涉及
	LastDays *int32 `json:"last_days,omitempty"`
}

func (o ListBackedupByHostIdRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListBackedupByHostIdRequest struct{}"
	}

	return strings.Join([]string{"ListBackedupByHostIdRequest", string(data)}, " ")
}

type ListBackedupByHostIdRequestStatus struct {
	value string
}

type ListBackedupByHostIdRequestStatusEnum struct {
	AVAILABLE       ListBackedupByHostIdRequestStatus
	PROTECTING      ListBackedupByHostIdRequestStatus
	DELETING        ListBackedupByHostIdRequestStatus
	RESTORING       ListBackedupByHostIdRequestStatus
	ERROR           ListBackedupByHostIdRequestStatus
	WAITING_PROTECT ListBackedupByHostIdRequestStatus
	WAITING_DELETE  ListBackedupByHostIdRequestStatus
	WAITING_RESTORE ListBackedupByHostIdRequestStatus
}

func GetListBackedupByHostIdRequestStatusEnum() ListBackedupByHostIdRequestStatusEnum {
	return ListBackedupByHostIdRequestStatusEnum{
		AVAILABLE: ListBackedupByHostIdRequestStatus{
			value: "available",
		},
		PROTECTING: ListBackedupByHostIdRequestStatus{
			value: "protecting",
		},
		DELETING: ListBackedupByHostIdRequestStatus{
			value: "deleting",
		},
		RESTORING: ListBackedupByHostIdRequestStatus{
			value: "restoring",
		},
		ERROR: ListBackedupByHostIdRequestStatus{
			value: "error",
		},
		WAITING_PROTECT: ListBackedupByHostIdRequestStatus{
			value: "waiting_protect",
		},
		WAITING_DELETE: ListBackedupByHostIdRequestStatus{
			value: "waiting_delete",
		},
		WAITING_RESTORE: ListBackedupByHostIdRequestStatus{
			value: "waiting_restore",
		},
	}
}

func (c ListBackedupByHostIdRequestStatus) Value() string {
	return c.value
}

func (c ListBackedupByHostIdRequestStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListBackedupByHostIdRequestStatus) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: string")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(string); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to string error")
	}
}
