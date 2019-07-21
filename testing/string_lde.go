
/*
 This file was autogenerated via
 ------------------------------------------------------------
 ldetool generate --package ldetesting --go-string string.lde
 ------------------------------------------------------------
 do not touch it with bare hands!
*/

package ldetesting

import (
	"fmt"
	ip "net"
	"strconv"
	"strings"
	time "time"
)

var constAbc = "abc"
var constAddrColonSpace = "addr: "
var constAmountColon = "Amount:"
var constChangeSpaceInternalSpaceStateSpace = "change internal state "
var constColonC = ":c"
var constLsbrck = "["
var constSpacePumpSpace = " Pump "
var constSpaceToSpace = " to "
var constStarsSpaceTimeColonSpace = "*** Time: "
var constStateSpaceChangeSpace = "State change "
var constUnrecognizedSequence = "ï»¿"

// Rule ...
type Rule struct {
	Rest     string
	Data     string
	Signed   int
	Unsigned uint
	Str      string
}

// Extract ...
func (p *Rule) Extract(line string) (bool, error) {
	p.Rest = line
	var err error
	var pos int
	var tmp string
	var tmpInt int64
	var tmpUint uint64

	// Checks if the rest starts with '[' and pass it
	if len(p.Rest) >= 1 && p.Rest[0] == '[' {
		p.Rest = p.Rest[1:]
	} else {
		return false, nil
	}

	// Take until ']' as Data(string)
	pos = strings.IndexByte(p.Rest, ']')
	if pos >= 0 {
		p.Data = p.Rest[:pos]
		p.Rest = p.Rest[pos+1:]
	} else {
		return false, nil
	}

	// Checks if the rest starts with ' ' and pass it
	if len(p.Rest) >= 1 && p.Rest[0] == ' ' {
		p.Rest = p.Rest[1:]
	} else {
		return false, nil
	}

	// Take until ' ' as Signed(int)
	pos = strings.IndexByte(p.Rest, ' ')
	if pos >= 0 {
		tmp = p.Rest[:pos]
		p.Rest = p.Rest[pos+1:]
	} else {
		return false, nil
	}
	if tmpInt, err = strconv.ParseInt(tmp, 10, 64); err != nil {
		return false, fmt.Errorf("parsing `%s` into field Signed(int): %s", tmp, err)
	}
	p.Signed = int(tmpInt)

	// Checks if the rest starts with ' ' and pass it
	if len(p.Rest) >= 1 && p.Rest[0] == ' ' {
		p.Rest = p.Rest[1:]
	} else {
		return false, nil
	}

	// Take until ' ' as Unsigned(uint)
	pos = strings.IndexByte(p.Rest, ' ')
	if pos >= 0 {
		tmp = p.Rest[:pos]
		p.Rest = p.Rest[pos+1:]
	} else {
		return false, nil
	}
	if tmpUint, err = strconv.ParseUint(tmp, 10, 64); err != nil {
		return false, fmt.Errorf("parsing `%s` into field Unsigned(uint): %s", tmp, err)
	}
	p.Unsigned = uint(tmpUint)

	// Take the rest as Str(str)
	p.Str = p.Rest
	p.Rest = p.Rest[len(p.Rest):]
	return true, nil
}

// RegressionCheck1 ...
type RegressionCheck1 struct {
	Rest   string
	Time   string
	Pump   int8
	PState struct {
		Valid bool
		State string
	}
	IState struct {
		Valid bool
		State string
	}
}

