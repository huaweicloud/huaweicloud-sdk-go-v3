package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateTransferResponseBodyLogTransferInfoLogAgencyTransfer 委托转储信息。若转储为委托转储，则会返回该参数
type CreateTransferResponseBodyLogTransferInfoLogAgencyTransfer struct {

	// 委托方账号ID
	AgencyDomainId string `json:"agency_domain_id"`

	// 委托方账号名称
	AgencyDomainName string `json:"agency_domain_name"`

	// 委托方配置的委托名称
	AgencyName string `json:"agency_name"`

	// 委托方项目ID
	AgencyProjectId string `json:"agency_project_id"`

	// 被委托方账号ID，实际配置转储的账号ID
	BeAgencyDomainId string `json:"be_agency_domain_id"`

	// 被委托方项目ID，实际配置转储的账号的项目ID
	BeAgencyProjectId string `json:"be_agency_project_id"`
}

func (o CreateTransferResponseBodyLogTransferInfoLogAgencyTransfer) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateTransferResponseBodyLogTransferInfoLogAgencyTransfer struct{}"
	}

	return strings.Join([]string{"CreateTransferResponseBodyLogTransferInfoLogAgencyTransfer", string(data)}, " ")
}
