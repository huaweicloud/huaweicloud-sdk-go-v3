package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateClusterConfResponse Response Object
type CreateClusterConfResponse struct {
	Body           map[string]interface{} `json:"body,omitempty"`
	HttpStatusCode int                    `json:"-"`
}

func (o CreateClusterConfResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateClusterConfResponse struct{}"
	}

	return strings.Join([]string{"CreateClusterConfResponse", string(data)}, " ")
}
