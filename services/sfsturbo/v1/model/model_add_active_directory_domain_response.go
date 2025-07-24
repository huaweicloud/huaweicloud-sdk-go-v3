package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddActiveDirectoryDomainResponse Response Object
type AddActiveDirectoryDomainResponse struct {

	// 加入AD域异步任务的id
	JobId *string `json:"job_id,omitempty"`

	XRequestId     *string `json:"X-request-id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o AddActiveDirectoryDomainResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddActiveDirectoryDomainResponse struct{}"
	}

	return strings.Join([]string{"AddActiveDirectoryDomainResponse", string(data)}, " ")
}
