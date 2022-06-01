package tagging

import (
	"fmt"
	"strings"
)

type Mandatorytags struct {
	Name           string
	Creator        string
	POC            string
	Venue          string
	Project        string
	ServiceArea    string
	Capability     string
	CapVersion     string
	Release        string
	Component      string
	SecurityPlanId string
	ExposedWeb     string
	Experimental   string
	UserFacing     string
	CritInfra      string
	SourceControl  string
	Alfa           string
	Bravo          string
	Charlie        string
	Delta          string
	Echo           string
	Foxtrot        string
}

func GenerateMandatoryTags(creator string, pocs []string, venue, project, servicearea, capability, component, capversion, release, securityplan, exposed, experimental, userfacing, critinfra, sourcecontrol string) Mandatorytags {

	mtags := Mandatorytags{
		Name:           fmt.Sprintf("%v-%v-%v-%v-%v", project, venue, servicearea, capability, component),
		Creator:        creator,
		POC:            strings.Join(pocs, ":"),
		Venue:          venue,
		Project:        project,
		ServiceArea:    servicearea,
		Capability:     capability,
		CapVersion:     capversion,
		Release:        release,
		Component:      component,
		SecurityPlanId: securityplan,
		ExposedWeb:     exposed,
		Experimental:   experimental,
		UserFacing:     userfacing,
		CritInfra:      critinfra,
		SourceControl:  sourcecontrol,
		Alfa:           "",
		Bravo:          "",
		Charlie:        "",
		Delta:          "",
		Echo:           "",
		Foxtrot:        "",
	}

	return mtags
}

func generateEC2Tags() {

}

func generateS3Tags() {

}

func generateALBTags() {

}

func generateTGTags() {

}

func generateRedisTags() {

}

func generateSSMTags() {

}

func generateKMSTags() {

}
