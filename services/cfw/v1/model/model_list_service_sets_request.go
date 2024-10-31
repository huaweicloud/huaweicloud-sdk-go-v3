package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListServiceSetsRequest Request Object
type ListServiceSetsRequest struct {

	// 防护对象id，是创建云防火墙后用于区分互联网边界防护和VPC边界防护的标志id，可通过调用[查询防火墙实例接口](ListFirewallDetail.xml)获得，通过返回值中的data.records.protect_objects.object_id（.表示各对象之间层级的区分）获得，注意type为0的为互联网边界防护对象id，type为1的为VPC边界防护对象id，type可通过data.records.protect_objects.type（.表示各对象之间层级的区分）获得
	ObjectId string `json:"object_id"`

	// 关键字，可使用功能服务组名称和服务组描述的一部分
	KeyWord *string `json:"key_word,omitempty"`

	// 每页查询个数，范围为1-1024
	Limit int32 `json:"limit"`

	// 偏移量：指定返回记录的开始位置，必须为数字，取值范围为大于或等于0，默认0
	Offset int32 `json:"offset"`

	// 企业项目ID，用户根据组织规划企业项目，对应的ID为企业项目ID，可通过[如何获取企业项目ID](cfw_02_0027.xml)获取，用户未开启企业项目时为0
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 防火墙id，可通过[防火墙ID获取方式](cfw_02_0028.xml)获取
	FwInstanceId *string `json:"fw_instance_id,omitempty"`

	// 查询服务组类型，0表示自定义服务组，1表示预定义服务组
	QueryServiceSetType *int32 `json:"query_service_set_type,omitempty"`
}

func (o ListServiceSetsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListServiceSetsRequest struct{}"
	}

	return strings.Join([]string{"ListServiceSetsRequest", string(data)}, " ")
}
