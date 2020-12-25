/*
 * DDS
 *
 * API v3
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

type ResizeInstanceVolumeRequestBody struct {
	Volume *ResizeInstanceVolumeOption `json:"volume"`
}

func (o ResizeInstanceVolumeRequestBody) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ResizeInstanceVolumeRequestBody", string(data)}, " ")
}
