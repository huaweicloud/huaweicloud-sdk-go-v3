package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// TenantRepositoryDto 租户下仓库列表
type TenantRepositoryDto struct {

	// **参数解释：** 仓库所有者。 **取值范围：** 字符串长度不少于1，不超过128。
	Owner *string `json:"owner,omitempty"`

	// **参数解释：** 仓库容量,单位:MB,保留2位小数。 **取值范围：** 不涉及。
	Capacity *float64 `json:"capacity,omitempty"`

	// **参数解释：** 仓库状态。 **取值范围：** - 0，正常。 - 3，冻结。 - 4，关闭。 - 5，清理中。 - 7，删除中。
	Status *TenantRepositoryDtoStatus `json:"status,omitempty"`

	// **参数解释：** 内容审核结果。 **取值范围：** true，审核通过。 false，审核未通过。
	ModerationResult *bool `json:"moderation_result,omitempty"`

	// **参数解释：** 创建时间。 **取值范围：** 不涉及。
	CreateTime *string `json:"create_time,omitempty"`

	// **参数解释：** 成员数量。 **取值范围：** 不涉及。
	MemberNumber *int32 `json:"member_number,omitempty"`

	// **参数解释：** 仓库Id。 **取值范围：** 不涉及。
	RepositoryId *int32 `json:"repository_id,omitempty"`

	// **参数解释：** 仓库名。 **取值范围：** 不涉及。
	RepositoryName *string `json:"repository_name,omitempty"`

	// **参数解释：** 项目名。 **取值范围：** 不涉及。
	ProjectName *string `json:"project_name,omitempty"`

	// **参数解释：** 项目Id。 **取值范围：** 不涉及。
	ProjectId *string `json:"project_id,omitempty"`
}

func (o TenantRepositoryDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TenantRepositoryDto struct{}"
	}

	return strings.Join([]string{"TenantRepositoryDto", string(data)}, " ")
}

type TenantRepositoryDtoStatus struct {
	value int32
}

type TenantRepositoryDtoStatusEnum struct {
	E_0 TenantRepositoryDtoStatus
	E_3 TenantRepositoryDtoStatus
	E_4 TenantRepositoryDtoStatus
	E_5 TenantRepositoryDtoStatus
	E_7 TenantRepositoryDtoStatus
}

func GetTenantRepositoryDtoStatusEnum() TenantRepositoryDtoStatusEnum {
	return TenantRepositoryDtoStatusEnum{
		E_0: TenantRepositoryDtoStatus{
			value: 0,
		}, E_3: TenantRepositoryDtoStatus{
			value: 3,
		}, E_4: TenantRepositoryDtoStatus{
			value: 4,
		}, E_5: TenantRepositoryDtoStatus{
			value: 5,
		}, E_7: TenantRepositoryDtoStatus{
			value: 7,
		},
	}
}

func (c TenantRepositoryDtoStatus) Value() int32 {
	return c.value
}

func (c TenantRepositoryDtoStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *TenantRepositoryDtoStatus) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("int32")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: int32")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(int32); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to int32 error")
	}
}
