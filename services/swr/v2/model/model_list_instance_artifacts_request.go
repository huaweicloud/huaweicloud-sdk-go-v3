package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListInstanceArtifactsRequest Request Object
type ListInstanceArtifactsRequest struct {

	// 企业仓库实例ID
	InstanceId string `json:"instance_id"`

	// 命名空间名称
	NamespaceName string `json:"namespace_name"`

	// 仓库名称
	RepositoryName string `json:"repository_name"`

	// 起始索引，默认值为0。**注意：offset和limit参数需要配套使用，offset必须为0或者为limit的倍数。**
	Offset *int32 `json:"offset,omitempty"`

	// 返回条数，默认为10，最大值为100。**注意：offset和limit参数需要配套使用，offset必须为0或者为limit的倍数。**
	Limit *int32 `json:"limit,omitempty"`

	// 制品类型
	Type *ListInstanceArtifactsRequestType `json:"type,omitempty"`

	// 制品包含版本，模糊匹配条件
	Tags *string `json:"tags,omitempty"`
}

func (o ListInstanceArtifactsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInstanceArtifactsRequest struct{}"
	}

	return strings.Join([]string{"ListInstanceArtifactsRequest", string(data)}, " ")
}

type ListInstanceArtifactsRequestType struct {
	value string
}

type ListInstanceArtifactsRequestTypeEnum struct {
	IMAGE ListInstanceArtifactsRequestType
	CHART ListInstanceArtifactsRequestType
}

func GetListInstanceArtifactsRequestTypeEnum() ListInstanceArtifactsRequestTypeEnum {
	return ListInstanceArtifactsRequestTypeEnum{
		IMAGE: ListInstanceArtifactsRequestType{
			value: "IMAGE",
		},
		CHART: ListInstanceArtifactsRequestType{
			value: "CHART",
		},
	}
}

func (c ListInstanceArtifactsRequestType) Value() string {
	return c.value
}

func (c ListInstanceArtifactsRequestType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListInstanceArtifactsRequestType) UnmarshalJSON(b []byte) error {
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
