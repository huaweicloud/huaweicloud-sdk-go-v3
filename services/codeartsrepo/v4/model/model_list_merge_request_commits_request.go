package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListMergeRequestCommitsRequest Request Object
type ListMergeRequestCommitsRequest struct {

	// **参数解释：** 仓库的ID，通过[[查询用户所有仓库](https://support.huaweicloud.com/api-codeartsrepo/ListUserAllRepositories.html)](tag:hws)[[查询用户所有仓库](https://support.huaweicloud.com/intl/en-us/api-codeartsrepo/ListUserAllRepositories.html)](tag:hws_hk)[[查询用户所有仓库](https://support.huaweicloud.com/eu/api-codeartsrepo/ListUserAllRepositories.html)](tag:hws_eu)[查询项目列表](tag:hcs,hcs_sm)接口查询项目列表获取。 **约束限制：** 不涉及。
	RepositoryId int32 `json:"repository_id"`

	// **参数解释：**  合并请求 iid。
	MergeRequestIid int32 `json:"merge_request_iid"`

	// **参数解释：** 是否以简单模式展示commit，传值为simple时以简单模式展示，否则正常展示。 **约束限制 ** - simple, 以简单模式展示commit。
	View *ListMergeRequestCommitsRequestView `json:"view,omitempty"`

	// **参数解释：** 偏移量，从0开始。
	Offset *int32 `json:"offset,omitempty"`

	// **参数解释：** 返回数量。
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListMergeRequestCommitsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListMergeRequestCommitsRequest struct{}"
	}

	return strings.Join([]string{"ListMergeRequestCommitsRequest", string(data)}, " ")
}

type ListMergeRequestCommitsRequestView struct {
	value string
}

type ListMergeRequestCommitsRequestViewEnum struct {
	SIMPLE ListMergeRequestCommitsRequestView
}

func GetListMergeRequestCommitsRequestViewEnum() ListMergeRequestCommitsRequestViewEnum {
	return ListMergeRequestCommitsRequestViewEnum{
		SIMPLE: ListMergeRequestCommitsRequestView{
			value: "simple",
		},
	}
}

func (c ListMergeRequestCommitsRequestView) Value() string {
	return c.value
}

func (c ListMergeRequestCommitsRequestView) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListMergeRequestCommitsRequestView) UnmarshalJSON(b []byte) error {
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
