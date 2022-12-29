package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type DeletePodResponse struct {

	// pod的ID
	PodId          *string `json:"pod_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeletePodResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeletePodResponse struct{}"
	}

	return strings.Join([]string{"DeletePodResponse", string(data)}, " ")
}
