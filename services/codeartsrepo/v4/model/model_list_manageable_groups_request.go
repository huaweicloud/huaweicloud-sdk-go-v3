package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListManageableGroupsRequest Request Object
type ListManageableGroupsRequest struct {

	// **参数解释：** 项目的32位uuid，项目唯一标识，通过[[查询项目列表](https://support.huaweicloud.com/api-projectman/ListProjectsV4.html)](tag:hws)[[查询项目列表](https://support.huaweicloud.com/intl/en-us/api-projectman/ListProjectsV4.html)](tag:hws_hk)[[查询项目列表](https://support.huaweicloud.com/eu/api-projectman/ListProjectsV4.html)](tag:hws_eu)[查询项目列表](tag:hcs,hcs_sm)接口查询项目列表获取。 **取值范围：** 字符串长度32。
	ProjectId string `json:"project_id"`

	// **参数解释：** 创建资源类型。 **约束限制：** - group 代码组。 - repository仓库。
	Scope *ListManageableGroupsRequestScope `json:"scope,omitempty"`

	// **参数解释：** 偏移量，从0开始。
	Offset *int32 `json:"offset,omitempty"`

	// **参数解释：** 返回数量。
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListManageableGroupsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListManageableGroupsRequest struct{}"
	}

	return strings.Join([]string{"ListManageableGroupsRequest", string(data)}, " ")
}

type ListManageableGroupsRequestScope struct {
	value string
}

type ListManageableGroupsRequestScopeEnum struct {
	GROUP      ListManageableGroupsRequestScope
	REPOSITORY ListManageableGroupsRequestScope
}

func GetListManageableGroupsRequestScopeEnum() ListManageableGroupsRequestScopeEnum {
	return ListManageableGroupsRequestScopeEnum{
		GROUP: ListManageableGroupsRequestScope{
			value: "group",
		},
		REPOSITORY: ListManageableGroupsRequestScope{
			value: "repository",
		},
	}
}

func (c ListManageableGroupsRequestScope) Value() string {
	return c.value
}

func (c ListManageableGroupsRequestScope) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListManageableGroupsRequestScope) UnmarshalJSON(b []byte) error {
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
