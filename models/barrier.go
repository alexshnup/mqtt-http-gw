package models

import (
	"errors"
	"strconv"
	"time"
)

var (
	Barriers map[string]*Barrier
)

type Barrier struct {
	BarrierId string
	Payload   string
	Timestamp string
}

func init() {
	Barriers = make(map[string]*Barrier)
	// Barriers["1"] = &Barrier{"1", 1, "one"}
	// Barriers["2"] = &Barrier{"2", 2, "two"}
}

func Add1(barrier Barrier) (BarrierId string) {
	barrier.Timestamp = strconv.FormatInt(time.Now().Unix(), 10)
	Barriers[barrier.BarrierId] = &barrier
	return barrier.BarrierId
}

func Get1(BarrierId string) (barrier *Barrier, err error) {
	if v, ok := Barriers[BarrierId]; ok {
		return v, nil
	}
	return nil, errors.New("BarrierId Not Exist")
}

func GetAllAll() map[string]*Barrier {
	return Barriers
}

func Update1(BarrierId string, Payload string) (err error) {
	if v, ok := Barriers[BarrierId]; ok {
		v.Payload = Payload
		v.Timestamp = strconv.FormatInt(time.Now().Unix(), 10)

		return nil
	}
	return errors.New("BarrierId Not Exist")
}

func Delete1(BarrierId string) {
	delete(Barriers, BarrierId)
}
