package main

import (
    "testing"
    "time"
)

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
    PrepareHoliday()

    isTrue(time.Date(2016,  1,  1, 12, 0, 0, 0, time.UTC), t)
    isTrue(time.Date(2016,  2, 11, 12, 0, 0, 0, time.UTC), t)
    isTrue(time.Date(2016, 11,  3, 12, 0, 0, 0, time.UTC), t)
    isTrue(time.Date(2016, 11, 23, 12, 0, 0, 0, time.UTC), t)
    isTrue(time.Date(2016, 12, 23, 12, 0, 0, 0, time.UTC), t)

    isFalse(time.Date(2016,  1, 2, 12, 0, 0, 0, time.UTC), t)
    isFalse(time.Date(2016, 11, 4, 12, 0, 0, 0, time.UTC), t)
}
