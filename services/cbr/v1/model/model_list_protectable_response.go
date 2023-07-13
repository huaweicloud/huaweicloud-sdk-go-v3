package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListProtectableResponse Response Object
type ListProtectableResponse struct {

	// 可保护性查询实例
	Instances      *[]ProtectablesResp `json:"instances,omitempty"`
	HttpStatusCode int                 `json:"-"`
}

func (o ListProtectableResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListProtectableResponse struct{}"
	}

	return strings.Join([]string{"ListProtectableResponse", string(data)}, " ")
}
