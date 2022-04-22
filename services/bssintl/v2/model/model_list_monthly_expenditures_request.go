package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListMonthlyExpendituresRequest struct {

	// 查询消费汇总账单所在的账期，格式为YYYY-MM。
	Cycle string `json:"cycle"`

	// 云服务类型编码，例如OBS的云服务类型编码为“hws.service.type.obs”。您可以调用查询云服务类型列表接口获取。当不传递此参数时，查询的账单是以云服务类型为维度的月度消费账单。当传递此参数时，查询的账单是该云服务类型下以资源类型为维度的月度消费账单。
	CloudServiceTypeCode *string `json:"cloud_service_type_code,omitempty"`

	// 0：华为云账户 1：伙伴设置预算账户，仅当客户关联合作伙伴且关联类型为转售模式时，才会存在伙伴拨款设置预算账户。不传此参数默认查询华为云账户下的消费汇总。
	Type *string `json:"type,omitempty"`

	// 企业项目ID
	EnterpriseProjectId *string `json:"enterpriseProjectId,omitempty"`

	// 用户的标识，由IAM统一分配
	DomainId string `json:"domain_id"`
}

func (o ListMonthlyExpendituresRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListMonthlyExpendituresRequest struct{}"
	}

	return strings.Join([]string{"ListMonthlyExpendituresRequest", string(data)}, " ")
}
