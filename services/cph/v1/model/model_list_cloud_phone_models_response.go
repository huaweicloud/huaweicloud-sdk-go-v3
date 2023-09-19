package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListCloudPhoneModelsResponse Response Object
type ListCloudPhoneModelsResponse struct {

	// 请求的唯一标识ID。
	RequestId *string `json:"request_id,omitempty"`

	// 云手机的规格信息。
	PhoneModels    *[]PhoneModel `json:"phone_models,omitempty"`
	HttpStatusCode int           `json:"-"`
}

func (o ListCloudPhoneModelsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCloudPhoneModelsResponse struct{}"
	}

	return strings.Join([]string{"ListCloudPhoneModelsResponse", string(data)}, " ")
}
