package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListRepositoryContributorsRequest Request Object
type ListRepositoryContributorsRequest struct {

	// **参数解释：** 仓库的ID，通过[[查询用户所有仓库](https://support.huaweicloud.com/api-codeartsrepo/ListUserAllRepositories.html)](tag:hws)[[查询用户所有仓库](https://support.huaweicloud.com/intl/en-us/api-codeartsrepo/ListUserAllRepositories.html)](tag:hws_hk)[[查询用户所有仓库](https://support.huaweicloud.com/eu/api-codeartsrepo/ListUserAllRepositories.html)](tag:hws_eu)[查询项目列表](tag:hcs,hcs_sm)接口查询项目列表获取。 **约束限制：** 不涉及。
	RepositoryId int32 `json:"repository_id"`

	// **参数解释：** 偏移量，从0开始。
	Offset *int32 `json:"offset,omitempty"`

	// **参数解释：** 返回数量。
	Limit *int32 `json:"limit,omitempty"`

	// **参数解释：** 排序方式。 **取值范围：** - name，名字。 - email，邮箱。 - commits，提交数量。
	OrderBy *ListRepositoryContributorsRequestOrderBy `json:"order_by,omitempty"`

	// **参数解释：** 返回排序 - asc 按提交数量正序返回 - desc 按提交数量倒序返回
	Sort *ListRepositoryContributorsRequestSort `json:"sort,omitempty"`

	// **参数解释：** 分支或者tag名称。 **约束限制：** 不支持以 - . refs/heads/ refs/remotes/ 开头，不支持空格 [ \\ < ~ ^ : ? * ! ( ) ' \" | 等特殊字符，不支持以. / .lock结尾。 **取值范围：** 字符串长度不少于1，不超过200。 **默认取值：** 仓库默认分支。
	RefName *string `json:"ref_name,omitempty"`

	// **参数解释：** 是否跳过合并。 **约束限制：** 不涉及。 **取值范围：** - false, 跳过合并 - true, 不跳过合并
	SkipMerge *bool `json:"skip_merge,omitempty"`

	// **参数解释：** 提交者。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	Author *string `json:"author,omitempty"`
}

func (o ListRepositoryContributorsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRepositoryContributorsRequest struct{}"
	}

	return strings.Join([]string{"ListRepositoryContributorsRequest", string(data)}, " ")
}

type ListRepositoryContributorsRequestOrderBy struct {
	value string
}

type ListRepositoryContributorsRequestOrderByEnum struct {
	NAME    ListRepositoryContributorsRequestOrderBy
	EMAIL   ListRepositoryContributorsRequestOrderBy
	COMMITS ListRepositoryContributorsRequestOrderBy
}

func GetListRepositoryContributorsRequestOrderByEnum() ListRepositoryContributorsRequestOrderByEnum {
	return ListRepositoryContributorsRequestOrderByEnum{
		NAME: ListRepositoryContributorsRequestOrderBy{
			value: "name",
		},
		EMAIL: ListRepositoryContributorsRequestOrderBy{
			value: "email",
		},
		COMMITS: ListRepositoryContributorsRequestOrderBy{
			value: "commits",
		},
	}
}

func (c ListRepositoryContributorsRequestOrderBy) Value() string {
	return c.value
}

func (c ListRepositoryContributorsRequestOrderBy) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListRepositoryContributorsRequestOrderBy) UnmarshalJSON(b []byte) error {
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

type ListRepositoryContributorsRequestSort struct {
	value string
}

type ListRepositoryContributorsRequestSortEnum struct {
	ASC  ListRepositoryContributorsRequestSort
	DESC ListRepositoryContributorsRequestSort
}

func GetListRepositoryContributorsRequestSortEnum() ListRepositoryContributorsRequestSortEnum {
	return ListRepositoryContributorsRequestSortEnum{
		ASC: ListRepositoryContributorsRequestSort{
			value: "asc",
		},
		DESC: ListRepositoryContributorsRequestSort{
			value: "desc",
		},
	}
}

func (c ListRepositoryContributorsRequestSort) Value() string {
	return c.value
}

func (c ListRepositoryContributorsRequestSort) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListRepositoryContributorsRequestSort) UnmarshalJSON(b []byte) error {
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
