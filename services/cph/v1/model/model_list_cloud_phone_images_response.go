package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListCloudPhoneImagesResponse Response Object
type ListCloudPhoneImagesResponse struct {

	// 请求的唯一标识ID。
	RequestId *string `json:"request_id,omitempty"`

	// 手机镜像信息
	PhoneImages    *[]PhoneImage `json:"phone_images,omitempty"`
	HttpStatusCode int           `json:"-"`
}

func (o ListCloudPhoneImagesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCloudPhoneImagesResponse struct{}"
	}

	return strings.Join([]string{"ListCloudPhoneImagesResponse", string(data)}, " ")
}
