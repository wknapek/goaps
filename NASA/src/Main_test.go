package main

import (
	"testing"
	"time"
)

func TestDate_correct(t *testing.T) {
	startDate := time.Now()
	startDate.AddDate(0, -1, 0)
	endDate := time.Now()
	if checkDate(startDate, endDate) {
		t.Fail()
	}
}

func TestDate_false(t *testing.T) {
	startDate := time.Now()
	endDate := time.Now()
	endDate.AddDate(-1, 0, 0)
	if checkDate(startDate, endDate) {
		t.Fail()
	}
}

func TestDate_future(t *testing.T) {
	startDate := time.Now()
	endDate := time.Now()
	endDate.AddDate(+1, 0, 0)
	if checkDate(startDate, endDate) {
		t.Fail()
	}
}

func TestIsEndDate(t *testing.T) {
	startDate := time.Now()
	endDate := time.Now()
	if isEndDate(startDate, endDate) {
		t.Fail()
	}
}

func TestIsEndDate_neg(t *testing.T) {
	startDate := time.Now()
	endDate := time.Now()
	endDate = endDate.AddDate(0, 1, 0)
	if !isEndDate(startDate, endDate) {
		t.Fail()
	}
}
