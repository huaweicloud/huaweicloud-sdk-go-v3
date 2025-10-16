package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SwitchWdrSnapshotStatusResponse Response Object
type SwitchWdrSnapshotStatusResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o SwitchWdrSnapshotStatusResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SwitchWdrSnapshotStatusResponse struct{}"
	}

	return strings.Join([]string{"SwitchWdrSnapshotStatusResponse", string(data)}, " ")
}
