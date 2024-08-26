package handler

import (
	"crud-customer/config"
	"crud-customer/internal/entity"
	"crud-customer/internal/service"
	mockservice "crud-customer/mocks/internal_/service"
	"crud-customer/util/typehelper"
	"crud-customer/util/validator"
	"encoding/json"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func Test_customerImpl_CreateCustomer(t *testing.T) {
	type fields struct {
		customerService service.Customer
		cfg             *config.Config
	}
	type testCase struct {
		name       string
		fields     fields
		req        *http.Request
		rec        *httptest.ResponseRecorder
		setupFunc  func(t *testing.T, tt *testCase)
		wantStatus int
		wantResp   string
		wantErr    assert.ErrorAssertionFunc
	}

	testCases := []testCase{
		{
			name: "success",
			fields: fields{
				customerService: mockservice.NewCustomer(t),
				cfg:             &config.Config{},
			},
			req: httptest.NewRequest(http.MethodPost, "/customers", strings.NewReader(`{"name":"test","age":20}`)),
			rec: httptest.NewRecorder(),
			setupFunc: func(t *testing.T, tt *testCase) {

				tt.req.Header.Set("Content-Type", "application/json")
				tt.fields.customerService.(*mockservice.Customer).EXPECT().CreateCustomer(mock.Anything, "test", uint(20)).Return(&entity.Customer{
					ID:   1,
					Name: typehelper.GetPointer("test"),
					Age:  typehelper.GetPointer(uint(20)),
				}, nil)
			},
			wantErr:    assert.NoError,
			wantStatus: http.StatusCreated,
			wantResp:   `{"success":true,"message":"customer created successfully","data":{"id":1,"name":"test","age":20}}`,
		},
		{
			name: "Cannot bind request",
			fields: fields{
				customerService: mockservice.NewCustomer(t),
				cfg:             &config.Config{},
			},
			req: httptest.NewRequest(http.MethodPost, "/customers", strings.NewReader(`{"name":"test"`)),
			rec: httptest.NewRecorder(),
			setupFunc: func(t *testing.T, tt *testCase) {

				tt.req.Header.Set("Content-Type", "application/json")
			},
			wantErr:    assert.Error,
			wantStatus: http.StatusBadRequest,
			wantResp:   `{"message":"unexpected EOF", "status_code":400, "success":false}`,
		},
		{
			name: "Cannot error validating request",
			fields: fields{
				customerService: mockservice.NewCustomer(t),
				cfg:             &config.Config{},
			},
			req: httptest.NewRequest(http.MethodPost, "/customers", strings.NewReader(`{"name":"test", "age":0}`)),
			rec: httptest.NewRecorder(),
			setupFunc: func(t *testing.T, tt *testCase) {

				tt.req.Header.Set("Content-Type", "application/json")
			},
			wantErr:    assert.Error,
			wantStatus: http.StatusBadRequest,
			wantResp:   `{"message":"Key: 'CreateCustomerRequest.Age' Error:Field validation for 'Age' failed on the 'required' tag", "status_code":400, "success":false}`,
		},
		{
			name: "Should return internal error when service return error",
			fields: fields{
				customerService: mockservice.NewCustomer(t),
				cfg:             &config.Config{},
			},
			req: httptest.NewRequest(http.MethodPost, "/customers", strings.NewReader(`{"name":"test", "age":20}`)),
			rec: httptest.NewRecorder(),
			setupFunc: func(t *testing.T, tt *testCase) {
				tt.req.Header.Set("Content-Type", "application/json")
				tt.fields.customerService.(*mockservice.Customer).EXPECT().CreateCustomer(mock.Anything, "test", uint(20)).Return(nil, fmt.Errorf("error"))
			},
			wantErr:    assert.Error,
			wantStatus: http.StatusInternalServerError,
			wantResp:   `{"message":"error creating customer: error", "status_code":500, "success":false}`,
		},
	}
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			tt.setupFunc(t, &tt)
			cu := &customerImpl{
				customerService: tt.fields.customerService,
				cfg:             tt.fields.cfg,
			}
			app := echo.New()
			app.Validator = validator.GetEchoValidator()
			c := app.NewContext(tt.req, tt.rec)
			c.SetPath("/customer/")
			err := cu.CreateCustomer(c)
			if tt.wantErr(t, err, "CreateCustomer() error = %v, wantErr %v") {
				if err == nil {
					assert.Equal(t, tt.wantStatus, tt.rec.Code, "CreateCustomer() Status got = %v, want %v", tt.rec.Code, tt.wantStatus)
					assert.JSONEq(t, tt.wantResp, tt.rec.Body.String(), "CreateCustomer() got = %v, want %v", tt.rec.Body.String(), tt.wantResp)
				} else {
					httpErr := err.(*echo.HTTPError)
					assert.Equal(t, tt.wantStatus, httpErr.Code, "CreateCustomer() Status got = %v, want %v", tt.rec.Code, tt.wantStatus)
					errResponse := httpErr.Message.(*ErrorResponse)
					jsonRes, err := json.Marshal(errResponse)
					assert.NoError(t, err)
					assert.JSONEq(t, tt.wantResp, string(jsonRes))
				}
			}
		})
	}
}

