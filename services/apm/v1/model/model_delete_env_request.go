package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type DeleteEnvRequest struct {

	// 环境id
	EnvId int64 `json:"env_id"`

	// 应用id，用于鉴权
	XBusinessId int64 `json:"x-business-id"`
}

func (o DeleteEnvRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteEnvRequest struct{}"
	}

	return strings.Join([]string{"DeleteEnvRequest", string(data)}, " ")
}
