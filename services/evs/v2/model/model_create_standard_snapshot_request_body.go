package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateStandardSnapshotRequestBody struct {
	Snapshot *CreateStandardSnapshot `json:"snapshot"`
}

func (o CreateStandardSnapshotRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateStandardSnapshotRequestBody struct{}"
	}

	return strings.Join([]string{"CreateStandardSnapshotRequestBody", string(data)}, " ")
}
