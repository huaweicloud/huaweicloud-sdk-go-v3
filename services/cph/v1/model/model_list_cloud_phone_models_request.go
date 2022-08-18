package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListCloudPhoneModelsRequest struct {
}

func (o ListCloudPhoneModelsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCloudPhoneModelsRequest struct{}"
	}

	return strings.Join([]string{"ListCloudPhoneModelsRequest", string(data)}, " ")
}
