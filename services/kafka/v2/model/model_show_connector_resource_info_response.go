package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowConnectorResourceInfoResponse Response Object
type ShowConnectorResourceInfoResponse struct {

	// **参数解释**： 是否售罄。 **取值范围**： - true：售罄。 - false：未售罄。
	SoldOut        *bool `json:"soldOut,omitempty"`
	HttpStatusCode int   `json:"-"`
}

func (o ShowConnectorResourceInfoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowConnectorResourceInfoResponse struct{}"
	}

	return strings.Join([]string{"ShowConnectorResourceInfoResponse", string(data)}, " ")
}
