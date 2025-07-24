package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateActiveDirectoryDomainResponse Response Object
type UpdateActiveDirectoryDomainResponse struct {

	// 修改AD域异步任务的id
	JobId *string `json:"job_id,omitempty"`

	XRequestId     *string `json:"X-request-id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateActiveDirectoryDomainResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateActiveDirectoryDomainResponse struct{}"
	}

	return strings.Join([]string{"UpdateActiveDirectoryDomainResponse", string(data)}, " ")
}
