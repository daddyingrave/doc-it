package main

import (
	"os"
	"strings"

	"doc-it/pkg/config"
	"doc-it/pkg/docit"
	"doc-it/pkg/errorutils"
	"doc-it/pkg/fsutils"
)

func main() {
	conf := config.Conf{
		IncludeFileTypes: []string{".yaml", ".yml"},
		MetaMarker:       "@doc-it",
		OutputDir:        "test-output",
	}

	yamls := docit.ReadYamls("test-yamls", conf)
	meta := make([]docit.Meta, 0)
	for _, y := range yamls {
		meta = append(meta, y.ToMeta(conf))
	}

	err := fsutils.CreateDirIfNotExist(conf.OutputDir)
	errorutils.Check(err)

	sb := strings.Builder{}

	for _, m := range meta {
		for _, ref := range m.Comments {
			sb.WriteString("Reference:\n" + ref.Reference + "\n")
			sb.WriteString("Path:\n" + ref.ObjectLink + "\n")
			if ref.BlockContent != "" {
				sb.WriteString("Block content:\n" + ref.BlockContent + "\n")
			}
		}

		err := os.WriteFile(
			conf.OutputDir+"/"+m.Path.FileName()+".md",
			[]byte(sb.String()),
			0666,
		)
		sb.Reset()
		errorutils.Check(err)
	}
}
