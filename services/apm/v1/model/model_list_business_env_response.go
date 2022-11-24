package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListBusinessEnvResponse struct {

	// 环境列表。
	EnvEntryList   *[]EnvEntry `json:"env_entry_list,omitempty"`
	HttpStatusCode int         `json:"-"`
}

func (o ListBusinessEnvResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListBusinessEnvResponse struct{}"
	}

	return strings.Join([]string{"ListBusinessEnvResponse", string(data)}, " ")
}
