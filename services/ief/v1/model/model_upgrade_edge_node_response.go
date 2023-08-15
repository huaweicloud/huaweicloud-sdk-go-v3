package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpgradeEdgeNodeResponse Response Object
type UpgradeEdgeNodeResponse struct {
	Body           *string `json:"body,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpgradeEdgeNodeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpgradeEdgeNodeResponse struct{}"
	}

	return strings.Join([]string{"UpgradeEdgeNodeResponse", string(data)}, " ")
}
