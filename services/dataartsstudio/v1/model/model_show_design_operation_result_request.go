package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowDesignOperationResultRequest Request Object
type ShowDesignOperationResultRequest struct {

	// 工作空间ID，获取方法请参见[实例ID和工作空间ID](dataartsstudio_02_0350.xml)。
	Workspace string `json:"workspace"`

	// 项目ID，获取方法请参见[项目ID和账号ID](projectid_accountid.xml)。  多project场景采用AK/SK认证的接口请求，则该字段必选。
	XProjectId *string `json:"X-Project-Id,omitempty"`

	// 默认值：application/json;charset=UTF-8 可选，有Body体的情况下必选，没有Body体则无需填写和校验。
	ContentType *string `json:"Content-Type,omitempty"`

	// 批量操作类型。 枚举值：   - PUBLISH: 发布   - ADD_TAGS: 打标签   - APPROVAL: 审批   - ER_REVERSE_DB: 关系建模逆向数据库   - CODETABLE_REVERSE_DB: 码表逆向数据库   - DIMENSION_REVERSE_DB: 维度逆向数据库   - FACT_LOGIC_TABLE_REVERSE_DB: 事实表逆向数据库   - SYNC_TABLES: 同步元模型   - IMPORT_STANDARD: 导入数据标准   - IMPORT_CODETABLE: 导入码表   - IMPORT_ER_TABLE: 导入关系建模（逻辑实体/物理表）   - IMPORT_BUSINESS: 导入业务分层（主题）   - TRANSFORM_LOGIC_MODEL: 逻辑模型转物理模型   - PUBLISH_CODETABLE: 发布码表   - PUBLISH_STANDARD: 发布数据标准   - TABLE_MODEL_RELOCATE: 关系建模业务表批量修改主题   - DIMENSION_RELOCATE: 维度批量修改主题   - FACT_LOGIC_TABLE_RELOCATE: 事实表批量修改主题   - AGGREGATION_LOGIC_TABLE_RELOCATE: 汇总表批量修改主题   - ATOMIC_INDEX_RELOCATE: 原子指标批量修改主题   - DERIVATIVE_INDEX_RELOCATE: 衍生指标批量修改主题   - COMPOUND_METRIC_RELOCATE: 复合指标批量修改主题   - BIZ_METRIC_RELOCATE: 业务指标批量修改流程   - CODE_TABLE_RELOCATE: 码表批量修改目录   - STANDARD_ELEMENT_RELOCATE: 数据标准批量修改目录   - INFO_ARCH_RELOCATE: 信息架构批量修改主题
	OperationType *ShowDesignOperationResultRequestOperationType `json:"operation_type,omitempty"`

	// 每页查询条数，即查询Y条数据。默认值50，取值范围[1,100]。
	Limit *int32 `json:"limit,omitempty"`

	// 查询起始坐标，即跳过X条数据，仅支持0或limit的整数倍，不满足则向下取整，默认值0。
	Offset *int32 `json:"offset,omitempty"`

	// 批量操作id，在逻辑模型转物理表时，填写的是逻辑模型的model_id，在逆向数据库时，填写的是目标模型的model_id。model_id可从接口[获取模型](ListWorkspaces.xml)中获取。
	OperationId *string `json:"operation_id,omitempty"`
}

func (o ShowDesignOperationResultRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDesignOperationResultRequest struct{}"
	}

	return strings.Join([]string{"ShowDesignOperationResultRequest", string(data)}, " ")
}

type ShowDesignOperationResultRequestOperationType struct {
	value string
}

type ShowDesignOperationResultRequestOperationTypeEnum struct {
	PUBLISH                          ShowDesignOperationResultRequestOperationType
	ADD_TAGS                         ShowDesignOperationResultRequestOperationType
	APPROVAL                         ShowDesignOperationResultRequestOperationType
	ER_REVERSE_DB                    ShowDesignOperationResultRequestOperationType
	CODETABLE_REVERSE_DB             ShowDesignOperationResultRequestOperationType
	DIMENSION_REVERSE_DB             ShowDesignOperationResultRequestOperationType
	FACT_LOGIC_TABLE_REVERSE_DB      ShowDesignOperationResultRequestOperationType
	SYNC_TABLES                      ShowDesignOperationResultRequestOperationType
	IMPORT_STANDARD                  ShowDesignOperationResultRequestOperationType
	IMPORT_CODETABLE                 ShowDesignOperationResultRequestOperationType
	IMPORT_ER_TABLE                  ShowDesignOperationResultRequestOperationType
	IMPORT_BUSINESS                  ShowDesignOperationResultRequestOperationType
	TRANSFORM_LOGIC_MODEL            ShowDesignOperationResultRequestOperationType
	PUBLISH_CODETABLE                ShowDesignOperationResultRequestOperationType
	PUBLISH_STANDARD                 ShowDesignOperationResultRequestOperationType
	TABLE_MODEL_RELOCATE             ShowDesignOperationResultRequestOperationType
	DIMENSION_RELOCATE               ShowDesignOperationResultRequestOperationType
	FACT_LOGIC_TABLE_RELOCATE        ShowDesignOperationResultRequestOperationType
	AGGREGATION_LOGIC_TABLE_RELOCATE ShowDesignOperationResultRequestOperationType
	ATOMIC_INDEX_RELOCATE            ShowDesignOperationResultRequestOperationType
	DERIVATIVE_INDEX_RELOCATE        ShowDesignOperationResultRequestOperationType
	COMPOUND_METRIC_RELOCATE         ShowDesignOperationResultRequestOperationType
	BIZ_METRIC_RELOCATE              ShowDesignOperationResultRequestOperationType
	CODE_TABLE_RELOCATE              ShowDesignOperationResultRequestOperationType
	STANDARD_ELEMENT_RELOCATE        ShowDesignOperationResultRequestOperationType
	INFO_ARCH_RELOCATE               ShowDesignOperationResultRequestOperationType
}

