package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowResultsResponse struct {
	// 漏洞总数

	Total *int32 `json:"total,omitempty"`
	// 漏洞信息列表

	Data *[]VulnItem `json:"data,omitempty"`

	Statistics     *VulnsLevel `json:"statistics,omitempty"`
	HttpStatusCode int         `json:"-"`
}

func (o ShowResultsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowResultsResponse struct{}"
	}

	return strings.Join([]string{"ShowResultsResponse", string(data)}, " ")
}
