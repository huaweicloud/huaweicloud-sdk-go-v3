package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListProjectTagsRequest Request Object
type ListProjectTagsRequest struct {

	// 资源类型。  - cph-server，云手机服务器
	ResourceType ListProjectTagsRequestResourceType `json:"resource_type"`

	// 每页返回的资源个数。取值范围：1~100（默认值为100），一般设置为10、20、50。
	Limit *int32 `json:"limit,omitempty"`

	// 索引位置，从第一条数据偏移offset条数据后开始查询，默认为0（偏移0条数据，表示从第一条数据开始查询）,必须为数字，不能为负数
	Offset *int32 `json:"offset,omitempty"`
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
	CPH_SERVER ListProjectTagsRequestResourceType
}

func GetListProjectTagsRequestResourceTypeEnum() ListProjectTagsRequestResourceTypeEnum {
	return ListProjectTagsRequestResourceTypeEnum{
		CPH_SERVER: ListProjectTagsRequestResourceType{
			value: "cph-server",
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
