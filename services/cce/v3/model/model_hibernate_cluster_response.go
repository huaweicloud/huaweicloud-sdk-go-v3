/*
 * CCE
 *
 * CCE开放API
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type HibernateClusterResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o HibernateClusterResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"HibernateClusterResponse", string(data)}, " ")
}
