package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RefreshPublicationSnapshotResponse Response Object
type RefreshPublicationSnapshotResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o RefreshPublicationSnapshotResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RefreshPublicationSnapshotResponse struct{}"
	}

	return strings.Join([]string{"RefreshPublicationSnapshotResponse", string(data)}, " ")
}
