package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ZoneData struct {

	// zone的ID，uuid形式的一个资源标识。
	Id *string `json:"id,omitempty"`

	// zone名称。
	Name *string `json:"name,omitempty"`

	// 对zone的描述信息。
	Description *string `json:"description,omitempty"`

	// 管理该zone的管理员邮箱。
	Email *string `json:"email,omitempty"`

	// zone类型，取值system。
	ZoneType *string `json:"zone_type,omitempty"`

	// 该zone下SOA记录中的ttl值。
	Ttl *string `json:"ttl,omitempty"`

	// 该zone下SOA记录中用于标识zone文件变更的序列值，用于主从节点同步。
	Serial *string `json:"serial,omitempty"`

	// 资源状态。
	Status *string `json:"status,omitempty"`

	// 该zone下的recordset个数。
	RecordNum *string `json:"record_num,omitempty"`

	// 托管该zone的pool，由系统分配。
	PoolId *string `json:"pool_id,omitempty"`

	// zone所属的项目ID。
	ProjectId *string `json:"project_id,omitempty"`

	// 创建时间。  采用UTC时间格式，格式为：YYYY-MM-DDTHH:MM:SSZ。
	CreatedAt *string `json:"created_at,omitempty"`

	// 更新时间。  采用UTC时间格式，格式为：YYYY-MM-DDTHH:MM:SSZ。
	UpdatedAt *string `json:"updated_at,omitempty"`

	// 域名所属的区域。
	Region *string `json:"region,omitempty"`

	Links *PageLink `json:"links,omitempty"`

	// 主从模式中，从DNS服务器用以获取DNS信息。
	Masters *string `json:"masters,omitempty"`
}

func (o ZoneData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ZoneData struct{}"
	}

	return strings.Join([]string{"ZoneData", string(data)}, " ")
}
