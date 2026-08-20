package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateIpdProcessInstanceRequest Request Object
type CreateIpdProcessInstanceRequest struct {

	// 项目32位ID，项目唯一标识。通过查询IPD项目列表获取，响应消息体中的id字段的值就是项目ID。
	ProjectId string `json:"project_id"`

	// 操作类型
	OperateType *string `json:"operate_type,omitempty"`

	// 提出项目的domainId
	DomainId *string `json:"domain_id,omitempty"`

	Body *CreateProcessInstanceReq `json:"body,omitempty"`
}

func (o CreateIpdProcessInstanceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateIpdProcessInstanceRequest struct{}"
	}

	return strings.Join([]string{"CreateIpdProcessInstanceRequest", string(data)}, " ")
}
