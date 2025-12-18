package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowRepositoryMergeRequestsStatisticRequest Request Object
type ShowRepositoryMergeRequestsStatisticRequest struct {

	// **参数解释：** 仓库的ID，通过[[查询用户所有仓库](https://support.huaweicloud.com/api-codeartsrepo/ListUserAllRepositories.html)](tag:hws)[[查询用户所有仓库](https://support.huaweicloud.com/intl/en-us/api-codeartsrepo/ListUserAllRepositories.html)](tag:hws_hk)[[查询用户所有仓库](https://support.huaweicloud.com/eu/api-codeartsrepo/ListUserAllRepositories.html)](tag:hws_eu)[查询项目列表](tag:hcs,hcs_sm)接口查询项目列表获取。 **约束限制：** 不涉及。
	RepositoryId int32 `json:"repository_id"`

	// **参数解释：** 合并请求iid。
	Iids string `json:"iids"`

	// **参数解释：** 统计字段。 **约束限制 ** - commits_count，统计提交数 - changed_files_count，文件变更数 - notes_count， 检视意见数 - changed_lines_count，代码变更行数
	Fields *ShowRepositoryMergeRequestsStatisticRequestFields `json:"fields,omitempty"`
}

func (o ShowRepositoryMergeRequestsStatisticRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowRepositoryMergeRequestsStatisticRequest struct{}"
	}

	return strings.Join([]string{"ShowRepositoryMergeRequestsStatisticRequest", string(data)}, " ")
}

type ShowRepositoryMergeRequestsStatisticRequestFields struct {
	value string
}

type ShowRepositoryMergeRequestsStatisticRequestFieldsEnum struct {
	COMMITS_COUNT       ShowRepositoryMergeRequestsStatisticRequestFields
	CHANGED_FILES_COUNT ShowRepositoryMergeRequestsStatisticRequestFields
	NOTES_COUNT         ShowRepositoryMergeRequestsStatisticRequestFields
	CHANGED_LINES_COUNT ShowRepositoryMergeRequestsStatisticRequestFields
}

func GetShowRepositoryMergeRequestsStatisticRequestFieldsEnum() ShowRepositoryMergeRequestsStatisticRequestFieldsEnum {
	return ShowRepositoryMergeRequestsStatisticRequestFieldsEnum{
		COMMITS_COUNT: ShowRepositoryMergeRequestsStatisticRequestFields{
			value: "commits_count",
		},
		CHANGED_FILES_COUNT: ShowRepositoryMergeRequestsStatisticRequestFields{
			value: "changed_files_count",
		},
		NOTES_COUNT: ShowRepositoryMergeRequestsStatisticRequestFields{
			value: "notes_count",
		},
		CHANGED_LINES_COUNT: ShowRepositoryMergeRequestsStatisticRequestFields{
			value: "changed_lines_count",
		},
	}
}

func (c ShowRepositoryMergeRequestsStatisticRequestFields) Value() string {
	return c.value
}

func (c ShowRepositoryMergeRequestsStatisticRequestFields) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowRepositoryMergeRequestsStatisticRequestFields) UnmarshalJSON(b []byte) error {
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
