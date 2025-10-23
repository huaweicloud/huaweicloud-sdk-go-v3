package model

import (
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"strings"
)

// ListTenantRepositoriesRequest Request Object
type ListTenantRepositoriesRequest struct {

	// **参数解释：** 仓库名称。 **取值范围：** 字符串长度不少于1，不超过128。
	RepositoryName *string `json:"repository_name,omitempty"`

	// **参数解释：** 成员数量。
	MemberNumber *int32 `json:"member_number,omitempty"`

	// **参数解释：** 仓库状态。 **取值范围：** - 0，正常。 - 3，冻结。 - 4，关闭。 - 5，清理中。 - 7，删除中。
	Status *ListTenantRepositoriesRequestStatus `json:"status,omitempty"`

	// **参数解释：** 仓库所有者。 **取值范围：** 字符串长度不少于1，不超过128。
	Owner *string `json:"owner,omitempty"`

	// **参数解释：** 筛选在该时间之后创建的仓库。
	CreatedAfter *sdktime.SdkTime `json:"created_after,omitempty"`

	// **参数解释：** 筛选在该时间之前创建的仓库。
	CreatedBefore *sdktime.SdkTime `json:"created_before,omitempty"`

	// **参数解释：** 结果集排序方式。 **约束限制：** 与sort_field搭配使用。  **取值范围：** - asc，正序返回。 - desc，倒序返回。
	Sort *ListTenantRepositoriesRequestSort `json:"sort,omitempty"`

	// **参数解释：** 用作排序的字段。 - owner，仓库所有者。 - capacity，使用空间。 - status，状态。 - create_time，创建时间。 - member_number，成员数量。 - repository_name，仓库名称。
	SortField *ListTenantRepositoriesRequestSortField `json:"sort_field,omitempty"`

	// **参数解释：** 偏移量，从0开始。
	Offset *int32 `json:"offset,omitempty"`

	// **参数解释：** 返回数量。
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListTenantRepositoriesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTenantRepositoriesRequest struct{}"
	}

	return strings.Join([]string{"ListTenantRepositoriesRequest", string(data)}, " ")
}

type ListTenantRepositoriesRequestStatus struct {
	value int32
}

type ListTenantRepositoriesRequestStatusEnum struct {
	E_0 ListTenantRepositoriesRequestStatus
	E_3 ListTenantRepositoriesRequestStatus
	E_4 ListTenantRepositoriesRequestStatus
	E_5 ListTenantRepositoriesRequestStatus
	E_7 ListTenantRepositoriesRequestStatus
}

func GetListTenantRepositoriesRequestStatusEnum() ListTenantRepositoriesRequestStatusEnum {
	return ListTenantRepositoriesRequestStatusEnum{
		E_0: ListTenantRepositoriesRequestStatus{
			value: 0,
		}, E_3: ListTenantRepositoriesRequestStatus{
			value: 3,
		}, E_4: ListTenantRepositoriesRequestStatus{
			value: 4,
		}, E_5: ListTenantRepositoriesRequestStatus{
			value: 5,
		}, E_7: ListTenantRepositoriesRequestStatus{
			value: 7,
		},
	}
}

func (c ListTenantRepositoriesRequestStatus) Value() int32 {
	return c.value
}

func (c ListTenantRepositoriesRequestStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListTenantRepositoriesRequestStatus) UnmarshalJSON(b []byte) error {
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

type ListTenantRepositoriesRequestSort struct {
	value string
}

type ListTenantRepositoriesRequestSortEnum struct {
	ASC  ListTenantRepositoriesRequestSort
	DESC ListTenantRepositoriesRequestSort
}

func GetListTenantRepositoriesRequestSortEnum() ListTenantRepositoriesRequestSortEnum {
	return ListTenantRepositoriesRequestSortEnum{
		ASC: ListTenantRepositoriesRequestSort{
			value: "asc",
		},
		DESC: ListTenantRepositoriesRequestSort{
			value: "desc",
		},
	}
}

func (c ListTenantRepositoriesRequestSort) Value() string {
	return c.value
}

func (c ListTenantRepositoriesRequestSort) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListTenantRepositoriesRequestSort) UnmarshalJSON(b []byte) error {
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

type ListTenantRepositoriesRequestSortField struct {
	value string
}

type ListTenantRepositoriesRequestSortFieldEnum struct {
	OWNER           ListTenantRepositoriesRequestSortField
	CAPACITY        ListTenantRepositoriesRequestSortField
	STATUS          ListTenantRepositoriesRequestSortField
	CREATE_TIME     ListTenantRepositoriesRequestSortField
	MEMBER_NUMBER   ListTenantRepositoriesRequestSortField
	REPOSITORY_NAME ListTenantRepositoriesRequestSortField
}

func GetListTenantRepositoriesRequestSortFieldEnum() ListTenantRepositoriesRequestSortFieldEnum {
	return ListTenantRepositoriesRequestSortFieldEnum{
		OWNER: ListTenantRepositoriesRequestSortField{
			value: "owner",
		},
		CAPACITY: ListTenantRepositoriesRequestSortField{
			value: "capacity",
		},
		STATUS: ListTenantRepositoriesRequestSortField{
			value: "status",
		},
		CREATE_TIME: ListTenantRepositoriesRequestSortField{
			value: "create_time",
		},
		MEMBER_NUMBER: ListTenantRepositoriesRequestSortField{
			value: "member_number",
		},
		REPOSITORY_NAME: ListTenantRepositoriesRequestSortField{
			value: "repository_name",
		},
	}
}

func (c ListTenantRepositoriesRequestSortField) Value() string {
	return c.value
}

func (c ListTenantRepositoriesRequestSortField) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListTenantRepositoriesRequestSortField) UnmarshalJSON(b []byte) error {
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
