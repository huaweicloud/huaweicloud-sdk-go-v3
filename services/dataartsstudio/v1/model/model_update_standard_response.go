package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateStandardResponse Response Object
type UpdateStandardResponse struct {

	// 返回的数据信息
	Data           *interface{} `json:"data,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o UpdateStandardResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateStandardResponse struct{}"
	}

	return strings.Join([]string{"UpdateStandardResponse", string(data)}, " ")
}
