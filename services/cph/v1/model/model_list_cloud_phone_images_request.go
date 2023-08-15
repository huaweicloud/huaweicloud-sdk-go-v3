package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListCloudPhoneImagesRequest Request Object
type ListCloudPhoneImagesRequest struct {
}

func (o ListCloudPhoneImagesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCloudPhoneImagesRequest struct{}"
	}

	return strings.Join([]string{"ListCloudPhoneImagesRequest", string(data)}, " ")
}
