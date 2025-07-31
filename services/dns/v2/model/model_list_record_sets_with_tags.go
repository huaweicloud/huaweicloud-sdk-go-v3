package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ListRecordSetsWithTags struct {

	// **参数解释：** 记录集的ID。 **取值范围：** 不涉及。
	Id *string `json:"id,omitempty"`

	// **参数解释：** 记录集的名称。 **取值范围：** 不涉及
	Name *string `json:"name,omitempty"`

	// **参数解释：** 记录集的描述信息。 **取值范围：** 长度不超过255个字符。
	Description *string `json:"description,omitempty"`

	// **参数解释：** 托管该记录的域名ID。 **取值范围：** 不涉及。
	ZoneId *string `json:"zone_id,omitempty"`

	// **参数解释：** 托管该记录的域名。 **取值范围：** 不涉及。
	ZoneName *string `json:"zone_name,omitempty"`

	// **参数解释：** 记录类型。 **取值范围：** - 公网域名的记录类型: A、AAAA、MX、CNAME、TXT、SRV、NS、SOA、CAA。 - 内网域名的记录类型: A、AAAA、MX、CNAME、TXT、PTR、SRV、NS、SOA。
	Type *string `json:"type,omitempty"`

	// **参数解释：** 解析记录在本地DNS服务器的缓存时间，缓存时间越长更新生效越慢，以秒为单位。 **取值范围：** 1~2147483647。
	Ttl *int32 `json:"ttl,omitempty"`

	// **参数解释：** 域名解析后的值。 **取值范围：** 不涉及。
	Records *[]string `json:"records,omitempty"`

	// **参数解释：** 记录集的创建时间。 格式：yyyy-MM-dd'T'HH:mm:ss.SSS。 **取值范围：** 不涉及。
	CreateAt *string `json:"create_at,omitempty"`

	// **参数解释：** 记录集的最近一次修改时间。 格式：yyyy-MM-dd'T'HH:mm:ss.SSS。 **取值范围：** 不涉及。
	UpdateAt *string `json:"update_at,omitempty"`

	// **参数解释：** 记录集状态。 **取值范围：** - ACTIVE：正常 - PENDING_CREATE：创建中 - PENDING_UPDATE：更新中 - PENDING_DELETE：删除中 - PENDING_FREEZE：冻结中 - FREEZE：冻结 - ILLEGAL：违规冻结 - POLICE：公安冻结 - PENDING_DISABLE：暂停中 - DISABLE：暂停 - ERROR：失败
	Status *string `json:"status,omitempty"`

	// **参数解释：** 标识是否由系统默认生成，系统默认生成的记录集不能删除。 **取值范围：** 不涉及。
	Default *bool `json:"default,omitempty"`

	// **参数解释：** 该记录集所属的项目ID。 **取值范围：** 不涉及。
	ProjectId *string `json:"project_id,omitempty"`

	Links *PageLink `json:"links,omitempty"`

	// **参数解释：** 资源标签。 **取值范围：** 不涉及。
	Tags *[]Tag `json:"tags,omitempty"`
}

func (o ListRecordSetsWithTags) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRecordSetsWithTags struct{}"
	}

	return strings.Join([]string{"ListRecordSetsWithTags", string(data)}, " ")
}
