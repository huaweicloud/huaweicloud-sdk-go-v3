package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreatePublicZoneResponse Response Object
type CreatePublicZoneResponse struct {

	// 域名的ID，UUID形式的一个资源标识。
	Id *string `json:"id,omitempty"`

	// 域名。
	Name *string `json:"name,omitempty"`

	// 对域名的描述信息。
	Description *string `json:"description,omitempty"`

	// 管理该域名的管理员邮箱，用于生成该域名的SOA记录。
	Email *string `json:"email,omitempty"`

	// 域名类型，公网（public）。
	ZoneType *string `json:"zone_type,omitempty"`

	// 该域名下SOA记录中的ttl值。
	Ttl *int32 `json:"ttl,omitempty"`

	// 该域名下SOA记录中用于标识域名文件变更的序列值，用于主从节点同步。
	Serial *int32 `json:"serial,omitempty"`

	// 资源状态。
	Status *string `json:"status,omitempty"`

	// 该域名下的记录集个数。
	RecordNum *int32 `json:"record_num,omitempty"`

	// 托管该域名的pool，由系统分配。
	PoolId *string `json:"pool_id,omitempty"`

	// 域名所属的项目ID。
	ProjectId *string `json:"project_id,omitempty"`

	// 创建时间。 格式：yyyy-MM-dd'T'HH:mm:ss.SSS。
	CreatedAt *string `json:"created_at,omitempty"`

	// 更新时间。 格式：yyyy-MM-dd'T'HH:mm:ss.SSS。
	UpdatedAt *string `json:"updated_at,omitempty"`

	Links *PageLink `json:"links,omitempty"`

	// 域名关联的企业项目ID，长度不超过36个字符。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 主从模式中，从DNS服务器获取DNS信息。目前暂未使用。
	Masters        *[]string `json:"masters,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o CreatePublicZoneResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePublicZoneResponse struct{}"
	}

	return strings.Join([]string{"CreatePublicZoneResponse", string(data)}, " ")
}
