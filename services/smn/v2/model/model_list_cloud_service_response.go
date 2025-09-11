package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListCloudServiceResponse Response Object
type ListCloudServiceResponse struct {

	// 请求的唯一标识ID。
	RequestId *string `json:"request_id,omitempty"`

	// 云服务信息列表。
	CloudServices  *[]ListCloudServiceResponseItemInfo `json:"cloud_services,omitempty"`
	HttpStatusCode int                                 `json:"-"`
}

func (o ListCloudServiceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCloudServiceResponse struct{}"
	}

	return strings.Join([]string{"ListCloudServiceResponse", string(data)}, " ")
}
