package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListNotificationMaskResourcesResponse Response Object
type ListNotificationMaskResourcesResponse struct {

	// **参数解释**： 通知屏蔽资源列表
	Resources *[]ResourceResp `json:"resources,omitempty"`

	// **参数解释**： 资源总数 **取值范围**： [0,100]
	Count          *int32 `json:"count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListNotificationMaskResourcesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListNotificationMaskResourcesResponse struct{}"
	}

	return strings.Join([]string{"ListNotificationMaskResourcesResponse", string(data)}, " ")
}
