package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExecuteAssetActionResponse Response Object
type ExecuteAssetActionResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o ExecuteAssetActionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExecuteAssetActionResponse struct{}"
	}

	return strings.Join([]string{"ExecuteAssetActionResponse", string(data)}, " ")
}
