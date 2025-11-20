package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type IamAccount struct {

	// iam账号。
	IamUser *string `json:"iam_user,omitempty"`

	// iam账号id。
	IamUserId *string `json:"iam_user_id,omitempty"`

	// iam账号权限。
	UserAccount *string `json:"user_account,omitempty"`

	// iam账号token。
	IamToken *string `json:"iam_token,omitempty"`

	// iamdomain账号。
	IamDomain *string `json:"iam_domain,omitempty"`

	// iamdomain账号id。
	IamDomainId *string `json:"iam_domain_id,omitempty"`

	// iamxdomain账号id。
	IamXDomainId *string `json:"iam_x_domain_id,omitempty"`

	// iamxdomain账号类型。
	IamXDomainType *string `json:"iam_x_domain_type,omitempty"`

	// iam项目id。
	IamProjectId *string `json:"iam_project_id,omitempty"`

	// iam项目名称。
	IamProjectName *string `json:"iam_project_name,omitempty"`

	// 语言。
	Language *string `json:"language,omitempty"`

	// 实例id。
	InstanceId *string `json:"instance_id,omitempty"`

	// assumed_by_domain_id。
	AssumedByDomainId *string `json:"assumed_by_domain_id,omitempty"`

	// assumed_by_user_id。
	AssumedByUserId *string `json:"assumed_by_user_id,omitempty"`

	// assumed_by_user_name。
	AssumedByUserName *string `json:"assumed_by_user_name,omitempty"`

	// 角色信息。
	Roles *[]string `json:"roles,omitempty"`

	// exclusive_project_id。
	ExclusiveProjectId *string `json:"exclusive_project_id,omitempty"`

	// 是否支持eps。
	EpsEnable *bool `json:"eps_enable,omitempty"`

	// shared_project_name。
	SharedProjectName *string `json:"shared_project_name,omitempty"`

	// authorization。
	Authorization *string `json:"authorization,omitempty"`

	// 文本属性。
	ContextAttributes *string `json:"context_attributes,omitempty"`

	// 用户文件。
	UserProfile *string `json:"user_profile,omitempty"`

	// sts_token。
	StsToken *string `json:"sts_token,omitempty"`

	// access_key_id。
	AccessKeyId *string `json:"access_key_id,omitempty"`

	// secret_access_key。
	SecretAccessKey *string `json:"secret_access_key,omitempty"`

	// source_account。
	SourceAccount *string `json:"source_account,omitempty"`

	// source_urn。
	SourceUrn *string `json:"source_urn,omitempty"`

	// request_proof。
	RequestProof *string `json:"request_proof,omitempty"`

	// x_project_id。
	XProjectId *string `json:"x_project_id,omitempty"`

	// fulfillment_agency。
	FulfillmentAgency *string `json:"fulfillment_agency,omitempty"`

	// operation_id。
	OperationId *string `json:"operation_id,omitempty"`
}

func (o IamAccount) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IamAccount struct{}"
	}

	return strings.Join([]string{"IamAccount", string(data)}, " ")
}
