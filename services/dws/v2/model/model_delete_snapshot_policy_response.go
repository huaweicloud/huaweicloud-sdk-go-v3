package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type DeleteSnapshotPolicyResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteSnapshotPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteSnapshotPolicyResponse struct{}"
	}

	return strings.Join([]string{"DeleteSnapshotPolicyResponse", string(data)}, " ")
}