func GetShowDesignOperationResultRequestOperationTypeEnum() ShowDesignOperationResultRequestOperationTypeEnum {
	return ShowDesignOperationResultRequestOperationTypeEnum{
		PUBLISH: ShowDesignOperationResultRequestOperationType{
			value: "PUBLISH",
		},
		ADD_TAGS: ShowDesignOperationResultRequestOperationType{
			value: "ADD_TAGS",
		},
		APPROVAL: ShowDesignOperationResultRequestOperationType{
			value: "APPROVAL",
		},
		ER_REVERSE_DB: ShowDesignOperationResultRequestOperationType{
			value: "ER_REVERSE_DB",
		},
		CODETABLE_REVERSE_DB: ShowDesignOperationResultRequestOperationType{
			value: "CODETABLE_REVERSE_DB",
		},
		DIMENSION_REVERSE_DB: ShowDesignOperationResultRequestOperationType{
			value: "DIMENSION_REVERSE_DB",
		},
		FACT_LOGIC_TABLE_REVERSE_DB: ShowDesignOperationResultRequestOperationType{
			value: "FACT_LOGIC_TABLE_REVERSE_DB",
		},
		SYNC_TABLES: ShowDesignOperationResultRequestOperationType{
			value: "SYNC_TABLES",
		},
		IMPORT_STANDARD: ShowDesignOperationResultRequestOperationType{
			value: "IMPORT_STANDARD",
		},
		IMPORT_CODETABLE: ShowDesignOperationResultRequestOperationType{
			value: "IMPORT_CODETABLE",
		},
		IMPORT_ER_TABLE: ShowDesignOperationResultRequestOperationType{
			value: "IMPORT_ER_TABLE",
		},
		IMPORT_BUSINESS: ShowDesignOperationResultRequestOperationType{
			value: "IMPORT_BUSINESS",
		},
		TRANSFORM_LOGIC_MODEL: ShowDesignOperationResultRequestOperationType{
			value: "TRANSFORM_LOGIC_MODEL",
		},
		PUBLISH_CODETABLE: ShowDesignOperationResultRequestOperationType{
			value: "PUBLISH_CODETABLE",
		},
		PUBLISH_STANDARD: ShowDesignOperationResultRequestOperationType{
			value: "PUBLISH_STANDARD",
		},
		TABLE_MODEL_RELOCATE: ShowDesignOperationResultRequestOperationType{
			value: "TABLE_MODEL_RELOCATE",
		},
		DIMENSION_RELOCATE: ShowDesignOperationResultRequestOperationType{
			value: "DIMENSION_RELOCATE",
		},
		FACT_LOGIC_TABLE_RELOCATE: ShowDesignOperationResultRequestOperationType{
			value: "FACT_LOGIC_TABLE_RELOCATE",
		},
		AGGREGATION_LOGIC_TABLE_RELOCATE: ShowDesignOperationResultRequestOperationType{
			value: "AGGREGATION_LOGIC_TABLE_RELOCATE",
		},
		ATOMIC_INDEX_RELOCATE: ShowDesignOperationResultRequestOperationType{
			value: "ATOMIC_INDEX_RELOCATE",
		},
		DERIVATIVE_INDEX_RELOCATE: ShowDesignOperationResultRequestOperationType{
			value: "DERIVATIVE_INDEX_RELOCATE",
		},
		COMPOUND_METRIC_RELOCATE: ShowDesignOperationResultRequestOperationType{
			value: "COMPOUND_METRIC_RELOCATE",
		},
		BIZ_METRIC_RELOCATE: ShowDesignOperationResultRequestOperationType{
			value: "BIZ_METRIC_RELOCATE",
		},
		CODE_TABLE_RELOCATE: ShowDesignOperationResultRequestOperationType{
			value: "CODE_TABLE_RELOCATE",
		},
		STANDARD_ELEMENT_RELOCATE: ShowDesignOperationResultRequestOperationType{
			value: "STANDARD_ELEMENT_RELOCATE",
		},
		INFO_ARCH_RELOCATE: ShowDesignOperationResultRequestOperationType{
			value: "INFO_ARCH_RELOCATE",
		},
	}
}

func (c ShowDesignOperationResultRequestOperationType) Value() string {
	return c.value
}

func (c ShowDesignOperationResultRequestOperationType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowDesignOperationResultRequestOperationType) UnmarshalJSON(b []byte) error {
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