func Test_customerImpl_UpdateCustomer(t *testing.T) {
	type fields struct {
		customerService service.Customer
		cfg             *config.Config
	}
	type testCase struct {
		name       string
		fields     fields
		req        *http.Request
		rec        *httptest.ResponseRecorder
		setupFunc  func(t *testing.T, tt *testCase)
		wantStatus int
		wantResp   string
		wantErr    assert.ErrorAssertionFunc
	}

	testCases := []testCase{
		{
			name: "success",
			fields: fields{
				customerService: mockservice.NewCustomer(t),
				cfg:             &config.Config{},
			},
			req: httptest.NewRequest(http.MethodPost, "/", strings.NewReader(`{"name":"test","age":20}`)),
			rec: httptest.NewRecorder(),
			setupFunc: func(t *testing.T, tt *testCase) {
				tt.req.Header.Set("Content-Type", "application/json")
				tt.fields.customerService.(*mockservice.Customer).EXPECT().GetCustomerByID(mock.Anything, uint(1)).Return(&entity.Customer{}, nil)
				tt.fields.customerService.(*mockservice.Customer).EXPECT().
					UpdateCustomer(mock.Anything, uint(1), &entity.Customer{
						Name: typehelper.GetPointer("test"),
						Age:  typehelper.GetPointer(uint(20)),
					}).
					Return(&entity.Customer{
						ID:   1,
						Name: typehelper.GetPointer("test"),
						Age:  typehelper.GetPointer(uint(20)),
					}, nil)
			},
			wantErr:    assert.NoError,
			wantStatus: http.StatusOK,
			wantResp:   `{"success":true,"message":"customer updated successfully","data":{"id":1,"name":"test","age":20}}`,
		},
		{
			name: "Cannot bind request",
			fields: fields{
				customerService: mockservice.NewCustomer(t),
				cfg:             &config.Config{},
			},
			req: httptest.NewRequest(http.MethodPost, "/", strings.NewReader(`{"name":"test"`)),
			rec: httptest.NewRecorder(),
			setupFunc: func(t *testing.T, tt *testCase) {
				tt.req.Header.Set("Content-Type", "application/json")
			},
			wantErr:    assert.Error,
			wantStatus: http.StatusBadRequest,
			wantResp:   `{"message":"unexpected EOF", "status_code":400, "success":false}`,
		},
		{
			name: "Cannot error validating request",
			fields: fields{
				customerService: mockservice.NewCustomer(t),
				cfg:             &config.Config{},
			},
			req: httptest.NewRequest(http.MethodPost, "/customers", strings.NewReader(`{"name":"test", "age":0}`)),
			rec: httptest.NewRecorder(),
			setupFunc: func(t *testing.T, tt *testCase) {

				tt.req.Header.Set("Content-Type", "application/json")
			},
			wantErr:    assert.Error,
			wantStatus: http.StatusBadRequest,
			wantResp:   `{"message":"Key: 'UpdateCustomerRequest.Age' Error:Field validation for 'Age' failed on the 'required' tag", "status_code":400, "success":false}`,
		},
		{
			name: "Cannot find customer by id",
			fields: fields{
				customerService: mockservice.NewCustomer(t),
				cfg:             &config.Config{},
			},
			req: httptest.NewRequest(http.MethodPost, "/", strings.NewReader(`{"name":"test","age":20}`)),
			rec: httptest.NewRecorder(),
			setupFunc: func(t *testing.T, tt *testCase) {
				tt.req.Header.Set("Content-Type", "application/json")
				tt.fields.customerService.(*mockservice.Customer).EXPECT().GetCustomerByID(mock.Anything, uint(1)).Return(nil, fmt.Errorf("not found"))
			},
			wantErr:    assert.Error,
			wantStatus: http.StatusNotFound,
			wantResp:   `{"message":"customer not found: not found","status_code":404, "success":false}`,
		},
		{
			name: "Should return internal error when service return error",
			fields: fields{
				customerService: mockservice.NewCustomer(t),
				cfg:             &config.Config{},
			},
			req: httptest.NewRequest(http.MethodPost, "/customers", strings.NewReader(`{"name":"test", "age":20}`)),
			rec: httptest.NewRecorder(),
			setupFunc: func(t *testing.T, tt *testCase) {
				tt.req.Header.Set("Content-Type", "application/json")
				tt.fields.customerService.(*mockservice.Customer).EXPECT().GetCustomerByID(mock.Anything, uint(1)).Return(&entity.Customer{}, nil)
				tt.fields.customerService.(*mockservice.Customer).EXPECT().
					UpdateCustomer(mock.Anything, uint(1), &entity.Customer{
						Name: typehelper.GetPointer("test"),
						Age:  typehelper.GetPointer(uint(20)),
					}).
					Return(nil, fmt.Errorf("internal error"))
			},
			wantErr:    assert.Error,
			wantStatus: http.StatusInternalServerError,
			wantResp:   `{"message":"error updating customer: internal error", "status_code":500, "success":false}`,
		},
	}
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			tt.setupFunc(t, &tt)
			cu := &customerImpl{
				customerService: tt.fields.customerService,
				cfg:             tt.fields.cfg,
			}
			app := echo.New()
			app.Validator = validator.GetEchoValidator()
			c := app.NewContext(tt.req, tt.rec)
			c.SetPath("/customer/1")
			c.SetParamNames("id")
			c.SetParamValues("1")
			err := cu.UpdateCustomer(c)
			if tt.wantErr(t, err, "CreateCustomer() error = %v, wantErr %v") {
				if err == nil {
					assert.Equal(t, tt.wantStatus, tt.rec.Code, "CreateCustomer() Status got = %v, want %v", tt.rec.Code, tt.wantStatus)
					assert.JSONEq(t, tt.wantResp, tt.rec.Body.String(), "CreateCustomer() got = %v, want %v", tt.rec.Body.String(), tt.wantResp)
				} else {
					httpErr := err.(*echo.HTTPError)
					assert.Equal(t, tt.wantStatus, httpErr.Code, "CreateCustomer() Status got = %v, want %v", tt.rec.Code, tt.wantStatus)
					errResponse := httpErr.Message.(*ErrorResponse)
					jsonRes, err := json.Marshal(errResponse)
					assert.NoError(t, err)
					assert.JSONEq(t, tt.wantResp, string(jsonRes))
				}
			}
		})
	}
}

