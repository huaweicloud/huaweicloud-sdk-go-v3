package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AffectedResource struct {

	// 被防护对象账户ID
	AffectedAttachedDomainId *string `json:"affected_attached_domain_id,omitempty"`

	// 被防护对象项目ID
	AffectedAttachedProjectId *string `json:"affected_attached_project_id,omitempty"`

	AffectedHead *DataResourceHead `json:"affected_head,omitempty"`

	// 资源扩展信息
	AffectedProperties *interface{} `json:"affected_properties,omitempty"`

	// 被防护(受影响）对象在防线系统内唯一ID
	AffectedProtectedId *string `json:"affected_protected_id,omitempty"`

	// 被防护(受影响）对象子类型: 固定为：DB
	AffectedSubtype *string `json:"affected_subtype,omitempty"`

	// 被防护(受影响）对象类型，数据库资产，固定为：Data
	AffectedType *string `json:"affected_type,omitempty"`

	// 被防护对象urn
	AffectedUrn *string `json:"affected_urn,omitempty"`

	// 被防护对象URN扩展
	AffectedUrnext *string `json:"affected_urnext,omitempty"`

	// 被防护(受影响）对象值，云下数据库同affectedProtectedId，云上不传
	AffectedValue *string `json:"affected_value,omitempty"`
}

func (o AffectedResource) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AffectedResource struct{}"
	}

	return strings.Join([]string{"AffectedResource", string(data)}, " ")
}
