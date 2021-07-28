package client

import (
	"context"
	"net/http"
	"reflect"
	"testing"

	"github.com/stretchr/testify/mock"
)

type MockGetter struct {
	mock.Mock
}

func (mg *MockGetter) Get(url string) (*http.Response, error) {
	ret := mg.Called(url)

	if ret.Get(0) != nil {
		return ret.Get(0).(*http.Response), ret.Error(1)
	}

	return nil, ret.Error(1)
}

func TestClient_Do(t *testing.T) {
	type fields struct {
		HTTP func() *MockGetter
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *http.Response
		wantErr bool
	}{
		{
			name:    "Calls Get",
			wantErr: false,
			want: &http.Response{
				StatusCode: http.StatusOK,
			},
			fields: fields{
				HTTP: func() *MockGetter {
					mg := new(MockGetter)
					mg.On("Get", mock.Anything).Return(&http.Response{
						StatusCode: http.StatusOK,
					}, nil)
					return mg
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mg := tt.fields.HTTP()
			c := &Client{
				HTTP: mg,
			}
			got, err := c.Do(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.Do() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Client.Do() = %v, want %v", got, tt.want)
			}
			mg.AssertExpectations(t)
		})
	}
}
