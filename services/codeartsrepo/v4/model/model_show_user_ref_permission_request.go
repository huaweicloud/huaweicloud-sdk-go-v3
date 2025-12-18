package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowUserRefPermissionRequest Request Object
type ShowUserRefPermissionRequest struct {

	// **参数解释：** 仓库的ID，通过[[查询用户所有仓库](https://support.huaweicloud.com/api-codeartsrepo/ListUserAllRepositories.html)](tag:hws)[[查询用户所有仓库](https://support.huaweicloud.com/intl/en-us/api-codeartsrepo/ListUserAllRepositories.html)](tag:hws_hk)[[查询用户所有仓库](https://support.huaweicloud.com/eu/api-codeartsrepo/ListUserAllRepositories.html)](tag:hws_eu)[查询项目列表](tag:hcs,hcs_sm)接口查询项目列表获取。 **约束限制：** 不涉及。
	RepositoryId int32 `json:"repository_id"`

	// **参数解释：** 分支或者标签名称。  **约束限制：** 不支持空格 [ \\ < ~ ^ : ? * ! ( ) ' \" | 等特殊字符，不支持以. / .lock结尾，分支以refs/head/开头，标签以refs/tag/开头  **取值范围：** 字符串长度不少于1，不超过210。 **默认取值：** 不涉及。
	TargetRef string `json:"target_ref"`

	// **参数解释：** 动作类型，可用于查询指定动作的权限 - read，查看 - review，检视 - approval, 审核 - create-change，创建变更请求 - merge，合并变更请求 - create-delete，创建/删除分支 - push, 推送
	Action *ShowUserRefPermissionRequestAction `json:"action,omitempty"`

	// **参数解释：** 变更请求在仓库内部的ID。
	ChangeRequestIid *int32 `json:"change_request_iid,omitempty"`
}

func (o ShowUserRefPermissionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowUserRefPermissionRequest struct{}"
	}

	return strings.Join([]string{"ShowUserRefPermissionRequest", string(data)}, " ")
}

type ShowUserRefPermissionRequestAction struct {
	value string
}

type ShowUserRefPermissionRequestActionEnum struct {
	READ          ShowUserRefPermissionRequestAction
	REVIEW        ShowUserRefPermissionRequestAction
	APPROVAL      ShowUserRefPermissionRequestAction
	CREATE_CHANGE ShowUserRefPermissionRequestAction
	MERGE         ShowUserRefPermissionRequestAction
	CREATE_DELETE ShowUserRefPermissionRequestAction
	PUSH          ShowUserRefPermissionRequestAction
}

func GetShowUserRefPermissionRequestActionEnum() ShowUserRefPermissionRequestActionEnum {
	return ShowUserRefPermissionRequestActionEnum{
		READ: ShowUserRefPermissionRequestAction{
			value: "read",
		},
		REVIEW: ShowUserRefPermissionRequestAction{
			value: "review",
		},
		APPROVAL: ShowUserRefPermissionRequestAction{
			value: "approval",
		},
		CREATE_CHANGE: ShowUserRefPermissionRequestAction{
			value: "create-change",
		},
		MERGE: ShowUserRefPermissionRequestAction{
			value: "merge",
		},
		CREATE_DELETE: ShowUserRefPermissionRequestAction{
			value: "create-delete",
		},
		PUSH: ShowUserRefPermissionRequestAction{
			value: "push",
		},
	}
}

func (c ShowUserRefPermissionRequestAction) Value() string {
	return c.value
}

func (c ShowUserRefPermissionRequestAction) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowUserRefPermissionRequestAction) UnmarshalJSON(b []byte) error {
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
