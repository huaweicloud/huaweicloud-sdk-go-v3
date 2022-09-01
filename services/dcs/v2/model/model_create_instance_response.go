package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type CreateInstanceResponse struct {

	// 订单ID，仅在创建包周期实例时返回。
	OrderId *string `json:"order_id,omitempty" xml:"order_id"`

	// 缓存实例ID和名称，如果批量创建实例，则会返回多个。
	Instances      *[]Instances `json:"instances,omitempty" xml:"instances"`
	HttpStatusCode int          `json:"-"`
}

func (o CreateInstanceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateInstanceResponse struct{}"
	}

	return strings.Join([]string{"CreateInstanceResponse", string(data)}, " ")
}
