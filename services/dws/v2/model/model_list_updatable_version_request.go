package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListUpdatableVersionRequest Request Object
type ListUpdatableVersionRequest struct {

	// 集群ID
	ClusterId string `json:"cluster_id"`

	// 偏移量
	Offset *int32 `json:"offset,omitempty"`

	// 条目数
	Limit *int32 `json:"limit,omitempty"`

	// 升级类型 cluster:集群升级 hotpatch:热补丁升级
	Type *ListUpdatableVersionRequestType `json:"type,omitempty"`
}

func (o ListUpdatableVersionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListUpdatableVersionRequest struct{}"
	}

	return strings.Join([]string{"ListUpdatableVersionRequest", string(data)}, " ")
}

type ListUpdatableVersionRequestType struct {
	value string
}

type ListUpdatableVersionRequestTypeEnum struct {
	CLUSTER  ListUpdatableVersionRequestType
	HOTPATCH ListUpdatableVersionRequestType
}

func GetListUpdatableVersionRequestTypeEnum() ListUpdatableVersionRequestTypeEnum {
	return ListUpdatableVersionRequestTypeEnum{
		CLUSTER: ListUpdatableVersionRequestType{
			value: "cluster",
		},
		HOTPATCH: ListUpdatableVersionRequestType{
			value: "hotpatch",
		},
	}
}

func (c ListUpdatableVersionRequestType) Value() string {
	return c.value
}

func (c ListUpdatableVersionRequestType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListUpdatableVersionRequestType) UnmarshalJSON(b []byte) error {
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
