package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StartAutoCreateSnapshotsResponse Response Object
type StartAutoCreateSnapshotsResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o StartAutoCreateSnapshotsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StartAutoCreateSnapshotsResponse struct{}"
	}

	return strings.Join([]string{"StartAutoCreateSnapshotsResponse", string(data)}, " ")
}
