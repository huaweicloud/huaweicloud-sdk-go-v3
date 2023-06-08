package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListParamGroupTemplatesResponse struct {

	// 参数模板数量。
	Count *int32 `json:"count,omitempty"`

	// 参数模板列表。
	Configurations *[]ConfigurationResult `json:"configurations,omitempty"`
	HttpStatusCode int                    `json:"-"`
}

func (o ListParamGroupTemplatesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListParamGroupTemplatesResponse struct{}"
	}

	return strings.Join([]string{"ListParamGroupTemplatesResponse", string(data)}, " ")
}
