package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdatePrivateZoneStatusResponse Response Object
type UpdatePrivateZoneStatusResponse struct {

	// **参数解释：** 域名的ID，UUID形式的一个资源标识。 **取值范围：** 不涉及。
	Id *string `json:"id,omitempty"`

	// **参数解释：** 域名。 **取值范围：** 不涉及。
	Name *string `json:"name,omitempty"`

	// **参数解释：** 域名的描述信息。 **取值范围：** 长度不超过255个字符。
	Description *string `json:"description,omitempty"`

	// **参数解释：** 管理该域名的管理员邮箱，用于生成该域名的SOA记录。 **取值范围：** 不涉及。
	Email *string `json:"email,omitempty"`

	// **参数解释：** 域名类型。 **取值范围：** public：公网域名。
	ZoneType *string `json:"zone_type,omitempty"`

	// **参数解释：** 该域名下SOA记录中的有效缓存时间，以秒为单位。 **取值范围：** 1~2147483647。
	Ttl *int32 `json:"ttl,omitempty"`

	// **参数解释：** 该域名下SOA记录中用于标识域名文件变更的序列值，用于主从节点同步。 **取值范围：** 不涉及。
	Serial *int32 `json:"serial,omitempty"`

	// **参数解释：** 公网域名状态。 **取值范围：** - ACTIVE：正常 - DISABLE：暂停
	Status *string `json:"status,omitempty"`

	// **参数解释：** 该域名下的记录集个数。 **取值范围：** 不涉及。
	RecordNum *int32 `json:"record_num,omitempty"`

	// **参数解释：** 托管该域名的pool，由系统分配。 **取值范围：** 不涉及。
	PoolId *string `json:"pool_id,omitempty"`

	// **参数解释：** 域名所属的项目ID。 **取值范围：** 不涉及。
	ProjectId *string `json:"project_id,omitempty"`

	// **参数解释：** 域名的创建时间。 格式：yyyy-MM-dd'T'HH:mm:ss.SSS。 **取值范围：** 不涉及。
	CreatedAt *string `json:"created_at,omitempty"`

	// **参数解释：** 域名的最近一次修改时间。 格式：yyyy-MM-dd'T'HH:mm:ss.SSS。 **取值范围：** 不涉及。
	UpdatedAt *string `json:"updated_at,omitempty"`

	Links *PageLink `json:"links,omitempty"`

	// **参数解释：** 主从模式中，从DNS服务器获取DNS信息。目前暂未使用。 **取值范围：** 不涉及。
	Masters        *[]string `json:"masters,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o UpdatePrivateZoneStatusResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePrivateZoneStatusResponse struct{}"
	}

	return strings.Join([]string{"UpdatePrivateZoneStatusResponse", string(data)}, " ")
}
