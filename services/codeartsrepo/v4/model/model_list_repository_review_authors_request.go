package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListRepositoryReviewAuthorsRequest Request Object
type ListRepositoryReviewAuthorsRequest struct {

	// **参数解释：** 仓库的ID，通过[[查询用户所有仓库](https://support.huaweicloud.com/api-codeartsrepo/ListUserAllRepositories.html)](tag:hws)[[查询用户所有仓库](https://support.huaweicloud.com/intl/en-us/api-codeartsrepo/ListUserAllRepositories.html)](tag:hws_hk)[[查询用户所有仓库](https://support.huaweicloud.com/eu/api-codeartsrepo/ListUserAllRepositories.html)](tag:hws_eu)[查询项目列表](tag:hcs,hcs_sm)接口查询项目列表获取。 **约束限制：** 不涉及。
	RepositoryId int32 `json:"repository_id"`

	// **参数解释：** 意见类型。 **取值范围：** - Commit，提交。 - MergeRequest，合并请求。
	NoteableType ListRepositoryReviewAuthorsRequestNoteableType `json:"noteable_type"`

	// **参数解释：** 解决状态。 **取值范围：** - resolved，已解决。 - unresolved，未解决。   - all，全部。
	ResolvedStatus ListRepositoryReviewAuthorsRequestResolvedStatus `json:"resolved_status"`

	// **参数解释：** 根据检视人名字或用户名筛选意见。
	ReviewersFilter *string `json:"reviewers_filter,omitempty"`

	// **参数解释：** 偏移量，从0开始。
	Offset *int32 `json:"offset,omitempty"`

	// **参数解释：** 返回数量。
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListRepositoryReviewAuthorsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRepositoryReviewAuthorsRequest struct{}"
	}

	return strings.Join([]string{"ListRepositoryReviewAuthorsRequest", string(data)}, " ")
}

type ListRepositoryReviewAuthorsRequestNoteableType struct {
	value string
}

type ListRepositoryReviewAuthorsRequestNoteableTypeEnum struct {
	COMMIT        ListRepositoryReviewAuthorsRequestNoteableType
	MERGE_REQUEST ListRepositoryReviewAuthorsRequestNoteableType
}

func GetListRepositoryReviewAuthorsRequestNoteableTypeEnum() ListRepositoryReviewAuthorsRequestNoteableTypeEnum {
	return ListRepositoryReviewAuthorsRequestNoteableTypeEnum{
		COMMIT: ListRepositoryReviewAuthorsRequestNoteableType{
			value: "Commit",
		},
		MERGE_REQUEST: ListRepositoryReviewAuthorsRequestNoteableType{
			value: "MergeRequest",
		},
	}
}

func (c ListRepositoryReviewAuthorsRequestNoteableType) Value() string {
	return c.value
}

func (c ListRepositoryReviewAuthorsRequestNoteableType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListRepositoryReviewAuthorsRequestNoteableType) UnmarshalJSON(b []byte) error {
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

type ListRepositoryReviewAuthorsRequestResolvedStatus struct {
	value string
}

type ListRepositoryReviewAuthorsRequestResolvedStatusEnum struct {
	RESOLVED   ListRepositoryReviewAuthorsRequestResolvedStatus
	UNRESOLVED ListRepositoryReviewAuthorsRequestResolvedStatus
	ALL        ListRepositoryReviewAuthorsRequestResolvedStatus
}

func GetListRepositoryReviewAuthorsRequestResolvedStatusEnum() ListRepositoryReviewAuthorsRequestResolvedStatusEnum {
	return ListRepositoryReviewAuthorsRequestResolvedStatusEnum{
		RESOLVED: ListRepositoryReviewAuthorsRequestResolvedStatus{
			value: "resolved",
		},
		UNRESOLVED: ListRepositoryReviewAuthorsRequestResolvedStatus{
			value: "unresolved",
		},
		ALL: ListRepositoryReviewAuthorsRequestResolvedStatus{
			value: "all",
		},
	}
}

func (c ListRepositoryReviewAuthorsRequestResolvedStatus) Value() string {
	return c.value
}

func (c ListRepositoryReviewAuthorsRequestResolvedStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListRepositoryReviewAuthorsRequestResolvedStatus) UnmarshalJSON(b []byte) error {
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
