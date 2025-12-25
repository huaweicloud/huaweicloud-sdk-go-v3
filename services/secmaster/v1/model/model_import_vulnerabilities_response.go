package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ImportVulnerabilitiesResponse Response Object
type ImportVulnerabilitiesResponse struct {

	// 错误码
	Code *string `json:"code,omitempty"`

	// 错误信息
	Message *string `json:"message,omitempty"`

	Data           *ImportVulnerabilitiesResponseBodyData `json:"data,omitempty"`
	HttpStatusCode int                                    `json:"-"`
}

func (o ImportVulnerabilitiesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImportVulnerabilitiesResponse struct{}"
	}

	return strings.Join([]string{"ImportVulnerabilitiesResponse", string(data)}, " ")
}
