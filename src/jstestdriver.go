package testing

import (
	"testing"

	"github.com/bazelbuild/rules_webtesting/go/webtest"
	"github.com/tebeka/selenium"
)

func TestWebApp(t *testing.T) {
	wd, err := webtest.NewWebDriverSession(selenium.Capabilities{})
	if err != nil {
		t.Fatal(err)
	}

	// your test here
	wd.Get("www.google,com")

	if err := wd.Quit(); err != nil {
		t.Logf("Error quitting webdriver: %v", err)
	}
}
