package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListRepositoryForksRequest Request Object
type ListRepositoryForksRequest struct {

	// **参数解释：** 仓库的ID，通过[[查询用户所有仓库](https://support.huaweicloud.com/api-codeartsrepo/ListUserAllRepositories.html)](tag:hws)[[查询用户所有仓库](https://support.huaweicloud.com/intl/en-us/api-codeartsrepo/ListUserAllRepositories.html)](tag:hws_hk)[[查询用户所有仓库](https://support.huaweicloud.com/eu/api-codeartsrepo/ListUserAllRepositories.html)](tag:hws_eu)[查询项目列表](tag:hcs,hcs_sm)接口查询项目列表获取。 **约束限制：** 不涉及。
	RepositoryId int32 `json:"repository_id"`

	// **参数解释：** 偏移量，从0开始。
	Offset *int32 `json:"offset,omitempty"`

	// **参数解释：** 返回数量。
	Limit *int32 `json:"limit,omitempty"`

	// **参数解释：**  排序字段。 **约束限制：**  必须为枚举值中的选项。 **取值范围：**  - created_at，创建时间。 - updated_at，更新时间。 **默认取值：**  created_at。
	OrderBy *ListRepositoryForksRequestOrderBy `json:"order_by,omitempty"`

	// **参数解释：** 排列顺序。 **约束限制：** 必须为枚举值中的选项。 **取值范围：**  - asc - desc **默认取值：** desc。
	Sort *ListRepositoryForksRequestSort `json:"sort,omitempty"`

	// **参数解释：**  视图。 **约束限制：**  必须为枚举值中的选项。 **取值范围：**  - basic，基本信息。  - least，最简信息。 **默认取值：**  least。
	View *ListRepositoryForksRequestView `json:"view,omitempty"`
}

func (o ListRepositoryForksRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRepositoryForksRequest struct{}"
	}

	return strings.Join([]string{"ListRepositoryForksRequest", string(data)}, " ")
}

type ListRepositoryForksRequestOrderBy struct {
	value string
}

type ListRepositoryForksRequestOrderByEnum struct {
	CREATED_AT ListRepositoryForksRequestOrderBy
	UPDATED_AT ListRepositoryForksRequestOrderBy
}

func GetListRepositoryForksRequestOrderByEnum() ListRepositoryForksRequestOrderByEnum {
	return ListRepositoryForksRequestOrderByEnum{
		CREATED_AT: ListRepositoryForksRequestOrderBy{
			value: "created_at",
		},
		UPDATED_AT: ListRepositoryForksRequestOrderBy{
			value: "updated_at",
		},
	}
}

func (c ListRepositoryForksRequestOrderBy) Value() string {
	return c.value
}

func (c ListRepositoryForksRequestOrderBy) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListRepositoryForksRequestOrderBy) UnmarshalJSON(b []byte) error {
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

type ListRepositoryForksRequestSort struct {
	value string
}

type ListRepositoryForksRequestSortEnum struct {
	ASC  ListRepositoryForksRequestSort
	DESC ListRepositoryForksRequestSort
}

func GetListRepositoryForksRequestSortEnum() ListRepositoryForksRequestSortEnum {
	return ListRepositoryForksRequestSortEnum{
		ASC: ListRepositoryForksRequestSort{
			value: "asc",
		},
		DESC: ListRepositoryForksRequestSort{
			value: "desc",
		},
	}
}

func (c ListRepositoryForksRequestSort) Value() string {
	return c.value
}

func (c ListRepositoryForksRequestSort) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListRepositoryForksRequestSort) UnmarshalJSON(b []byte) error {
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

type ListRepositoryForksRequestView struct {
	value string
}

type ListRepositoryForksRequestViewEnum struct {
	BASIC ListRepositoryForksRequestView
	LEAST ListRepositoryForksRequestView
}

func GetListRepositoryForksRequestViewEnum() ListRepositoryForksRequestViewEnum {
	return ListRepositoryForksRequestViewEnum{
		BASIC: ListRepositoryForksRequestView{
			value: "basic",
		},
		LEAST: ListRepositoryForksRequestView{
			value: "least",
		},
	}
}

func (c ListRepositoryForksRequestView) Value() string {
	return c.value
}

func (c ListRepositoryForksRequestView) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListRepositoryForksRequestView) UnmarshalJSON(b []byte) error {
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
