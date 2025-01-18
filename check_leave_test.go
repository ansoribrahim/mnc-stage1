package main

import (
	"testing"
	"time"
)

func TestCheckLeave(t *testing.T) {
	tests := []struct {
		name             string
		cutiBersamaTotal int
		joinDate         time.Time
		leaveDate        time.Time
		leaveDuration    int
		expectedResult   bool
		expectedReason   string
	}{
		{
			name:             "Cuti tidak boleh diambil karena belum 180 hari",
			cutiBersamaTotal: 7,
			joinDate:         time.Date(2021, 5, 1, 0, 0, 0, 0, time.UTC),
			leaveDate:        time.Date(2021, 7, 5, 0, 0, 0, 0, time.UTC),
			leaveDuration:    1,
			expectedResult:   false,
			expectedReason:   "Karena belum 180 hari sejak tanggal join karyawan",
		},
		{
			name:             "Cuti lebih dari 3 hari berturut-turut",
			cutiBersamaTotal: 7,
			joinDate:         time.Date(2021, 5, 1, 0, 0, 0, 0, time.UTC),
			leaveDate:        time.Date(2021, 11, 5, 0, 0, 0, 0, time.UTC),
			leaveDuration:    3,
			expectedResult:   false,
			expectedReason:   "Karena hanya boleh mengambil 1 hari cuti",
		},
		{
			name:             "Cuti bisa diambil di tahun pertama",
			cutiBersamaTotal: 7,
			joinDate:         time.Date(2021, 1, 5, 0, 0, 0, 0, time.UTC),
			leaveDate:        time.Date(2021, 12, 18, 0, 0, 0, 0, time.UTC),
			leaveDuration:    1,
			expectedResult:   true,
			expectedReason:   "",
		},
		{
			name:             "Cuti bisa diambil di tahun pertama dengan durasi 3 hari",
			cutiBersamaTotal: 7,
			joinDate:         time.Date(2021, 1, 5, 0, 0, 0, 0, time.UTC),
			leaveDate:        time.Date(2021, 12, 18, 0, 0, 0, 0, time.UTC),
			leaveDuration:    3,
			expectedResult:   true,
			expectedReason:   "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, alasan := CheckLeave(tt.cutiBersamaTotal, tt.joinDate, tt.leaveDate, tt.leaveDuration)
			if result != tt.expectedResult || alasan != tt.expectedReason {
				t.Errorf("expected result: %v, alasan: %s, but got result: %v, alasan: %s", tt.expectedResult, tt.expectedReason, result, alasan)
			}
		})
	}
}
