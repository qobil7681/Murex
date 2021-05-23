package profile_test

import (
	"os"
	"strings"
	"testing"

	_ "github.com/lmorg/murex/builtins"
	"github.com/lmorg/murex/config/profile"
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/test"
	"github.com/lmorg/murex/test/count"
	"github.com/lmorg/murex/utils/home"
)

func TestProfilePaths(t *testing.T) {
	count.Tests(t, 10)

	var path string
	home := home.MyDir
	temp, err := test.TempDir()
	if err != nil {
		t.Fatalf(err.Error())
	}

	// unset env vars (default paths)

	os.Unsetenv(profile.PreloadEnvVar) // don't care about errors
	path = profile.PreloadPath()
	if !strings.HasPrefix(path, home) {
		t.Error("Unexpected PreloadPath():")
		t.Logf("Expected prefix:  '%s'", home)
		t.Logf("Actual full path: '%s'", path)
	}

	os.Unsetenv(profile.ModuleEnvVar) // don't care about errors
	path = profile.ModulePath()
	if !strings.HasPrefix(path, home) {
		t.Error("Unexpected ModulePath():")
		t.Logf("Expected prefix:  '%s'", home)
		t.Logf("Actual full path: '%s'", path)
	}

	os.Unsetenv(profile.ProfileEnvVar) // don't care about errors
	path = profile.ProfilePath()
	if !strings.HasPrefix(path, home) {
		t.Error("Unexpected ProfilePath():")
		t.Logf("Expected prefix:  '%s'", home)
		t.Logf("Actual full path: '%s'", path)
	}

	// set env vars (custom paths)

	if os.Setenv(profile.PreloadEnvVar, temp) != nil {
		t.Errorf("Unable to set env var %s: %s", profile.PreloadEnvVar, err.Error())
	}
	path = profile.PreloadPath()
	if !strings.HasPrefix(path, temp) {
		t.Error("Unexpected PreloadPath():")
		t.Logf("Expected prefix:  '%s'", temp)
		t.Logf("Actual full path: '%s'", path)
	}

	if os.Setenv(profile.ModuleEnvVar, temp) != nil {
		t.Errorf("Unable to set env var %s: %s", profile.ModuleEnvVar, err.Error())
	}
	path = profile.ModulePath()
	if !strings.HasPrefix(path, temp) {
		t.Error("Unexpected ModulePath():")
		t.Logf("Expected prefix:  '%s'", temp)
		t.Logf("Actual full path: '%s'", path)
	}

	if os.Setenv(profile.ProfileEnvVar, temp) != nil {
		t.Errorf("Unable to set env var %s: %s", profile.ProfileEnvVar, err.Error())
	}
	path = profile.ProfilePath()
	if !strings.HasPrefix(path, temp) {
		t.Error("Unexpected ProfilePath():")
		t.Logf("Expected prefix:  '%s'", temp)
		t.Logf("Actual full path: '%s'", path)
	}

	// set env vars (exact custom file names)

	if os.Setenv(profile.PreloadEnvVar, temp+"foobar") != nil {
		t.Errorf("Unable to set env var %s: %s", profile.PreloadEnvVar, err.Error())
	}
	path = profile.PreloadPath()
	if path != temp+"foobar" {
		t.Error("Unexpected PreloadPath():")
		t.Logf("Expected path:  '%s'", temp+"foobar")
		t.Logf("Actual path:    '%s'", path)
	}

	if os.Setenv(profile.ProfileEnvVar, temp+"foobar") != nil {
		t.Errorf("Unable to set env var %s: %s", profile.ProfileEnvVar, err.Error())
	}
	path = profile.ProfilePath()
	if path != temp+"foobar" {
		t.Error("Unexpected ProfilePath():")
		t.Logf("Expected path:  '%s'", temp+"foobar")
		t.Logf("Actual path:    '%s'", path)
	}

	// as above but negative test

	if os.Setenv(profile.PreloadEnvVar, temp) != nil {
		t.Errorf("Unable to set env var %s: %s", profile.PreloadEnvVar, err.Error())
	}
	path = profile.PreloadPath()
	if path == temp {
		t.Error("Unexpected PreloadPath():")
		t.Logf("Expected prefix:  '%s'", temp)
		t.Logf("Actual full path: '%s'", path)
	}

	if os.Setenv(profile.ProfileEnvVar, temp) != nil {
		t.Errorf("Unable to set env var %s: %s", profile.ProfileEnvVar, err.Error())
	}
	path = profile.ProfilePath()
	if path == temp+"foobar" {
		t.Error("Unexpected ProfilePath():")
		t.Logf("Expected prefix:  '%s'", temp)
		t.Logf("Actual full path: '%s'", path)
	}
}

// TestProfileAndCustomPaths tests a range of functionality from the env var
// custom paths to a lot of the code surrounding loading, enabling and disabling
// modules and packages
func TestProfileAndCustomPaths(t *testing.T) {
	var (
		preloadFileName = "preload.mx"
		modulesPathName = ""
		profileFileName = "profile.mx"
	)

	path, err := test.TempDir()
	if err != nil {
		t.Fatalf(err.Error())
	}

	// set env vars

	if os.Setenv(profile.PreloadEnvVar, path+preloadFileName) != nil {
		t.Errorf("Unable to set env var %s: %s", profile.PreloadEnvVar, err.Error())
	}

	if os.Setenv(profile.ModuleEnvVar, path+modulesPathName) != nil {
		t.Errorf("Unable to set env var %s: %s", profile.ModuleEnvVar, err.Error())
	}

	if os.Setenv(profile.ProfileEnvVar, path+profileFileName) != nil {
		t.Errorf("Unable to set env var %s: %s", profile.ProfileEnvVar, err.Error())
	}

	// initialize preload

	file, err := os.OpenFile(path+preloadFileName, os.O_CREATE|os.O_TRUNC|os.O_RDWR, 0640)
	if err != nil {
		t.Fatalf("Error initializing %s: %s", preloadFileName, err.Error())
	}

	_, err = file.WriteString("function: test_preload {}\n")
	if err != nil {
		t.Fatalf("Error initializing %s: %s", preloadFileName, err.Error())
	}

	if file.Close() != nil {
		t.Fatalf("Error closing %s: %s", preloadFileName, err.Error())
	}

	// initialize profile

	file, err = os.OpenFile(path+profileFileName, os.O_CREATE|os.O_TRUNC|os.O_RDWR, 0640)
	if err != nil {
		t.Fatalf("Error initializing %s: %s", profileFileName, err.Error())
	}

	_, err = file.WriteString("function: test_profile {}\n")
	if err != nil {
		t.Fatalf("Error initializing %s: %s", profileFileName, err.Error())
	}

	if file.Close() != nil {
		t.Fatalf("Error closing %s: %s", profileFileName, err.Error())
	}

	// run tests

	count.Tests(t, 2)

	lang.InitEnv()
	profile.Execute()

	if !lang.MxFunctions.Exists("test_preload") {
		t.Errorf("test_preload failed to be defined. Reason: unknown")
		t.Logf("  %v", lang.MxFunctions.Dump())
	}

	if !lang.MxFunctions.Exists("test_profile") {
		t.Errorf("test_profile failed to be defined. Reason: unknown")
		t.Logf("  %v", lang.MxFunctions.Dump())
	}

	// TODO: enable/disable modules
}
