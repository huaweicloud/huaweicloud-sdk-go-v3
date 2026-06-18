package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListTenantEncryptedRepositoriesRequest Request Object
type ListTenantEncryptedRepositoriesRequest struct {

	// **参数解释：** 租户id
	TenantId string `json:"tenant_id"`

	// **参数解释：** 偏移量，从0开始。
	Offset *int32 `json:"offset,omitempty"`

	// **参数解释：** 返回数量。
	Limit *int32 `json:"limit,omitempty"`

	// **参数解释：** 排序字段 repoName 仓库名称  ownerName | 仓库所有者名称 ownerName, 不传该字段时不进行排序
	OrderBy *ListTenantEncryptedRepositoriesRequestOrderBy `json:"order_by,omitempty"`

	// **参数解释：** 排序顺序 asc顺序 desc逆序
	Sort *ListTenantEncryptedRepositoriesRequestSort `json:"sort,omitempty"`
}

func (o ListTenantEncryptedRepositoriesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTenantEncryptedRepositoriesRequest struct{}"
	}

	return strings.Join([]string{"ListTenantEncryptedRepositoriesRequest", string(data)}, " ")
}

type ListTenantEncryptedRepositoriesRequestOrderBy struct {
	value string
}

type ListTenantEncryptedRepositoriesRequestOrderByEnum struct {
	REPO_NAME  ListTenantEncryptedRepositoriesRequestOrderBy
	OWNER_NAME ListTenantEncryptedRepositoriesRequestOrderBy
}

func GetListTenantEncryptedRepositoriesRequestOrderByEnum() ListTenantEncryptedRepositoriesRequestOrderByEnum {
	return ListTenantEncryptedRepositoriesRequestOrderByEnum{
		REPO_NAME: ListTenantEncryptedRepositoriesRequestOrderBy{
			value: "repoName",
		},
		OWNER_NAME: ListTenantEncryptedRepositoriesRequestOrderBy{
			value: "ownerName",
		},
	}
}

func (c ListTenantEncryptedRepositoriesRequestOrderBy) Value() string {
	return c.value
}

func (c ListTenantEncryptedRepositoriesRequestOrderBy) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListTenantEncryptedRepositoriesRequestOrderBy) UnmarshalJSON(b []byte) error {
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

type ListTenantEncryptedRepositoriesRequestSort struct {
	value string
}

type ListTenantEncryptedRepositoriesRequestSortEnum struct {
	ASC  ListTenantEncryptedRepositoriesRequestSort
	DESC ListTenantEncryptedRepositoriesRequestSort
}

func GetListTenantEncryptedRepositoriesRequestSortEnum() ListTenantEncryptedRepositoriesRequestSortEnum {
	return ListTenantEncryptedRepositoriesRequestSortEnum{
		ASC: ListTenantEncryptedRepositoriesRequestSort{
			value: "asc",
		},
		DESC: ListTenantEncryptedRepositoriesRequestSort{
			value: "desc",
		},
	}
}

func (c ListTenantEncryptedRepositoriesRequestSort) Value() string {
	return c.value
}

func (c ListTenantEncryptedRepositoriesRequestSort) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListTenantEncryptedRepositoriesRequestSort) UnmarshalJSON(b []byte) error {
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
