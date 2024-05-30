package model

import (
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"strings"
)

type WorkspaceVo struct {

	// 编号，填写String类型替代Long类型。
	Id *string `json:"id,omitempty"`

	// 工作区名字。
	Name string `json:"name"`

	// 描述。
	Description *string `json:"description,omitempty"`

	// 是否为物理表。
	IsPhysical *bool `json:"is_physical,omitempty"`

	// 是否为常用。
	Frequent *bool `json:"frequent,omitempty"`

	// 分层治理。
	Top *bool `json:"top,omitempty"`

	Level *ModelLevel `json:"level,omitempty"`

	// 数据连接类型，对应表所在的数仓类型，取值可以为DLI、DWS、MRS_HIVE、POSTGRESQL、MRS_SPARK、CLICKHOUSE、MYSQL、ORACLE和DORIS等。
	DwType *string `json:"dw_type,omitempty"`

	// 创建时间，只读，格式遵循RFC3339，精确到秒，UTC时区，即yyyy-mm-ddTHH:MM:SSZ，如1970-01-01T00:00:00Z。
	CreateTime *sdktime.SdkTime `json:"create_time,omitempty"`

	// 更新时间，只读，格式遵循RFC3339，精确到秒，UTC时区，即yyyy-mm-ddTHH:MM:SSZ，如1970-01-01T00:00:00Z。
	UpdateTime *sdktime.SdkTime `json:"update_time,omitempty"`

	// 创建人。
	CreateBy *string `json:"create_by,omitempty"`

	// 更新人。
	UpdateBy *string `json:"update_by,omitempty"`

	// 工作区类型枚举。 枚举值：   - THIRD_NF: 关系建模   - DIMENSION: 维度建模
	Type WorkspaceVoType `json:"type"`

	// 关联的业务分层的ID列表 {\"l1Ids\":[],\"l2Ids\":[],\"l3Ids\":[]}。
	BizCatalogIds *string `json:"biz_catalog_ids,omitempty"`

	// 数据库名称数组。
	Databases *[]string `json:"databases,omitempty"`

	// 模型校验前缀，长度不超过100，数字字母下划线组成，字母开头
	TableModelPrefix *string `json:"table_model_prefix,omitempty"`
}

func (o WorkspaceVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WorkspaceVo struct{}"
	}

	return strings.Join([]string{"WorkspaceVo", string(data)}, " ")
}

type WorkspaceVoType struct {
	value string
}

type WorkspaceVoTypeEnum struct {
	THIRD_NF  WorkspaceVoType
	DIMENSION WorkspaceVoType
}

func GetWorkspaceVoTypeEnum() WorkspaceVoTypeEnum {
	return WorkspaceVoTypeEnum{
		THIRD_NF: WorkspaceVoType{
			value: "THIRD_NF",
		},
		DIMENSION: WorkspaceVoType{
			value: "DIMENSION",
		},
	}
}

func (c WorkspaceVoType) Value() string {
	return c.value
}

func (c WorkspaceVoType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *WorkspaceVoType) UnmarshalJSON(b []byte) error {
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
