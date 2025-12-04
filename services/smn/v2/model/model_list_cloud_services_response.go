package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListCloudServicesResponse Response Object
type ListCloudServicesResponse struct {

	// 请求的唯一标识ID。
	RequestId *string `json:"request_id,omitempty"`

	// 云服务信息列表。
	CloudServices  *[]ListCloudServiceResponseItemInfo `json:"cloud_services,omitempty"`
	HttpStatusCode int                                 `json:"-"`
}

func (o ListCloudServicesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCloudServicesResponse struct{}"
	}

	return strings.Join([]string{"ListCloudServicesResponse", string(data)}, " ")
}
