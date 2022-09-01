package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowRecordSetResponse struct {

	// Record Set的ID。
	Id *string `json:"id,omitempty" xml:"id"`

	// Record Set的名称。
	Name *string `json:"name,omitempty" xml:"name"`

	// Record Set的描述信息。
	Description *string `json:"description,omitempty" xml:"description"`

	// 托管该记录的zone_id。
	ZoneId *string `json:"zone_id,omitempty" xml:"zone_id"`

	// 托管该记录的zone_name。
	ZoneName *string `json:"zone_name,omitempty" xml:"zone_name"`

	// 记录类型。
	Type *string `json:"type,omitempty" xml:"type"`

	// 解析记录在本地DNS服务器的缓存时间，缓存时间越长更新生效越慢，以秒为单位。
	Ttl *int32 `json:"ttl,omitempty" xml:"ttl"`

	// 域名解析后的值。
	Records *[]string `json:"records,omitempty" xml:"records"`

	// 创建时间。
	CreateAt *string `json:"create_at,omitempty" xml:"create_at"`

	// 更新时间。
	UpdateAt *string `json:"update_at,omitempty" xml:"update_at"`

	// 资源状态。
	Status *string `json:"status,omitempty" xml:"status"`

	// 标识是否由系统默认生成，系统默认生成的Record Set不能删除。
	Default *bool `json:"default,omitempty" xml:"default"`

	// 该Record Set所属的项目ID。
	ProjectId *string `json:"project_id,omitempty" xml:"project_id"`

	Links          *PageLink `json:"links,omitempty" xml:"links"`
	HttpStatusCode int       `json:"-"`
}

func (o ShowRecordSetResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowRecordSetResponse struct{}"
	}

	return strings.Join([]string{"ShowRecordSetResponse", string(data)}, " ")
}
