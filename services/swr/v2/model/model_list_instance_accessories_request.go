package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListInstanceAccessoriesRequest Request Object
type ListInstanceAccessoriesRequest struct {

	// 企业仓库实例ID
	InstanceId string `json:"instance_id"`

	// 命名空间名称
	NamespaceName string `json:"namespace_name"`

	// 制品名称
	RepositoryName string `json:"repository_name"`

	// 制品摘要
	Reference string `json:"reference"`

	// 起始索引，默认值为0。**注意：offset和limit参数需要配套使用，offset必须为0或者为limit的倍数。**
	Offset *int32 `json:"offset,omitempty"`

	// 返回条数，默认为10，最大值为100。**注意：offset和limit参数需要配套使用，offset必须为0或者为limit的倍数。**
	Limit *int32 `json:"limit,omitempty"`

	// 附件类型，signature.cosign
	Type *ListInstanceAccessoriesRequestType `json:"type,omitempty"`
}

func (o ListInstanceAccessoriesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInstanceAccessoriesRequest struct{}"
	}

	return strings.Join([]string{"ListInstanceAccessoriesRequest", string(data)}, " ")
}

type ListInstanceAccessoriesRequestType struct {
	value string
}

type ListInstanceAccessoriesRequestTypeEnum struct {
	SIGNATURE_COSIGN ListInstanceAccessoriesRequestType
}

func GetListInstanceAccessoriesRequestTypeEnum() ListInstanceAccessoriesRequestTypeEnum {
	return ListInstanceAccessoriesRequestTypeEnum{
		SIGNATURE_COSIGN: ListInstanceAccessoriesRequestType{
			value: "signature.cosign",
		},
	}
}

func (c ListInstanceAccessoriesRequestType) Value() string {
	return c.value
}

func (c ListInstanceAccessoriesRequestType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListInstanceAccessoriesRequestType) UnmarshalJSON(b []byte) error {
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
