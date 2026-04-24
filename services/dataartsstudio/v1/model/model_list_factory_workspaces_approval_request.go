package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListFactoryWorkspacesApprovalRequest Request Object
type ListFactoryWorkspacesApprovalRequest struct {

	// 工作空间ID，获取方法请参见[实例ID和工作空间ID](dataartsstudio_02_0350.xml)。
	Workspace string `json:"workspace"`

	// 项目ID，获取方法请参见[项目ID和账号ID](projectid_accountid.xml)。  多project场景采用AK/SK认证的接口请求，则该字段必选。
	XProjectId *string `json:"X-Project-Id,omitempty"`

	// 默认值：application/json;charset=UTF-8 可选，有Body体的情况下必选，没有Body体则无需填写和校验。
	ContentType *string `json:"Content-Type,omitempty"`

	// 申请开始时间，13位时间戳。
	BeginTime *int64 `json:"begin_time,omitempty"`

	// 申请结束时间，13位时间戳。当开始时间和结束时间都不传时，默认查询前7天到当天的数据。
	EndTime *int64 `json:"end_time,omitempty"`

	// 审批状态：  - DEVELOPING: 待审批  - APPROVED: 已通过  - REJECT: 已驳回  默认查询全部的状态
	Status *string `json:"status,omitempty"`

	// 审批类型:  - DEVELOPING: 查询待审批信息  - FINISHED: 查询已审批信息  - APPLY: 查询我的申请页面  默认值：APPLY
	Type *string `json:"type,omitempty"`

	// 申请单号。
	ApplyId *string `json:"apply_id,omitempty"`

	// 审批人。
	ApproverName *string `json:"approver_name,omitempty"`

	// 申请人，该参数只支持在待审批和已审批页面使用。
	CreateUser *string `json:"create_user,omitempty"`

	// 对象名。
	ObjectName *string `json:"object_name,omitempty"`

	// 审批对象： - JOB: 作业 - SCRIPT: 脚本  默认审批全部对象。
	ObjectType *string `json:"object_type,omitempty"`

	// 分页的起始页，取值范围大于等于0。样例: offset=0 默认值: 0。
	Offset *int32 `json:"offset,omitempty"`

	// 分页返回结果，指定每页最大记录数，范围[1,100]。样例: limit=10 默认值: 10。
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListFactoryWorkspacesApprovalRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListFactoryWorkspacesApprovalRequest struct{}"
	}

	return strings.Join([]string{"ListFactoryWorkspacesApprovalRequest", string(data)}, " ")
}
