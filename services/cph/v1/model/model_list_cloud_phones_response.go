package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListCloudPhonesResponse struct {

	// 请求的唯一标识ID
	RequestId string `json:"request_id" xml:"request_id"`

	// 云手机信息
	Phones         []interface{} `json:"phones" xml:"phones"`
	HttpStatusCode int           `json:"-"`
}

func (o ListCloudPhonesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCloudPhonesResponse struct{}"
	}

	return strings.Join([]string{"ListCloudPhonesResponse", string(data)}, " ")
}
