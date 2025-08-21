package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListMergeRequestValidAssignedCandidatesRequest Request Object
type ListMergeRequestValidAssignedCandidatesRequest struct {

	// **参数解释：** 仓库的ID，通过[[查询用户所有仓库](https://support.huaweicloud.com/api-codeartsrepo/ListUserAllRepositories.html)](tag:hws)[[查询用户所有仓库](https://support.huaweicloud.com/intl/en-us/api-codeartsrepo/ListUserAllRepositories.html)](tag:hws_hk)[查询项目列表](tag:hcs,hcs_sm)接口查询项目列表获取。 **约束限制：** 不涉及。
	RepositoryId int32 `json:"repository_id"`

	// **参数解释：** 目标分支。创建MR时，代码将要合入的分支。
	TargetBranch *string `json:"target_branch,omitempty"`

	// **参数解释：**  合并请求 iid。
	MergeRequestIid *int32 `json:"merge_request_iid,omitempty"`

	// **参数解释：** 偏移量，从0开始。
	Offset *int32 `json:"offset,omitempty"`

	// **参数解释：** 返回数量。
	Limit *int32 `json:"limit,omitempty"`

	// **参数解释：** 查询关键字，可模糊匹配用户名称、用户昵称、租户名称。
	Search *string `json:"search,omitempty"`

	// **参数解释：** Search user list by name list。
	SearchByNameList *string `json:"search_by_name_list,omitempty"`

	// **参数解释：** 项目的32位uuid，项目唯一标识，通过[[查询项目列表](https://support.huaweicloud.com/api-projectman/ListProjectsV4.html)](tag:hws)[[查询项目列表](https://support.huaweicloud.com/intl/en-us/api-projectman/ListProjectsV4.html)](tag:hws_hk)[查询项目列表](tag:hcs,hcs_sm)接口查询项目列表获取。 **取值范围：** 字符串长度32。
	TargetProjectId *string `json:"target_project_id,omitempty"`

	// **参数解释：** The type of assignee, merge user or approver
	View *ListMergeRequestValidAssignedCandidatesRequestView `json:"view,omitempty"`

	// **参数解释：** The type of assignee, merge user or approver
	Mode *ListMergeRequestValidAssignedCandidatesRequestMode `json:"mode,omitempty"`

	// **参数解释：** The type of memeber, developer
	OnlyDevelopers *bool `json:"only_developers,omitempty"`
}

func (o ListMergeRequestValidAssignedCandidatesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListMergeRequestValidAssignedCandidatesRequest struct{}"
	}

	return strings.Join([]string{"ListMergeRequestValidAssignedCandidatesRequest", string(data)}, " ")
}

type ListMergeRequestValidAssignedCandidatesRequestView struct {
	value string
}

type ListMergeRequestValidAssignedCandidatesRequestViewEnum struct {
	APPROVER ListMergeRequestValidAssignedCandidatesRequestView
	ASSIGNEE ListMergeRequestValidAssignedCandidatesRequestView
}

func GetListMergeRequestValidAssignedCandidatesRequestViewEnum() ListMergeRequestValidAssignedCandidatesRequestViewEnum {
	return ListMergeRequestValidAssignedCandidatesRequestViewEnum{
		APPROVER: ListMergeRequestValidAssignedCandidatesRequestView{
			value: "approver",
		},
		ASSIGNEE: ListMergeRequestValidAssignedCandidatesRequestView{
			value: "assignee",
		},
	}
}

func (c ListMergeRequestValidAssignedCandidatesRequestView) Value() string {
	return c.value
}

func (c ListMergeRequestValidAssignedCandidatesRequestView) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListMergeRequestValidAssignedCandidatesRequestView) UnmarshalJSON(b []byte) error {
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

type ListMergeRequestValidAssignedCandidatesRequestMode struct {
	value string
}

type ListMergeRequestValidAssignedCandidatesRequestModeEnum struct {
	APPROVER ListMergeRequestValidAssignedCandidatesRequestMode
	ASSIGNEE ListMergeRequestValidAssignedCandidatesRequestMode
}

func GetListMergeRequestValidAssignedCandidatesRequestModeEnum() ListMergeRequestValidAssignedCandidatesRequestModeEnum {
	return ListMergeRequestValidAssignedCandidatesRequestModeEnum{
		APPROVER: ListMergeRequestValidAssignedCandidatesRequestMode{
			value: "approver",
		},
		ASSIGNEE: ListMergeRequestValidAssignedCandidatesRequestMode{
			value: "assignee",
		},
	}
}

func (c ListMergeRequestValidAssignedCandidatesRequestMode) Value() string {
	return c.value
}

func (c ListMergeRequestValidAssignedCandidatesRequestMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListMergeRequestValidAssignedCandidatesRequestMode) UnmarshalJSON(b []byte) error {
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
