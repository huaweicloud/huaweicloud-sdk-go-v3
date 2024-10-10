package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateReplicationJobRequest Request Object
type CreateReplicationJobRequest struct {

	// 请求语言类型。
	XLanguage *string `json:"X-Language,omitempty"`

	Body *CreateOfflineTaskReq `json:"body,omitempty"`
}

func (o CreateReplicationJobRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateReplicationJobRequest struct{}"
	}

	return strings.Join([]string{"CreateReplicationJobRequest", string(data)}, " ")
}
