package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SwitchWdrSnapshotResponseBody struct {
}

func (o SwitchWdrSnapshotResponseBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SwitchWdrSnapshotResponseBody struct{}"
	}

	return strings.Join([]string{"SwitchWdrSnapshotResponseBody", string(data)}, " ")
}
