package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListParameterGroupTemplatesResponse Response Object
type ListParameterGroupTemplatesResponse struct {

	// **参数解释**: 参数模板总记录数。 **取值范围**: [0, 2147483647]，取决于实际查询大小。
	Count *int32 `json:"count,omitempty"`

	// **参数解释**: 参数模板信息，具体参数请参考ConfigurationsResult。
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
