package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// HibernateClusterResponse Response Object
type HibernateClusterResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o HibernateClusterResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HibernateClusterResponse struct{}"
	}

	return strings.Join([]string{"HibernateClusterResponse", string(data)}, " ")
}
