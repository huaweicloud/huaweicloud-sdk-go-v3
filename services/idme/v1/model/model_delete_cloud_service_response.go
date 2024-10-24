package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteCloudServiceResponse Response Object
type DeleteCloudServiceResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteCloudServiceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteCloudServiceResponse struct{}"
	}

	return strings.Join([]string{"DeleteCloudServiceResponse", string(data)}, " ")
}
