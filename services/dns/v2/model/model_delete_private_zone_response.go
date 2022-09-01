package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type DeletePrivateZoneResponse struct {

	// zone的ID，uuid形式的一个资源标识。
	Id *string `json:"id,omitempty" xml:"id"`

	// zone名称。
	Name *string `json:"name,omitempty" xml:"name"`

	// 对zone的描述信息。
	Description *string `json:"description,omitempty" xml:"description"`

	// 管理该zone的管理员邮箱。
	Email *string `json:"email,omitempty" xml:"email"`

	// zone类型，公网（public）或者内网（private）。
	ZoneType *string `json:"zone_type,omitempty" xml:"zone_type"`

	// 该zone下SOA记录中的ttl值。
	Ttl *int32 `json:"ttl,omitempty" xml:"ttl"`

	// 该zone下SOA记录中用于标识zone文件变更的序列值，用于主从节点同步。
	Serial *int32 `json:"serial,omitempty" xml:"serial"`

	// 状态
	Status *string `json:"status,omitempty" xml:"status"`

	// 该zone下的recordset个数。
	RecordNum *int32 `json:"record_num,omitempty" xml:"record_num"`

	// 托管该zone的pool，由系统分配。
	PoolId *string `json:"pool_id,omitempty" xml:"pool_id"`

	// zone所属的项目ID。
	ProjectId *string `json:"project_id,omitempty" xml:"project_id"`

	// 创建时间。
	CreatedAt *string `json:"created_at,omitempty" xml:"created_at"`

	// 更新时间。
	UpdatedAt *string `json:"updated_at,omitempty" xml:"updated_at"`

	Links *PageLink `json:"links,omitempty" xml:"links"`

	// 主从模式中，从DNS服务器用以获取DNS信息。
	Masters *[]string `json:"masters,omitempty" xml:"masters"`

	// 与该zone关联的Router(VPC)列表。
	Routers        *[]RouterWithStatus `json:"routers,omitempty" xml:"routers"`
	HttpStatusCode int                 `json:"-"`
}

func (o DeletePrivateZoneResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeletePrivateZoneResponse struct{}"
	}

	return strings.Join([]string{"DeletePrivateZoneResponse", string(data)}, " ")
}
