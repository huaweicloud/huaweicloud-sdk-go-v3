package model

import (
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"strings"
)

type DimensionVo struct {

	// 编码，填写String类型替代Long类型。
	Id *string `json:"id,omitempty"`

	// 字段名。
	NameEn string `json:"name_en"`

	// 维度类型。 枚举值：   - COMMON: 普通维度   - LOOKUP: 码表维度   - HIERARCHIES: 层级维度
	DimensionType DimensionVoDimensionType `json:"dimension_type"`

	// 业务属性。
	NameCh string `json:"name_ch"`

	// 描述。
	Description *string `json:"description,omitempty"`

	// 创建人。
	CreateBy *string `json:"create_by,omitempty"`

	// 更新人。
	UpdateBy *string `json:"update_by,omitempty"`

	// 引用码表ID，填写String类型替代Long类型。
	CodeTableId *string `json:"code_table_id,omitempty"`

	CodeTable *CodeTableVo `json:"code_table,omitempty"`

	// 主题域分组ID，只读，填写String类型替代Long类型。
	L1Id *string `json:"l1_id,omitempty"`

	// 主题域ID，只读，创建和更新时无需填写。
	L2Id *string `json:"l2_id,omitempty"`

	// 业务对象ID，填写String类型替代Long类型。
	L3Id string `json:"l3_id"`

	// 层级属性。
	Hierarchies *[]DimensionHierarchiesVo `json:"hierarchies,omitempty"`

	Status *BizStatusEnum `json:"status,omitempty"`

	// 创建时间，只读，格式遵循RFC3339，精确到秒，UTC时区，即yyyy-mm-ddTHH:MM:SSZ，如1970-01-01T00:00:00Z。
	CreateTime *sdktime.SdkTime `json:"create_time,omitempty"`

	// 更新时间，只读，格式遵循RFC3339，精确到秒，UTC时区，即yyyy-mm-ddTHH:MM:SSZ，如1970-01-01T00:00:00Z。
	UpdateTime *sdktime.SdkTime `json:"update_time,omitempty"`

	ApprovalInfo *ApprovalVo `json:"approval_info,omitempty"`

	NewBiz *BizVersionManageVo `json:"new_biz,omitempty"`

	// 主题域分组中文名，只读，创建和更新时无需填写。
	L1 *string `json:"l1,omitempty"`

	// 主题域中文名，只读，创建和更新时无需填写。
	L2 *string `json:"l2,omitempty"`

	// 业务对象中文名，只读，创建和更新时无需填写。
	L3 *string `json:"l3,omitempty"`

	// 维度属性信息。
	Attributes []DimensionAttributeVo `json:"attributes"`

	// 表映射信息。
	Mappings *[]TableMappingVo `json:"mappings,omitempty"`

	Datasource *BizDatasourceRelationVo `json:"datasource,omitempty"`

	// 资产责任人。
	Owner *string `json:"owner,omitempty"`

	// 外表路径
	ObsLocation *string `json:"obs_location,omitempty"`

	// 表类型。
	TableType *string `json:"table_type,omitempty"`

	// DISTRIBUTE BY [HASH(column)|REPLICATION]。 枚举值：   - HASH: 对指定的列进行Hash，通过映射，把数据分布到指定DN   - REPLICATION: 表的每一行存在所有数据节点（DN）中，即每个数据节点都有完整的表数据
	Distribute *DimensionVoDistribute `json:"distribute,omitempty"`

	// DISTRIBUTE BY HASH column.
	DistributeColumn *string `json:"distribute_column,omitempty"`

	// 别名。
	Alias *string `json:"alias,omitempty"`

	// 自定义项。
	SelfDefinedFields *[]SelfDefinedFieldVo `json:"self_defined_fields,omitempty"`

	// 其他配置
	Configs *string `json:"configs,omitempty"`

	// 开发环境版本，填写String类型替代Long类型。
	DevVersion *string `json:"dev_version,omitempty"`

	// 生产环境版本，填写String类型替代Long类型。
	ProdVersion *string `json:"prod_version,omitempty"`

	// 开发环境版本名称
	DevVersionName *string `json:"dev_version_name,omitempty"`

	// 生产环境版本名称
	ProdVersionName *string `json:"prod_version_name,omitempty"`

	EnvType *EnvTypeEnum `json:"env_type,omitempty"`
}

func (o DimensionVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DimensionVo struct{}"
	}

	return strings.Join([]string{"DimensionVo", string(data)}, " ")
}

type DimensionVoDimensionType struct {
	value string
}

type DimensionVoDimensionTypeEnum struct {
	COMMON      DimensionVoDimensionType
	LOOKUP      DimensionVoDimensionType
	HIERARCHIES DimensionVoDimensionType
}

func GetDimensionVoDimensionTypeEnum() DimensionVoDimensionTypeEnum {
	return DimensionVoDimensionTypeEnum{
		COMMON: DimensionVoDimensionType{
			value: "COMMON",
		},
		LOOKUP: DimensionVoDimensionType{
			value: "LOOKUP",
		},
		HIERARCHIES: DimensionVoDimensionType{
			value: "HIERARCHIES",
		},
	}
}

func (c DimensionVoDimensionType) Value() string {
	return c.value
}

func (c DimensionVoDimensionType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DimensionVoDimensionType) UnmarshalJSON(b []byte) error {
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

type DimensionVoDistribute struct {
	value string
}

type DimensionVoDistributeEnum struct {
	HASH        DimensionVoDistribute
	REPLICATION DimensionVoDistribute
}

func GetDimensionVoDistributeEnum() DimensionVoDistributeEnum {
	return DimensionVoDistributeEnum{
		HASH: DimensionVoDistribute{
			value: "HASH",
		},
		REPLICATION: DimensionVoDistribute{
			value: "REPLICATION",
		},
	}
}

func (c DimensionVoDistribute) Value() string {
	return c.value
}

func (c DimensionVoDistribute) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DimensionVoDistribute) UnmarshalJSON(b []byte) error {
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