func Test_customerImpl_DeleteCustomer(t *testing.T) {
	type fields struct {
		customerService service.Customer
		cfg             *config.Config
	}
	type testCase struct {
		name       string
		fields     fields
		c          *echo.Context
		req        *http.Request
		rec        *httptest.ResponseRecorder
		setupFunc  func(t *testing.T, tt *testCase)
		wantStatus int
		wantResp   string
		wantErr    assert.ErrorAssertionFunc
	}

	testCases := []testCase{
		{
			name: "success",
			fields: fields{
				customerService: mockservice.NewCustomer(t),
				cfg:             &config.Config{},
			},
			req: httptest.NewRequest(http.MethodDelete, "/", nil),
			rec: httptest.NewRecorder(),
			setupFunc: func(t *testing.T, tt *testCase) {
				app := echo.New()
				app.Validator = validator.GetEchoValidator()
				c := app.NewContext(tt.req, tt.rec)
				c.SetPath("/customer/1")
				c.SetParamNames("id")
				c.SetParamValues("1")
				tt.c = &c
				tt.fields.customerService.(*mockservice.Customer).EXPECT().GetCustomerByID(mock.Anything, uint(1)).Return(&entity.Customer{}, nil)
				tt.fields.customerService.(*mockservice.Customer).EXPECT().DeleteCustomer(mock.Anything, uint(1)).Return(nil)
			},
			wantErr:    assert.NoError,
			wantStatus: http.StatusOK,
			wantResp:   `{"success":true,"message":"customer deleted successfully"}`,
		},
		{
			name: "Should return internal error when service return error",
			fields: fields{
				customerService: mockservice.NewCustomer(t),
				cfg:             &config.Config{},
			},
			req: httptest.NewRequest(http.MethodDelete, "/", nil),
			rec: httptest.NewRecorder(),
			setupFunc: func(t *testing.T, tt *testCase) {
				app := echo.New()
				app.Validator = validator.GetEchoValidator()
				c := app.NewContext(tt.req, tt.rec)
				c.SetPath("/customer/1")
				c.SetParamNames("id")
				c.SetParamValues("1")
				tt.c = &c
				tt.fields.customerService.(*mockservice.Customer).EXPECT().GetCustomerByID(mock.Anything, uint(1)).Return(&entity.Customer{}, nil)
				tt.fields.customerService.(*mockservice.Customer).EXPECT().DeleteCustomer(mock.Anything, uint(1)).Return(fmt.Errorf("internal error"))
			},
			wantErr:    assert.Error,
			wantStatus: http.StatusInternalServerError,
			wantResp:   `{"message":"error deleting customer: internal error", "status_code":500, "success":false}`,
		},
		{
			name: "Cannot find customer by id",
			fields: fields{
				customerService: mockservice.NewCustomer(t),
				cfg:             &config.Config{},
			},
			req: httptest.NewRequest(http.MethodDelete, "/", nil),
			rec: httptest.NewRecorder(),
			setupFunc: func(t *testing.T, tt *testCase) {
				app := echo.New()
				app.Validator = validator.GetEchoValidator()
				c := app.NewContext(tt.req, tt.rec)
				c.SetPath("/customer/1")
				c.SetParamNames("id")
				c.SetParamValues("1")
				tt.c = &c
				tt.fields.customerService.(*mockservice.Customer).EXPECT().GetCustomerByID(mock.Anything, uint(1)).Return(nil, fmt.Errorf("not found"))
			},
			wantErr:    assert.Error,
			wantStatus: http.StatusNotFound,
			wantResp:   `{"message":"customer not found: not found","status_code":404, "success":false}`,
		},
		{
			name: "Cannot bind request",
			fields: fields{
				customerService: mockservice.NewCustomer(t),
				cfg:             &config.Config{},
			},
			req: httptest.NewRequest(http.MethodDelete, "/", nil),
			rec: httptest.NewRecorder(),
			setupFunc: func(t *testing.T, tt *testCase) {
				app := echo.New()
				app.Validator = validator.GetEchoValidator()
				c := app.NewContext(tt.req, tt.rec)
				c.SetPath("/customer/1")
				c.SetParamNames("id")
				c.SetParamValues("asdf")
				tt.c = &c
			},
			wantErr:    assert.Error,
			wantStatus: http.StatusBadRequest,
			wantResp:   `{"message":"strconv.ParseUint: parsing \"asdf\": invalid syntax", "status_code":400, "success":false}`,
		},
	}
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			tt.setupFunc(t, &tt)
			cu := &customerImpl{
				customerService: tt.fields.customerService,
				cfg:             tt.fields.cfg,
			}
			err := cu.DeleteCustomer(*tt.c)
			if tt.wantErr(t, err, "CreateCustomer() error = %v, wantErr %v") {
				if err == nil {
					assert.Equal(t, tt.wantStatus, tt.rec.Code, "CreateCustomer() Status got = %v, want %v", tt.rec.Code, tt.wantStatus)
					assert.JSONEq(t, tt.wantResp, tt.rec.Body.String(), "CreateCustomer() got = %v, want %v", tt.rec.Body.String(), tt.wantResp)
				} else {
					httpErr := err.(*echo.HTTPError)
					assert.Equal(t, tt.wantStatus, httpErr.Code, "CreateCustomer() Status got = %v, want %v", tt.rec.Code, tt.wantStatus)
					errResponse := httpErr.Message.(*ErrorResponse)
					jsonRes, err := json.Marshal(errResponse)
					assert.NoError(t, err)
					assert.JSONEq(t, tt.wantResp, string(jsonRes))
				}
			}
		})
	}
}

