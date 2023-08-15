package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListClustersResponse Response Object
type ListClustersResponse struct {

	// 集群列表，请参见clusters参数说明
	Clusters       *[]Clusters `json:"clusters,omitempty"`
	HttpStatusCode int         `json:"-"`
}

func (o ListClustersResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListClustersResponse struct{}"
	}

	return strings.Join([]string{"ListClustersResponse", string(data)}, " ")
}
