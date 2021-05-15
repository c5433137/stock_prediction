package logic

import (
	"reflect"
	"testing"
)

func TestGetStockData(t *testing.T) {
	type args struct {
		stockId string
		start   string
		end     string
	}
	tests := []struct {
		name string
		args args
		want []CodeList
	}{
		// TODO: Add test cases.
		{
			args:args{
				stockId:"600004",//"002156"
				start:"2021-04-30",
				end:"2021-05-14",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetStockData(tt.args.stockId, tt.args.start, tt.args.end); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetStockData() = %v, want %v", got, tt.want)
			}
		})
	}
}
