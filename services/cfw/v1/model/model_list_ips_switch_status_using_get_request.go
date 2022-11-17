package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListIpsSwitchStatusUsingGetRequest struct {

	// 防护对象id
	ObjectId string `json:"object_id"`
}

func (o ListIpsSwitchStatusUsingGetRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListIpsSwitchStatusUsingGetRequest struct{}"
	}

	return strings.Join([]string{"ListIpsSwitchStatusUsingGetRequest", string(data)}, " ")
}
