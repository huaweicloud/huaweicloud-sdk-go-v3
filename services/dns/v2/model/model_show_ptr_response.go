package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowPtrResponse Response Object
type ShowPtrResponse struct {
	Publicip *PublicIpRes `json:"publicip,omitempty"`

	// 反向解析记录对应的域名列表,最大个数不超过10个。
	Ptrdnames *[]string `json:"ptrdnames,omitempty"`

	// 反向解析记录的ID。
	Id *string `json:"id,omitempty"`

	// 对反向解析记录的描述。
	Description *string `json:"description,omitempty"`

	// 反向解析记录在本地DNS服务器的缓存时间，缓存时间越长更新生效越慢，以秒为单位。
	Ttl *int32 `json:"ttl,omitempty"`

	// **参数解释：** 资源状态。 **取值范围：** - ACTIVE：正常 - PENDING_CREATE：创建中 - PENDING_UPDATE：更新中 - PENDING_DELETE：删除中 - PENDING_FREEZE：冻结中 - FREEZE：冻结 - ILLEGAL：违规冻结 - POLICE：公安冻结 - ERROR：失败
	Status *string `json:"status,omitempty"`

	Links *PageLink `json:"links,omitempty"`

	// 反向解析关联的企业项目ID，长度不超过36个字符。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`
	HttpStatusCode      int     `json:"-"`
}

func (o ShowPtrResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPtrResponse struct{}"
	}

	return strings.Join([]string{"ShowPtrResponse", string(data)}, " ")
}
