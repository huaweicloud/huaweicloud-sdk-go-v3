package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAddressSetDetailRequest Request Object
type ListAddressSetDetailRequest struct {

	// 地址组id，可通过[查询地址组列表接口](ListAddressSets.xml)查询获得，通过返回值中的data.records.set_id（.表示各对象之间层级的区分）获得。
	SetId string `json:"set_id"`

	// 企业项目ID，用户根据组织规划企业项目，对应的ID为企业项目ID，可通过[如何获取企业项目ID](cfw_02_0027.xml)获取，用户未开启企业项目时为0
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 防火墙ID，可通过[防火墙ID获取方式](cfw_02_0028.xml)获取
	FwInstanceId *string `json:"fw_instance_id,omitempty"`

	// 查询地址组类型，0表示自定义地址组，1表示预定义地址组
	QueryAddressSetType *int32 `json:"query_address_set_type,omitempty"`
}

func (o ListAddressSetDetailRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAddressSetDetailRequest struct{}"
	}

	return strings.Join([]string{"ListAddressSetDetailRequest", string(data)}, " ")
}
