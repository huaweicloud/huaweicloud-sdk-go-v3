package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type RestartOrFlushInstancesResponse struct {

	// 删除/重启/清空实例的结果。
	Results        *[]BatchOpsResult `json:"results,omitempty"`
	HttpStatusCode int               `json:"-"`
}

func (o RestartOrFlushInstancesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RestartOrFlushInstancesResponse struct{}"
	}

	return strings.Join([]string{"RestartOrFlushInstancesResponse", string(data)}, " ")
}
