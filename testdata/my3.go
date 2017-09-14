package testdata

//go:generate truepack  -omit-clue

type OmitClueTestStruct struct {
	S string `zid:"0"`
	N int64  `zid:"1"`
}
