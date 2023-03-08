package lib

import (
	"time"

	"github.com/sin3degrees/tengo/v2"
)

var timesModule = map[string]interface{}{
	"ANSIC":               time.ANSIC,
	"UnixDate":            time.UnixDate,
	"RubyDate":            time.RubyDate,
	"RFC822":              time.RFC822,
	"RFC822Z":             time.RFC822Z,
	"RFC850":              time.RFC850,
	"RFC1123":             time.RFC1123,
	"RFC1123Z":            time.RFC1123Z,
	"RFC3339":             time.RFC3339,
	"RFC3339Nano":         time.RFC3339Nano,
	"Kitchen":             time.Kitchen,
	"Stamp":               time.Stamp,
	"StampMilli":          time.StampMilli,
	"StampMicro":          time.StampMicro,
	"StampNano":           time.StampNano,
	"Nanosecond":          int64(time.Nanosecond),
	"Microsecond":         int64(time.Microsecond),
	"Millisecond":         int64(time.Millisecond),
	"Second":              int64(time.Second),
	"Minute":              int64(time.Minute),
	"Hour":                int64(time.Hour),
	"January":             int64(time.January),
	"February":            int64(time.February),
	"March":               int64(time.March),
	"April":               int64(time.April),
	"May":                 int64(time.May),
	"June":                int64(time.June),
	"July":                int64(time.July),
	"August":              int64(time.August),
	"September":           int64(time.September),
	"October":             int64(time.October),
	"November":            int64(time.November),
	"December":            int64(time.December),
	"Sleep":               timesSleep,
	"ParseDuration":       timesParseDuration,
	"Since":               timesSince,
	"Until":               timesUntil,
	"DurationHours":       timesDurationHours,
	"DurationMinutes":     timesDurationMinutes,
	"DurationNanoseconds": timesDurationNanoseconds,
	"DurationSeconds":     timesDurationSeconds,
	"DurationString":      timesDurationString,
	"MonthString":         timesMonthString,
	"Date":                timesDate,
	"Now":                 timesNow,
	"Parse":               timesParse,
	"Unix":                timesUnix,
	"Add":                 timesAdd,
	"AddDate":             timesAddDate,
	"Sub":                 timesSub,
	"After":               timesAfter,
	"Before":              timesBefore,
	"TimeYear":            timesTimeYear,
	"TimeMonth":           timesTimeMonth,
	"TimeDay":             timesTimeDay,
	"TimeWeekday":         timesTimeWeekday,
	"TimeHour":            timesTimeHour,
	"TimeMinute":          timesTimeMinute,
	"TimeSecond":          timesTimeSecond,
	"TimeNanosecond":      timesTimeNanosecond,
	"TimeUnix":            timesTimeUnix,
	"TimeUnixNano":        timesTimeUnixNano,
	"TimeFormat":          timesTimeFormat,
	"TimeLocation":        timesTimeLocation,
	"TimeString":          timesTimeString,
	"IsZero":              timesIsZero,
	"ToLocal":             timesToLocal,
	"ToUTC":               timesToUTC,
	"InLocation":          timesInLocation,
}

