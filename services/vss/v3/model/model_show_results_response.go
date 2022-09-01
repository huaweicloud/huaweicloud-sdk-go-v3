package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowResultsResponse struct {

	// 漏洞总数
	Total *int32 `json:"total,omitempty" xml:"total"`

	// 漏洞信息列表
	Data *[]VulnItem `json:"data,omitempty" xml:"data"`

	Statistics     *VulnsLevel `json:"statistics,omitempty" xml:"statistics"`
	HttpStatusCode int         `json:"-"`
}

func (o ShowResultsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowResultsResponse struct{}"
	}

	return strings.Join([]string{"ShowResultsResponse", string(data)}, " ")
}
