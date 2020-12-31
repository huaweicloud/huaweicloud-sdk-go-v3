package v1

import (
	http_client "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/dgc/v1/model"
)

type DgcClient struct {
	HcClient *http_client.HcHttpClient
}

func NewDgcClient(hcClient *http_client.HcHttpClient) *DgcClient {
	return &DgcClient{HcClient: hcClient}
}

func DgcClientBuilder() *http_client.HcHttpClientBuilder {
	builder := http_client.NewHcHttpClientBuilder()
	return builder
}

//
func (c *DgcClient) CreateConnection(request *model.CreateConnectionRequest) (*model.CreateConnectionResponse, error) {
	requestDef := GenReqDefForCreateConnection()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateConnectionResponse), nil
	}
}

//
func (c *DgcClient) DeleteConnction(request *model.DeleteConnctionRequest) (*model.DeleteConnctionResponse, error) {
	requestDef := GenReqDefForDeleteConnction()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteConnctionResponse), nil
	}
}

//
func (c *DgcClient) ExportConnections(request *model.ExportConnectionsRequest) (*model.ExportConnectionsResponse, error) {
	requestDef := GenReqDefForExportConnections()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ExportConnectionsResponse), nil
	}
}

//
func (c *DgcClient) ImportConnections(request *model.ImportConnectionsRequest) (*model.ImportConnectionsResponse, error) {
	requestDef := GenReqDefForImportConnections()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ImportConnectionsResponse), nil
	}
}

//
func (c *DgcClient) ListConnections(request *model.ListConnectionsRequest) (*model.ListConnectionsResponse, error) {
	requestDef := GenReqDefForListConnections()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListConnectionsResponse), nil
	}
}

//
func (c *DgcClient) ShowConnection(request *model.ShowConnectionRequest) (*model.ShowConnectionResponse, error) {
	requestDef := GenReqDefForShowConnection()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowConnectionResponse), nil
	}
}

//
func (c *DgcClient) UpdateConnection(request *model.UpdateConnectionRequest) (*model.UpdateConnectionResponse, error) {
	requestDef := GenReqDefForUpdateConnection()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateConnectionResponse), nil
	}
}
