package main

import ("testing"

"fyne.io/fyne/v2/test"
)

func Test_makeUI(t *testing.T){
	var testCfg config
	edit, preview := testCfg.makeUI()

	test.Type(edit, "Hello")

	if preview.String() != "Hello"{
		t.Error("Failed! We did not get the expected value!")
	}
}