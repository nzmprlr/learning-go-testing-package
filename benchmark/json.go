package benchmark

import (
	_ "embed"
	stdjson "encoding/json"

	sonic "github.com/bytedance/sonic"
	gojson "github.com/goccy/go-json"
)

var (
	//go:embed json.json
	a []byte
)

type Json struct {
	Ca int     `json:"ca"`
	Cl float64 `json:"cl"`
	Ct float64 `json:"ct"`
	I  float64 `json:"i"`
	L  int     `json:"l"`
}

func Std() {
	v := map[string]Json{}
	err := stdjson.Unmarshal(a, &v)
	if err != nil {
		panic(err)
	}

	_ = v
}

func GoJson() {
	v := map[string]Json{}
	err := gojson.Unmarshal(a, &v)
	if err != nil {
		panic(err)
	}

	_ = v
}

func Sonic() {
	v := map[string]Json{}
	err := sonic.Unmarshal(a, &v)
	if err != nil {
		panic(err)
	}

	_ = v
}
