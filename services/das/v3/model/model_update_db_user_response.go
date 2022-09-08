package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type UpdateDbUserResponse struct {

	// 修改结果，修改成功返回OK
	Result         *string `json:"result,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateDbUserResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDbUserResponse struct{}"
	}

	return strings.Join([]string{"UpdateDbUserResponse", string(data)}, " ")
}
