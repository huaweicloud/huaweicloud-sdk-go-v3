package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type UpdatePtrRecordResponse struct {

	// PTR记录的ID，格式形如{region}:{floatingip_id}。
	Id *string `json:"id,omitempty" xml:"id"`

	// PTR记录对应的域名。
	Ptrdname *string `json:"ptrdname,omitempty" xml:"ptrdname"`

	// 对PTR记录的描述。
	Description *string `json:"description,omitempty" xml:"description"`

	// PTR记录在本地DNS服务器的缓存时间，缓存时间越长更新生效越慢，以秒为单位。
	Ttl *int32 `json:"ttl,omitempty" xml:"ttl"`

	// 弹性IP的IP地址。
	Address *string `json:"address,omitempty" xml:"address"`

	// 资源状态。
	Status *string `json:"status,omitempty" xml:"status"`

	// 对该资源的当前操作。  取值范围：  CREATE：表示创建 UPDATE：表示更新 DELETE：表示删除 NONE：表示无操作
	Action *string `json:"action,omitempty" xml:"action"`

	Links          *PageLink `json:"links,omitempty" xml:"links"`
	HttpStatusCode int       `json:"-"`
}

func (o UpdatePtrRecordResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePtrRecordResponse struct{}"
	}

	return strings.Join([]string{"UpdatePtrRecordResponse", string(data)}, " ")
}
