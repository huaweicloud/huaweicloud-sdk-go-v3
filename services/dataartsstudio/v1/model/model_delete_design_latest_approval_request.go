package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// DeleteDesignLatestApprovalRequest Request Object
type DeleteDesignLatestApprovalRequest struct {

	// 工作空间ID，获取方法请参见[实例ID和工作空间ID](dataartsstudio_02_0350.xml)。
	Workspace string `json:"workspace"`

	// 项目ID，获取方法请参见[项目ID和账号ID](projectid_accountid.xml)。  多project场景采用AK/SK认证的接口请求，则该字段必选。
	XProjectId *string `json:"X-Project-Id,omitempty"`

	// 默认值：application/json;charset=UTF-8 可选，有Body体的情况下必选，没有Body体则无需填写和校验。
	ContentType *string `json:"Content-Type,omitempty"`

	// 待删除下展的实体ID，填写String类型替代Long类型。
	BizId string `json:"biz_id"`

	// 待删除下展的实体类型。 枚举值：   - ATOMIC_INDEX: 原子指标   - DERIVATIVE_INDEX: 衍生指标   - DIMENSION: 维度   - FACT_LOGIC_TABLE: 事实表   - TABLE_MODEL: 关系建模：逻辑实体/物理表   - STANDARD_ELEMENT: 数据标准   - AGGREGATION_LOGIC_TABLE: 汇总表   - CODE_TABLE: 码表   - BIZ_METRIC: 业务指标   - COMPOUND_METRIC: 复合指标
	BizType DeleteDesignLatestApprovalRequestBizType `json:"biz_type"`
}

func (o DeleteDesignLatestApprovalRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteDesignLatestApprovalRequest struct{}"
	}

	return strings.Join([]string{"DeleteDesignLatestApprovalRequest", string(data)}, " ")
}

type DeleteDesignLatestApprovalRequestBizType struct {
	value string
}

type DeleteDesignLatestApprovalRequestBizTypeEnum struct {
	ATOMIC_INDEX            DeleteDesignLatestApprovalRequestBizType
	DERIVATIVE_INDEX        DeleteDesignLatestApprovalRequestBizType
	DIMENSION               DeleteDesignLatestApprovalRequestBizType
	FACT_LOGIC_TABLE        DeleteDesignLatestApprovalRequestBizType
	TABLE_MODEL             DeleteDesignLatestApprovalRequestBizType
	STANDARD_ELEMENT        DeleteDesignLatestApprovalRequestBizType
	AGGREGATION_LOGIC_TABLE DeleteDesignLatestApprovalRequestBizType
	CODE_TABLE              DeleteDesignLatestApprovalRequestBizType
	BIZ_METRIC              DeleteDesignLatestApprovalRequestBizType
	COMPOUND_METRIC         DeleteDesignLatestApprovalRequestBizType
}

func GetDeleteDesignLatestApprovalRequestBizTypeEnum() DeleteDesignLatestApprovalRequestBizTypeEnum {
	return DeleteDesignLatestApprovalRequestBizTypeEnum{
		ATOMIC_INDEX: DeleteDesignLatestApprovalRequestBizType{
			value: "ATOMIC_INDEX",
		},
		DERIVATIVE_INDEX: DeleteDesignLatestApprovalRequestBizType{
			value: "DERIVATIVE_INDEX",
		},
		DIMENSION: DeleteDesignLatestApprovalRequestBizType{
			value: "DIMENSION",
		},
		FACT_LOGIC_TABLE: DeleteDesignLatestApprovalRequestBizType{
			value: "FACT_LOGIC_TABLE",
		},
		TABLE_MODEL: DeleteDesignLatestApprovalRequestBizType{
			value: "TABLE_MODEL",
		},
		STANDARD_ELEMENT: DeleteDesignLatestApprovalRequestBizType{
			value: "STANDARD_ELEMENT",
		},
		AGGREGATION_LOGIC_TABLE: DeleteDesignLatestApprovalRequestBizType{
			value: "AGGREGATION_LOGIC_TABLE",
		},
		CODE_TABLE: DeleteDesignLatestApprovalRequestBizType{
			value: "CODE_TABLE",
		},
		BIZ_METRIC: DeleteDesignLatestApprovalRequestBizType{
			value: "BIZ_METRIC",
		},
		COMPOUND_METRIC: DeleteDesignLatestApprovalRequestBizType{
			value: "COMPOUND_METRIC",
		},
	}
}

func (c DeleteDesignLatestApprovalRequestBizType) Value() string {
	return c.value
}

func (c DeleteDesignLatestApprovalRequestBizType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DeleteDesignLatestApprovalRequestBizType) UnmarshalJSON(b []byte) error {
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
