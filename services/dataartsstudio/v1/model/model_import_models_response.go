package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ImportModelsResponse Response Object
type ImportModelsResponse struct {

	// 返回的数据信息
	Data           *interface{} `json:"data,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o ImportModelsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImportModelsResponse struct{}"
	}

	return strings.Join([]string{"ImportModelsResponse", string(data)}, " ")
}
