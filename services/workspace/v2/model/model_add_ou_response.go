package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddOuResponse Response Object
type AddOuResponse struct {

	// 创建OU的id
	Id             *string `json:"id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o AddOuResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddOuResponse struct{}"
	}

	return strings.Join([]string{"AddOuResponse", string(data)}, " ")
}
