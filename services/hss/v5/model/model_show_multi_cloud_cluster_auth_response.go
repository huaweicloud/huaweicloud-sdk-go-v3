package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowMultiCloudClusterAuthResponse Response Object
type ShowMultiCloudClusterAuthResponse struct {

	// **参数解释**： 账号总数 **取值范围**： 0-100
	TotalNum *int32 `json:"total_num,omitempty"`

	// **参数解释**： 多云集群的账号权限列表 **取值范围**： 列表项数量0-100
	DataList       *[]MultiCloudClusterAuthInfo `json:"data_list,omitempty"`
	HttpStatusCode int                          `json:"-"`
}

func (o ShowMultiCloudClusterAuthResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowMultiCloudClusterAuthResponse struct{}"
	}

	return strings.Join([]string{"ShowMultiCloudClusterAuthResponse", string(data)}, " ")
}
