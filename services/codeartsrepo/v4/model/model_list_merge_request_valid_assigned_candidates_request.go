package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListMergeRequestValidAssignedCandidatesRequest Request Object
type ListMergeRequestValidAssignedCandidatesRequest struct {

	// **参数解释：** 仓库的ID，通过[[查询用户所有仓库](https://support.huaweicloud.com/api-codeartsrepo/ListUserAllRepositories.html)](tag:hws)[[查询用户所有仓库](https://support.huaweicloud.com/intl/en-us/api-codeartsrepo/ListUserAllRepositories.html)](tag:hws_hk)[[查询用户所有仓库](https://support.huaweicloud.com/intl/zh-cn/api-codeartsrepo/ListUserAllRepositories.html)](tag:hws_hk_ch)[[查询用户所有仓库](https://support.huaweicloud.com/eu/api-codeartsrepo/ListUserAllRepositories.html)](tag:hws_eu)[查询项目列表](tag:hcs,hcs_sm)接口查询项目列表获取。 **约束限制：** 不涉及。 **默认取值：** 不涉及。
	RepositoryId int32 `json:"repository_id"`

	// **参数解释：** 目标分支。创建MR时，代码将要合入的分支。
	TargetBranch string `json:"target_branch"`

	// **参数解释：**  合并请求 iid。
	MergeRequestIid *int32 `json:"merge_request_iid,omitempty"`

	// **参数解释：** 偏移量，从0开始。
	Offset *int32 `json:"offset,omitempty"`

	// **参数解释：** 返回数量。
	Limit *int32 `json:"limit,omitempty"`

	// **参数解释：** 查询关键字，可模糊匹配用户名称、用户昵称、租户名称。
	Search *string `json:"search,omitempty"`

	// **参数解释：** 批量匹配用户，一次可传多个用户的用户名和昵称，用\", \"分隔，最多50个。示例：\"zhangsan, lisi, wangwu\"
	SearchByNameList *string `json:"search_by_name_list,omitempty"`

	// **参数解释：** 目标仓库id。创建MR时，代码将要合入的仓库。
	TargetRepositoryId *string `json:"target_repository_id,omitempty"`

	// **参数解释：** approver: 获取审核人 assingee: 获取合并人
	View *ListMergeRequestValidAssignedCandidatesRequestView `json:"view,omitempty"`

	// **参数解释：** true: 仅返回开发者。
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
