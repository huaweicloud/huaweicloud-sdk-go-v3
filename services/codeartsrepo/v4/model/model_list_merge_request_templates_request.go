package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListMergeRequestTemplatesRequest Request Object
type ListMergeRequestTemplatesRequest struct {

	// **参数解释：** 仓库的ID，通过[[查询用户所有仓库](https://support.huaweicloud.com/api-codeartsrepo/ListUserAllRepositories.html)](tag:hws)[[查询用户所有仓库](https://support.huaweicloud.com/intl/en-us/api-codeartsrepo/ListUserAllRepositories.html)](tag:hws_hk)[查询项目列表](tag:hcs,hcs_sm)接口查询项目列表获取。 **约束限制：** 不涉及。
	RepositoryId int32 `json:"repository_id"`

	// **参数解释：** 结果集返回的范围 **取值范围：** - default，返回模板名称为“default template”的模板。 - list，返回模板列表。
	View ListMergeRequestTemplatesRequestView `json:"view"`

	// **参数解释：** 搜索的模板名称 **取值范围：** 字符串长度不少于1，不超过100000。
	TemplateName *string `json:"template_name,omitempty"`

	// **参数解释：** 偏移量，从0开始。
	Offset *int32 `json:"offset,omitempty"`

	// **参数解释：** 返回数量。
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListMergeRequestTemplatesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListMergeRequestTemplatesRequest struct{}"
	}

	return strings.Join([]string{"ListMergeRequestTemplatesRequest", string(data)}, " ")
}

type ListMergeRequestTemplatesRequestView struct {
	value string
}

type ListMergeRequestTemplatesRequestViewEnum struct {
	DEFAULT ListMergeRequestTemplatesRequestView
	LIST    ListMergeRequestTemplatesRequestView
}

func GetListMergeRequestTemplatesRequestViewEnum() ListMergeRequestTemplatesRequestViewEnum {
	return ListMergeRequestTemplatesRequestViewEnum{
		DEFAULT: ListMergeRequestTemplatesRequestView{
			value: "default",
		},
		LIST: ListMergeRequestTemplatesRequestView{
			value: "list",
		},
	}
}

func (c ListMergeRequestTemplatesRequestView) Value() string {
	return c.value
}

func (c ListMergeRequestTemplatesRequestView) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListMergeRequestTemplatesRequestView) UnmarshalJSON(b []byte) error {
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
