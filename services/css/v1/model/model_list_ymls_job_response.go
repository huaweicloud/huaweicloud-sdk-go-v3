package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListYmlsJobResponse struct {
	// 配置列表。

	ConfigList     *[]ConfigList `json:"configList,omitempty"`
	HttpStatusCode int           `json:"-"`
}

func (o ListYmlsJobResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListYmlsJobResponse struct{}"
	}

	return strings.Join([]string{"ListYmlsJobResponse", string(data)}, " ")
}
