package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListServicesResponse Response Object
type ListServicesResponse struct {

	// 云服务名称。
	Services       *[]string `json:"services,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListServicesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListServicesResponse struct{}"
	}

	return strings.Join([]string{"ListServicesResponse", string(data)}, " ")
}
