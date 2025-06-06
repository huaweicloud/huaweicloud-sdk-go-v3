// Copyright 2022 Huawei Technologies Co.,Ltd.
//
// Licensed to the Apache Software Foundation (ASF) under one
// or more contributor license agreements.  See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership.  The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License.  You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package basic

import (
	"fmt"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/auth"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/auth/cache"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/auth/internal"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/auth/signer"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/impl"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/request"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdkerr"
	"io/ioutil"
	"os"
	"strings"
	"time"
)

const (
	ProjectIdInHeader                 = "X-Project-Id"
	SecurityTokenInHeader             = "X-Security-Token"
	AuthTokenInHeader                 = "X-Auth-Token"
	emptyAk                           = "EMPTY_AK"
	emptySK                           = "EMPTY_SK"
	defaultExpirationThresholdSeconds = 2 * 60 * 60 // 2h
)

var DefaultDerivedPredicate = auth.GetDefaultDerivedPredicate()

type Credentials struct {
	IamEndpoint      string
	AK               string
	SK               string
	ProjectId        string
	SecurityToken    string
	IdpId            string
	IdTokenFile      string
	DerivedPredicate func(*request.DefaultHttpRequest) bool
	MetadataAccessor *internal.MetadataAccessor

	derivedAuthServiceName string
	regionId               string
	authToken              string
	expiredAt              int64
}

// ProcessAuthParams This function may panic under certain circumstances.
func (s *Credentials) ProcessAuthParams(client *impl.DefaultHttpClient, region string) auth.ICredential {
	if s.ProjectId != "" {
		return s
	}

	authCache := cache.GetCache()
	akWithName := s.AK + region
	if projectId, ok := authCache.GetAuth(akWithName); ok {
		s.ProjectId = projectId
		return s
	}

	derivedPredicate := s.DerivedPredicate
	s.DerivedPredicate = nil

	var iamEndpoint string
	if s.IamEndpoint != "" {
		iamEndpoint = s.IamEndpoint
	} else {
		iamEndpoint = internal.GetIamEndpointById(region)
	}
	r := internal.GetKeystoneListProjectsRequest(iamEndpoint, region, client.GetHttpConfig())
	req, err := s.ProcessAuthRequest(client, r)
	if err != nil {
		panic(fmt.Errorf("failed to get project id of region '%s' automatically: %s", region, err.Error()))
	}

	resp, err := internal.KeystoneListProjects(client, req)
	if err != nil {
		panic(fmt.Errorf("failed to get project id of region '%s' automatically, %s", region, err.Error()))
	}
	projects := *resp.Projects
	if len(projects) < 1 {
		panic(fmt.Errorf("failed to get project id of region '%s' automatically, X-IAM-Trace-Id=%s,"+
			" confirm that the project exists in your account, or set project id manually:"+
			" basic.NewCredentialsBuilder().WithAk(ak).WithSk(sk).WithProjectId(projectId).SafeBuild()", region, resp.TraceId))
	} else if len(projects) > 1 {
		projectIds := make([]string, 0, len(projects))
		for _, project := range projects {
			projectIds = append(projectIds, project.Id)
		}
		panic(fmt.Errorf("multiple project ids found: [%s], X-IAM-Trace-Id=%s, please select one when initializing the credentials:"+
			" basic.NewCredentialsBuilder().WithAk(ak).WithSk(sk).WithProjectId(projectId).SafeBuild()", strings.Join(projectIds, ","), resp.TraceId))
	}

	id := projects[0].Id
	s.ProjectId = id
	authCache.PutAuth(akWithName, id)

	s.DerivedPredicate = derivedPredicate

	return s
}

