package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateRecordSetResponse Response Object
type CreateRecordSetResponse struct {

	// 记录集的ID。
	Id *string `json:"id,omitempty"`

	// 记录集的名称。
	Name *string `json:"name,omitempty"`

	// 记录集的描述信息。
	Description *string `json:"description,omitempty"`

	// 托管该记录的域名ID。
	ZoneId *string `json:"zone_id,omitempty"`

	// 托管该记录的域名。
	ZoneName *string `json:"zone_name,omitempty"`

	// 记录类型。  公网域名场景的记录类型: A、AAAA、MX、CNAME、TXT、NS、SRV、CAA。  内网域名场景的记录类型: A、AAAA、MX、CNAME、TXT、SRV。
	Type *string `json:"type,omitempty"`

	// 解析记录在本地DNS服务器的缓存时间，缓存时间越长更新生效越慢，以秒为单位。 如果您的服务地址经常更换，建议TTL值设置相对小些，反之，建议设置相对大些。
	Ttl *int32 `json:"ttl,omitempty"`

	// 域名解析后的值。
	Records *[]string `json:"records,omitempty"`

	// 创建时间。 格式：yyyy-MM-dd'T'HH:mm:ss.SSS。
	CreateAt *string `json:"create_at,omitempty"`

	// 更新时间。 格式：yyyy-MM-dd'T'HH:mm:ss.SSS。
	UpdateAt *string `json:"update_at,omitempty"`

	// 资源状态。
	Status *string `json:"status,omitempty"`

	// 标识是否由系统默认生成，系统默认生成的记录集不能删除。
	Default *bool `json:"default,omitempty"`

	// 该记录集所属的项目ID。
	ProjectId *string `json:"project_id,omitempty"`

	Links          *PageLink `json:"links,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o CreateRecordSetResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateRecordSetResponse struct{}"
	}

	return strings.Join([]string{"CreateRecordSetResponse", string(data)}, " ")
}