// Extract ...
func (p *RegressionCheck1) Extract(line string) (bool, error) {
	p.Rest = line
	var err error
	var pos int
	var rest1 string
	var tmp string
	var tmpInt int64

	// Take until " Pump " as Time(string)
	pos = strings.Index(p.Rest, constSpacePumpSpace)
	if pos >= 0 {
		p.Time = p.Rest[:pos]
		p.Rest = p.Rest[pos+len(constSpacePumpSpace):]
	} else {
		return false, nil
	}

	// Take until ' ' as Pump(int8)
	pos = -1
	for i, char := range p.Rest {
		if char == ' ' {
			pos = i
			break
		}
	}
	if pos >= 0 {
		tmp = p.Rest[:pos]
		p.Rest = p.Rest[pos+1:]
	} else {
		return false, nil
	}
	if tmpInt, err = strconv.ParseInt(tmp, 10, 8); err != nil {
		return false, fmt.Errorf("parsing `%s` into field Pump(int8): %s", tmp, err)
	}
	p.Pump = int8(tmpInt)
	rest1 = p.Rest

	// Checks if the rest starts with `"State change "` and pass it
	if strings.HasPrefix(rest1, constStateSpaceChangeSpace) {
		rest1 = rest1[len(constStateSpaceChangeSpace):]
	} else {
		p.PState.Valid = false
		goto regressioncheck1PStateLabel
	}

	// Looking for " to " and then pass it
	pos = strings.Index(rest1, constSpaceToSpace)
	if pos >= 0 {
		rest1 = rest1[pos+len(constSpaceToSpace):]
	} else {
		p.PState.Valid = false
		goto regressioncheck1PStateLabel
	}

	// Take until "[" as State(string)
	pos = strings.Index(rest1, constLsbrck)
	if pos >= 0 {
		p.PState.State = rest1[:pos]
		rest1 = rest1[pos+len(constLsbrck):]
	} else {
		p.PState.Valid = false
		goto regressioncheck1PStateLabel
	}

	p.PState.Valid = true
	p.Rest = rest1
regressioncheck1PStateLabel:
	rest1 = p.Rest

	// Checks if the rest starts with `"change internal state "` and pass it
	if strings.HasPrefix(rest1, constChangeSpaceInternalSpaceStateSpace) {
		rest1 = rest1[len(constChangeSpaceInternalSpaceStateSpace):]
	} else {
		p.IState.Valid = false
		goto regressioncheck1IStateLabel
	}

	// Looking for " to " and then pass it
	pos = strings.Index(rest1, constSpaceToSpace)
	if pos >= 0 {
		rest1 = rest1[pos+len(constSpaceToSpace):]
	} else {
		p.IState.Valid = false
		goto regressioncheck1IStateLabel
	}

	// Take the rest as State(string)
	p.IState.State = rest1
	rest1 = rest1[len(rest1):]
	p.IState.Valid = true
	p.Rest = rest1
regressioncheck1IStateLabel:

	return true, nil
}

// GetPStateState ...
func (p *RegressionCheck1) GetPStateState() (res string) {
	if p.PState.Valid {
		res = p.PState.State
	}
	return
}

// GetIStateState ...
func (p *RegressionCheck1) GetIStateState() (res string) {
	if p.IState.Valid {
		res = p.IState.State
	}
	return
}

// RegressionCheck2 ...
type RegressionCheck2 struct {
	Rest string
	Time string
}

// Extract ...
func (p *RegressionCheck2) Extract(line string) (bool, error) {
	p.Rest = line

	// Checks if the rest starts with `"ï»¿"` and pass it
	if strings.HasPrefix(p.Rest, constUnrecognizedSequence) {
		p.Rest = p.Rest[len(constUnrecognizedSequence):]
	} else {
		return false, nil
	}

	// Checks if the rest starts with `"*** Time: "` and pass it
	if strings.HasPrefix(p.Rest, constStarsSpaceTimeColonSpace) {
		p.Rest = p.Rest[len(constStarsSpaceTimeColonSpace):]
	} else {
		return false, nil
	}

	// Take the rest as Time(string)
	p.Time = p.Rest
	p.Rest = p.Rest[len(p.Rest):]
	return true, nil
}

// RegressionCheck3 ...
type RegressionCheck3 struct {
	Rest string
}

// Extract ...
func (p *RegressionCheck3) Extract(line string) (bool, error) {
	p.Rest = line

	// Checks if rest[1:] starts with 'b'
	if len(p.Rest) < 1+1 || p.Rest[1] != 'b' {
		return false, nil
	}

	// Checks if rest[2:] starts with `":c"`
	if len(p.Rest) < 2+len(constColonC) || !strings.HasPrefix(p.Rest[2:], constColonC) {
		return false, nil
	}

	return true, nil
}

// BeforeLookup ...
type BeforeLookup struct {
	Rest string
	Data string
}

