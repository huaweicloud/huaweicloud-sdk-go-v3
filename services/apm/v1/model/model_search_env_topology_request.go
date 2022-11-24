package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type SearchEnvTopologyRequest struct {

	// 应用id。
	XBusinessId int64 `json:"x-business-id"`

	Body *EnvTopoRequest `json:"body,omitempty"`
}

func (o SearchEnvTopologyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SearchEnvTopologyRequest struct{}"
	}

	return strings.Join([]string{"SearchEnvTopologyRequest", string(data)}, " ")
}
