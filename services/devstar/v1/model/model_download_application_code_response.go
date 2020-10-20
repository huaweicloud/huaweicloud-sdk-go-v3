/*
 * DevStar
 *
 * DevStar API
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type DownloadApplicationCodeResponse struct {
}

func (o DownloadApplicationCodeResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"DownloadApplicationCodeResponse", string(data)}, " ")
}
