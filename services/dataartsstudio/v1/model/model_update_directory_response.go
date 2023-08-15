package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateDirectoryResponse Response Object
type UpdateDirectoryResponse struct {

	// 返回的数据信息
	Data           *interface{} `json:"data,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o UpdateDirectoryResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDirectoryResponse struct{}"
	}

	return strings.Join([]string{"UpdateDirectoryResponse", string(data)}, " ")
}
