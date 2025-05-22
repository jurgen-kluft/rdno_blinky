package rdno_blinky

import (
	denv "github.com/jurgen-kluft/ccode/denv"
	cunittest "github.com/jurgen-kluft/cunittest/package"
	rdno_core "github.com/jurgen-kluft/rdno_core/package"
	rdno_wifi "github.com/jurgen-kluft/rdno_wifi/package"
)

// GetPackage returns the package object of 'rdno_blinky'
func GetPackage() *denv.Package {
	// Dependencies
	cunittestpkg := cunittest.GetPackage()
	ucorepkg := rdno_core.GetPackage()
	uwifipkg := rdno_wifi.GetPackage()

	// The main (rdno_blinky) package
	mainpkg := denv.NewPackage("rdno_blinky")
	mainpkg.AddPackage(ucorepkg)
	mainpkg.AddPackage(uwifipkg)

	// 'rdno_blinky' library
	mainlib := denv.SetupCppLibProject("rdno_blinky"+"lib", "github.com\\jurgen-kluft\\rdno_blinky")
	mainlib.AddDependencies(uwifipkg.GetMainLib()...)
	mainlib.AddDependencies(ucorepkg.GetMainLib()...)

	// 'rdno_blinky' application
	mainapp := denv.SetupCppAppProject("rdno_blinky", "github.com\\jurgen-kluft\\rdno_blinky")
	mainapp.AddDependency(mainlib)

	// 'rdno_blinky' unittest project
	// The unittest project only supports Windows, Mac and Linux
	maintest := denv.SetupCppTestProjectForDesktop("rdno_blinky"+"test", "github.com\\jurgen-kluft\\rdno_blinky")
	maintest.AddDependency(mainlib)
	maintest.AddDependencies(cunittestpkg.GetMainLib()...)

	mainpkg.AddMainApp(mainapp)
	mainpkg.AddMainLib(mainlib)
	mainpkg.AddUnittest(maintest)
	return mainpkg
}
