package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdatePublicZoneResponse Response Object
type UpdatePublicZoneResponse struct {

	// zone的ID，uuid形式的一个资源标识。
	Id *string `json:"id,omitempty"`

	// zone名称。
	Name *string `json:"name,omitempty"`

	// 对zone的描述信息。
	Description *string `json:"description,omitempty"`

	// 管理该zone的管理员邮箱，用于生成该Zone的SOA记录。
	Email *string `json:"email,omitempty"`

	// zone类型，公网（public）。
	ZoneType *string `json:"zone_type,omitempty"`

	// 该zone下SOA记录中的ttl值。
	Ttl *int32 `json:"ttl,omitempty"`

	// 该zone下SOA记录中用于标识zone文件变更的序列值，用于主从节点同步。  该参数暂未使用。
	Serial *int32 `json:"serial,omitempty"`

	// 资源状态。
	Status *string `json:"status,omitempty"`

	// 该zone下的recordset个数。
	RecordNum *int32 `json:"record_num,omitempty"`

	// 托管该zone的pool，由系统分配。
	PoolId *string `json:"pool_id,omitempty"`

	// zone所属的项目ID。
	ProjectId *string `json:"project_id,omitempty"`

	// 创建时间。  采用UTC时间格式，格式为：YYYY-MM-DDTHH:MM:SSZ
	CreatedAt *string `json:"created_at,omitempty"`

	// 更新时间。  采用UTC时间格式，格式为：YYYY-MM-DDTHH:MM:SSZ
	UpdatedAt *string `json:"updated_at,omitempty"`

	Links *PageLink `json:"links,omitempty"`

	// 主从模式中，从DNS服务器获取DNS信息。  目前暂未使用。
	Masters        *[]string `json:"masters,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o UpdatePublicZoneResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePublicZoneResponse struct{}"
	}

	return strings.Join([]string{"UpdatePublicZoneResponse", string(data)}, " ")
}
