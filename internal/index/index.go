package index

import (
	"encoding/base64"
	"encoding/json"
	"github.com/golangee/gotrino-tutorial/nestor"
)

var Tutorial nestor.Fragment

func init() {
	buf, err := base64.StdEncoding.DecodeString(tutorials)
	if err != nil {
		panic(err)
	}

	tmp := make([]nestor.Fragment, 0, 1)
	err = json.Unmarshal(buf, &tmp)
	if err != nil {
		panic(err)
	}

	Tutorial = tmp[0]
}
