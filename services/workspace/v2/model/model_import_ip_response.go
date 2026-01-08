package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ImportIpResponse Response Object
type ImportIpResponse struct {

	// ip成功列表。
	IpList *[]ImportIpInfo `json:"ip_list,omitempty"`

	// ip失败列表。
	FailedIpList *[]ImportIpInfo `json:"failed_ip_list,omitempty"`

	// 错误码。
	FailedCode *string `json:"failed_code,omitempty"`

	// 错误码。
	FailedMessage *string `json:"failed_message,omitempty"`

	// ip列表数量。
	TotalCount     *int32 `json:"total_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ImportIpResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImportIpResponse struct{}"
	}

	return strings.Join([]string{"ImportIpResponse", string(data)}, " ")
}
