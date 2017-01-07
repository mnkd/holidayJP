package holidayJP

import (
    "testing"
    "time"
)

var JST = time.FixedZone("JST", 3600*9)

func isTrue(ti time.Time, t *testing.T) {
    if IsHoliday(ti) != true {
        t.Error(ti, "is Holiday")
    }
}

func isFalse(ti time.Time, t *testing.T) {
    if IsHoliday(ti) != false {
        t.Error(ti, "is not Holiday")
    }
}

func TestHolidays(t *testing.T) {
    isTrue(time.Date(2016,  1,  1, 12, 0, 0, 0, time.UTC).In(JST), t)
    isTrue(time.Date(2016,  2, 11, 12, 0, 0, 0, time.UTC).In(JST), t)
    isTrue(time.Date(2016, 11,  3, 12, 0, 0, 0, time.UTC).In(JST), t)
    isTrue(time.Date(2016, 11, 23, 12, 0, 0, 0, time.UTC).In(JST), t)
    isTrue(time.Date(2016, 12, 23, 12, 0, 0, 0, time.UTC).In(JST), t)

    isFalse(time.Date(2016,  1, 2, 12, 0, 0, 0, time.UTC).In(JST), t)
    isFalse(time.Date(2016, 11, 4, 12, 0, 0, 0, time.UTC).In(JST), t)
}
