package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddCicdConfigurationResponse Response Object
type AddCicdConfigurationResponse struct {

	// **参数解释**： cicd标识 **取值范围**： 字符长度1-128位
	CicdId         *interface{} `json:"cicd_id,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o AddCicdConfigurationResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddCicdConfigurationResponse struct{}"
	}

	return strings.Join([]string{"AddCicdConfigurationResponse", string(data)}, " ")
}
