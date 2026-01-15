package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowPtrRecordSetResponse Response Object
type ShowPtrRecordSetResponse struct {

	// 反向解析记录的ID，格式形如{region}:{floatingip_id}。
	Id *string `json:"id,omitempty"`

	// 反向解析记录对应的域名。
	Ptrdname *string `json:"ptrdname,omitempty"`

	// 对反向解析记录的描述。
	Description *string `json:"description,omitempty"`

	// 反向解析记录在本地DNS服务器的缓存时间，缓存时间越长更新生效越慢，以秒为单位。
	Ttl *int32 `json:"ttl,omitempty"`

	// 弹性公网IP的IP地址。
	Address *string `json:"address,omitempty"`

	// **参数解释：** 资源状态。 **取值范围：** - ACTIVE：正常 - PENDING_CREATE：创建中 - PENDING_UPDATE：更新中 - PENDING_DELETE：删除中 - PENDING_FREEZE：冻结中 - FREEZE：冻结 - ILLEGAL：违规冻结 - POLICE：公安冻结 - ERROR：失败
	Status *string `json:"status,omitempty"`

	// **参数解释：** 对该资源的当前操作。 **取值范围：** - CREATE：创建操作 - UPDATE：更新操作 - DELETE：删除操作 - NONE：无操作
	Action *string `json:"action,omitempty"`

	Links *PageLink `json:"links,omitempty"`

	// 反向解析关联的企业项目ID，长度不超过36个字符。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`
	HttpStatusCode      int     `json:"-"`
}

func (o ShowPtrRecordSetResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPtrRecordSetResponse struct{}"
	}

	return strings.Join([]string{"ShowPtrRecordSetResponse", string(data)}, " ")
}
