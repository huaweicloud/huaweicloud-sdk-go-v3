package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type BatchImportConfigsResponse struct {
	// 已成功导入的配置项id

	Ids            *interface{} `json:"ids,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o BatchImportConfigsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchImportConfigsResponse struct{}"
	}

	return strings.Join([]string{"BatchImportConfigsResponse", string(data)}, " ")
}
