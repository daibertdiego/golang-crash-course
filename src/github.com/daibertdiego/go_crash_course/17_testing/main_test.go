package main

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestAdd(t *testing.T) {
	type args struct {
		num1 float64
		num2 float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{"5 + 6", args{5, 6}, 11},
		{"1 + 1", args{1, 1}, 2},
		{"5.5 + 5.5", args{5.5, 5.5}, 11},
		{"-5 + 5", args{-5.0, 5}, 0},
		{"10 + 5 (Invalid)", args{10, 5}, 15},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, error := Add(tt.args.num1, tt.args.num2); got != tt.want && error == nil {
				t.Errorf("Add() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAddTestify(t *testing.T) {
	type args struct {
		num1 float64
		num2 float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{"5 + 6", args{5, 6}, 11},
		{"1 + 1", args{1, 1}, 2},
		{"5.5 + 5.5", args{5.5, 5.5}, 11},
		{"-5 + 5", args{-5.0, 5}, 0},
		{"10 + 6 (Invalid)", args{10, 6}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := Add(tt.args.num1, tt.args.num2)
			if err != nil {
				assert.EqualError(t, err, "result number too large for this function")
			}
			assert.Equal(t, tt.want, result)
		})
	}
}

func TestMockingLayerSuccess(t *testing.T) {
	mockLibraryDB := new(MockLibraryDB)

	// Setup expectations
	mockLibraryDB.On("GetBook").Return(DBBook{ID: "GALO"})
	mockLibraryDB.On("CheckOutBook").Return(nil)

	libraryService := NewLibraryService(mockLibraryDB)

	error := libraryService.CheckOutBook("GALO", "1313")

	mockLibraryDB.AssertExpectations(t)

	assert.Nil(t, error)
}

func TestMockingLayerFailure(t *testing.T) {
	mockLibraryDB := new(MockLibraryDB)

	// Setup expectations
	userId := "1313"
	mockLibraryDB.On("GetBook").Return(DBBook{ID: "GALO", CheckOutByUserID: &userId})

	libraryService := NewLibraryService(mockLibraryDB)

	error := libraryService.CheckOutBook("GALO", "1313")

	mockLibraryDB.AssertExpectations(t)

	assert.EqualError(t, error, "book with id GALO is currently checked out")
}

type MockLibraryDB struct {
	mock.Mock
}

func (mock *MockLibraryDB) GetBook(bookID string) DBBook {
	args := mock.Called()
	result := args.Get(0)
	return result.(DBBook)
}

func (mock *MockLibraryDB) CheckOutBook(bookID string, userID string) error {
	args := mock.Called()
	return args.Error(0)
}
