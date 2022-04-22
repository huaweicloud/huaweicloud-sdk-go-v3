package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListYmlsResponse struct {

	// 配置列表。
	ConfigList     *[]ConfigList `json:"configList,omitempty"`
	HttpStatusCode int           `json:"-"`
}

func (o ListYmlsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListYmlsResponse struct{}"
	}

	return strings.Join([]string{"ListYmlsResponse", string(data)}, " ")
}
