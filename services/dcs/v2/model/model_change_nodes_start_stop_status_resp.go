package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ChangeNodesStartStopStatusResp struct {
}

func (o ChangeNodesStartStopStatusResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeNodesStartStopStatusResp struct{}"
	}

	return strings.Join([]string{"ChangeNodesStartStopStatusResp", string(data)}, " ")
}
