package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RunDrugReceptorPreprocessResponse Response Object
type RunDrugReceptorPreprocessResponse struct {
	Body           *string `json:"body,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o RunDrugReceptorPreprocessResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RunDrugReceptorPreprocessResponse struct{}"
	}

	return strings.Join([]string{"RunDrugReceptorPreprocessResponse", string(data)}, " ")
}
