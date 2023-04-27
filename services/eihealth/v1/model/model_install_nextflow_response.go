package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type InstallNextflowResponse struct {

	// 引擎ID
	Id             *string `json:"id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o InstallNextflowResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InstallNextflowResponse struct{}"
	}

	return strings.Join([]string{"InstallNextflowResponse", string(data)}, " ")
}
