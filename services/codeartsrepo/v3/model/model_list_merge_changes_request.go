package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListMergeChangesRequest Request Object
type ListMergeChangesRequest struct {

	// 仓库主键id
	RepositoryId int32 `json:"repository_id"`

	// MR长id
	MergeRequestIid string `json:"merge_request_iid"`

	// 文件路径
	FilePath *string `json:"file_path,omitempty"`

	// 是否无视空格
	IgnoreWhitespaceChange *bool `json:"ignore_whitespace_change,omitempty"`

	// 是否需要编码
	ForceEncode *bool `json:"force_encode,omitempty"`

	// 是否为建议视图
	View *ListMergeChangesRequestView `json:"view,omitempty"`

	// commit的id
	CommitId *string `json:"commit_id,omitempty"`
}

func (o ListMergeChangesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListMergeChangesRequest struct{}"
	}

	return strings.Join([]string{"ListMergeChangesRequest", string(data)}, " ")
}

type ListMergeChangesRequestView struct {
	value string
}

type ListMergeChangesRequestViewEnum struct {
	SIMPLE ListMergeChangesRequestView
}

func GetListMergeChangesRequestViewEnum() ListMergeChangesRequestViewEnum {
	return ListMergeChangesRequestViewEnum{
		SIMPLE: ListMergeChangesRequestView{
			value: "simple",
		},
	}
}

func (c ListMergeChangesRequestView) Value() string {
	return c.value
}

func (c ListMergeChangesRequestView) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListMergeChangesRequestView) UnmarshalJSON(b []byte) error {
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
