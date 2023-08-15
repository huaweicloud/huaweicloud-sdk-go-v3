package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowNotebookTokenResponse Response Object
type ShowNotebookTokenResponse struct {

	// notebook鉴权信息
	Token          *string `json:"token,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowNotebookTokenResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowNotebookTokenResponse struct{}"
	}

	return strings.Join([]string{"ShowNotebookTokenResponse", string(data)}, " ")
}
