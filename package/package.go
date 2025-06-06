package rdno_blinky

import (
	denv "github.com/jurgen-kluft/ccode/denv"
	cunittest "github.com/jurgen-kluft/cunittest/package"
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
	//ucorepkg := rdno_core.GetPackage()
	uwifipkg := rdno_wifi.GetPackage()

	// main package
	mainpkg := denv.NewPackage(repo_path, repo_name)
	mainpkg.AddPackage(uwifipkg)

	// main library
	mainlib := denv.SetupCppLibProject(mainpkg, name)
	//mainlib.AddDependencies(ucorepkg.GetMainLib()...)
	mainlib.AddDependencies(uwifipkg.GetMainLib()...)

	// test library
	testlib := denv.SetupCppTestLibProject(mainpkg, name)
	//testlib.AddDependencies(ucorepkg.GetTestLib()...)
	testlib.AddDependencies(uwifipkg.GetTestLib()...)

	// unittest project
	maintest := denv.SetupCppTestProject(mainpkg, name)
	maintest.AddDependencies(cunittestpkg.GetMainLib()...)
	maintest.AddDependency(testlib)

	// application
	mainapp := denv.SetupCppAppProject(mainpkg, name, "app")
	mainapp.AddDependency(mainlib)

	mainpkg.AddMainApp(mainapp)
	mainpkg.AddMainLib(mainlib)
	mainpkg.AddTestLib(testlib)
	mainpkg.AddUnittest(maintest)
	return mainpkg
}
