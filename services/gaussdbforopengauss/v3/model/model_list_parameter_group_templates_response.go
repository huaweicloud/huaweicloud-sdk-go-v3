package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListParameterGroupTemplatesResponse Response Object
type ListParameterGroupTemplatesResponse struct {

	// 参数模板数量。
	Count *int32 `json:"count,omitempty"`

	// 参数模板列表。
	Configurations *[]ConfigurationsResult `json:"configurations,omitempty"`
	HttpStatusCode int                     `json:"-"`
}

func (o ListParameterGroupTemplatesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListParameterGroupTemplatesResponse struct{}"
	}

	return strings.Join([]string{"ListParameterGroupTemplatesResponse", string(data)}, " ")
}
