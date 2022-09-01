package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListCloudPhoneImagesResponse struct {

	// 请求的唯一标识ID
	RequestId string `json:"request_id" xml:"request_id"`

	// 手机镜像信息
	PhoneImages    []interface{} `json:"phone_images" xml:"phone_images"`
	HttpStatusCode int           `json:"-"`
}

func (o ListCloudPhoneImagesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCloudPhoneImagesResponse struct{}"
	}

	return strings.Join([]string{"ListCloudPhoneImagesResponse", string(data)}, " ")
}
