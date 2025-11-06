package model

import (
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"strings"
)

// ShowGroupResponse Response Object
type ShowGroupResponse struct {

	// 创建者id
	CreatorId *int32 `json:"creator_id,omitempty"`

	// 项目id
	ProjectId *string `json:"project_id,omitempty"`

	// 创建时间
	CreatedAt *sdktime.SdkTime `json:"created_at,omitempty"`

	// 描述
	Description *string `json:"description,omitempty"`

	// 代码组全名
	FullName *string `json:"full_name,omitempty"`

	// 代码组层级路径id
	AncestorIds *[]int32 `json:"ancestor_ids,omitempty"`

	// 代码组层级路径名称
	AncestorNames *[]string `json:"ancestor_names,omitempty"`

	// 代码组id
	Id *int32 `json:"id,omitempty"`

	// 代码组成员计数
	MembersCount *int32 `json:"members_count,omitempty"`

	// 代码组名
	Name *string `json:"name,omitempty"`

	// 仓库计数
	RepositoryCount *int32 `json:"repository_count,omitempty"`

	// 关注计数
	StarCount *int32 `json:"star_count,omitempty"`

	// 是否关注
	Starred *bool `json:"starred,omitempty"`

	// 子组计数
	SubgroupCount *int32 `json:"subgroup_count,omitempty"`

	// 可见性, private public
	Visibility *ShowGroupResponseVisibility `json:"visibility,omitempty"`

	Sum            *GroupSumDto `json:"sum,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o ShowGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowGroupResponse struct{}"
	}

	return strings.Join([]string{"ShowGroupResponse", string(data)}, " ")
}

type ShowGroupResponseVisibility struct {
	value string
}

type ShowGroupResponseVisibilityEnum struct {
	PUBLIC  ShowGroupResponseVisibility
	PRIVATE ShowGroupResponseVisibility
}

func GetShowGroupResponseVisibilityEnum() ShowGroupResponseVisibilityEnum {
	return ShowGroupResponseVisibilityEnum{
		PUBLIC: ShowGroupResponseVisibility{
			value: "public",
		},
		PRIVATE: ShowGroupResponseVisibility{
			value: "private",
		},
	}
}

func (c ShowGroupResponseVisibility) Value() string {
	return c.value
}

func (c ShowGroupResponseVisibility) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowGroupResponseVisibility) UnmarshalJSON(b []byte) error {
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
