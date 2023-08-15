package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListCloudPhonesResponse Response Object
type ListCloudPhonesResponse struct {

	// 云手机信息
	Phones *[]Phone `json:"phones,omitempty"`

	// 请求的唯一标识ID。
	RequestId      *string `json:"request_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListCloudPhonesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCloudPhonesResponse struct{}"
	}

	return strings.Join([]string{"ListCloudPhonesResponse", string(data)}, " ")
}
