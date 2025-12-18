package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowRefCompareRequest Request Object
type ShowRefCompareRequest struct {

	// **参数解释：** 仓库的ID，通过[[查询用户所有仓库](https://support.huaweicloud.com/api-codeartsrepo/ListUserAllRepositories.html)](tag:hws)[[查询用户所有仓库](https://support.huaweicloud.com/intl/en-us/api-codeartsrepo/ListUserAllRepositories.html)](tag:hws_hk)[[查询用户所有仓库](https://support.huaweicloud.com/eu/api-codeartsrepo/ListUserAllRepositories.html)](tag:hws_eu)[查询项目列表](tag:hcs,hcs_sm)接口查询项目列表获取。 **约束限制：** 不涉及。
	RepositoryId int32 `json:"repository_id"`

	// **参数解释：** 要开始比较的分支名称、标签名称或者commitid。 **取值范围：** 字符串长度不少于1，不超过100000。
	From string `json:"from"`

	// **参数解释：** 要停止比较的分支名称、标签名称或者commitid。 **取值范围：** 字符串长度不少于1，不超过100000。
	To string `json:"to"`

	// **参数解释：** 对比类型。 **取值范围：** - branch，分支。 - commit，提交。 - tag，标签。
	CompareType *ShowRefCompareRequestCompareType `json:"compare_type,omitempty"`

	// **参数解释：** 合并请求的目标仓库，默认为仓库ID。
	TargetId *int32 `json:"target_id,omitempty"`

	// **参数解释：** 比较方法。 **取值范围：** - true，用于`from`和`to`之间的直接比较(`from`..`to`)。 - false，使用merge base进行比较(`from`...`to`)。
	Straight *bool `json:"straight,omitempty"`

	// **参数解释：** 是否忽略空白数量更改。 **取值范围：** - true，忽略空白数量的更改。 - false，不会忽略空白数量的更改。
	IgnoreWhitespaceChange *bool `json:"ignore_whitespace_change,omitempty"`

	// **参数解释：** 是否忽略空白数量更改。 **取值范围：** - count，数量。 - commits，提交信息。 - diffs，差异信息。 - files，文件信息。 - commits,diffs，提交信息和差异信息。
	View *ShowRefCompareRequestView `json:"view,omitempty"`

	// **参数解释：** 是否仅返回带有提交计数和diffs计数的结果。 **取值范围：** - true，仅返回带有提交计数和diffs计数的结果。 - false，按照compare_view参数返回结果信息。
	OnlyCount *bool `json:"only_count,omitempty"`

	// **参数解释：** 根据给定的文件路径返回不同的结果。如果文件已重命名，则file_path应包括old_path和new_path，如“file_path=old_path,new_path”。 **取值范围：** 字符串长度不少于1，不超过100000。
	FilePath *string `json:"file_path,omitempty"`

	// **参数解释：** 根据给定的参数返回不同的结果。 **取值范围：** - change_lines，变更行数。
	AdditionalFields *ShowRefCompareRequestAdditionalFields `json:"additional_fields,omitempty"`

	// **参数解释：** 偏移量，从0开始。
	Offset *int32 `json:"offset,omitempty"`

	// **参数解释：** 返回数量。
	Limit *int32 `json:"limit,omitempty"`
}

func (o ShowRefCompareRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowRefCompareRequest struct{}"
	}

	return strings.Join([]string{"ShowRefCompareRequest", string(data)}, " ")
}

type ShowRefCompareRequestCompareType struct {
	value string
}

type ShowRefCompareRequestCompareTypeEnum struct {
	BRANCH ShowRefCompareRequestCompareType
	COMMIT ShowRefCompareRequestCompareType
	TAG    ShowRefCompareRequestCompareType
}

func GetShowRefCompareRequestCompareTypeEnum() ShowRefCompareRequestCompareTypeEnum {
	return ShowRefCompareRequestCompareTypeEnum{
		BRANCH: ShowRefCompareRequestCompareType{
			value: "branch",
		},
		COMMIT: ShowRefCompareRequestCompareType{
			value: "commit",
		},
		TAG: ShowRefCompareRequestCompareType{
			value: "tag",
		},
	}
}

func (c ShowRefCompareRequestCompareType) Value() string {
	return c.value
}

func (c ShowRefCompareRequestCompareType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowRefCompareRequestCompareType) UnmarshalJSON(b []byte) error {
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

type ShowRefCompareRequestView struct {
	value string
}

type ShowRefCompareRequestViewEnum struct {
	COUNT        ShowRefCompareRequestView
	COMMITS      ShowRefCompareRequestView
	DIFFS        ShowRefCompareRequestView
	FILES        ShowRefCompareRequestView
	COMMITSDIFFS ShowRefCompareRequestView
}

func GetShowRefCompareRequestViewEnum() ShowRefCompareRequestViewEnum {
	return ShowRefCompareRequestViewEnum{
		COUNT: ShowRefCompareRequestView{
			value: "count",
		},
		COMMITS: ShowRefCompareRequestView{
			value: "commits",
		},
		DIFFS: ShowRefCompareRequestView{
			value: "diffs",
		},
		FILES: ShowRefCompareRequestView{
			value: "files",
		},
		COMMITSDIFFS: ShowRefCompareRequestView{
			value: "commits,diffs",
		},
	}
}

func (c ShowRefCompareRequestView) Value() string {
	return c.value
}

func (c ShowRefCompareRequestView) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowRefCompareRequestView) UnmarshalJSON(b []byte) error {
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

type ShowRefCompareRequestAdditionalFields struct {
	value string
}

type ShowRefCompareRequestAdditionalFieldsEnum struct {
	CHANGE_LINES ShowRefCompareRequestAdditionalFields
}

func GetShowRefCompareRequestAdditionalFieldsEnum() ShowRefCompareRequestAdditionalFieldsEnum {
	return ShowRefCompareRequestAdditionalFieldsEnum{
		CHANGE_LINES: ShowRefCompareRequestAdditionalFields{
			value: "change_lines",
		},
	}
}

func (c ShowRefCompareRequestAdditionalFields) Value() string {
	return c.value
}

func (c ShowRefCompareRequestAdditionalFields) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowRefCompareRequestAdditionalFields) UnmarshalJSON(b []byte) error {
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
