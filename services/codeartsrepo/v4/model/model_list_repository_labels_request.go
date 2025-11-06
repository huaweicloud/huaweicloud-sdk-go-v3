package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListRepositoryLabelsRequest Request Object
type ListRepositoryLabelsRequest struct {

	// **参数解释：** 仓库的ID，通过[[查询用户所有仓库](https://support.huaweicloud.com/api-codeartsrepo/ListUserAllRepositories.html)](tag:hws)[[查询用户所有仓库](https://support.huaweicloud.com/intl/en-us/api-codeartsrepo/ListUserAllRepositories.html)](tag:hws_hk)[查询项目列表](tag:hcs,hcs_sm)接口查询项目列表获取。 **约束限制：** 不涉及。
	RepositoryId int32 `json:"repository_id"`

	// **参数解释：** 偏移量，从0开始。
	Offset *int32 `json:"offset,omitempty"`

	// **参数解释：** 返回数量。
	Limit *int32 `json:"limit,omitempty"`

	// **参数解释：** 查询关键字，可模糊匹配标签名。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	Search *string `json:"search,omitempty"`

	// **参数解释：**  排序。 **约束限制：** 不涉及。 **取值范围：** - name_asc，按标签名升序。 - name_desc，按标签名降序。 - created_asc，按标签创建时间升序。 - created_desc，按标签创建时间降序。 - updated_asc，按标签更新时间升序。 - updated_desc，按标签更新时间降序。 **默认取值：** name_asc
	Sort *ListRepositoryLabelsRequestSort `json:"sort,omitempty"`

	// **参数解释：** 是否包含失效的标签。 **约束限制：** 不涉及。 **取值范围：** - true，包含。 - false，不包含。 **默认取值：** true
	IncludeExpired *bool `json:"include_expired,omitempty"`

	// **参数解释：** 结果集属性，根据给定的参数返回不同的结果。 **约束限制：** 不涉及。 **取值范围：** - simple，返回简略信息。 - basic，返回基本信息。 - detail，返回详细信息。 **默认取值：** basic
	View *ListRepositoryLabelsRequestView `json:"view,omitempty"`
}

func (o ListRepositoryLabelsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRepositoryLabelsRequest struct{}"
	}

	return strings.Join([]string{"ListRepositoryLabelsRequest", string(data)}, " ")
}

type ListRepositoryLabelsRequestSort struct {
	value string
}

type ListRepositoryLabelsRequestSortEnum struct {
	NAME_ASC     ListRepositoryLabelsRequestSort
	NAME_DESC    ListRepositoryLabelsRequestSort
	CREATED_ASC  ListRepositoryLabelsRequestSort
	CREATED_DESC ListRepositoryLabelsRequestSort
	UPDATED_ASC  ListRepositoryLabelsRequestSort
	UPDATED_DESC ListRepositoryLabelsRequestSort
}

func GetListRepositoryLabelsRequestSortEnum() ListRepositoryLabelsRequestSortEnum {
	return ListRepositoryLabelsRequestSortEnum{
		NAME_ASC: ListRepositoryLabelsRequestSort{
			value: "name_asc",
		},
		NAME_DESC: ListRepositoryLabelsRequestSort{
			value: "name_desc",
		},
		CREATED_ASC: ListRepositoryLabelsRequestSort{
			value: "created_asc",
		},
		CREATED_DESC: ListRepositoryLabelsRequestSort{
			value: "created_desc",
		},
		UPDATED_ASC: ListRepositoryLabelsRequestSort{
			value: "updated_asc",
		},
		UPDATED_DESC: ListRepositoryLabelsRequestSort{
			value: "updated_desc",
		},
	}
}

func (c ListRepositoryLabelsRequestSort) Value() string {
	return c.value
}

func (c ListRepositoryLabelsRequestSort) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListRepositoryLabelsRequestSort) UnmarshalJSON(b []byte) error {
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

type ListRepositoryLabelsRequestView struct {
	value string
}

type ListRepositoryLabelsRequestViewEnum struct {
	SIMPLE ListRepositoryLabelsRequestView
	BASIC  ListRepositoryLabelsRequestView
	DETAIL ListRepositoryLabelsRequestView
}

func GetListRepositoryLabelsRequestViewEnum() ListRepositoryLabelsRequestViewEnum {
	return ListRepositoryLabelsRequestViewEnum{
		SIMPLE: ListRepositoryLabelsRequestView{
			value: "simple",
		},
		BASIC: ListRepositoryLabelsRequestView{
			value: "basic",
		},
		DETAIL: ListRepositoryLabelsRequestView{
			value: "detail",
		},
	}
}

func (c ListRepositoryLabelsRequestView) Value() string {
	return c.value
}

func (c ListRepositoryLabelsRequestView) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListRepositoryLabelsRequestView) UnmarshalJSON(b []byte) error {
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
