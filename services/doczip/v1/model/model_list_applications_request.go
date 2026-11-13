package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListApplicationsRequest Request Object
type ListApplicationsRequest struct {

	// **参数解释**： 页码。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 1。
	Offset *int32 `json:"offset,omitempty"`

	// **参数解释**： 单页数量。 **约束限制**： 不涉及。 **取值范围**： [1-100]。 **默认取值**： 20。
	Limit *int32 `json:"limit,omitempty"`

	// **参数解释**： 应用状态，不传入则返回全部未失效应用。 **约束限制**： 不涉及。 **取值范围**： - 0：待发布。 - 1：已发布，生效状态。 - 2：已发布，失效状态。 **默认取值**： 不涉及。
	Status *ListApplicationsRequestStatus `json:"status,omitempty"`

	// **参数解释**： 应用名称，不传入则返回全部未失效应用。 **约束限制**： 不涉及。 **取值范围**： [1-200]个字符。 **默认取值**： 不涉及。
	ApplicationName *string `json:"application_name,omitempty"`

	// **参数解释**： 租户ID。 **约束限制**： 不涉及。 **取值范围**： [1-50]个字符。 **默认取值**： 不涉及。
	TenantId string `json:"tenant_id"`
}

func (o ListApplicationsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListApplicationsRequest struct{}"
	}

	return strings.Join([]string{"ListApplicationsRequest", string(data)}, " ")
}

type ListApplicationsRequestStatus struct {
	value int32
}

type ListApplicationsRequestStatusEnum struct {
	E_0 ListApplicationsRequestStatus
	E_1 ListApplicationsRequestStatus
	E_2 ListApplicationsRequestStatus
}

func GetListApplicationsRequestStatusEnum() ListApplicationsRequestStatusEnum {
	return ListApplicationsRequestStatusEnum{
		E_0: ListApplicationsRequestStatus{
			value: 0,
		}, E_1: ListApplicationsRequestStatus{
			value: 1,
		}, E_2: ListApplicationsRequestStatus{
			value: 2,
		},
	}
}

func (c ListApplicationsRequestStatus) Value() int32 {
	return c.value
}

func (c ListApplicationsRequestStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListApplicationsRequestStatus) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("int32")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: int32")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(int32); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to int32 error")
	}
}
