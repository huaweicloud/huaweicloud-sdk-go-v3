package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type UpdateSnapshotSettingResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateSnapshotSettingResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateSnapshotSettingResponse struct{}"
	}

	return strings.Join([]string{"UpdateSnapshotSettingResponse", string(data)}, " ")
}
