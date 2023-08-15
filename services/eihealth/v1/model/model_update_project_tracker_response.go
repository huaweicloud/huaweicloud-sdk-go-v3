package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateProjectTrackerResponse Response Object
type UpdateProjectTrackerResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateProjectTrackerResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateProjectTrackerResponse struct{}"
	}

	return strings.Join([]string{"UpdateProjectTrackerResponse", string(data)}, " ")
}
