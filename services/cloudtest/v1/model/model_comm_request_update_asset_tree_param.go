package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CommRequestUpdateAssetTreeParam struct {
	Params *UpdateAssetTreeParam `json:"params"`
}

func (o CommRequestUpdateAssetTreeParam) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CommRequestUpdateAssetTreeParam struct{}"
	}

	return strings.Join([]string{"CommRequestUpdateAssetTreeParam", string(data)}, " ")
}
