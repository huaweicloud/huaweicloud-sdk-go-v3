package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteAssetEditTaskResponse Response Object
type DeleteAssetEditTaskResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteAssetEditTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteAssetEditTaskResponse struct{}"
	}

	return strings.Join([]string{"DeleteAssetEditTaskResponse", string(data)}, " ")
}
