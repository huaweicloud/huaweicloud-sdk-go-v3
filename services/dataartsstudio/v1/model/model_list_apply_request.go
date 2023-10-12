package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListApplyRequest Request Object
type ListApplyRequest struct {

	// 工作空间id
	Workspace string `json:"workspace"`

	// dlm版本类型
	DlmType ListApplyRequestDlmType `json:"Dlm-Type"`

	// 资源类型
	ContentType string `json:"Content-Type"`

	// 查询起始坐标, 即跳过前X条数据。仅支持0或limit的整数倍，不满足则向下取整。
	Offset *int32 `json:"offset,omitempty"`

	// 查询条数, 即查询Y条数据
	Limit *int32 `json:"limit,omitempty"`

	// api名称
	ApiName *string `json:"api_name,omitempty"`

	// 查询类型, 0:收到的申请(待审核), 1:收到的申请(已审核), 2:发出的申请(开发), 3:发出的申请(调用)
	QueryType *int32 `json:"query_type,omitempty"`
}

func (o ListApplyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListApplyRequest struct{}"
	}

	return strings.Join([]string{"ListApplyRequest", string(data)}, " ")
}

type ListApplyRequestDlmType struct {
	value string
}

type ListApplyRequestDlmTypeEnum struct {
	SHARED    ListApplyRequestDlmType
	EXCLUSIVE ListApplyRequestDlmType
}

func GetListApplyRequestDlmTypeEnum() ListApplyRequestDlmTypeEnum {
	return ListApplyRequestDlmTypeEnum{
		SHARED: ListApplyRequestDlmType{
			value: "SHARED",
		},
		EXCLUSIVE: ListApplyRequestDlmType{
			value: "EXCLUSIVE",
		},
	}
}

func (c ListApplyRequestDlmType) Value() string {
	return c.value
}

func (c ListApplyRequestDlmType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListApplyRequestDlmType) UnmarshalJSON(b []byte) error {
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
