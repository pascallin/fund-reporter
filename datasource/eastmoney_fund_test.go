package datasource

import (
	"testing"
)

func TestGetFundsData(t *testing.T) {
	t.Run("GetFundsData while wrong code", func(t *testing.T) {
		codes := []string{"xxx"}
		result := GetFundsData(codes)
		if result[0] != nil {
			t.Errorf("expect receve nil")
		}
	})
	t.Run("GetFundsData succeed", func(t *testing.T) {
		codes := []string{"161725"}
		result := GetFundsData(codes)
		if result[0] == nil {
			t.Errorf("expect receve fund data")
		}
		if result[0].Code != "161725" {
			t.Errorf("expect receve same code")
		}
	})
}

func TestGetFundsDataWithQueue(t *testing.T) {
	t.Run("GetFundsDataWithQueue while wrong code", func(t *testing.T) {
		codes := []string{"xxx"}
		result := GetFundsDataWithQueue(codes, 1)
		if result[0] != nil {
			t.Errorf("expect receve nil")
		}
	})
	t.Run("GetFundsDataWithQueue succeed", func(t *testing.T) {
		codes := []string{"161725"}
		result := GetFundsData(codes)
		if result[0] == nil {
			t.Errorf("expect receve fund data")
		}
		if result[0].Code != "161725" {
			t.Errorf("expect receve same code")
		}
	})
}
