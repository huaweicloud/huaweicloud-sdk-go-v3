package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RegisterDomainResponse Response Object
type RegisterDomainResponse struct {

	// 任务ID。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o RegisterDomainResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RegisterDomainResponse struct{}"
	}

	return strings.Join([]string{"RegisterDomainResponse", string(data)}, " ")
}
