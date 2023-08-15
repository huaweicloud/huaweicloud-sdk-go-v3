package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateSnapshotPolicyResponse Response Object
type CreateSnapshotPolicyResponse struct {
	Body           *string `json:"body,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateSnapshotPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSnapshotPolicyResponse struct{}"
	}

	return strings.Join([]string{"CreateSnapshotPolicyResponse", string(data)}, " ")
}
