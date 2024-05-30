package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AtomicIndexVo struct {

	// 编码，填写String类型替代Long类型。
	Id *string `json:"id,omitempty"`

	// 原子指标英文名。
	NameEn string `json:"name_en"`

	// 原子指标英文名。
	NameCh string `json:"name_ch"`

	// 描述。
	Description *string `json:"description,omitempty"`

	// 创建人。
	CreateBy *string `json:"create_by,omitempty"`

	// 计算表达式，形如'sum(${fact_column_id})'，其中fact_column_id表示引用事实表中的字段ID
	CalExp string `json:"cal_exp"`

	// 引用函数ID，填写String类型替代Long类型。
	CalFnIds *[]string `json:"cal_fn_ids,omitempty"`

	// 主题域分组ID，只读，填写String类型替代Long类型。
	L1Id *string `json:"l1_id,omitempty"`

	// 主题域ID，只读，创建和更新时无需填写。
	L2Id *string `json:"l2_id,omitempty"`

	// 业务对象guid，填写String类型替代Long类型。
	L3Id string `json:"l3_id"`

	// 事实表ID，填写String类型替代Long类型。
	TableId string `json:"table_id"`

	// 事实表名称。
	TbName *string `json:"tb_name,omitempty"`

	// 数据连接类型，对应表所在的数仓类型，取值可以为DLI、DWS、MRS_HIVE、POSTGRESQL、MRS_SPARK、CLICKHOUSE、MYSQL、ORACLE和DORIS等。
	DwType *string `json:"dw_type,omitempty"`

	// 字段ID信息，填写String类型替代Long类型。
	FieldIds []string `json:"field_ids"`

	// 字段名称信息。
	FieldNames *[]string `json:"field_names,omitempty"`

	Status *BizStatusEnum `json:"status,omitempty"`

	BizType *BizTypeEnum `json:"biz_type,omitempty"`

	// 创建时间，只读，格式遵循RFC3339，精确到秒，UTC时区，即yyyy-mm-ddTHH:MM:SSZ，如1970-01-01T00:00:00Z。
	CreateTime *sdktime.SdkTime `json:"create_time,omitempty"`

	// 更新时间，只读，格式遵循RFC3339，精确到秒，UTC时区，即yyyy-mm-ddTHH:MM:SSZ，如1970-01-01T00:00:00Z。
	UpdateTime *sdktime.SdkTime `json:"update_time,omitempty"`

	// 主题域分组中文名，只读，创建和更新时无需填写。
	L1 *string `json:"l1,omitempty"`

	// 主题域中文名，只读，创建和更新时无需填写。
	L2 *string `json:"l2,omitempty"`

	// 业务对象中文名，只读，创建和更新时无需填写。
	L3 *string `json:"l3,omitempty"`

	ApprovalInfo *ApprovalVo `json:"approval_info,omitempty"`

	NewBiz *BizVersionManageVo `json:"new_biz,omitempty"`
}

func (o AtomicIndexVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AtomicIndexVo struct{}"
	}

	return strings.Join([]string{"AtomicIndexVo", string(data)}, " ")
}