func timesSleep(args ...tengo.Object) (ret tengo.Object, err error) {
	if len(args) != 1 {
		err = tengo.ErrWrongNumArguments
		return
	}

	i1, ok := tengo.ToInt64(args[0])
	if !ok {
		err = tengo.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "int(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	time.Sleep(time.Duration(i1))
	ret = tengo.UndefinedValue

	return
}

func timesParseDuration(args ...tengo.Object) (
	ret tengo.Object,
	err error,
) {
	if len(args) != 1 {
		err = tengo.ErrWrongNumArguments
		return
	}

	s1, ok := tengo.ToString(args[0])
	if !ok {
		err = tengo.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "string(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	dur, err := time.ParseDuration(s1)
	if err != nil {
		ret = wrapError(err)
		return
	}

	ret = &tengo.Int{Value: int64(dur)}

	return
}

func timesSince(args ...tengo.Object) (
	ret tengo.Object,
	err error,
) {
	if len(args) != 1 {
		err = tengo.ErrWrongNumArguments
		return
	}

	t1, ok := tengo.ToTime(args[0])
	if !ok {
		err = tengo.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "time(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	ret = &tengo.Int{Value: int64(time.Since(t1))}

	return
}

func timesUntil(args ...tengo.Object) (
	ret tengo.Object,
	err error,
) {
	if len(args) != 1 {
		err = tengo.ErrWrongNumArguments
		return
	}

	t1, ok := tengo.ToTime(args[0])
	if !ok {
		err = tengo.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "time(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	ret = &tengo.Int{Value: int64(time.Until(t1))}

	return
}

func timesDurationHours(args ...tengo.Object) (
	ret tengo.Object,
	err error,
) {
	if len(args) != 1 {
		err = tengo.ErrWrongNumArguments
		return
	}

	i1, ok := tengo.ToInt64(args[0])
	if !ok {
		err = tengo.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "int(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	ret = &tengo.Float{Value: time.Duration(i1).Hours()}

	return
}

func timesDurationMinutes(args ...tengo.Object) (
	ret tengo.Object,
	err error,
) {
	if len(args) != 1 {
		err = tengo.ErrWrongNumArguments
		return
	}

	i1, ok := tengo.ToInt64(args[0])
	if !ok {
		err = tengo.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "int(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	ret = &tengo.Float{Value: time.Duration(i1).Minutes()}

	return
}

func timesDurationNanoseconds(args ...tengo.Object) (
	ret tengo.Object,
	err error,
) {
	if len(args) != 1 {
		err = tengo.ErrWrongNumArguments
		return
	}

	i1, ok := tengo.ToInt64(args[0])
	if !ok {
		err = tengo.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "int(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	ret = &tengo.Int{Value: time.Duration(i1).Nanoseconds()}

	return
}

func timesDurationSeconds(args ...tengo.Object) (
	ret tengo.Object,
	err error,
) {
	if len(args) != 1 {
		err = tengo.ErrWrongNumArguments
		return
	}

	i1, ok := tengo.ToInt64(args[0])
	if !ok {
		err = tengo.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "int(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	ret = &tengo.Float{Value: time.Duration(i1).Seconds()}

	return
}

func timesDurationString(args ...tengo.Object) (
	ret tengo.Object,
	err error,
) {
	if len(args) != 1 {
		err = tengo.ErrWrongNumArguments
		return
	}

	i1, ok := tengo.ToInt64(args[0])
	if !ok {
		err = tengo.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "int(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	ret = &tengo.String{Value: time.Duration(i1).String()}

	return
}

func timesMonthString(args ...tengo.Object) (
	ret tengo.Object,
	err error,
) {
	if len(args) != 1 {
		err = tengo.ErrWrongNumArguments
		return
	}

	i1, ok := tengo.ToInt64(args[0])
	if !ok {
		err = tengo.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "int(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	ret = &tengo.String{Value: time.Month(i1).String()}

	return
}

func timesDate(args ...tengo.Object) (
	ret tengo.Object,
	err error,
) {
	if len(args) < 7 || len(args) > 8 {
		err = tengo.ErrWrongNumArguments
		return
	}

	i1, ok := tengo.ToInt(args[0])
	if !ok {
		err = tengo.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "int(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}
	i2, ok := tengo.ToInt(args[1])
	if !ok {
		err = tengo.ErrInvalidArgumentType{
			Name:     "second",
			Expected: "int(compatible)",
			Found:    args[1].TypeName(),
		}
		return
	}
	i3, ok := tengo.ToInt(args[2])
	if !ok {
		err = tengo.ErrInvalidArgumentType{
			Name:     "third",
			Expected: "int(compatible)",
			Found:    args[2].TypeName(),
		}
		return
	}
	i4, ok := tengo.ToInt(args[3])
	if !ok {
		err = tengo.ErrInvalidArgumentType{
			Name:     "fourth",
			Expected: "int(compatible)",
			Found:    args[3].TypeName(),
		}
		return
	}
	i5, ok := tengo.ToInt(args[4])
	if !ok {
		err = tengo.ErrInvalidArgumentType{
			Name:     "fifth",
			Expected: "int(compatible)",
			Found:    args[4].TypeName(),
		}
		return
	}
	i6, ok := tengo.ToInt(args[5])
	if !ok {
		err = tengo.ErrInvalidArgumentType{
			Name:     "sixth",
			Expected: "int(compatible)",
			Found:    args[5].TypeName(),
		}
		return
	}
	i7, ok := tengo.ToInt(args[6])
	if !ok {
		err = tengo.ErrInvalidArgumentType{
			Name:     "seventh",
			Expected: "int(compatible)",
			Found:    args[6].TypeName(),
		}
		return
	}

	var loc *time.Location
	if len(args) == 8 {
		i8, ok := tengo.ToString(args[7])
		if !ok {
			err = tengo.ErrInvalidArgumentType{
				Name:     "eighth",
				Expected: "string(compatible)",
				Found:    args[7].TypeName(),
			}
			return
		}
		loc, err = time.LoadLocation(i8)
		if err != nil {
			ret = wrapError(err)
			return
		}
	} else {
		loc = time.Now().Location()
	}

	ret = &tengo.Time{
		Value: time.Date(i1,
			time.Month(i2), i3, i4, i5, i6, i7, loc),
	}

	return
}

func timesNow(args ...tengo.Object) (ret tengo.Object, err error) {
	if len(args) != 0 {
		err = tengo.ErrWrongNumArguments
		return
	}

	ret = &tengo.Time{Value: time.Now()}

	return
}

func timesParse(args ...tengo.Object) (ret tengo.Object, err error) {
	if len(args) != 2 {
		err = tengo.ErrWrongNumArguments
		return
	}

	s1, ok := tengo.ToString(args[0])
	if !ok {
		err = tengo.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "string(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	s2, ok := tengo.ToString(args[1])
	if !ok {
		err = tengo.ErrInvalidArgumentType{
			Name:     "second",
			Expected: "string(compatible)",
			Found:    args[1].TypeName(),
		}
		return
	}

	parsed, err := time.Parse(s1, s2)
	if err != nil {
		ret = wrapError(err)
		return
	}

	ret = &tengo.Time{Value: parsed}

	return
}

func timesUnix(args ...tengo.Object) (ret tengo.Object, err error) {
	if len(args) != 2 {
		err = tengo.ErrWrongNumArguments
		return
	}

	i1, ok := tengo.ToInt64(args[0])
	if !ok {
		err = tengo.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "int(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	i2, ok := tengo.ToInt64(args[1])
	if !ok {
		err = tengo.ErrInvalidArgumentType{
			Name:     "second",
			Expected: "int(compatible)",
			Found:    args[1].TypeName(),
		}
		return
	}

	ret = &tengo.Time{Value: time.Unix(i1, i2)}

	return
}

func timesAdd(args ...tengo.Object) (ret tengo.Object, err error) {
	if len(args) != 2 {
		err = tengo.ErrWrongNumArguments
		return
	}

	t1, ok := tengo.ToTime(args[0])
	if !ok {
		err = tengo.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "time(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	i2, ok := tengo.ToInt64(args[1])
	if !ok {
		err = tengo.ErrInvalidArgumentType{
			Name:     "second",
			Expected: "int(compatible)",
			Found:    args[1].TypeName(),
		}
		return
	}

	ret = &tengo.Time{Value: t1.Add(time.Duration(i2))}

	return
}

func timesSub(args ...tengo.Object) (ret tengo.Object, err error) {
	if len(args) != 2 {
		err = tengo.ErrWrongNumArguments
		return
	}

	t1, ok := tengo.ToTime(args[0])
	if !ok {
		err = tengo.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "time(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	t2, ok := tengo.ToTime(args[1])
	if !ok {
		err = tengo.ErrInvalidArgumentType{
			Name:     "second",
			Expected: "time(compatible)",
			Found:    args[1].TypeName(),
		}
		return
	}

	ret = &tengo.Int{Value: int64(t1.Sub(t2))}

	return
}

func timesAddDate(args ...tengo.Object) (ret tengo.Object, err error) {
	if len(args) != 4 {
		err = tengo.ErrWrongNumArguments
		return
	}

	t1, ok := tengo.ToTime(args[0])
	if !ok {
		err = tengo.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "time(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	i2, ok := tengo.ToInt(args[1])
	if !ok {
		err = tengo.ErrInvalidArgumentType{
			Name:     "second",
			Expected: "int(compatible)",
			Found:    args[1].TypeName(),
		}
		return
	}

	i3, ok := tengo.ToInt(args[2])
	if !ok {
		err = tengo.ErrInvalidArgumentType{
			Name:     "third",
			Expected: "int(compatible)",
			Found:    args[2].TypeName(),
		}
		return
	}

	i4, ok := tengo.ToInt(args[3])
	if !ok {
		err = tengo.ErrInvalidArgumentType{
			Name:     "fourth",
			Expected: "int(compatible)",
			Found:    args[3].TypeName(),
		}
		return
	}

	ret = &tengo.Time{Value: t1.AddDate(i2, i3, i4)}

	return
}

func timesAfter(args ...tengo.Object) (ret tengo.Object, err error) {
	if len(args) != 2 {
		err = tengo.ErrWrongNumArguments
		return
	}

	t1, ok := tengo.ToTime(args[0])
	if !ok {
		err = tengo.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "time(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	t2, ok := tengo.ToTime(args[1])
	if !ok {
		err = tengo.ErrInvalidArgumentType{
			Name:     "second",
			Expected: "time(compatible)",
			Found:    args[1].TypeName(),
		}
		return
	}

	if t1.After(t2) {
		ret = tengo.TrueValue
	} else {
		ret = tengo.FalseValue
	}

	return
}

func timesBefore(args ...tengo.Object) (ret tengo.Object, err error) {
	if len(args) != 2 {
		err = tengo.ErrWrongNumArguments
		return
	}

	t1, ok := tengo.ToTime(args[0])
	if !ok {
		err = tengo.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "time(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	t2, ok := tengo.ToTime(args[1])
	if !ok {
		err = tengo.ErrInvalidArgumentType{
			Name:     "second",
			Expected: "time(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	if t1.Before(t2) {
		ret = tengo.TrueValue
	} else {
		ret = tengo.FalseValue
	}

	return
}

func timesTimeYear(args ...tengo.Object) (ret tengo.Object, err error) {
	if len(args) != 1 {
		err = tengo.ErrWrongNumArguments
		return
	}

	t1, ok := tengo.ToTime(args[0])
	if !ok {
		err = tengo.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "time(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	ret = &tengo.Int{Value: int64(t1.Year())}

	return
}

func timesTimeMonth(args ...tengo.Object) (ret tengo.Object, err error) {
	if len(args) != 1 {
		err = tengo.ErrWrongNumArguments
		return
	}

	t1, ok := tengo.ToTime(args[0])
	if !ok {
		err = tengo.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "time(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	ret = &tengo.Int{Value: int64(t1.Month())}

	return
}

func timesTimeDay(args ...tengo.Object) (ret tengo.Object, err error) {
	if len(args) != 1 {
		err = tengo.ErrWrongNumArguments
		return
	}

	t1, ok := tengo.ToTime(args[0])
	if !ok {
		err = tengo.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "time(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	ret = &tengo.Int{Value: int64(t1.Day())}

	return
}

func timesTimeWeekday(args ...tengo.Object) (ret tengo.Object, err error) {
	if len(args) != 1 {
		err = tengo.ErrWrongNumArguments
		return
	}

	t1, ok := tengo.ToTime(args[0])
	if !ok {
		err = tengo.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "time(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	ret = &tengo.Int{Value: int64(t1.Weekday())}

	return
}

func timesTimeHour(args ...tengo.Object) (ret tengo.Object, err error) {
	if len(args) != 1 {
		err = tengo.ErrWrongNumArguments
		return
	}

	t1, ok := tengo.ToTime(args[0])
	if !ok {
		err = tengo.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "time(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	ret = &tengo.Int{Value: int64(t1.Hour())}

	return
}

func timesTimeMinute(args ...tengo.Object) (ret tengo.Object, err error) {
	if len(args) != 1 {
		err = tengo.ErrWrongNumArguments
		return
	}

	t1, ok := tengo.ToTime(args[0])
	if !ok {
		err = tengo.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "time(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	ret = &tengo.Int{Value: int64(t1.Minute())}

	return
}

func timesTimeSecond(args ...tengo.Object) (ret tengo.Object, err error) {
	if len(args) != 1 {
		err = tengo.ErrWrongNumArguments
		return
	}

	t1, ok := tengo.ToTime(args[0])
	if !ok {
		err = tengo.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "time(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	ret = &tengo.Int{Value: int64(t1.Second())}

	return
}

func timesTimeNanosecond(args ...tengo.Object) (
	ret tengo.Object,
	err error,
) {
	if len(args) != 1 {
		err = tengo.ErrWrongNumArguments
		return
	}

	t1, ok := tengo.ToTime(args[0])
	if !ok {
		err = tengo.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "time(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	ret = &tengo.Int{Value: int64(t1.Nanosecond())}

	return
}

func timesTimeUnix(args ...tengo.Object) (ret tengo.Object, err error) {
	if len(args) != 1 {
		err = tengo.ErrWrongNumArguments
		return
	}

	t1, ok := tengo.ToTime(args[0])
	if !ok {
		err = tengo.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "time(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	ret = &tengo.Int{Value: t1.Unix()}

	return
}

func timesTimeUnixNano(args ...tengo.Object) (
	ret tengo.Object,
	err error,
) {
	if len(args) != 1 {
		err = tengo.ErrWrongNumArguments
		return
	}

	t1, ok := tengo.ToTime(args[0])
	if !ok {
		err = tengo.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "time(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	ret = &tengo.Int{Value: t1.UnixNano()}

	return
}

func timesTimeFormat(args ...tengo.Object) (ret tengo.Object, err error) {
	if len(args) != 2 {
		err = tengo.ErrWrongNumArguments
		return
	}

	t1, ok := tengo.ToTime(args[0])
	if !ok {
		err = tengo.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "time(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	s2, ok := tengo.ToString(args[1])
	if !ok {
		err = tengo.ErrInvalidArgumentType{
			Name:     "second",
			Expected: "string(compatible)",
			Found:    args[1].TypeName(),
		}
		return
	}

	s := t1.Format(s2)
	if len(s) > tengo.MaxStringLen {

		return nil, tengo.ErrStringLimit
	}

	ret = &tengo.String{Value: s}

	return
}

func timesIsZero(args ...tengo.Object) (ret tengo.Object, err error) {
	if len(args) != 1 {
		err = tengo.ErrWrongNumArguments
		return
	}

	t1, ok := tengo.ToTime(args[0])
	if !ok {
		err = tengo.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "time(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	if t1.IsZero() {
		ret = tengo.TrueValue
	} else {
		ret = tengo.FalseValue
	}

	return
}

func timesToLocal(args ...tengo.Object) (ret tengo.Object, err error) {
	if len(args) != 1 {
		err = tengo.ErrWrongNumArguments
		return
	}

	t1, ok := tengo.ToTime(args[0])
	if !ok {
		err = tengo.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "time(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	ret = &tengo.Time{Value: t1.Local()}

	return
}

func timesToUTC(args ...tengo.Object) (ret tengo.Object, err error) {
	if len(args) != 1 {
		err = tengo.ErrWrongNumArguments
		return
	}

	t1, ok := tengo.ToTime(args[0])
	if !ok {
		err = tengo.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "time(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	ret = &tengo.Time{Value: t1.UTC()}

	return
}

func timesTimeLocation(args ...tengo.Object) (
	ret tengo.Object,
	err error,
) {
	if len(args) != 1 {
		err = tengo.ErrWrongNumArguments
		return
	}

	t1, ok := tengo.ToTime(args[0])
	if !ok {
		err = tengo.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "time(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	ret = &tengo.String{Value: t1.Location().String()}

	return
}

func timesInLocation(args ...tengo.Object) (
	ret tengo.Object,
	err error,
) {
	if len(args) != 2 {
		err = tengo.ErrWrongNumArguments
		return
	}

	t1, ok := tengo.ToTime(args[0])
	if !ok {
		err = tengo.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "time(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	s2, ok := tengo.ToString(args[1])
	if !ok {
		err = tengo.ErrInvalidArgumentType{
			Name:     "second",
			Expected: "string(compatible)",
			Found:    args[1].TypeName(),
		}
		return
	}

	location, err := time.LoadLocation(s2)
	if err != nil {
		ret = wrapError(err)
		return
	}

	ret = &tengo.Time{Value: t1.In(location)}

	return
}

func timesTimeString(args ...tengo.Object) (ret tengo.Object, err error) {
	if len(args) != 1 {
		err = tengo.ErrWrongNumArguments
		return
	}

	t1, ok := tengo.ToTime(args[0])
	if !ok {
		err = tengo.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "time(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	ret = &tengo.String{Value: t1.String()}

	return
}
