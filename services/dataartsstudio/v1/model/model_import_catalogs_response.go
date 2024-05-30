package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ImportCatalogsResponse Response Object
type ImportCatalogsResponse struct {
	Data           *ExportDesignModelsResultData `json:"data,omitempty"`
	HttpStatusCode int                           `json:"-"`
}

func (o ImportCatalogsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImportCatalogsResponse struct{}"
	}

	return strings.Join([]string{"ImportCatalogsResponse", string(data)}, " ")
}
