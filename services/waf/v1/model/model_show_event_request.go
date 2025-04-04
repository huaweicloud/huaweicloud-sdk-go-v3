package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowEventRequest Request Object
type ShowEventRequest struct {

	// 语言，默认值为en-us。zh-cn（中文）/en-us（英文）
	XLanguage *string `json:"X-Language,omitempty"`

	// 您可以通过调用企业项目管理服务（EPS）的查询企业项目列表接口（ListEnterpriseProject）查询企业项目id。若需要查询当前用户所有企业项目绑定的资源信息，请传参all_granted_eps。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 防护事件id,通过调用查询攻击事件列表(ListEvent)接口获取防护事件id
	Eventid string `json:"eventid"`
}

func (o ShowEventRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowEventRequest struct{}"
	}

	return strings.Join([]string{"ShowEventRequest", string(data)}, " ")
}
