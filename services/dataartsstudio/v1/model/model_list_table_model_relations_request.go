package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListTableModelRelationsRequest Request Object
type ListTableModelRelationsRequest struct {

	// 工作空间ID，获取方法请参见[实例ID和工作空间ID](dataartsstudio_02_0350.xml)。
	Workspace string `json:"workspace"`

	// 项目ID，获取方法请参见[项目ID和账号ID](projectid_accountid.xml)。  多project场景采用AK/SK认证的接口请求，则该字段必选。
	XProjectId *string `json:"X-Project-Id,omitempty"`

	// 默认值：application/json;charset=UTF-8 可选，有Body体的情况下必选，没有Body体则无需填写和校验。
	ContentType *string `json:"Content-Type,omitempty"`

	// 所属关系建模的模型ID，填写String类型替代Long类型。
	ModelId string `json:"model_id"`

	// 表模型ids，填写String类型替代Long类型。
	TableIds *string `json:"table_ids,omitempty"`

	// 表类型。 枚举值：   - TABLE_MODEL: 关系建模：逻辑实体/物理表   - FACT_LOGIC_TABLE: 事实表
	BizType *ListTableModelRelationsRequestBizType `json:"biz_type,omitempty"`

	// 每页查询条数，即查询Y条数据。默认值50，取值范围[1,100]。
	Limit *int32 `json:"limit,omitempty"`

	// 查询起始坐标，即跳过X条数据，仅支持0或limit的整数倍，不满足则向下取整，默认值0。
	Offset *int32 `json:"offset,omitempty"`
}

func (o ListTableModelRelationsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTableModelRelationsRequest struct{}"
	}

	return strings.Join([]string{"ListTableModelRelationsRequest", string(data)}, " ")
}

type ListTableModelRelationsRequestBizType struct {
	value string
}

type ListTableModelRelationsRequestBizTypeEnum struct {
	TABLE_MODEL      ListTableModelRelationsRequestBizType
	FACT_LOGIC_TABLE ListTableModelRelationsRequestBizType
}

func GetListTableModelRelationsRequestBizTypeEnum() ListTableModelRelationsRequestBizTypeEnum {
	return ListTableModelRelationsRequestBizTypeEnum{
		TABLE_MODEL: ListTableModelRelationsRequestBizType{
			value: "TABLE_MODEL",
		},
		FACT_LOGIC_TABLE: ListTableModelRelationsRequestBizType{
			value: "FACT_LOGIC_TABLE",
		},
	}
}

func (c ListTableModelRelationsRequestBizType) Value() string {
	return c.value
}

func (c ListTableModelRelationsRequestBizType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListTableModelRelationsRequestBizType) UnmarshalJSON(b []byte) error {
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
