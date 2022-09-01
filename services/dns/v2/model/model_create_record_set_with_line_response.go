package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type CreateRecordSetWithLineResponse struct {

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

	// 记录类型。  取值范围：A、AAAA、MX、CNAME、TXT、NS、SRV、CAA。
	Type *string `json:"type,omitempty" xml:"type"`

	// 解析记录在本地DNS服务器的缓存时间，缓存时间越长更新生效越慢，以秒为单位。
	Ttl *int32 `json:"ttl,omitempty" xml:"ttl"`

	// 域名解析后的值。
	Records *[]string `json:"records,omitempty" xml:"records"`

	// 创建时间。
	CreatedAt *string `json:"created_at,omitempty" xml:"created_at"`

	// 更新时间。
	UpdatedAt *string `json:"updated_at,omitempty" xml:"updated_at"`

	// 资源状态。
	Status *string `json:"status,omitempty" xml:"status"`

	// 标识是否由系统默认生成，系统默认生成的Record Set不能删除。
	Default *bool `json:"default,omitempty" xml:"default"`

	// 该Record Set所属的项目ID。
	ProjectId *string `json:"project_id,omitempty" xml:"project_id"`

	Links *PageLink `json:"links,omitempty" xml:"links"`

	// 解析线路ID。
	Line *string `json:"line,omitempty" xml:"line"`

	// 解析记录的权重。
	Weight *int32 `json:"weight,omitempty" xml:"weight"`

	// 健康检查ID。
	HealthCheckId *string `json:"health_check_id,omitempty" xml:"health_check_id"`

	AliasTarget    *AliasTarget `json:"alias_target,omitempty" xml:"alias_target"`
	HttpStatusCode int          `json:"-"`
}

func (o CreateRecordSetWithLineResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateRecordSetWithLineResponse struct{}"
	}

	return strings.Join([]string{"CreateRecordSetWithLineResponse", string(data)}, " ")
}
