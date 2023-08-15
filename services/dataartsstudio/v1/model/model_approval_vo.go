package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ApprovalVo struct {

	// ID
	Id *int64 `json:"id,omitempty"`

	// tenant id
	TenantId *string `json:"tenant_id,omitempty"`

	// 业务中文名
	NameCh *string `json:"name_ch,omitempty"`

	// 业务英文名
	NameEn *string `json:"name_en,omitempty"`

	// 业务ID
	BizId *int64 `json:"biz_id,omitempty"`

	BizType *BizTypeEnum `json:"biz_type,omitempty"`

	// 业务详情
	BizInfo *string `json:"biz_info,omitempty"`

	// 业务详情
	BizInfoObj *interface{} `json:"biz_info_obj,omitempty"`

	// 业务版本
	BizVersion *int32 `json:"biz_version,omitempty"`

	BizStatus *BizStatusEnum `json:"biz_status,omitempty"`

	ApprovalStatus *ApprovalStatusEnum `json:"approval_status,omitempty"`

	ApprovalType *ApprovalTypeEnum `json:"approval_type,omitempty"`

	// 提交时间
	SubmitTime *sdktime.SdkTime `json:"submit_time,omitempty"`

	// 创建者
	CreateBy *string `json:"create_by,omitempty"`

	// 主题域分组
	L1 *string `json:"l1,omitempty"`

	// 主题域
	L2 *string `json:"l2,omitempty"`

	// 业务对象
	L3 *string `json:"l3,omitempty"`

	// 审核时间
	ApprovalTime *sdktime.SdkTime `json:"approval_time,omitempty"`

	// 审核人
	Approver *string `json:"approver,omitempty"`

	// 审核人邮箱
	Email *string `json:"email,omitempty"`

	// 审核信息
	Msg *string `json:"msg,omitempty"`

	// 目录树
	DirectoryPath *string `json:"directory_path,omitempty"`
}

func (o ApprovalVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ApprovalVo struct{}"
	}

	return strings.Join([]string{"ApprovalVo", string(data)}, " ")
}
