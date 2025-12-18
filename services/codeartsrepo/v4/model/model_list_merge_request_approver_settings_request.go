package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListMergeRequestApproverSettingsRequest Request Object
type ListMergeRequestApproverSettingsRequest struct {

	// **参数解释：** 仓库的ID，通过[[查询用户所有仓库](https://support.huaweicloud.com/api-codeartsrepo/ListUserAllRepositories.html)](tag:hws)[[查询用户所有仓库](https://support.huaweicloud.com/intl/en-us/api-codeartsrepo/ListUserAllRepositories.html)](tag:hws_hk)[[查询用户所有仓库](https://support.huaweicloud.com/eu/api-codeartsrepo/ListUserAllRepositories.html)](tag:hws_eu)[查询项目列表](tag:hcs,hcs_sm)接口查询项目列表获取。 **约束限制：** 不涉及。
	RepositoryId int32 `json:"repository_id"`

	// **参数解释：** 审核配置类型。 **约束限制 ** - branch, 分支配置。 - file, 文件配置 （未启用）。 - project, 项目配置 （未启用）。
	TargetType *ListMergeRequestApproverSettingsRequestTargetType `json:"target_type,omitempty"`

	// **参数解释：** 偏移量，从0开始。
	Offset *int32 `json:"offset,omitempty"`

	// **参数解释：** 返回数量。
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListMergeRequestApproverSettingsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListMergeRequestApproverSettingsRequest struct{}"
	}

	return strings.Join([]string{"ListMergeRequestApproverSettingsRequest", string(data)}, " ")
}

type ListMergeRequestApproverSettingsRequestTargetType struct {
	value string
}

type ListMergeRequestApproverSettingsRequestTargetTypeEnum struct {
	BRANCH  ListMergeRequestApproverSettingsRequestTargetType
	FILE    ListMergeRequestApproverSettingsRequestTargetType
	PROJECT ListMergeRequestApproverSettingsRequestTargetType
}

func GetListMergeRequestApproverSettingsRequestTargetTypeEnum() ListMergeRequestApproverSettingsRequestTargetTypeEnum {
	return ListMergeRequestApproverSettingsRequestTargetTypeEnum{
		BRANCH: ListMergeRequestApproverSettingsRequestTargetType{
			value: "branch",
		},
		FILE: ListMergeRequestApproverSettingsRequestTargetType{
			value: "file",
		},
		PROJECT: ListMergeRequestApproverSettingsRequestTargetType{
			value: "project",
		},
	}
}

func (c ListMergeRequestApproverSettingsRequestTargetType) Value() string {
	return c.value
}

func (c ListMergeRequestApproverSettingsRequestTargetType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListMergeRequestApproverSettingsRequestTargetType) UnmarshalJSON(b []byte) error {
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
