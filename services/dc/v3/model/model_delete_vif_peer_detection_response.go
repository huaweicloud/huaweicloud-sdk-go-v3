package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteVifPeerDetectionResponse Response Object
type DeleteVifPeerDetectionResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteVifPeerDetectionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteVifPeerDetectionResponse struct{}"
	}

	return strings.Join([]string{"DeleteVifPeerDetectionResponse", string(data)}, " ")
}
