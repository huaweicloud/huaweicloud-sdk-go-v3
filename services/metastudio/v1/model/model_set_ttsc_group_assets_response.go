package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SetTtscGroupAssetsResponse Response Object
type SetTtscGroupAssetsResponse struct {

	// 配置项id。
	Id             *string `json:"id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o SetTtscGroupAssetsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetTtscGroupAssetsResponse struct{}"
	}

	return strings.Join([]string{"SetTtscGroupAssetsResponse", string(data)}, " ")
}
