package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdatePromInstanceResponse Response Object
type UpdatePromInstanceResponse struct {

	// 普罗实例总数。
	Count *int64 `json:"count,omitempty"`

	// 普罗实例列表名称。
	Prometheus     *[]PromInstanceEpsModel `json:"prometheus,omitempty"`
	HttpStatusCode int                     `json:"-"`
}

func (o UpdatePromInstanceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePromInstanceResponse struct{}"
	}

	return strings.Join([]string{"UpdatePromInstanceResponse", string(data)}, " ")
}
