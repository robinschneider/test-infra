package rocketchat

import (
	prowapi "k8s.io/test-infra/prow/apis/prowjobs/v1"
	"k8s.io/test-infra/prow/config"
)

const (
	reporterName    = "rocketchatreporter"
	DefaultHostName = "*"
)

type rocketChatClient interface {
	WriteMessage(text, channel string) error
}

type rocketChatReporter struct {
	clients map[string]rocketChatClient
	config  func(*prowapi.Refs) config.RocketChatReporter
	dryRun  bool
}

func hostAndChannel(cfg *prowapi.RocketChatReporterConfig) (string, string) {
	host, channel := cfg.Host, cfg.Channel
	if host == "" {
		host = DefaultHostName
	}
	return host, channel
}

func (sr *rocketChatReporter) getConfig(pj *prowapi.ProwJob) (*config.RocketChatReporter, *prowapi.RocketChatReporterConfig) {
	refs := pj.Spec.Refs
	if refs == nil && len(pj.Spec.ExtraRefs) > 0 {
		refs = &pj.Spec.ExtraRefs[0]
	}
	globalConfig := sr.config(refs)
	var jobRocketChatConfig *prowapi.RocketChatReporterConfig
	if pj.Spec.ReporterConfig != nil && pj.Spec.ReporterConfig.RocketChat != nil {
		jobRocketChatConfig = pj.Spec.ReporterConfig.RocketChat
	}
	return &globalConfig, jobRocketChatConfig
}
