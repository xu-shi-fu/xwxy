package httpping

import (
	"embed"

	"github.com/starter-go/application"
	"github.com/starter-go/libgin/modules/libgin"
	"github.com/starter-go/starter"
	"github.com/xu-shi-fu/xwxy/tools/httpping/gen/main4httppingclient"
	"github.com/xu-shi-fu/xwxy/tools/httpping/gen/main4httppingserver"
	"github.com/xu-shi-fu/xwxy/tools/httpping/gen/test4httpping"
)

////////////////////////////////////////////////////////////////////////////////

const (
	theModuleName     = "github.com/xu-shi-fu/xwxy/tools/httpping"
	theModuleVersion  = "v0.0.1"
	theModuleRevision = 1
)

////////////////////////////////////////////////////////////////////////////////

const (
	theMainModuleResPath = "src/main/resources"
	theTestModuleResPath = "src/test/resources"
)

//go:embed "src/main/resources"
var theMainModuleResFS embed.FS

//go:embed "src/test/resources"
var theTestModuleResFS embed.FS

////////////////////////////////////////////////////////////////////////////////

func ModuleForClient() application.Module {
	mb := new(application.ModuleBuilder)

	mb.Name(theModuleName + "#main-client")
	mb.Version(theModuleVersion)
	mb.Revision(theModuleRevision)

	mb.EmbedResources(theMainModuleResFS, theMainModuleResPath)

	mb.Components(main4httppingclient.ExportComponents)

	mb.Depend(starter.Module())
	// mb.Depend(libgin.Module())

	return mb.Create()
}

func ModuleForServer() application.Module {
	mb := new(application.ModuleBuilder)

	mb.Name(theModuleName + "#main-server")
	mb.Version(theModuleVersion)
	mb.Revision(theModuleRevision)

	mb.EmbedResources(theMainModuleResFS, theMainModuleResPath)

	mb.Components(main4httppingserver.ExportComponents)

	mb.Depend(starter.Module())
	mb.Depend(libgin.Module())

	return mb.Create()
}

func ModuleForTest() application.Module {
	mb := new(application.ModuleBuilder)

	mb.Name(theModuleName + "#test")
	mb.Version(theModuleVersion)
	mb.Revision(theModuleRevision)

	mb.EmbedResources(theTestModuleResFS, theTestModuleResPath)

	mb.Components(test4httpping.ExportComponents)

	mb.Depend(ModuleForClient())

	return mb.Create()
}

////////////////////////////////////////////////////////////////////////////////
// EOF
