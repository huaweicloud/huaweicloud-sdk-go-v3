package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type UpdateFuncSnapshotResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateFuncSnapshotResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateFuncSnapshotResponse struct{}"
	}

	return strings.Join([]string{"UpdateFuncSnapshotResponse", string(data)}, " ")
}
