package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Request Object
type DeleteTagResourceRequest struct {

	// 资源类型 policy服务策略 ou组织OU account帐号信息 root根
	ResourceType DeleteTagResourceRequestResourceType `json:"resource_type"`

	// 根、组织单元、帐号或策略的唯一标识符（ID）。
	ResourceId string `json:"resource_id"`

	Body *TagResourceReqBody `json:"body,omitempty"`
}

func (o DeleteTagResourceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteTagResourceRequest struct{}"
	}

	return strings.Join([]string{"DeleteTagResourceRequest", string(data)}, " ")
}

type DeleteTagResourceRequestResourceType struct {
	value string
}

type DeleteTagResourceRequestResourceTypeEnum struct {
	POLICY  DeleteTagResourceRequestResourceType
	OU      DeleteTagResourceRequestResourceType
	ACCOUNT DeleteTagResourceRequestResourceType
	ROOT    DeleteTagResourceRequestResourceType
}

func GetDeleteTagResourceRequestResourceTypeEnum() DeleteTagResourceRequestResourceTypeEnum {
	return DeleteTagResourceRequestResourceTypeEnum{
		POLICY: DeleteTagResourceRequestResourceType{
			value: "policy",
		},
		OU: DeleteTagResourceRequestResourceType{
			value: "ou",
		},
		ACCOUNT: DeleteTagResourceRequestResourceType{
			value: "account",
		},
		ROOT: DeleteTagResourceRequestResourceType{
			value: "root",
		},
	}
}

func (c DeleteTagResourceRequestResourceType) Value() string {
	return c.value
}

func (c DeleteTagResourceRequestResourceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DeleteTagResourceRequestResourceType) UnmarshalJSON(b []byte) error {
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