func (s *Credentials) ProcessAuthRequest(client *impl.DefaultHttpClient, req *request.DefaultHttpRequest) (*request.DefaultHttpRequest, error) {
	if s.needUpdateFederalAuthToken() {
		err := s.updateAuthTokenByIdToken(client)
		if err != nil {
			return nil, err
		}
	} else if s.needUpdateSecurityTokenFromMetadata() {
		err := s.UpdateSecurityTokenFromMetadata()
		if err != nil {
			return nil, err
		}
	}

	reqBuilder := req.Builder()
	if s.ProjectId != "" {
		reqBuilder = reqBuilder.AddAutoFilledPathParam("project_id", s.ProjectId).AddHeaderParam(ProjectIdInHeader, s.ProjectId)
	}

	if s.authToken != "" {
		req := reqBuilder.Build()
		req.AddHeaderParam(AuthTokenInHeader, s.authToken)
		return req, nil
	}

	if s.SecurityToken != "" {
		reqBuilder.AddHeaderParam(SecurityTokenInHeader, s.SecurityToken)
	}

	var additionalHeaders map[string]string
	var err error
	if s.IsDerivedAuth(req) {
		additionalHeaders, err = signer.GetDerivedSigner().Sign(reqBuilder.Build(), s.AK, s.SK, s.derivedAuthServiceName, s.regionId)
	} else {
		sn, err := signer.GetSigner(req.GetSigningAlgorithm())
		if err != nil {
			return nil, err
		}
		additionalHeaders, err = sn.Sign(reqBuilder.Build(), s.AK, s.SK)
	}

	if err != nil {
		return nil, err
	}

	for key, value := range additionalHeaders {
		req.AddHeaderParam(key, value)
	}
	return req, nil
}

func (s *Credentials) ProcessDerivedAuthParams(derivedAuthServiceName, regionId string) auth.ICredential {
	if s.derivedAuthServiceName == "" {
		s.derivedAuthServiceName = derivedAuthServiceName
	}

	if s.regionId == "" {
		s.regionId = regionId
	}

	return s
}

func (s *Credentials) IsDerivedAuth(httpRequest *request.DefaultHttpRequest) bool {
	if s.DerivedPredicate == nil {
		return false
	}

	return s.DerivedPredicate(httpRequest)
}

func (s *Credentials) needUpdateSecurityTokenFromMetadata() bool {
	if s.authToken != "" {
		return false
	}
	if s.AK == "" && s.SK == "" {
		return true
	}
	if s.expiredAt == 0 || s.SecurityToken == "" {
		return false
	}
	return s.expiredAt-time.Now().Unix() < defaultExpirationThresholdSeconds
}

func (s *Credentials) UpdateSecurityTokenFromMetadata() error {
	if s.MetadataAccessor == nil {
		s.MetadataAccessor = internal.NewMetadataAccessor()
	}
	credential, err := s.MetadataAccessor.GetCredentials()
	if err != nil {
		return err
	}

	s.AK = credential.Access
	s.SK = credential.Secret
	s.SecurityToken = credential.Securitytoken
	location, err := time.ParseInLocation(`2006-01-02T15:04:05Z`, credential.ExpiresAt, time.UTC)
	if err != nil {
		return err
	}
	s.expiredAt = location.Unix()

	return nil
}

func (s *Credentials) needUpdateFederalAuthToken() bool {
	if s.IdpId == "" || s.IdTokenFile == "" {
		return false
	}
	if s.authToken == "" {
		return true
	}
	return s.expiredAt-time.Now().Unix() < defaultExpirationThresholdSeconds
}

func (s *Credentials) updateAuthTokenByIdToken(client *impl.DefaultHttpClient) error {
	idToken, err := s.getIdToken()
	if err != nil {
		return err
	}

	var iamEndpoint string
	if s.IamEndpoint != "" {
		iamEndpoint = s.IamEndpoint
	} else {
		iamEndpoint = internal.GetIamEndpoint()
	}
	req := internal.GetProjectTokenWithIdTokenRequest(iamEndpoint, s.IdpId, idToken, s.ProjectId, client.GetHttpConfig())
	resp, err := internal.CreateTokenWithIdToken(client, req)
	if err != nil {
		return err
	}

	location, err := time.ParseInLocation(`2006-01-02T15:04:05Z`, resp.Token.ExpiresAt, time.UTC)
	if err != nil {
		return err
	}
	s.expiredAt = location.Unix()
	s.authToken = resp.XSubjectToken
	return nil
}

