package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AtomicIndexVo struct {

	// 编码
	Id *int64 `json:"id,omitempty"`

	// 字段名
	NameEn string `json:"name_en"`

	// 业务属性
	NameCh string `json:"name_ch"`

	Description *string `json:"description,omitempty"`

	// 创建人
	CreateBy *string `json:"create_by,omitempty"`

	// 计算表达式
	CalExp string `json:"cal_exp"`

	// 计算表达式id
	CalFnIds *[]int64 `json:"cal_fn_ids,omitempty"`

	// 主题域分组id
	L1Id *int64 `json:"l1_id,omitempty"`

	L2Id *string `json:"l2_id,omitempty"`

	// 业务对象guid
	L3Id int64 `json:"l3_id"`

	// 表id
	TableId int64 `json:"table_id"`

	// 表名称
	TbName *string `json:"tb_name,omitempty"`

	// 数据连接类型
	DwType *string `json:"dw_type,omitempty"`

	// 字段id信息
	FieldIds []int64 `json:"field_ids"`

	// 字段名称信息
	FieldNames *[]string `json:"field_names,omitempty"`

	Status *BizStatusEnum `json:"status,omitempty"`

	BizType *BizTypeEnum `json:"biz_type,omitempty"`

	// 创建时间
	CreateTime *sdktime.SdkTime `json:"create_time,omitempty"`

	// 更新时间
	UpdateTime *sdktime.SdkTime `json:"update_time,omitempty"`

	// 主题域分组中文名
	L1 *string `json:"l1,omitempty"`

	// 主题域中文名
	L2 *string `json:"l2,omitempty"`

	// 业务对象中文名
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
