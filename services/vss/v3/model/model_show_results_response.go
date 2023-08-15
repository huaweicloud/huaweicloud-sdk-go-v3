package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowResultsResponse Response Object
type ShowResultsResponse struct {

	// 网站漏洞总数
	Total *int32 `json:"total,omitempty"`

	// 网站漏洞信息列表
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