func Test_customerImpl_GetCustomerByID(t *testing.T) {
	type fields struct {
		customerService service.Customer
		cfg             *config.Config
	}
	type testCase struct {
		name       string
		fields     fields
		req        *http.Request
		rec        *httptest.ResponseRecorder
		c          *echo.Context
		setupFunc  func(t *testing.T, tt *testCase)
		wantStatus int
		wantResp   string
		wantErr    assert.ErrorAssertionFunc
	}

	testCases := []testCase{
		{
			name: "success",
			fields: fields{
				customerService: mockservice.NewCustomer(t),
				cfg:             &config.Config{},
			},
			req: httptest.NewRequest(http.MethodGet, "/", nil),
			rec: httptest.NewRecorder(),
			setupFunc: func(t *testing.T, tt *testCase) {
				app := echo.New()
				app.Validator = validator.GetEchoValidator()
				c := app.NewContext(tt.req, tt.rec)
				c.SetPath("/customer/1")
				c.SetParamNames("id")
				c.SetParamValues("1")
				tt.c = &c
				tt.fields.customerService.(*mockservice.Customer).EXPECT().GetCustomerByID(mock.Anything, uint(1)).Return(&entity.Customer{
					ID:   1,
					Name: typehelper.GetPointer("test"),
					Age:  typehelper.GetPointer(uint(20)),
				}, nil)
			},
			wantErr:    assert.NoError,
			wantStatus: http.StatusOK,
			wantResp:   `{"success":true,"message":"customer found","data":{"id":1,"name":"test","age":20}}`,
		},
		{
			name: "Should return internal error when service return error",
			fields: fields{
				customerService: mockservice.NewCustomer(t),
				cfg:             &config.Config{},
			},
			req: httptest.NewRequest(http.MethodGet, "/", nil),
			rec: httptest.NewRecorder(),
			setupFunc: func(t *testing.T, tt *testCase) {
				app := echo.New()
				app.Validator = validator.GetEchoValidator()
				c := app.NewContext(tt.req, tt.rec)
				c.SetPath("/customer/1")
				c.SetParamNames("id")
				c.SetParamValues("1")
				tt.c = &c
				tt.fields.customerService.(*mockservice.Customer).EXPECT().GetCustomerByID(mock.Anything, uint(1)).Return(nil, fmt.Errorf("internal error"))
			},
			wantErr:    assert.Error,
			wantStatus: http.StatusNotFound,
			wantResp:   `{"message":"error getting customer: internal error", "status_code":404, "success":false}`,
		},
		{
			name: "Cannot find customer by id",
			fields: fields{
				customerService: mockservice.NewCustomer(t),
				cfg:             &config.Config{},
			},
			req: httptest.NewRequest(http.MethodGet, "/", nil),
			rec: httptest.NewRecorder(),
			setupFunc: func(t *testing.T, tt *testCase) {
				app := echo.New()
				app.Validator = validator.GetEchoValidator()
				c := app.NewContext(tt.req, tt.rec)
				c.SetPath("/customer/1")
				c.SetParamNames("id")
				c.SetParamValues("1")
				tt.c = &c
				tt.fields.customerService.(*mockservice.Customer).EXPECT().GetCustomerByID(mock.Anything, uint(1)).Return(nil, fmt.Errorf("not found"))
			},
			wantErr:    assert.Error,
			wantStatus: http.StatusNotFound,
			wantResp:   `{"message":"error getting customer: not found","status_code":404, "success":false}`,
		},
		{
			name: "cannot bind request",
			fields: fields{
				customerService: mockservice.NewCustomer(t),
				cfg:             &config.Config{},
			},
			req: httptest.NewRequest(http.MethodGet, "/", nil),
			rec: httptest.NewRecorder(),
			setupFunc: func(t *testing.T, tt *testCase) {
				app := echo.New()
				app.Validator = validator.GetEchoValidator()
				c := app.NewContext(tt.req, tt.rec)
				c.SetPath("/customer/1")
				c.SetParamNames("id")
				c.SetParamValues("asdf")
				tt.c = &c
			},
			wantErr:    assert.Error,
			wantStatus: http.StatusBadRequest,
			wantResp:   `{"message":"strconv.ParseUint: parsing \"asdf\": invalid syntax", "status_code":400, "success":false}`,
		},
	}
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			tt.setupFunc(t, &tt)
			cu := &customerImpl{
				customerService: tt.fields.customerService,
				cfg:             tt.fields.cfg,
			}
			err := cu.GetCustomerByID(*tt.c)
			if tt.wantErr(t, err, "CreateCustomer() error = %v, wantErr %v") {
				if err == nil {
					assert.Equal(t, tt.wantStatus, tt.rec.Code, "CreateCustomer() Status got = %v, want %v", tt.rec.Code, tt.wantStatus)
					assert.JSONEq(t, tt.wantResp, tt.rec.Body.String(), "CreateCustomer() got = %v, want %v", tt.rec.Body.String(), tt.wantResp)
				} else {
					httpErr := err.(*echo.HTTPError)
					assert.Equal(t, tt.wantStatus, httpErr.Code, "CreateCustomer() Status got = %v, want %v", tt.rec.Code, tt.wantStatus)
					errResponse := httpErr.Message.(*ErrorResponse)
					jsonRes, err := json.Marshal(errResponse)
					assert.NoError(t, err)
					assert.JSONEq(t, tt.wantResp, string(jsonRes))
				}
			}
		})
	}
}

