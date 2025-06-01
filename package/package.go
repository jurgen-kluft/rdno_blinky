package rdno_blinky

import (
	denv "github.com/jurgen-kluft/ccode/denv"
	cunittest "github.com/jurgen-kluft/cunittest/package"
	rdno_core "github.com/jurgen-kluft/rdno_core/package"
	rdno_wifi "github.com/jurgen-kluft/rdno_wifi/package"
)

const (
	repo_path = "github.com\\jurgen-kluft"
	repo_name = "rdno_blinky"
)

func GetPackage() *denv.Package {
	name := repo_name

	// dependencies
	cunittestpkg := cunittest.GetPackage()
	ucorepkg := rdno_core.GetPackage()
	uwifipkg := rdno_wifi.GetPackage()

	// main package
	mainpkg := denv.NewPackage(repo_path, repo_name)
	mainpkg.AddPackage(ucorepkg)
	mainpkg.AddPackage(uwifipkg)

	// main library
	mainlib := denv.SetupCppLibProject(mainpkg, name)
	mainlib.AddDependencies(uwifipkg.GetMainLib()...)
	mainlib.AddDependencies(ucorepkg.GetMainLib()...)

	// application
	mainapp := denv.SetupCppAppProject(mainpkg, name)
	mainapp.AddDependency(mainlib)

	// test library
	testlib := denv.SetupCppTestLibProject(mainpkg, name)
	testlib.AddDependencies(uwifipkg.GetTestLib()...)
	testlib.AddDependencies(ucorepkg.GetTestLib()...)
	testlib.AddDependencies(cunittestpkg.GetTestLib()...)

	// unittest project
	maintest := denv.SetupCppTestProjectForDesktop(mainpkg, name)
	maintest.AddDependencies(cunittestpkg.GetMainLib()...)
	maintest.AddDependency(testlib)

	mainpkg.AddMainApp(mainapp)
	mainpkg.AddMainLib(mainlib)
	mainpkg.AddTestLib(testlib)
	mainpkg.AddUnittest(maintest)
	return mainpkg
}
