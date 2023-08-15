package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListConfigurationDatastoresResponse Response Object
type ListConfigurationDatastoresResponse struct {

	// 引擎信息。
	Datastores     *[]DataStoreList `json:"datastores,omitempty"`
	HttpStatusCode int              `json:"-"`
}

func (o ListConfigurationDatastoresResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListConfigurationDatastoresResponse struct{}"
	}

	return strings.Join([]string{"ListConfigurationDatastoresResponse", string(data)}, " ")
}
