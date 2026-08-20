package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListIpdLabelsRequest Request Object
type ListIpdLabelsRequest struct {

	// 项目32位ID，项目唯一标识。通过查询IPD项目列表获取，响应消息体中的id字段的值就是项目ID。
	ProjectId string `json:"project_id"`

	// 标签归属的工作项分类，不传该参数时默认查询所有类型下的标签。不推荐使用此参数，建议使用category_types参数。
	LabelType *ListIpdLabelsRequestLabelType `json:"label_type,omitempty"`

	// 标签名称
	Title *string `json:"title,omitempty"`

	// 工作项类型编码。
	CategoryTypes *string `json:"category_types,omitempty"`
}

func (o ListIpdLabelsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListIpdLabelsRequest struct{}"
	}

	return strings.Join([]string{"ListIpdLabelsRequest", string(data)}, " ")
}

type ListIpdLabelsRequestLabelType struct {
	value string
}

type ListIpdLabelsRequestLabelTypeEnum struct {
	FEATURE         ListIpdLabelsRequestLabelType
	RAW_REQUIREMENT ListIpdLabelsRequestLabelType
	REQUIREMENT     ListIpdLabelsRequestLabelType
	TASK            ListIpdLabelsRequestLabelType
	BUG             ListIpdLabelsRequestLabelType
}

func GetListIpdLabelsRequestLabelTypeEnum() ListIpdLabelsRequestLabelTypeEnum {
	return ListIpdLabelsRequestLabelTypeEnum{
		FEATURE: ListIpdLabelsRequestLabelType{
			value: "feature",
		},
		RAW_REQUIREMENT: ListIpdLabelsRequestLabelType{
			value: "raw requirement",
		},
		REQUIREMENT: ListIpdLabelsRequestLabelType{
			value: "requirement",
		},
		TASK: ListIpdLabelsRequestLabelType{
			value: "task",
		},
		BUG: ListIpdLabelsRequestLabelType{
			value: "bug",
		},
	}
}

func (c ListIpdLabelsRequestLabelType) Value() string {
	return c.value
}

func (c ListIpdLabelsRequestLabelType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListIpdLabelsRequestLabelType) UnmarshalJSON(b []byte) error {
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