// Extract ...
func (p *BeforeLookup) Extract(line string) (bool, error) {
	p.Rest = line
	var pos int

	// Looking for "abc" and then pass it
	pos = strings.Index(p.Rest, constAbc)
	if pos >= 0 {
		p.Rest = p.Rest[pos:]
	} else {
		return false, nil
	}

	// Take the rest as Data(string)
	p.Data = p.Rest
	p.Rest = p.Rest[len(p.Rest):]
	return true, nil
}

// CheckPrefix ...
type CheckPrefix struct {
	Rest string
	Data string
}

// Extract ...
func (p *CheckPrefix) Extract(line string) (bool, error) {
	p.Rest = line

	// Checks if the rest starts with `"abc"`
	if !strings.HasPrefix(p.Rest, constAbc) {
		return false, nil
	}

	// Take the rest as Data(string)
	p.Data = p.Rest
	p.Rest = p.Rest[len(p.Rest):]
	return true, nil
}

// PassHeadingStringRegression ...
type PassHeadingStringRegression struct {
	Rest string
	Data string
}

// Extract ...
func (p *PassHeadingStringRegression) Extract(line string) (bool, error) {
	p.Rest = line
	var headPassCounter int
	var headPassValue rune

	// Pass all characters '#' at the rest start
	for headPassCounter, headPassValue = range string(p.Rest) {
		if headPassValue != '#' {
			break
		}
	}
	if headPassCounter > 0 {
		p.Rest = p.Rest[headPassCounter:]
	}

	// Take the rest as Data(string)
	p.Data = p.Rest
	p.Rest = p.Rest[len(p.Rest):]
	return true, nil
}

// Custom ...
type Custom struct {
	Rest string
	Time time.Time
	Addr struct {
		Valid bool
		IP    ip.IP
	}
}

// Extract ...
func (p *Custom) Extract(line string) (bool, error) {
	p.Rest = line
	var err error
	var pos int
	var rest1 string
	var tmp string

	// Take until ' ' as Time(time.Time)
	pos = strings.IndexByte(p.Rest, ' ')
	if pos >= 0 {
		tmp = p.Rest[:pos]
		p.Rest = p.Rest[pos+1:]
	} else {
		return false, nil
	}
	if p.Time, err = p.unmarshalTime(tmp); err != nil {
		return false, fmt.Errorf("parsing `%s` into field Time(time.Time): %s", tmp, err)
	}
	rest1 = p.Rest

	// Checks if the rest starts with `"addr: "` and pass it
	if strings.HasPrefix(rest1, constAddrColonSpace) {
		rest1 = rest1[len(constAddrColonSpace):]
	} else {
		p.Addr.Valid = false
		goto customAddrLabel
	}

	// Take until ' ' as IP(ip.IP)
	pos = strings.IndexByte(rest1, ' ')
	if pos >= 0 {
		tmp = rest1[:pos]
		rest1 = rest1[pos+1:]
	} else {
		p.Addr.Valid = false
		goto customAddrLabel
	}
	if p.Addr.IP, err = p.unmarshalAddrIP(tmp); err != nil {
		return false, fmt.Errorf("parsing `%s` into field Addr.IP(ip.IP): %s", tmp, err)
	}

	p.Addr.Valid = true
	p.Rest = rest1
customAddrLabel:

	return true, nil
}

// GetAddrIP ...
func (p *Custom) GetAddrIP() (res ip.IP) {
	if p.Addr.Valid {
		res = p.Addr.IP
	}
	return
}

// CustomBuiltin ...
type CustomBuiltin struct {
	Rest  string
	Field int
}

// Extract ...
func (p *CustomBuiltin) Extract(line string) (bool, error) {
	p.Rest = line
	var err error

	// Take the rest as Field($int)
	if p.Field, err = p.unmarshalField(p.Rest); err != nil {
		return false, fmt.Errorf("parsing `%s` into field Field(int): %s", p.Rest, err)
	}
	p.Rest = p.Rest[len(p.Rest):]
	return true, nil
}

// Boolean ...
type Boolean struct {
	Rest  string
	Check bool
}

// Extract ...
func (p *Boolean) Extract(line string) (bool, error) {
	p.Rest = line
	var err error

	// Take the rest as Check(bool)
	if p.Check, err = p.unmarshalCheck(p.Rest); err != nil {
		return false, fmt.Errorf("parsing `%s` into field Check(bool): %s", p.Rest, err)
	}
	p.Rest = p.Rest[len(p.Rest):]
	return true, nil
}

