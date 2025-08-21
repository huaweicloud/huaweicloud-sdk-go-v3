package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListCesInstanceResponse Response Object
type ListCesInstanceResponse struct {

	// 实例
	Instances *[]ListCesInstanceRspBodyInstances `json:"instances,omitempty"`

	// 总数
	Total          *int32 `json:"total,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListCesInstanceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCesInstanceResponse struct{}"
	}

	return strings.Join([]string{"ListCesInstanceResponse", string(data)}, " ")
}
