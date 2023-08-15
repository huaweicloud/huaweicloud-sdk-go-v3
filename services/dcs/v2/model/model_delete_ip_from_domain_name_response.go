package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteIpFromDomainNameResponse Response Object
type DeleteIpFromDomainNameResponse struct {

	// 域名摘除ip的任务ID。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteIpFromDomainNameResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteIpFromDomainNameResponse struct{}"
	}

	return strings.Join([]string{"DeleteIpFromDomainNameResponse", string(data)}, " ")
}
