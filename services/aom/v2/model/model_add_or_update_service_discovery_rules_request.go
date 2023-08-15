package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddOrUpdateServiceDiscoveryRulesRequest Request Object
type AddOrUpdateServiceDiscoveryRulesRequest struct {
	Body *AppRulesBody `json:"body,omitempty"`
}

func (o AddOrUpdateServiceDiscoveryRulesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddOrUpdateServiceDiscoveryRulesRequest struct{}"
	}

	return strings.Join([]string{"AddOrUpdateServiceDiscoveryRulesRequest", string(data)}, " ")
}
