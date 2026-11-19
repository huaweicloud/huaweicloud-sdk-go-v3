package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowVersionNewServiceResponse Response Object
type ShowVersionNewServiceResponse struct {

	// **参数解释**: 租户最高主机版本 **取值范围**: - hss.version.basic：基础版 - hss.version.advanced：高级版 - hss.version.enterprise：企业版 - hss.version.premium：旗舰版 - hss.version.wtp：网页防篡改版
	Version *string `json:"version,omitempty"`

	// **参数解释**: 租户最高容器版本 **取值范围**: - hss.version.container：容器版
	ContainerVersion *string `json:"container_version,omitempty"`
	HttpStatusCode   int     `json:"-"`
}

func (o ShowVersionNewServiceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowVersionNewServiceResponse struct{}"
	}

	return strings.Join([]string{"ShowVersionNewServiceResponse", string(data)}, " ")
}
