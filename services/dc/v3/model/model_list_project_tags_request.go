package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListProjectTagsRequest Request Object
type ListProjectTagsRequest struct {

	// - 专线服务资源类型，包括dc-directconnect/dc-vgw/dc-vif - dc-directconnect: 专线物理连接 - dc-vgw： 虚拟网关 - dc-vif： 虚拟接口
	ResourceType ListProjectTagsRequestResourceType `json:"resource_type"`
}

func (o ListProjectTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListProjectTagsRequest struct{}"
	}

	return strings.Join([]string{"ListProjectTagsRequest", string(data)}, " ")
}

type ListProjectTagsRequestResourceType struct {
	value string
}

type ListProjectTagsRequestResourceTypeEnum struct {
	DC_DIRECTCONNECT ListProjectTagsRequestResourceType
	DC_VGW           ListProjectTagsRequestResourceType
	DC_VIF           ListProjectTagsRequestResourceType
}

func GetListProjectTagsRequestResourceTypeEnum() ListProjectTagsRequestResourceTypeEnum {
	return ListProjectTagsRequestResourceTypeEnum{
		DC_DIRECTCONNECT: ListProjectTagsRequestResourceType{
			value: "dc-directconnect",
		},
		DC_VGW: ListProjectTagsRequestResourceType{
			value: "dc-vgw",
		},
		DC_VIF: ListProjectTagsRequestResourceType{
			value: "dc-vif",
		},
	}
}

func (c ListProjectTagsRequestResourceType) Value() string {
	return c.value
}

func (c ListProjectTagsRequestResourceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListProjectTagsRequestResourceType) UnmarshalJSON(b []byte) error {
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
