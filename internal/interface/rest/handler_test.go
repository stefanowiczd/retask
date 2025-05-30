//go:build unit

package rest

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/stretchr/testify/require"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"go.uber.org/mock/gomock"

	"github.com/stefanowiczd/retask/internal/interface/rest/mock"
)

func Test_calculatePackages(t *testing.T) {
	type testCaseParams struct {
		req                       calculatePackagesReq
		reqBody                   func(r calculatePackagesReq) io.Reader
		mockServicePackageManager func(*gomock.Controller) *mock.MockServicePackageManger
	}

	type testCaseExpected struct {
		statusCode int
		wantError  bool
	}

	type testCase struct {
		name     string
		params   testCaseParams
		expected testCaseExpected
	}

	tests := []testCase{
		{
			name: "invalid request body",
			params: testCaseParams{
				reqBody: func(_ calculatePackagesReq) io.Reader {
					return bytes.NewBuffer([]byte(`{ ... invalid json ... `))
				},
				mockServicePackageManager: func(controller *gomock.Controller) *mock.MockServicePackageManger {
					return mock.NewMockServicePackageManger(controller)
				},
			},
			expected: testCaseExpected{
				statusCode: http.StatusBadRequest,
				wantError:  true,
			},
		},
		{
			name: "invalid request body - param small less than zero",
			params: testCaseParams{
				req: calculatePackagesReq{
					Small:  -1,
					Medium: -1,
					Large:  -1,
				},
				reqBody: func(r calculatePackagesReq) io.Reader {
					body, _ := json.Marshal(r)
					return bytes.NewBuffer(body)
				},
				mockServicePackageManager: func(controller *gomock.Controller) *mock.MockServicePackageManger {
					return mock.NewMockServicePackageManger(controller)
				},
			},
			expected: testCaseExpected{
				statusCode: http.StatusBadRequest,
				wantError:  true,
			},
		},
		{
			name: "invalid request body - param medium less than zero",
			params: testCaseParams{
				req: calculatePackagesReq{
					Small:  10,
					Medium: -1,
					Large:  -1,
				},
				reqBody: func(r calculatePackagesReq) io.Reader {
					body, _ := json.Marshal(r)
					return bytes.NewBuffer(body)
				},
				mockServicePackageManager: func(controller *gomock.Controller) *mock.MockServicePackageManger {
					return mock.NewMockServicePackageManger(controller)
				},
			},
			expected: testCaseExpected{
				statusCode: http.StatusBadRequest,
				wantError:  true,
			},
		},
		{
			name: "invalid request body - param medium large than zero",
			params: testCaseParams{
				req: calculatePackagesReq{
					Small:  10,
					Medium: 10,
					Large:  -1,
				},
				reqBody: func(r calculatePackagesReq) io.Reader {
					body, _ := json.Marshal(r)
					return bytes.NewBuffer(body)
				},
				mockServicePackageManager: func(controller *gomock.Controller) *mock.MockServicePackageManger {
					return mock.NewMockServicePackageManger(controller)
				},
			},
			expected: testCaseExpected{
				statusCode: http.StatusBadRequest,
				wantError:  true,
			},
		},
		{
			name: "service package manager returns internal server error",
			params: testCaseParams{
				req: calculatePackagesReq{
					Small:  10,
					Medium: 10,
					Large:  10,
				},
				reqBody: func(r calculatePackagesReq) io.Reader {
					body, _ := json.Marshal(r)
					return bytes.NewBuffer(body)
				},
				mockServicePackageManager: func(controller *gomock.Controller) *mock.MockServicePackageManger {
					m := mock.NewMockServicePackageManger(controller)
					m.EXPECT().CalculateOptimumPackagesNumber(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(-1, -1, -1, errors.New("internal server error"))

					return m
				},
			},
			expected: testCaseExpected{
				statusCode: http.StatusBadRequest,
				wantError:  true,
			},
		},
		{
			name: "ok",
			params: testCaseParams{
				req: calculatePackagesReq{
					Small:  10,
					Medium: 10,
					Large:  10,
				},
				reqBody: func(r calculatePackagesReq) io.Reader {
					body, _ := json.Marshal(r)
					return bytes.NewBuffer(body)
				},
				mockServicePackageManager: func(controller *gomock.Controller) *mock.MockServicePackageManger {
					m := mock.NewMockServicePackageManger(controller)
					m.EXPECT().CalculateOptimumPackagesNumber(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(-1, -1, -1, nil)

					return m
				},
			},
			expected: testCaseExpected{
				statusCode: http.StatusOK,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			handler := NewHandlerPackageManager(tt.params.mockServicePackageManager(ctrl))

			req := httptest.NewRequest(http.MethodPost, "/packages", tt.params.reqBody(tt.params.req))
			w := httptest.NewRecorder()

			handler.calculatePackages(w, req)

			if tt.expected.wantError {
				require.Equal(t, tt.expected.statusCode, w.Code)
			} else {
				require.Equal(t, tt.expected.statusCode, w.Code)

				// TODO: add further validation of response body
			}
		})
	}
}
