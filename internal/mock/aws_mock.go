// Code generated by MockGen. DO NOT EDIT.
// Source: internal/client/aws.go
//
// Generated by this command:
//
//	mockgen -source internal/client/aws.go -destination internal/mock/aws_mock.go -package mock
//

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	iam "github.com/aws/aws-sdk-go-v2/service/iam"
	s3 "github.com/aws/aws-sdk-go-v2/service/s3"
	client "github.com/kkb0318/irsa-manager/internal/client"
	gomock "go.uber.org/mock/gomock"
)

// MockAwsIamAPI is a mock of AwsIamAPI interface.
type MockAwsIamAPI struct {
	ctrl     *gomock.Controller
	recorder *MockAwsIamAPIMockRecorder
}

// MockAwsIamAPIMockRecorder is the mock recorder for MockAwsIamAPI.
type MockAwsIamAPIMockRecorder struct {
	mock *MockAwsIamAPI
}

// NewMockAwsIamAPI creates a new mock instance.
func NewMockAwsIamAPI(ctrl *gomock.Controller) *MockAwsIamAPI {
	mock := &MockAwsIamAPI{ctrl: ctrl}
	mock.recorder = &MockAwsIamAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAwsIamAPI) EXPECT() *MockAwsIamAPIMockRecorder {
	return m.recorder
}

// CreateOpenIDConnectProvider mocks base method.
func (m *MockAwsIamAPI) CreateOpenIDConnectProvider(ctx context.Context, params *iam.CreateOpenIDConnectProviderInput, optFns ...func(*iam.Options)) (*iam.CreateOpenIDConnectProviderOutput, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, params}
	for _, a := range optFns {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateOpenIDConnectProvider", varargs...)
	ret0, _ := ret[0].(*iam.CreateOpenIDConnectProviderOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateOpenIDConnectProvider indicates an expected call of CreateOpenIDConnectProvider.
func (mr *MockAwsIamAPIMockRecorder) CreateOpenIDConnectProvider(ctx, params any, optFns ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, params}, optFns...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateOpenIDConnectProvider", reflect.TypeOf((*MockAwsIamAPI)(nil).CreateOpenIDConnectProvider), varargs...)
}

// MockAwsS3API is a mock of AwsS3API interface.
type MockAwsS3API struct {
	ctrl     *gomock.Controller
	recorder *MockAwsS3APIMockRecorder
}

// MockAwsS3APIMockRecorder is the mock recorder for MockAwsS3API.
type MockAwsS3APIMockRecorder struct {
	mock *MockAwsS3API
}

// NewMockAwsS3API creates a new mock instance.
func NewMockAwsS3API(ctrl *gomock.Controller) *MockAwsS3API {
	mock := &MockAwsS3API{ctrl: ctrl}
	mock.recorder = &MockAwsS3APIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAwsS3API) EXPECT() *MockAwsS3APIMockRecorder {
	return m.recorder
}

// CreateBucket mocks base method.
func (m *MockAwsS3API) CreateBucket(ctx context.Context, params *s3.CreateBucketInput, optFns ...func(*s3.Options)) (*s3.CreateBucketOutput, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, params}
	for _, a := range optFns {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateBucket", varargs...)
	ret0, _ := ret[0].(*s3.CreateBucketOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateBucket indicates an expected call of CreateBucket.
func (mr *MockAwsS3APIMockRecorder) CreateBucket(ctx, params any, optFns ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, params}, optFns...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateBucket", reflect.TypeOf((*MockAwsS3API)(nil).CreateBucket), varargs...)
}

// DeletePublicAccessBlock mocks base method.
func (m *MockAwsS3API) DeletePublicAccessBlock(ctx context.Context, params *s3.DeletePublicAccessBlockInput, optFns ...func(*s3.Options)) (*s3.DeletePublicAccessBlockOutput, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, params}
	for _, a := range optFns {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeletePublicAccessBlock", varargs...)
	ret0, _ := ret[0].(*s3.DeletePublicAccessBlockOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeletePublicAccessBlock indicates an expected call of DeletePublicAccessBlock.
func (mr *MockAwsS3APIMockRecorder) DeletePublicAccessBlock(ctx, params any, optFns ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, params}, optFns...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeletePublicAccessBlock", reflect.TypeOf((*MockAwsS3API)(nil).DeletePublicAccessBlock), varargs...)
}

// PutBucketOwnershipControls mocks base method.
func (m *MockAwsS3API) PutBucketOwnershipControls(ctx context.Context, params *s3.PutBucketOwnershipControlsInput, optFns ...func(*s3.Options)) (*s3.PutBucketOwnershipControlsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, params}
	for _, a := range optFns {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "PutBucketOwnershipControls", varargs...)
	ret0, _ := ret[0].(*s3.PutBucketOwnershipControlsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PutBucketOwnershipControls indicates an expected call of PutBucketOwnershipControls.
func (mr *MockAwsS3APIMockRecorder) PutBucketOwnershipControls(ctx, params any, optFns ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, params}, optFns...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PutBucketOwnershipControls", reflect.TypeOf((*MockAwsS3API)(nil).PutBucketOwnershipControls), varargs...)
}

// PutObject mocks base method.
func (m *MockAwsS3API) PutObject(ctx context.Context, params *s3.PutObjectInput, optFns ...func(*s3.Options)) (*s3.PutObjectOutput, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, params}
	for _, a := range optFns {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "PutObject", varargs...)
	ret0, _ := ret[0].(*s3.PutObjectOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PutObject indicates an expected call of PutObject.
func (mr *MockAwsS3APIMockRecorder) PutObject(ctx, params any, optFns ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, params}, optFns...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PutObject", reflect.TypeOf((*MockAwsS3API)(nil).PutObject), varargs...)
}

// MockAwsClient is a mock of AwsClient interface.
type MockAwsClient struct {
	ctrl     *gomock.Controller
	recorder *MockAwsClientMockRecorder
}

// MockAwsClientMockRecorder is the mock recorder for MockAwsClient.
type MockAwsClientMockRecorder struct {
	mock *MockAwsClient
}

// NewMockAwsClient creates a new mock instance.
func NewMockAwsClient(ctrl *gomock.Controller) *MockAwsClient {
	mock := &MockAwsClient{ctrl: ctrl}
	mock.recorder = &MockAwsClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAwsClient) EXPECT() *MockAwsClientMockRecorder {
	return m.recorder
}

// IamCient mocks base method.
func (m *MockAwsClient) IamCient() *client.AwsIamClient {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IamCient")
	ret0, _ := ret[0].(*client.AwsIamClient)
	return ret0
}

// IamCient indicates an expected call of IamCient.
func (mr *MockAwsClientMockRecorder) IamCient() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IamCient", reflect.TypeOf((*MockAwsClient)(nil).IamCient))
}

// S3Cient mocks base method.
func (m *MockAwsClient) S3Cient(region, bucketName string) *client.AwsS3Client {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "S3Cient", region, bucketName)
	ret0, _ := ret[0].(*client.AwsS3Client)
	return ret0
}

// S3Cient indicates an expected call of S3Cient.
func (mr *MockAwsClientMockRecorder) S3Cient(region, bucketName any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "S3Cient", reflect.TypeOf((*MockAwsClient)(nil).S3Cient), region, bucketName)
}
