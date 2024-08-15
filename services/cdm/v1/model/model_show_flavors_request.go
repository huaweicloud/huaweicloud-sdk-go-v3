package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowFlavorsRequest Request Object
type ShowFlavorsRequest struct {

	// 版本ID，获取方法请参见[CDM支持的版本](ShowDatastores.xml)。
	DatastoreId string `json:"datastore_id"`
}

func (o ShowFlavorsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowFlavorsRequest struct{}"
	}

	return strings.Join([]string{"ShowFlavorsRequest", string(data)}, " ")
}