func (s *Credentials) getIdToken() (string, error) {
	_, err := os.Stat(s.IdTokenFile)
	if err != nil {
		return "", err
	}

	bytes, err := ioutil.ReadFile(s.IdTokenFile)
	if err != nil {
		return "", err
	}
	idToken := string(bytes)
	if idToken == "" {
		return "", sdkerr.NewCredentialsTypeError("id token is empty")
	}
	return idToken, nil
}

type CredentialsBuilder struct {
	Credentials *Credentials
	errMap      map[string]string
}

func NewCredentialsBuilder() *CredentialsBuilder {
	return &CredentialsBuilder{
		Credentials: &Credentials{},
		errMap:      make(map[string]string),
	}
}

func (builder *CredentialsBuilder) WithIamEndpointOverride(endpoint string) *CredentialsBuilder {
	builder.Credentials.IamEndpoint = endpoint
	return builder
}

func (builder *CredentialsBuilder) WithAk(ak string) *CredentialsBuilder {
	if ak == "" {
		builder.errMap[emptyAk] = "input ak cannot be an empty string"
	} else {
		builder.Credentials.AK = ak
		delete(builder.errMap, emptyAk)
	}
	return builder
}

func (builder *CredentialsBuilder) WithSk(sk string) *CredentialsBuilder {
	if sk == "" {
		builder.errMap[emptySK] = "input sk cannot be an empty string"
	} else {
		builder.Credentials.SK = sk
		delete(builder.errMap, emptySK)
	}
	return builder
}

func (builder *CredentialsBuilder) WithProjectId(projectId string) *CredentialsBuilder {
	builder.Credentials.ProjectId = projectId
	return builder
}

func (builder *CredentialsBuilder) WithSecurityToken(token string) *CredentialsBuilder {
	builder.Credentials.SecurityToken = token
	return builder
}

func (builder *CredentialsBuilder) WithDerivedPredicate(derivedPredicate func(*request.DefaultHttpRequest) bool) *CredentialsBuilder {
	builder.Credentials.DerivedPredicate = derivedPredicate
	return builder
}

func (builder *CredentialsBuilder) WithIdpId(idpId string) *CredentialsBuilder {
	builder.Credentials.IdpId = idpId
	return builder
}

func (builder *CredentialsBuilder) WithIdTokenFile(idTokenFile string) *CredentialsBuilder {
	builder.Credentials.IdTokenFile = idTokenFile
	return builder
}

// Deprecated: This function may panic under certain circumstances. Use SafeBuild instead.
func (builder *CredentialsBuilder) Build() *Credentials {
	credentials, err := builder.SafeBuild()
	if err != nil {
		panic(err)
	}
	return credentials
}

func (builder *CredentialsBuilder) SafeBuild() (*Credentials, error) {
	if builder.errMap != nil && len(builder.errMap) != 0 {
		errMsg := "build credentials failed: "
		for _, msg := range builder.errMap {
			errMsg += msg + "; "
		}
		return nil, sdkerr.NewCredentialsTypeError(errMsg)
	}

	if builder.Credentials.IdpId != "" || builder.Credentials.IdTokenFile != "" {
		if builder.Credentials.IdpId == "" {
			return nil, sdkerr.NewCredentialsTypeError("IdpId is required when using IdpId&IdTokenFile")
		}
		if builder.Credentials.IdTokenFile == "" {
			return nil, sdkerr.NewCredentialsTypeError("IdTokenFile is required when using IdpId&IdTokenFile")
		}
		if builder.Credentials.ProjectId == "" {
			return nil, sdkerr.NewCredentialsTypeError("ProjectId is required when using IdpId&IdTokenFile")
		}
	}
	return builder.Credentials, nil
}
