package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListMergeChangesTreesRequest Request Object
type ListMergeChangesTreesRequest struct {

	// 仓库主键id
	RepositoryId int32 `json:"repository_id"`

	// MR长id
	MergeRequestIid string `json:"merge_request_iid"`

	// 是否为简易视图
	View *ListMergeChangesTreesRequestView `json:"view,omitempty"`

	// commit的id
	CommitId *string `json:"commit_id,omitempty"`
}

func (o ListMergeChangesTreesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListMergeChangesTreesRequest struct{}"
	}

	return strings.Join([]string{"ListMergeChangesTreesRequest", string(data)}, " ")
}

type ListMergeChangesTreesRequestView struct {
	value string
}

type ListMergeChangesTreesRequestViewEnum struct {
	SIMPLE ListMergeChangesTreesRequestView
}

func GetListMergeChangesTreesRequestViewEnum() ListMergeChangesTreesRequestViewEnum {
	return ListMergeChangesTreesRequestViewEnum{
		SIMPLE: ListMergeChangesTreesRequestView{
			value: "simple",
		},
	}
}

func (c ListMergeChangesTreesRequestView) Value() string {
	return c.value
}

func (c ListMergeChangesTreesRequestView) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListMergeChangesTreesRequestView) UnmarshalJSON(b []byte) error {
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
