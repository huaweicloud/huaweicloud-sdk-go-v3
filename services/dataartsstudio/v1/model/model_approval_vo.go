package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ApprovalVo 审批信息,只读参数。业务对象最近一次的审批信息，包括审批的业务详情、审核人信息、审核时间等。
type ApprovalVo struct {

	// 审批单ID，填写String类型替代Long类型。
	Id *string `json:"id,omitempty"`

	// 项目ID，获取方式参考接口路径参数“project_id”。
	TenantId *string `json:"tenant_id,omitempty"`

	// 业务中文名。
	NameCh *string `json:"name_ch,omitempty"`

	// 业务英文名。
	NameEn *string `json:"name_en,omitempty"`

	// 业务ID，填写String类型替代Long类型。
	BizId *string `json:"biz_id,omitempty"`

	BizType *BizTypeEnum `json:"biz_type,omitempty"`

	// 序列化之后的业务详情，类型是string。
	BizInfo *string `json:"biz_info,omitempty"`

	// 业务详情，类型是object。
	BizInfoObj *interface{} `json:"biz_info_obj,omitempty"`

	// 业务版本。
	BizVersion *int32 `json:"biz_version,omitempty"`

	BizStatus *BizStatusEnum `json:"biz_status,omitempty"`

	ApprovalStatus *ApprovalStatusEnum `json:"approval_status,omitempty"`

	ApprovalType *ApprovalTypeEnum `json:"approval_type,omitempty"`

	// 提交时间。
	SubmitTime *sdktime.SdkTime `json:"submit_time,omitempty"`

	// 创建者。
	CreateBy *string `json:"create_by,omitempty"`

	// 主题域分组中文名，只读，创建和更新时无需填写。
	L1 *string `json:"l1,omitempty"`

	// 主题域中文名，只读，创建和更新时无需填写。
	L2 *string `json:"l2,omitempty"`

	// 业务对象中文名，只读，创建和更新时无需填写。
	L3 *string `json:"l3,omitempty"`

	// 审核时间。
	ApprovalTime *sdktime.SdkTime `json:"approval_time,omitempty"`

	// 审核人。
	Approver *string `json:"approver,omitempty"`

	// 审核人邮箱。
	Email *string `json:"email,omitempty"`

	// 审核信息。
	Msg *string `json:"msg,omitempty"`

	// 目录树。
	DirectoryPath *string `json:"directory_path,omitempty"`
}

func (o ApprovalVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ApprovalVo struct{}"
	}

	return strings.Join([]string{"ApprovalVo", string(data)}, " ")
}