// SilentAreas ...
type SilentAreas struct {
	Rest string
	Alt1 struct {
		Valid  bool
		Amount int
	}
	Alt2 struct {
		Valid  bool
		Amount string
	}
}

// Extract ...
func (p *SilentAreas) Extract(line string) (bool, error) {
	p.Rest = line
	var err error
	var rest1 string
	var tmpInt int64
	rest1 = p.Rest

	// Checks if the rest starts with `"Amount:"` and pass it
	if strings.HasPrefix(rest1, constAmountColon) {
		rest1 = rest1[len(constAmountColon):]
	} else {
		p.Alt1.Valid = false
		goto silentareasAlt1Label
	}

	// Take the rest as Amount(int)
	if tmpInt, err = strconv.ParseInt(rest1, 10, 64); err != nil {
		p.Alt1.Valid = false
		goto silentareasAlt1Label
	}
	p.Alt1.Amount = int(tmpInt)
	rest1 = rest1[len(rest1):]
	if len(rest1) != 0 {
		p.Alt1.Valid = false
		goto silentareasAlt1Label
	}

	p.Alt1.Valid = true
	p.Rest = rest1
silentareasAlt1Label:
	rest1 = p.Rest

	// Checks if the rest starts with `"Amount:"` and pass it
	if strings.HasPrefix(rest1, constAmountColon) {
		rest1 = rest1[len(constAmountColon):]
	} else {
		p.Alt2.Valid = false
		goto silentareasAlt2Label
	}

	// Take the rest as Amount(string)
	p.Alt2.Amount = rest1
	rest1 = rest1[len(rest1):]
	if len(rest1) != 0 {
		p.Alt2.Valid = false
		goto silentareasAlt2Label
	}

	p.Alt2.Valid = true
	p.Rest = rest1
silentareasAlt2Label:

	return true, nil
}

// GetAlt1Amount ...
func (p *SilentAreas) GetAlt1Amount() (res int) {
	if p.Alt1.Valid {
		res = p.Alt1.Amount
	}
	return
}

// GetAlt2Amount ...
func (p *SilentAreas) GetAlt2Amount() (res string) {
	if p.Alt2.Valid {
		res = p.Alt2.Amount
	}
	return
}

// TrickyDateParsing ...
type TrickyDateParsing struct {
	Rest string
	Full struct {
		Valid     bool
		Day       int
		Month     int
		Year      int
		Hour      int
		Minute    int
		Second    int
		Microsecs int
	}
	Hour struct {
		Valid     bool
		Hour      int
		Minute    int
		Second    int
		Microsecs int
	}
	Seconds struct {
		Valid     bool
		Second    int
		Microsecs int
	}
}

