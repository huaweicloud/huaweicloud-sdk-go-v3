package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListCustomerBillsMonthlyBreakDownRequest struct {
	// |忽略大小写，默认 zh_cn：中文 en_us：英文|

	XLanguage *string `json:"X-Language,omitempty"`
	// |参数名称：分摊月| |参数的约束及描述：格式：YYYY-MM。|

	SharedMonth string `json:"shared_month"`
	// |参数名称：计费模式| |参数的约束及描述：枚举型，1、包周期；3、按需；10、预留实例,不传查询全部|

	ChargingMode *int32 `json:"charging_mode,omitempty"`
	// |参数名称：云服务类型编码| |参数的约束及描述：例如ECS的云服务类型编码为“hws.service.type.ec2”。您可以调用查询云服务类型列表接口获取。|

	ServiceTypeCode *string `json:"service_type_code,omitempty"`
	// |参数名称：资源类型编码| |参数的约束及描述：|

	ResourceTypeCode *string `json:"resource_type_code,omitempty"`
	// |参数名称：云服务区编码，| |参数的约束及描述：例如：“cn-north-1”。具体请参见地区和终端节点地区对应云服务的“区域”列的值。|

	RegionCode *string `json:"region_code,omitempty"`
	// |参数名称：账单类型。1：消费-新购2：消费-续订3：消费-变更4：退款-退订5：消费-使用8：消费-自动续订9：调账-补偿12：消费-按时计费13：消费-退订手续费14：消费-服务支持计划月末扣费15：消费-税金16：调账-扣费17：消费-保底差额保底差额=客户签约保底合同后，如果没有达到保底消费，客户需要补交的费用，仅限于直销或者伙伴推荐类子客户，且为后付费用户。100：退款-退订税金101：调账-补偿税金102：调账-扣费税金| |参数的约束及描述：|

	BillType *int32 `json:"bill_type,omitempty"`
	// |参数名称：偏移量，从0开始。默认值为0| |参数的约束及描述：|

	Offset *int32 `json:"offset,omitempty"`
	// |参数名称：每次查询的数量限制。默认值为10。| |参数的约束及描述：|

	Limit *int32 `json:"limit,omitempty"`
	// |参数名称：资源ID。| |参数的约束及描述：|

	ResourceId *string `json:"resource_id,omitempty"`
	// |参数名称：资源名称| |参数的约束及描述：|

	ResourceName *string `json:"resource_name,omitempty"`
	// |参数名称：企业项目标识（企业项目ID）。default项目对应ID：0未归集（表示该云服务不支持企业项目管理能力）项目对应ID：-1其余项目对应ID获取方法请参见如何获取企业项目ID。| |参数的约束及描述：|

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`
	// |参数名称：查询资源消费记录的方式。oneself：客户自己sub_customer：企业子客户all：客户自己和企业子客户默认为all，如果没有企业子客户，取值为all时查询的是客户自己的资源消费记录。| |参数的约束及描述：|

	Method *string `json:"method,omitempty"`
	// |参数名称：企业子账号ID。如果method取值不为sub_customer，则该参数无效。如果method取值为sub_customer，则该参数不能为空。| |参数的约束及描述：|

	SubCustomerId *string `json:"sub_customer_id,omitempty"`
}

func (o ListCustomerBillsMonthlyBreakDownRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCustomerBillsMonthlyBreakDownRequest struct{}"
	}

	return strings.Join([]string{"ListCustomerBillsMonthlyBreakDownRequest", string(data)}, " ")
}
