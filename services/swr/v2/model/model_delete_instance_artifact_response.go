package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteInstanceArtifactResponse Response Object
type DeleteInstanceArtifactResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteInstanceArtifactResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteInstanceArtifactResponse struct{}"
	}

	return strings.Join([]string{"DeleteInstanceArtifactResponse", string(data)}, " ")
}
