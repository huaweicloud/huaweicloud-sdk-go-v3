package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListMembersRequest Request Object
type ListMembersRequest struct {

	// **参数解释：** 仓库的ID，通过[[查询用户所有仓库](https://support.huaweicloud.com/api-codeartsrepo/ListUserAllRepositories.html)](tag:hws)[[查询用户所有仓库](https://support.huaweicloud.com/intl/en-us/api-codeartsrepo/ListUserAllRepositories.html)](tag:hws_hk)[查询项目列表](tag:hcs,hcs_sm)接口查询项目列表获取。 **约束限制：** 不涉及。
	RepositoryId int32 `json:"repository_id"`

	// **参数解释：** 查询关键字，可模糊匹配用户名称、用户昵称、租户名称。
	Search *string `json:"search,omitempty"`

	// **参数解释：** 偏移量，从0开始。
	Offset *int32 `json:"offset,omitempty"`

	// **参数解释：** 返回数量。
	Limit *int32 `json:"limit,omitempty"`

	// **参数解释：** 权限点。 **约束限制：** - repository，仓库。 - code，代码。 - member，成员。 - branch，分支。 - tag，Tag。 - mr，MR。 - label，标签。
	Permission *ListMembersRequestPermission `json:"permission,omitempty"`

	// **参数解释：** 权限动作点, 不同权限点有不同的动作。 **约束限制：** - repository：create,fork,delete,setting - code：push,download - member：create,update,delete - branch：create,delete - tag：create,delete - mr：create,update,comment,review,approve,merge,close,reopen - label：create,delete,update
	Action *ListMembersRequestAction `json:"action,omitempty"`
}

func (o ListMembersRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListMembersRequest struct{}"
	}

	return strings.Join([]string{"ListMembersRequest", string(data)}, " ")
}

type ListMembersRequestPermission struct {
	value string
}

type ListMembersRequestPermissionEnum struct {
	REPOSITORY ListMembersRequestPermission
	CODE       ListMembersRequestPermission
	MEMBER     ListMembersRequestPermission
	BRANCH     ListMembersRequestPermission
	TAG        ListMembersRequestPermission
	MR         ListMembersRequestPermission
	LABEL      ListMembersRequestPermission
}

func GetListMembersRequestPermissionEnum() ListMembersRequestPermissionEnum {
	return ListMembersRequestPermissionEnum{
		REPOSITORY: ListMembersRequestPermission{
			value: "repository",
		},
		CODE: ListMembersRequestPermission{
			value: "code",
		},
		MEMBER: ListMembersRequestPermission{
			value: "member",
		},
		BRANCH: ListMembersRequestPermission{
			value: "branch",
		},
		TAG: ListMembersRequestPermission{
			value: "tag",
		},
		MR: ListMembersRequestPermission{
			value: "mr",
		},
		LABEL: ListMembersRequestPermission{
			value: "label",
		},
	}
}

func (c ListMembersRequestPermission) Value() string {
	return c.value
}

func (c ListMembersRequestPermission) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListMembersRequestPermission) UnmarshalJSON(b []byte) error {
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

type ListMembersRequestAction struct {
	value string
}

type ListMembersRequestActionEnum struct {
	REPOSITORYCREATEFORKDELETESETTING                  ListMembersRequestAction
	CODEPUSHDOWNLOAD                                   ListMembersRequestAction
	MEMBERCREATEUPDATEDELETE                           ListMembersRequestAction
	BRANCHCREATEDELETE                                 ListMembersRequestAction
	TAGCREATEDELETE                                    ListMembersRequestAction
	MRCREATEUPDATECOMMENTREVIEWAPPROVEMERGECLOSEREOPEN ListMembersRequestAction
	LABELCREATEDELETEUPDATE                            ListMembersRequestAction
}

func GetListMembersRequestActionEnum() ListMembersRequestActionEnum {
	return ListMembersRequestActionEnum{
		REPOSITORYCREATEFORKDELETESETTING: ListMembersRequestAction{
			value: "repository：create,fork,delete,setting",
		},
		CODEPUSHDOWNLOAD: ListMembersRequestAction{
			value: "code：push,download",
		},
		MEMBERCREATEUPDATEDELETE: ListMembersRequestAction{
			value: "member：create,update,delete",
		},
		BRANCHCREATEDELETE: ListMembersRequestAction{
			value: "branch：create,delete",
		},
		TAGCREATEDELETE: ListMembersRequestAction{
			value: "tag：create,delete",
		},
		MRCREATEUPDATECOMMENTREVIEWAPPROVEMERGECLOSEREOPEN: ListMembersRequestAction{
			value: "mr：create,update,comment,review,approve,merge,close,reopen",
		},
		LABELCREATEDELETEUPDATE: ListMembersRequestAction{
			value: "label：create,delete,update",
		},
	}
}

func (c ListMembersRequestAction) Value() string {
	return c.value
}

func (c ListMembersRequestAction) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListMembersRequestAction) UnmarshalJSON(b []byte) error {
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