// Extract ...
func (p *TrickyDateParsing) Extract(line string) (bool, error) {
	p.Rest = line
	var err error
	var pos int
	var rest1 string
	var tmp string
	var tmpInt int64
	rest1 = p.Rest

	// Take until 3rd character if it is'/' as Day(int)
	if len(rest1) >= 2+1 && rest1[2] == '/' {
		pos = 2
	} else {
		pos = -1
	}
	if pos >= 0 {
		tmp = rest1[:pos]
		rest1 = rest1[pos+1:]
	} else {
		p.Full.Valid = false
		goto trickydateparsingFullLabel
	}
	if tmpInt, err = strconv.ParseInt(tmp, 10, 64); err != nil {
		p.Full.Valid = false
		goto trickydateparsingFullLabel
	}
	p.Full.Day = int(tmpInt)

	// Take until 3rd character if it is'/' as Month(int)
	if len(rest1) >= 2+1 && rest1[2] == '/' {
		pos = 2
	} else {
		pos = -1
	}
	if pos >= 0 {
		tmp = rest1[:pos]
		rest1 = rest1[pos+1:]
	} else {
		p.Full.Valid = false
		goto trickydateparsingFullLabel
	}
	if tmpInt, err = strconv.ParseInt(tmp, 10, 64); err != nil {
		p.Full.Valid = false
		goto trickydateparsingFullLabel
	}
	p.Full.Month = int(tmpInt)

	// Take until 5th character if it is' ' as Year(int)
	if len(rest1) >= 4+1 && rest1[4] == ' ' {
		pos = 4
	} else {
		pos = -1
	}
	if pos >= 0 {
		tmp = rest1[:pos]
		rest1 = rest1[pos+1:]
	} else {
		p.Full.Valid = false
		goto trickydateparsingFullLabel
	}
	if tmpInt, err = strconv.ParseInt(tmp, 10, 64); err != nil {
		p.Full.Valid = false
		goto trickydateparsingFullLabel
	}
	p.Full.Year = int(tmpInt)

	// Take until 3rd character if it is':' as Hour(int)
	if len(rest1) >= 2+1 && rest1[2] == ':' {
		pos = 2
	} else {
		pos = -1
	}
	if pos >= 0 {
		tmp = rest1[:pos]
		rest1 = rest1[pos+1:]
	} else {
		p.Full.Valid = false
		goto trickydateparsingFullLabel
	}
	if tmpInt, err = strconv.ParseInt(tmp, 10, 64); err != nil {
		p.Full.Valid = false
		goto trickydateparsingFullLabel
	}
	p.Full.Hour = int(tmpInt)

	// Take until 3rd character if it is':' as Minute(int)
	if len(rest1) >= 2+1 && rest1[2] == ':' {
		pos = 2
	} else {
		pos = -1
	}
	if pos >= 0 {
		tmp = rest1[:pos]
		rest1 = rest1[pos+1:]
	} else {
		p.Full.Valid = false
		goto trickydateparsingFullLabel
	}
	if tmpInt, err = strconv.ParseInt(tmp, 10, 64); err != nil {
		p.Full.Valid = false
		goto trickydateparsingFullLabel
	}
	p.Full.Minute = int(tmpInt)

	// Take until 3rd character if it is'.' as Second(int)
	if len(rest1) >= 2+1 && rest1[2] == '.' {
		pos = 2
	} else {
		pos = -1
	}
	if pos >= 0 {
		tmp = rest1[:pos]
		rest1 = rest1[pos+1:]
	} else {
		p.Full.Valid = false
		goto trickydateparsingFullLabel
	}
	if tmpInt, err = strconv.ParseInt(tmp, 10, 64); err != nil {
		p.Full.Valid = false
		goto trickydateparsingFullLabel
	}
	p.Full.Second = int(tmpInt)

	// Take the rest as Microsecs(int)
	if tmpInt, err = strconv.ParseInt(rest1, 10, 64); err != nil {
		p.Full.Valid = false
		goto trickydateparsingFullLabel
	}
	p.Full.Microsecs = int(tmpInt)
	rest1 = rest1[len(rest1):]
	p.Full.Valid = true
	p.Rest = rest1
trickydateparsingFullLabel:
	rest1 = p.Rest

	// Take until 3rd character if it is':' as Hour(int)
	if len(rest1) >= 2+1 && rest1[2] == ':' {
		pos = 2
	} else {
		pos = -1
	}
	if pos >= 0 {
		tmp = rest1[:pos]
		rest1 = rest1[pos+1:]
	} else {
		p.Hour.Valid = false
		goto trickydateparsingHourLabel
	}
	if tmpInt, err = strconv.ParseInt(tmp, 10, 64); err != nil {
		p.Hour.Valid = false
		goto trickydateparsingHourLabel
	}
	p.Hour.Hour = int(tmpInt)

	// Take until 3rd character if it is':' as Minute(int)
	if len(rest1) >= 2+1 && rest1[2] == ':' {
		pos = 2
	} else {
		pos = -1
	}
	if pos >= 0 {
		tmp = rest1[:pos]
		rest1 = rest1[pos+1:]
	} else {
		p.Hour.Valid = false
		goto trickydateparsingHourLabel
	}
	if tmpInt, err = strconv.ParseInt(tmp, 10, 64); err != nil {
		p.Hour.Valid = false
		goto trickydateparsingHourLabel
	}
	p.Hour.Minute = int(tmpInt)

	// Take until 3rd character if it is'.' as Second(int)
	if len(rest1) >= 2+1 && rest1[2] == '.' {
		pos = 2
	} else {
		pos = -1
	}
	if pos >= 0 {
		tmp = rest1[:pos]
		rest1 = rest1[pos+1:]
	} else {
		p.Hour.Valid = false
		goto trickydateparsingHourLabel
	}
	if tmpInt, err = strconv.ParseInt(tmp, 10, 64); err != nil {
		p.Hour.Valid = false
		goto trickydateparsingHourLabel
	}
	p.Hour.Second = int(tmpInt)

	// Take the rest as Microsecs(int)
	if tmpInt, err = strconv.ParseInt(rest1, 10, 64); err != nil {
		p.Hour.Valid = false
		goto trickydateparsingHourLabel
	}
	p.Hour.Microsecs = int(tmpInt)
	rest1 = rest1[len(rest1):]
	p.Hour.Valid = true
	p.Rest = rest1
trickydateparsingHourLabel:
	rest1 = p.Rest

	// Take until 3rd character if it is'.' as Second(int)
	if len(rest1) >= 2+1 && rest1[2] == '.' {
		pos = 2
	} else {
		pos = -1
	}
	if pos >= 0 {
		tmp = rest1[:pos]
		rest1 = rest1[pos+1:]
	} else {
		p.Seconds.Valid = false
		goto trickydateparsingSecondsLabel
	}
	if tmpInt, err = strconv.ParseInt(tmp, 10, 64); err != nil {
		p.Seconds.Valid = false
		goto trickydateparsingSecondsLabel
	}
	p.Seconds.Second = int(tmpInt)

	// Take the rest as Microsecs(int)
	if tmpInt, err = strconv.ParseInt(rest1, 10, 64); err != nil {
		p.Seconds.Valid = false
		goto trickydateparsingSecondsLabel
	}
	p.Seconds.Microsecs = int(tmpInt)
	rest1 = rest1[len(rest1):]
	p.Seconds.Valid = true
	p.Rest = rest1
trickydateparsingSecondsLabel:

	if len(p.Rest) != 0 {
		return false, nil
	}

	return true, nil
}

