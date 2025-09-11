package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateWdrSnapshotResponseBody struct {
}

func (o CreateWdrSnapshotResponseBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateWdrSnapshotResponseBody struct{}"
	}

	return strings.Join([]string{"CreateWdrSnapshotResponseBody", string(data)}, " ")
}
