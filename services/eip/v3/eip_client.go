package v3

import (
	http_client "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/eip/v3/model"
)

type EipClient struct {
	hcClient *http_client.HcHttpClient
}

func NewEipClient(hcClient *http_client.HcHttpClient) *EipClient {
	return &EipClient{hcClient: hcClient}
}

func EipClientBuilder() *http_client.HcHttpClientBuilder {
	builder := http_client.NewHcHttpClientBuilder()
	return builder
}

//绑定弹性公网IP
func (c *EipClient) AssociatePublicips(request *model.AssociatePublicipsRequest) (*model.AssociatePublicipsResponse, error) {
	requestDef := GenReqDefForAssociatePublicips(request)
	resp, responseDef := GenRespForAssociatePublicips()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//解绑弹性公网IP
func (c *EipClient) DisassociatePublicips(request *model.DisassociatePublicipsRequest) (*model.DisassociatePublicipsResponse, error) {
	requestDef := GenReqDefForDisassociatePublicips(request)
	resp, responseDef := GenRespForDisassociatePublicips()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//查询弹性公网IP列表信息
func (c *EipClient) ListPublicips(request *model.ListPublicipsRequest) (*model.ListPublicipsResponse, error) {
	requestDef := GenReqDefForListPublicips(request)
	resp, responseDef := GenRespForListPublicips()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//查询弹性公网IP详情
func (c *EipClient) ShowPublicip(request *model.ShowPublicipRequest) (*model.ShowPublicipResponse, error) {
	requestDef := GenReqDefForShowPublicip(request)
	resp, responseDef := GenRespForShowPublicip()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}
