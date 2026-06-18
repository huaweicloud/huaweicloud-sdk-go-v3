package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListCommitAssociatedRefsRequest Request Object
type ListCommitAssociatedRefsRequest struct {

	// **参数解释：** 仓库的ID，通过[[查询用户所有仓库](https://support.huaweicloud.com/api-codeartsrepo/ListUserAllRepositories.html)](tag:hws)[[查询用户所有仓库](https://support.huaweicloud.com/intl/en-us/api-codeartsrepo/ListUserAllRepositories.html)](tag:hws_hk)[[查询用户所有仓库](https://support.huaweicloud.com/intl/zh-cn/api-codeartsrepo/ListUserAllRepositories.html)](tag:hws_hk_ch)[[查询用户所有仓库](https://support.huaweicloud.com/eu/api-codeartsrepo/ListUserAllRepositories.html)](tag:hws_eu)[查询项目列表](tag:hcs,hcs_sm)接口查询项目列表获取。 **约束限制：** 不涉及。 **默认取值：** 不涉及。
	RepositoryId int32 `json:"repository_id"`

	// **参数解释：** 提交ID。
	Sha string `json:"sha"`

	// **参数解释：** 偏移量，从0开始。
	Offset *int32 `json:"offset,omitempty"`

	// **参数解释：** 返回数量。
	Limit *int32 `json:"limit,omitempty"`

	// **参数解释：** 分支或者tag。 **取值范围：** 只能为branch或者tag。 **约束限制：** 只能为branch或者tag。
	Type ListCommitAssociatedRefsRequestType `json:"type"`
}

func (o ListCommitAssociatedRefsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCommitAssociatedRefsRequest struct{}"
	}

	return strings.Join([]string{"ListCommitAssociatedRefsRequest", string(data)}, " ")
}

type ListCommitAssociatedRefsRequestType struct {
	value string
}

type ListCommitAssociatedRefsRequestTypeEnum struct {
	BRANCH ListCommitAssociatedRefsRequestType
	TAG    ListCommitAssociatedRefsRequestType
}

func GetListCommitAssociatedRefsRequestTypeEnum() ListCommitAssociatedRefsRequestTypeEnum {
	return ListCommitAssociatedRefsRequestTypeEnum{
		BRANCH: ListCommitAssociatedRefsRequestType{
			value: "branch",
		},
		TAG: ListCommitAssociatedRefsRequestType{
			value: "tag",
		},
	}
}

func (c ListCommitAssociatedRefsRequestType) Value() string {
	return c.value
}

func (c ListCommitAssociatedRefsRequestType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListCommitAssociatedRefsRequestType) UnmarshalJSON(b []byte) error {
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
