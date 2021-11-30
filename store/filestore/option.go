package filestore

import (
	"regexp"
	"strings"
)

type Option func(h *fileStore)

func WithBaseURI(baseURI string) Option {
	return func(s *fileStore) {
		baseURI = "/" + strings.Trim(baseURI, "/")
		if baseURI != "/" {
			baseURI += "/"
		}
		s.BaseURI = baseURI
	}
}

func WithBlacklist(blacklist *regexp.Regexp) Option {
	return func(s *fileStore) {
		if blacklist != nil {
			s.Blacklists = append(s.Blacklists, blacklist)
		}
	}
}