// GetFullDay ...
func (p *TrickyDateParsing) GetFullDay() (res int) {
	if p.Full.Valid {
		res = p.Full.Day
	}
	return
}

// GetFullMonth ...
func (p *TrickyDateParsing) GetFullMonth() (res int) {
	if p.Full.Valid {
		res = p.Full.Month
	}
	return
}

// GetFullYear ...
func (p *TrickyDateParsing) GetFullYear() (res int) {
	if p.Full.Valid {
		res = p.Full.Year
	}
	return
}

// GetFullHour ...
func (p *TrickyDateParsing) GetFullHour() (res int) {
	if p.Full.Valid {
		res = p.Full.Hour
	}
	return
}

// GetFullMinute ...
func (p *TrickyDateParsing) GetFullMinute() (res int) {
	if p.Full.Valid {
		res = p.Full.Minute
	}
	return
}

// GetFullSecond ...
func (p *TrickyDateParsing) GetFullSecond() (res int) {
	if p.Full.Valid {
		res = p.Full.Second
	}
	return
}

// GetFullMicrosecs ...
func (p *TrickyDateParsing) GetFullMicrosecs() (res int) {
	if p.Full.Valid {
		res = p.Full.Microsecs
	}
	return
}

// GetHourHour ...
func (p *TrickyDateParsing) GetHourHour() (res int) {
	if p.Hour.Valid {
		res = p.Hour.Hour
	}
	return
}

// GetHourMinute ...
func (p *TrickyDateParsing) GetHourMinute() (res int) {
	if p.Hour.Valid {
		res = p.Hour.Minute
	}
	return
}

// GetHourSecond ...
func (p *TrickyDateParsing) GetHourSecond() (res int) {
	if p.Hour.Valid {
		res = p.Hour.Second
	}
	return
}

// GetHourMicrosecs ...
func (p *TrickyDateParsing) GetHourMicrosecs() (res int) {
	if p.Hour.Valid {
		res = p.Hour.Microsecs
	}
	return
}

// GetSecondsSecond ...
func (p *TrickyDateParsing) GetSecondsSecond() (res int) {
	if p.Seconds.Valid {
		res = p.Seconds.Second
	}
	return
}

// GetSecondsMicrosecs ...
func (p *TrickyDateParsing) GetSecondsMicrosecs() (res int) {
	if p.Seconds.Valid {
		res = p.Seconds.Microsecs
	}
	return
}