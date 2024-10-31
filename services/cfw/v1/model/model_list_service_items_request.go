package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListServiceItemsRequest Request Object
type ListServiceItemsRequest struct {

	// 服务组id，可通过[获取服务组列表接口](ListServiceSets.xml)查询获得，通过返回值中的data.records.set_id（.表示各对象之间层级的区分）获得。
	SetId string `json:"set_id"`

	// 查询字段，可为服务组成员名称或服务组成员描述的一部分。
	KeyWord *string `json:"key_word,omitempty"`

	// 每页显示个数，范围为1-1024
	Limit int32 `json:"limit"`

	// 偏移量：指定返回记录的开始位置，必须为数字，取值范围为大于或等于0，默认0
	Offset int32 `json:"offset"`

	// 企业项目ID，用户根据组织规划企业项目，对应的ID为企业项目ID，可通过[如何获取企业项目ID](cfw_02_0027.xml)获取，用户未开启企业项目时为0
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 防火墙id，可通过[防火墙ID获取方式](cfw_02_0028.xml)获取
	FwInstanceId *string `json:"fw_instance_id,omitempty"`

	// 查询服务组类型，0表示自定义服务组，1表示预定义服务组。仅当set_id为预定义服务组id时生效
	QueryServiceSetType *int32 `json:"query_service_set_type,omitempty"`
}

func (o ListServiceItemsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListServiceItemsRequest struct{}"
	}

	return strings.Join([]string{"ListServiceItemsRequest", string(data)}, " ")
}
