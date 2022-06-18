package server

import (
	"time"
)

type Geo struct {
	tarih *time.Time
}

func (g Geo) SetDate(tarih string) error {
	zaman, err := time.Parse("2006-01-02", tarih)
	if err != nil {
		return err
	}

	g.tarih = &zaman

	return nil
}

func (g Geo) ClearDate() {

}
