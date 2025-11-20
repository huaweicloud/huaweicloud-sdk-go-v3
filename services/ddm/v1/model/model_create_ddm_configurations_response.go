package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateDdmConfigurationsResponse Response Object
type CreateDdmConfigurationsResponse struct {

	// 参数组id。
	Id *string `json:"id,omitempty"`

	// 名称。
	Name           *string `json:"name,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateDdmConfigurationsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDdmConfigurationsResponse struct{}"
	}

	return strings.Join([]string{"CreateDdmConfigurationsResponse", string(data)}, " ")
}
