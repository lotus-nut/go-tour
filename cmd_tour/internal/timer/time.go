package timer

import "time"

func GetNowTime() time.Time {
	loc, _ := time.LoadLocation("Asia/Shanghai")
	return time.Now().In(loc)
}

func GetCalculateTime(currentTimer time.Time, d string) (time.Time, error) {
	duration, err := time.ParseDuration(d)
	if err != nil {
		return time.Time{}, err
	}
	return currentTimer.Add(duration), nil
}