func Test_customerImpl_GetAllCustomer(t *testing.T) {
	type fields struct {
		customerService service.Customer
		cfg             *config.Config
	}
	type testCase struct {
		name       string
		fields     fields
		req        *http.Request
		rec        *httptest.ResponseRecorder
		setupFunc  func(t *testing.T, tt *testCase)
		wantStatus int
		wantResp   string
		wantErr    assert.ErrorAssertionFunc
	}

	testCases := []testCase{
		{
			name: "success",
			fields: fields{
				customerService: mockservice.NewCustomer(t),
				cfg:             &config.Config{},
			},
			req: httptest.NewRequest(http.MethodGet, "/", nil),
			rec: httptest.NewRecorder(),
			setupFunc: func(t *testing.T, tt *testCase) {
				tt.fields.customerService.(*mockservice.Customer).EXPECT().GetAllCustomer(mock.Anything).Return([]*entity.Customer{
					{
						ID:   1,
						Name: typehelper.GetPointer("test"),
						Age:  typehelper.GetPointer(uint(20)),
					},
				}, nil)
			},
			wantErr:    assert.NoError,
			wantStatus: http.StatusOK,
			wantResp:   `{"success":true,"data":[{"id":1,"name":"test","age":20}],"message":"customers found"}`,
		},
		{
			name: "Should return internal error when service return error",
			fields: fields{
				customerService: mockservice.NewCustomer(t),
				cfg:             &config.Config{},
			},
			req: httptest.NewRequest(http.MethodGet, "/", nil),
			rec: httptest.NewRecorder(),
			setupFunc: func(t *testing.T, tt *testCase) {
				tt.fields.customerService.(*mockservice.Customer).EXPECT().GetAllCustomer(mock.Anything).Return(nil, fmt.Errorf("internal error"))
			},
			wantErr:    assert.Error,
			wantStatus: http.StatusInternalServerError,
			wantResp:   `{"message":"error getting all customers: internal error", "status_code":500, "success":false}`,
		},
	}
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			tt.setupFunc(t, &tt)
			cu := &customerImpl{
				customerService: tt.fields.customerService,
				cfg:             tt.fields.cfg,
			}
			app := echo.New()
			app.Validator = validator.GetEchoValidator()
			c := app.NewContext(tt.req, tt.rec)
			err := cu.GetAllCustomer(c)
			if tt.wantErr(t, err, "CreateCustomer() error = %v, wantErr %v") {
				if err == nil {
					assert.Equal(t, tt.wantStatus, tt.rec.Code, "CreateCustomer() Status got = %v, want %v", tt.rec.Code, tt.wantStatus)
					assert.JSONEq(t, tt.wantResp, tt.rec.Body.String(), "CreateCustomer() got = %v, want %v", tt.rec.Body.String(), tt.wantResp)
				} else {
					httpErr := err.(*echo.HTTPError)
					assert.Equal(t, tt.wantStatus, httpErr.Code, "CreateCustomer() Status got = %v, want %v", tt.rec.Code, tt.wantStatus)
					errResponse := httpErr.Message.(*ErrorResponse)
					jsonRes, err := json.Marshal(errResponse)
					assert.NoError(t, err)
					assert.JSONEq(t, tt.wantResp, string(jsonRes))
				}
			}
		})
	}
}

func TestNewCustomer(t *testing.T) {
	cfg := &config.Config{}
	customerService := mockservice.NewCustomer(t)

	want := &customerImpl{
		customerService: customerService,
		cfg:             cfg,
	}

	got := NewCustomer(cfg, customerService)
	assert.Equalf(t, want, got, "NewCustomer() = %v, want %v", got, want)
}
