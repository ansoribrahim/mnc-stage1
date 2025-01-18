package main

import (
	"fmt"
	"math"
	"time"
)

func CheckLeave(cutiBersamaTotal int, joinDate time.Time, leaveDate time.Time, leaveDuration int) (bool, string) {

	startLeaveFromDate := joinDate.AddDate(0, 0, 180)

	daysInYear := int(time.Date(joinDate.Year(), 12, 31, 0, 0, 0, 0, joinDate.Location()).Sub(startLeaveFromDate).Hours()/24) + 1
	yearProportion := float64(daysInYear) / 365
	maxLeave := int(math.Floor(yearProportion * float64(cutiBersamaTotal)))

	switch {
	case leaveDate.Before(startLeaveFromDate):
		return false, "Karena belum 180 hari sejak tanggal join karyawan"
	case leaveDuration > maxLeave:
		return false, fmt.Sprintf("Karena hanya boleh mengambil %d hari cuti", maxLeave)
	case leaveDuration > 3:
		return false, "Durasi cuti pribadi maksimal 3 hari berturut-turut."
	}

	return true, ""
}
