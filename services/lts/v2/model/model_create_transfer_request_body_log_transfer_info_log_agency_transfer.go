package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 委托转储信息。若配置委托转储，则需要输入该参数
type CreateTransferRequestBodyLogTransferInfoLogAgencyTransfer struct {

	// 委托方账号ID
	AgencyDomainId string `json:"agency_domain_id" xml:"agency_domain_id"`

	// 委托方账号名称
	AgencyDomainName string `json:"agency_domain_name" xml:"agency_domain_name"`

	// 委托方配置的委托名称
	AgencyName string `json:"agency_name" xml:"agency_name"`

	// 委托方项目ID
	AgencyProjectId string `json:"agency_project_id" xml:"agency_project_id"`

	// 被委托方账号ID，实际配置转储的账号ID
	BeAgencyDomainId string `json:"be_agency_domain_id" xml:"be_agency_domain_id"`

	// 被委托方项目ID，实际配置转储的账号的项目ID
	BeAgencyProjectId string `json:"be_agency_project_id" xml:"be_agency_project_id"`
}

func (o CreateTransferRequestBodyLogTransferInfoLogAgencyTransfer) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateTransferRequestBodyLogTransferInfoLogAgencyTransfer struct{}"
	}

	return strings.Join([]string{"CreateTransferRequestBodyLogTransferInfoLogAgencyTransfer", string(data)}, " ")
}
