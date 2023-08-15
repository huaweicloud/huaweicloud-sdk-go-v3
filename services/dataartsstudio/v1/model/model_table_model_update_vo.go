package model

import (
	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"
	"strings"
)

type TableModelUpdateVo struct {

	// 编码
	Id int64 `json:"id"`

	// 模型id
	ModelId int64 `json:"model_id"`

	// 父表id
	ParentTableId *int64 `json:"parent_table_id,omitempty"`

	// 父表名称
	ParentTableName *string `json:"parent_table_name,omitempty"`

	// 父表编码
	ParentTableCode *string `json:"parent_table_code,omitempty"`

	Model *WorkspaceVo `json:"model,omitempty"`

	// 数据格式
	DataFormat *string `json:"data_format,omitempty"`

	// obs桶
	ObsBucket *string `json:"obs_bucket,omitempty"`

	// obs路径
	ObsLocation *string `json:"obs_location,omitempty"`

	// 其他配置
	Configs *string `json:"configs,omitempty"`

	// 表类型
	TableType *string `json:"table_type,omitempty"`

	Owner *string `json:"owner,omitempty"`

	// 表名
	TbName string `json:"tb_name"`

	// 数据连接id
	DwId *string `json:"dw_id,omitempty"`

	// 数据库名
	DbName *string `json:"db_name,omitempty"`

	// dli数据连接执行sql所需的队列，数据连接类型为DLI时必须
	QueueName *string `json:"queue_name,omitempty"`

	// DWS类型需要
	Schema *string `json:"schema,omitempty"`

	// 扩展信息
	ExtendInfo *string `json:"extend_info,omitempty"`

	// 表物化后的guid
	TbGuid *string `json:"tb_guid,omitempty"`

	// 数据表id
	TbId *string `json:"tb_id,omitempty"`

	// 逻辑实体名
	LogicTbName string `json:"logic_tb_name"`

	// 逻辑实体的guid
	LogicTbGuid *string `json:"logic_tb_guid,omitempty"`

	// 描述
	Description string `json:"description"`

	Status *BizStatusEnum `json:"status,omitempty"`

	// 逻辑实体的id
	LogicTbId *int64 `json:"logic_tb_id,omitempty"`

	// 归属的业务分类的id
	BizCatalogId *int64 `json:"biz_catalog_id,omitempty"`

	// 归属的业务分类的路径 {\"l1Id\":\"\",\"l2Id\":\"\",\"l3Id\":\"\"}
	CatalogPath *string `json:"catalog_path,omitempty"`

	// 创建人
	CreateBy *string `json:"create_by,omitempty"`

	// 更新人
	UpdateBy *string `json:"update_by,omitempty"`

	// 创建时间
	CreateTime *sdktime.SdkTime `json:"create_time,omitempty"`

	// 更新时间
	UpdateTime *sdktime.SdkTime `json:"update_time,omitempty"`

	// 表标签
	Tags *[]TagVo `json:"tags,omitempty"`

	ApprovalInfo *ApprovalVo `json:"approval_info,omitempty"`

	NewBiz *BizVersionManageVo `json:"new_biz,omitempty"`

	// 表属性信息
	Attributes []TableModelAttributeVo `json:"attributes"`

	// 表映射信息
	Mappings *[]TableMappingVo `json:"mappings,omitempty"`

	// 关系
	Relations *[]RelationVo `json:"relations,omitempty"`

	// 数据连接类型
	DwType string `json:"dw_type"`

	// 数据连接名称
	DwName *string `json:"dw_name,omitempty"`

	// 主题域分组中文名
	L1 *string `json:"l1,omitempty"`

	// 主题域中文名
	L2 *string `json:"l2,omitempty"`

	// 业务对象中文名
	L3 *string `json:"l3,omitempty"`

	// 主题域分组id
	L1Id *int64 `json:"l1_id,omitempty"`

	L2Id *string `json:"l2_id,omitempty"`

	// 业务对象id
	L3Id *int64 `json:"l3_id,omitempty"`

	// 分区表达式
	PartitionConf *string `json:"partition_conf,omitempty"`

	// DLF 作业 ID
	DlfTaskId *string `json:"dlf_task_id,omitempty"`

	// 是否使用最新分区
	UseRecentlyPartition *bool `json:"use_recently_partition,omitempty"`

	// 是否是逆向的
	Reversed *bool `json:"reversed,omitempty"`

	// 异常数据输出开关
	DirtyOutSwitch *bool `json:"dirty_out_switch,omitempty"`

	// 异常数据输出库
	DirtyOutDatabase *string `json:"dirty_out_database,omitempty"`

	// 异常表前缀
	DirtyOutPrefix *string `json:"dirty_out_prefix,omitempty"`

	// 异常表后缀
	DirtyOutSuffix *string `json:"dirty_out_suffix,omitempty"`

	// 质量责任人
	QualityOwner *string `json:"quality_owner,omitempty"`

	// 质量id
	QualityId *int64 `json:"quality_id,omitempty"`

	// DISTRIBUTE BY [HASH(column)|REPLICATION]
	Distribute *TableModelUpdateVoDistribute `json:"distribute,omitempty"`

	// DISTRIBUTE BY HASH column
	DistributeColumn *string `json:"distribute_column,omitempty"`

	// 是否分区表
	IsPartition *bool `json:"is_partition,omitempty"`

	PhysicalTable *SyncStatusEnum `json:"physical_table,omitempty"`

	TechnicalAsset *SyncStatusEnum `json:"technical_asset,omitempty"`

	BusinessAsset *SyncStatusEnum `json:"business_asset,omitempty"`

	MetaDataLink *SyncStatusEnum `json:"meta_data_link,omitempty"`

	DataQuality *SyncStatusEnum `json:"data_quality,omitempty"`

	SummaryStatus *SyncStatusEnum `json:"summary_status,omitempty"`

	// 别名
	Alias *string `json:"alias,omitempty"`

	// 自定义项
	SelfDefinedFields *[]SelfDefinedFieldVo `json:"self_defined_fields,omitempty"`
}

func (o TableModelUpdateVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TableModelUpdateVo struct{}"
	}

	return strings.Join([]string{"TableModelUpdateVo", string(data)}, " ")
}

type TableModelUpdateVoDistribute struct {
	value string
}

type TableModelUpdateVoDistributeEnum struct {
	HASH        TableModelUpdateVoDistribute
	REPLICATION TableModelUpdateVoDistribute
}

func GetTableModelUpdateVoDistributeEnum() TableModelUpdateVoDistributeEnum {
	return TableModelUpdateVoDistributeEnum{
		HASH: TableModelUpdateVoDistribute{
			value: "HASH",
		},
		REPLICATION: TableModelUpdateVoDistribute{
			value: "REPLICATION",
		},
	}
}

func (c TableModelUpdateVoDistribute) Value() string {
	return c.value
}

func (c TableModelUpdateVoDistribute) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *TableModelUpdateVoDistribute) UnmarshalJSON(b []byte) error {
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
