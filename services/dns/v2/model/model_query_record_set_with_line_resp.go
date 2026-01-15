package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type QueryRecordSetWithLineResp struct {

	// **参数解释：** 记录集的ID。 **取值范围：** 不涉及。
	Id *string `json:"id,omitempty"`

	// **参数解释：** 记录集的名称。 **取值范围：** 不涉及。
	Name *string `json:"name,omitempty"`

	// **参数解释：** 记录集的描述信息。 **取值范围：** 长度不超过255个字符。
	Description *string `json:"description,omitempty"`

	// **参数解释：** 托管该记录的域名ID。 **取值范围：** 不涉及。
	ZoneId *string `json:"zone_id,omitempty"`

	// **参数解释：** 托管该记录的域名。 **取值范围：** 不涉及。
	ZoneName *string `json:"zone_name,omitempty"`

	// **参数解释：** 记录类型。 **取值范围：** - A：将域名解析到指定的IPv4地址。 - AAAA：将域名解析到指定的IPv6地址。 - MX：指定域名对应的邮件服务器。 - CNAME：将域名解析到另一域名，或者多个域名映射到同一域名上。 - TXT：用于对域名进行标识和说明。 - SRV：用于记录某台服务器对外提供了哪些服务。 - NS：指定域名的权威DNS服务器。 - SOA：提供域名的基本信息和权威服务器的详细信息。 - CAA：指定为域名颁发HTTPS证书的授权CA机构，仅支持公网域名解析。 - PTR：指定IP地址反向解析记录，仅支持内网域名解析。
	Type *string `json:"type,omitempty"`

	// **参数解释：** 解析记录在本地DNS服务器的缓存时间，缓存时间越长更新生效越慢，以秒为单位。 **取值范围：** 1~2147483647。
	Ttl *int32 `json:"ttl,omitempty"`

	// **参数解释：** 域名解析后的值。 **取值范围：** 不涉及。
	Records *[]string `json:"records,omitempty"`

	// **参数解释：** 记录集的创建时间。 格式：yyyy-MM-dd'T'HH:mm:ss.SSS。 **取值范围：** 不涉及。
	CreatedAt *string `json:"created_at,omitempty"`

	// **参数解释：** 记录集的最近一次修改时间。 格式：yyyy-MM-dd'T'HH:mm:ss.SSS。 **取值范围：** 不涉及。
	UpdatedAt *string `json:"updated_at,omitempty"`

	// **参数解释：** 记录集状态。 **取值范围：** - ACTIVE：正常 - PENDING_CREATE：创建中 - PENDING_UPDATE：更新中 - PENDING_DELETE：删除中 - PENDING_FREEZE：冻结中 - FREEZE：冻结 - ILLEGAL：违规冻结 - POLICE：公安冻结 - PENDING_DISABLE：暂停中 - DISABLE：暂停 - ERROR：失败
	Status *string `json:"status,omitempty"`

	// **参数解释：** 标识是否由系统默认生成，系统默认生成的记录集不能删除。 **取值范围：** 不涉及。
	Default *bool `json:"default,omitempty"`

	// **参数解释：** 该记录集所属的项目ID。 **取值范围：** 不涉及。
	ProjectId *string `json:"project_id,omitempty"`

	Links *PageLink `json:"links,omitempty"`

	// **参数解释：** 解析线路ID。 **取值范围：** 不涉及。
	Line *string `json:"line,omitempty"`

	// **参数解释：** 解析记录的权重。 **取值范围：** 0~1000。
	Weight *int32 `json:"weight,omitempty"`

	// **参数解释：** 健康检查ID。 **取值范围：** 不涉及。
	HealthCheckId *string `json:"health_check_id,omitempty"`

	AliasTarget *AliasTarget `json:"alias_target,omitempty"`

	// **参数解释：** 规格，默认规格，保留字段。 **取值范围：** 不涉及。
	Bundle *string `json:"bundle,omitempty"`
}

func (o QueryRecordSetWithLineResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QueryRecordSetWithLineResp struct{}"
	}

	return strings.Join([]string{"QueryRecordSetWithLineResp", string(data)}, " ")
}
