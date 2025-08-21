package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListReportProfilesRequest Request Object
type ListReportProfilesRequest struct {

	// **参数解释**： 报告类型 **约束限制**： 不涉及 **取值范围**： daily 日报 weekly 周报 custom 自定义报告 **默认取值**： 不涉及
	Category *ListReportProfilesRequestCategory `json:"category,omitempty"`

	// **参数解释**： 防火墙ID，用户创建防火墙实例后产生的唯一ID，配置后可区分不同防火墙，可通过[防火墙ID获取方式](cfw_02_0028.xml)获取 **约束限制**： 不涉及 **取值范围**： 32位UUID **默认取值**： 不涉及
	FwInstanceId string `json:"fw_instance_id"`

	// **参数解释**： 查询返回记录的数量限制 **约束限制**： 不涉及 **取值范围**： 1-1024 **默认取值**： 不涉及
	Limit int32 `json:"limit"`

	// **参数解释**： 偏移量，表示查询该偏移量后面的记录 **约束限制**： 不涉及 **取值范围**： 0 - 1024 **默认取值**： 不涉及
	Offset int32 `json:"offset"`
}

func (o ListReportProfilesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListReportProfilesRequest struct{}"
	}

	return strings.Join([]string{"ListReportProfilesRequest", string(data)}, " ")
}

type ListReportProfilesRequestCategory struct {
	value string
}

type ListReportProfilesRequestCategoryEnum struct {
	DAILY  ListReportProfilesRequestCategory
	WEEKLY ListReportProfilesRequestCategory
	CUSTOM ListReportProfilesRequestCategory
}

func GetListReportProfilesRequestCategoryEnum() ListReportProfilesRequestCategoryEnum {
	return ListReportProfilesRequestCategoryEnum{
		DAILY: ListReportProfilesRequestCategory{
			value: "daily",
		},
		WEEKLY: ListReportProfilesRequestCategory{
			value: "weekly",
		},
		CUSTOM: ListReportProfilesRequestCategory{
			value: "custom",
		},
	}
}

func (c ListReportProfilesRequestCategory) Value() string {
	return c.value
}

func (c ListReportProfilesRequestCategory) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListReportProfilesRequestCategory) UnmarshalJSON(b []byte) error {
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
