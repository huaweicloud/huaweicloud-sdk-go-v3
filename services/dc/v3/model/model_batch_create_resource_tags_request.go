package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Request Object
type BatchCreateResourceTagsRequest struct {

	// 资源实例ID
	ResourceId string `json:"resource_id"`

	// - 专线服务资源类型，包括dc-directconnect/dc-vgw/dc-vif - dc-directconnect: 专线物理连接 - dc-vgw： 虚拟网关 - dc-vif： 虚拟接口
	ResourceType BatchCreateResourceTagsRequestResourceType `json:"resource_type"`

	Body *BatchOperateResourceTagsRequestBody `json:"body,omitempty"`
}

func (o BatchCreateResourceTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateResourceTagsRequest struct{}"
	}

	return strings.Join([]string{"BatchCreateResourceTagsRequest", string(data)}, " ")
}

type BatchCreateResourceTagsRequestResourceType struct {
	value string
}

type BatchCreateResourceTagsRequestResourceTypeEnum struct {
	DC_DIRECTCONNECT BatchCreateResourceTagsRequestResourceType
	DC_VGW           BatchCreateResourceTagsRequestResourceType
	DC_VIF           BatchCreateResourceTagsRequestResourceType
	DC_LAG           BatchCreateResourceTagsRequestResourceType
}

func GetBatchCreateResourceTagsRequestResourceTypeEnum() BatchCreateResourceTagsRequestResourceTypeEnum {
	return BatchCreateResourceTagsRequestResourceTypeEnum{
		DC_DIRECTCONNECT: BatchCreateResourceTagsRequestResourceType{
			value: "dc-directconnect",
		},
		DC_VGW: BatchCreateResourceTagsRequestResourceType{
			value: "dc-vgw",
		},
		DC_VIF: BatchCreateResourceTagsRequestResourceType{
			value: "dc-vif",
		},
		DC_LAG: BatchCreateResourceTagsRequestResourceType{
			value: "dc-lag",
		},
	}
}

func (c BatchCreateResourceTagsRequestResourceType) Value() string {
	return c.value
}

func (c BatchCreateResourceTagsRequestResourceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BatchCreateResourceTagsRequestResourceType) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}
