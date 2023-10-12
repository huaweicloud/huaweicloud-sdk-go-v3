package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowFactoryEnvResponse Response Object
type ShowFactoryEnvResponse struct {

	// 环境变量实体信息
	Params         *[]EnvRespParams `json:"params,omitempty"`
	HttpStatusCode int              `json:"-"`
}

func (o ShowFactoryEnvResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowFactoryEnvResponse struct{}"
	}

	return strings.Join([]string{"ShowFactoryEnvResponse", string(data)}, " ")
}
