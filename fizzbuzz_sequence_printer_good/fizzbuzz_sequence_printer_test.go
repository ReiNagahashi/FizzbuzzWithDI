package fizzbuzzsequenceprintergood

import (
	// "testing"

	"github.com/stretchr/testify/mock"
	// "go.uber.org/mock/gomock"
)




type MockOutputObject struct{
	mock.Mock
}

func (o *MockOutputObject) Write(data string){
	o.Called(data)
} 

// func TestPrintRange(t *testing.T){
// 	tests := []struct {
// 		name 		string
// 		start, end 	int
// 		expectCalls int
// 	}{
// 		{"Normal Case", 1, 5, 5},
// 		{"No Op Case (Start > End)", 0, -1, 0},
// 		{"Single Item", 1, 1, 1},
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T){
// 			ctrl := gomock.NewController(t)
// 			defer ctrl.Finish()

// 			mockConverter := NewMockNumberConverter(ctrl)
// 			mockOutput := NewMockOutputInterface(ctrl)

// 			// テーブル定義に基づいて期待値を設定
// 			mockConverter.EXPECT().Convert(gomock.Any()).Times(tt.expectCalls)
// 			mockOutput.EXPECT().write(gomock.Any()).Times(tt.expectCalls)

// 			p := NewFizzBuzzSequencePrinter(mockConverter, mockOutput)
// 			p.PrintRange(tt.start, tt.end)
// 		})
// 	}
// }