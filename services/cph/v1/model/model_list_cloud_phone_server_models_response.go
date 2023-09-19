package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListCloudPhoneServerModelsResponse Response Object
type ListCloudPhoneServerModelsResponse struct {

	// 请求的唯一标识ID。
	RequestId *string `json:"request_id,omitempty"`

	// 云手机服务器的规格信息。
	ServerModels   *[]ServerModel `json:"server_models,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o ListCloudPhoneServerModelsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCloudPhoneServerModelsResponse struct{}"
	}

	return strings.Join([]string{"ListCloudPhoneServerModelsResponse", string(data)}, " ")
}
