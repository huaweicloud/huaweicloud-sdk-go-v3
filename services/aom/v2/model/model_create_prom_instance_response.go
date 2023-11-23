package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreatePromInstanceResponse Response Object
type CreatePromInstanceResponse struct {

	// 普罗实例列表名称
	Prometheus     *[]PromInstanceEpsModel `json:"prometheus,omitempty"`
	HttpStatusCode int                     `json:"-"`
}

func (o CreatePromInstanceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePromInstanceResponse struct{}"
	}

	return strings.Join([]string{"CreatePromInstanceResponse", string(data)}, " ")
}
