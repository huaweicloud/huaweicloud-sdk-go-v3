package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteAddressItemRequest Request Object
type DeleteAddressItemRequest struct {

	// 地址组成员id，可通过[查询地址组成员接口](ListAddressItems.xml)查询获得，通过返回值中的data.records.item_id（.表示各对象之间层级的区分）获得。
	ItemId string `json:"item_id"`

	// 企业项目ID，用户根据组织规划企业项目，对应的ID为企业项目ID，可通过[如何获取企业项目ID](cfw_02_0027.xml)获取，用户未开启企业项目时为0
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 防火墙id，可通过[防火墙ID获取方式](cfw_02_0028.xml)获取
	FwInstanceId *string `json:"fw_instance_id,omitempty"`
}

func (o DeleteAddressItemRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteAddressItemRequest struct{}"
	}

	return strings.Join([]string{"DeleteAddressItemRequest", string(data)}, " ")
}
