package main

import (
	"github.com/adobe-experience-manager-aem-api/mcp-server/config"
	"github.com/adobe-experience-manager-aem-api/mcp-server/models"
	tools_sling "github.com/adobe-experience-manager-aem-api/mcp-server/tools/sling"
	tools_console "github.com/adobe-experience-manager-aem-api/mcp-server/tools/console"
	tools_custom "github.com/adobe-experience-manager-aem-api/mcp-server/tools/custom"
	tools_cq "github.com/adobe-experience-manager-aem-api/mcp-server/tools/cq"
	tools_crx "github.com/adobe-experience-manager-aem-api/mcp-server/tools/crx"
)

func GetAll(cfg *config.APIConfig) []models.Tool {
	return []models.Tool{
		tools_sling.CreateGetauthorizablekeystoreTool(cfg),
		tools_sling.CreateGetpackagefilterTool(cfg),
		tools_sling.CreateGettruststoreinfoTool(cfg),
		tools_console.CreateGetbundleinfoTool(cfg),
		tools_console.CreatePostbundleTool(cfg),
		tools_sling.CreateGetqueryTool(cfg),
		tools_sling.CreatePostqueryTool(cfg),
		tools_console.CreatePostsamlconfigurationTool(cfg),
		tools_sling.CreatePosttreeactivationTool(cfg),
		tools_sling.CreatePostconfigadobegranitesamlauthenticationhandlerTool(cfg),
		tools_sling.CreatePostconfigpropertyTool(cfg),
		tools_custom.CreatePostconfigaempasswordresetTool(cfg),
		tools_sling.CreatePostconfigapacheslingreferrerfilterTool(cfg),
		tools_sling.CreatePostconfigapacheslinggetservletTool(cfg),
		tools_sling.CreatePostnoderwTool(cfg),
		tools_custom.CreateGetaemhealthcheckTool(cfg),
		tools_sling.CreateGetagentsTool(cfg),
		tools_sling.CreatePostpathTool(cfg),
		tools_cq.CreatePostcqactionsTool(cfg),
		tools_sling.CreatePostconfigapachefelixjettybasedhttpserviceTool(cfg),
		tools_sling.CreateDeleteagentTool(cfg),
		tools_sling.CreateGetagentTool(cfg),
		tools_sling.CreatePostagentTool(cfg),
		tools_sling.CreatePostconfigapachehttpcomponentsproxyconfigurationTool(cfg),
		tools_crx.CreateGetinstallstatusTool(cfg),
		tools_console.CreateGetaemproductinfoTool(cfg),
		tools_custom.CreatePostconfigaemhealthcheckservletTool(cfg),
		tools_sling.CreateDeletenodeTool(cfg),
		tools_sling.CreateGetnodeTool(cfg),
		tools_console.CreatePostjmxrepositoryTool(cfg),
		tools_sling.CreatePostconfigapacheslingdavexservletTool(cfg),
		tools_crx.CreatePostsetpasswordTool(cfg),
		tools_crx.CreatePostpackageupdateTool(cfg),
	}
}
