package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteActiveDirectoryDomainResponse Response Object
type DeleteActiveDirectoryDomainResponse struct {

	// 退出AD域异步任务的id
	JobId *string `json:"job_id,omitempty"`

	XRequestId     *string `json:"X-request-id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteActiveDirectoryDomainResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteActiveDirectoryDomainResponse struct{}"
	}

	return strings.Join([]string{"DeleteActiveDirectoryDomainResponse", string(data)}, " ")
}
