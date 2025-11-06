package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListGroupMembersRequest Request Object
type ListGroupMembersRequest struct {

	// **参数解释：** 代码组id，代码组首页，Group ID后的数字Id
	GroupId int32 `json:"group_id"`

	// **参数解释：** 项目的32位uuid，项目唯一标识，通过[[查询项目列表](https://support.huaweicloud.com/api-projectman/ListProjectsV4.html)](tag:hws)[[查询项目列表](https://support.huaweicloud.com/intl/en-us/api-projectman/ListProjectsV4.html)](tag:hws_hk)[查询项目列表](tag:hcs,hcs_sm)接口查询项目列表获取。 **取值范围：** 字符串长度32。
	ProjectId string `json:"project_id"`

	// **参数解释：** 成员搜索字符串
	Query *string `json:"query,omitempty"`

	// **参数解释：** 成员加入方式 domain 租户 normal 普通 inherit 继承
	JoinWay *ListGroupMembersRequestJoinWay `json:"join_way,omitempty"`

	// **参数解释：** 过滤成员的access level， 10待审核 20浏览者 30开发者 40管理员 50所有者
	AccessLevel *int32 `json:"access_level,omitempty"`

	// **参数解释：** 偏移量，从0开始。
	Offset *int32 `json:"offset,omitempty"`

	// **参数解释：** 返回数量。
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListGroupMembersRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListGroupMembersRequest struct{}"
	}

	return strings.Join([]string{"ListGroupMembersRequest", string(data)}, " ")
}

type ListGroupMembersRequestJoinWay struct {
	value string
}

type ListGroupMembersRequestJoinWayEnum struct {
	DOMAIN  ListGroupMembersRequestJoinWay
	NORMAL  ListGroupMembersRequestJoinWay
	INHERIT ListGroupMembersRequestJoinWay
}

func GetListGroupMembersRequestJoinWayEnum() ListGroupMembersRequestJoinWayEnum {
	return ListGroupMembersRequestJoinWayEnum{
		DOMAIN: ListGroupMembersRequestJoinWay{
			value: "domain",
		},
		NORMAL: ListGroupMembersRequestJoinWay{
			value: "normal",
		},
		INHERIT: ListGroupMembersRequestJoinWay{
			value: "inherit",
		},
	}
}

func (c ListGroupMembersRequestJoinWay) Value() string {
	return c.value
}

func (c ListGroupMembersRequestJoinWay) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListGroupMembersRequestJoinWay) UnmarshalJSON(b []byte) error {
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
