package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchMigrateHostsServerRequest Request Object
type BatchMigrateHostsServerRequest struct {
	Body *BatchMigrateServerReq `json:"body,omitempty"`
}

func (o BatchMigrateHostsServerRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchMigrateHostsServerRequest struct{}"
	}

	return strings.Join([]string{"BatchMigrateHostsServerRequest", string(data)}, " ")
}
