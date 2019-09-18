package manifest

import (
	"fmt"
	"regexp"
)

type Stemcell struct {
	OS      string
	Version string
}

type Release struct {
	Name     string
	SHA1     string
	Stemcell Stemcell
	Version  string
	URL      string
}

func UpdateStemcellSection(manifestContent []byte, stemcell Stemcell) ([]byte, error) {
	if manifestContent == nil {
		return manifestContent, fmt.Errorf("manifest file has no content")
	}
	stemcellPattern := regexp.MustCompile(`(?s)stemcells:.*- alias: ([\w\-]*).*os: .* version: .*`)

	stemcellsTemplate := `stemcells:
- alias: $1
  os: %s
  version: "%s"
`
	updatedManifestContent := stemcellPattern.ReplaceAll(manifestContent,
		[]byte(fmt.Sprintf(stemcellsTemplate, stemcell.OS, stemcell.Version)))

	return updatedManifestContent, nil
}
