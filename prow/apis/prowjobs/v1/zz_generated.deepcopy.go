//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1

import (
	url "net/url"

	v1beta1 "github.com/tektoncd/pipeline/pkg/apis/pipeline/v1beta1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CensoringOptions) DeepCopyInto(out *CensoringOptions) {
	*out = *in
	if in.CensoringConcurrency != nil {
		in, out := &in.CensoringConcurrency, &out.CensoringConcurrency
		*out = new(int64)
		**out = **in
	}
	if in.CensoringBufferSize != nil {
		in, out := &in.CensoringBufferSize, &out.CensoringBufferSize
		*out = new(int)
		**out = **in
	}
	if in.IncludeDirectories != nil {
		in, out := &in.IncludeDirectories, &out.IncludeDirectories
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ExcludeDirectories != nil {
		in, out := &in.ExcludeDirectories, &out.ExcludeDirectories
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CensoringOptions.
func (in *CensoringOptions) DeepCopy() *CensoringOptions {
	if in == nil {
		return nil
	}
	out := new(CensoringOptions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DecorationConfig) DeepCopyInto(out *DecorationConfig) {
	*out = *in
	if in.Timeout != nil {
		in, out := &in.Timeout, &out.Timeout
		*out = new(Duration)
		**out = **in
	}
	if in.GracePeriod != nil {
		in, out := &in.GracePeriod, &out.GracePeriod
		*out = new(Duration)
		**out = **in
	}
	if in.UtilityImages != nil {
		in, out := &in.UtilityImages, &out.UtilityImages
		*out = new(UtilityImages)
		**out = **in
	}
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = new(Resources)
		(*in).DeepCopyInto(*out)
	}
	if in.GCSConfiguration != nil {
		in, out := &in.GCSConfiguration, &out.GCSConfiguration
		*out = new(GCSConfiguration)
		(*in).DeepCopyInto(*out)
	}
	if in.GCSCredentialsSecret != nil {
		in, out := &in.GCSCredentialsSecret, &out.GCSCredentialsSecret
		*out = new(string)
		**out = **in
	}
	if in.S3CredentialsSecret != nil {
		in, out := &in.S3CredentialsSecret, &out.S3CredentialsSecret
		*out = new(string)
		**out = **in
	}
	if in.DefaultServiceAccountName != nil {
		in, out := &in.DefaultServiceAccountName, &out.DefaultServiceAccountName
		*out = new(string)
		**out = **in
	}
	if in.SSHKeySecrets != nil {
		in, out := &in.SSHKeySecrets, &out.SSHKeySecrets
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.SSHHostFingerprints != nil {
		in, out := &in.SSHHostFingerprints, &out.SSHHostFingerprints
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.SkipCloning != nil {
		in, out := &in.SkipCloning, &out.SkipCloning
		*out = new(bool)
		**out = **in
	}
	if in.CookiefileSecret != nil {
		in, out := &in.CookiefileSecret, &out.CookiefileSecret
		*out = new(string)
		**out = **in
	}
	if in.OauthTokenSecret != nil {
		in, out := &in.OauthTokenSecret, &out.OauthTokenSecret
		*out = new(OauthTokenSecret)
		**out = **in
	}
	if in.GitHubAPIEndpoints != nil {
		in, out := &in.GitHubAPIEndpoints, &out.GitHubAPIEndpoints
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.GitHubAppPrivateKeySecret != nil {
		in, out := &in.GitHubAppPrivateKeySecret, &out.GitHubAppPrivateKeySecret
		*out = new(GitHubAppPrivateKeySecret)
		**out = **in
	}
	if in.CensorSecrets != nil {
		in, out := &in.CensorSecrets, &out.CensorSecrets
		*out = new(bool)
		**out = **in
	}
	if in.CensoringOptions != nil {
		in, out := &in.CensoringOptions, &out.CensoringOptions
		*out = new(CensoringOptions)
		(*in).DeepCopyInto(*out)
	}
	if in.UploadIgnoresInterrupts != nil {
		in, out := &in.UploadIgnoresInterrupts, &out.UploadIgnoresInterrupts
		*out = new(bool)
		**out = **in
	}
	if in.SetLimitEqualsMemoryRequest != nil {
		in, out := &in.SetLimitEqualsMemoryRequest, &out.SetLimitEqualsMemoryRequest
		*out = new(bool)
		**out = **in
	}
	if in.DefaultMemoryRequest != nil {
		in, out := &in.DefaultMemoryRequest, &out.DefaultMemoryRequest
		x := (*in).DeepCopy()
		*out = &x
	}
	if in.PodPendingTimeout != nil {
		in, out := &in.PodPendingTimeout, &out.PodPendingTimeout
		*out = new(metav1.Duration)
		**out = **in
	}
	if in.PodRunningTimeout != nil {
		in, out := &in.PodRunningTimeout, &out.PodRunningTimeout
		*out = new(metav1.Duration)
		**out = **in
	}
	if in.PodUnscheduledTimeout != nil {
		in, out := &in.PodUnscheduledTimeout, &out.PodUnscheduledTimeout
		*out = new(metav1.Duration)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DecorationConfig.
func (in *DecorationConfig) DeepCopy() *DecorationConfig {
	if in == nil {
		return nil
	}
	out := new(DecorationConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Duration) DeepCopyInto(out *Duration) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Duration.
func (in *Duration) DeepCopy() *Duration {
	if in == nil {
		return nil
	}
	out := new(Duration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GCSConfiguration) DeepCopyInto(out *GCSConfiguration) {
	*out = *in
	if in.MediaTypes != nil {
		in, out := &in.MediaTypes, &out.MediaTypes
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GCSConfiguration.
func (in *GCSConfiguration) DeepCopy() *GCSConfiguration {
	if in == nil {
		return nil
	}
	out := new(GCSConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GitHubAppPrivateKeySecret) DeepCopyInto(out *GitHubAppPrivateKeySecret) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GitHubAppPrivateKeySecret.
func (in *GitHubAppPrivateKeySecret) DeepCopy() *GitHubAppPrivateKeySecret {
	if in == nil {
		return nil
	}
	out := new(GitHubAppPrivateKeySecret)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GitHubTeamSlug) DeepCopyInto(out *GitHubTeamSlug) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GitHubTeamSlug.
func (in *GitHubTeamSlug) DeepCopy() *GitHubTeamSlug {
	if in == nil {
		return nil
	}
	out := new(GitHubTeamSlug)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JenkinsSpec) DeepCopyInto(out *JenkinsSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JenkinsSpec.
func (in *JenkinsSpec) DeepCopy() *JenkinsSpec {
	if in == nil {
		return nil
	}
	out := new(JenkinsSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OauthTokenSecret) DeepCopyInto(out *OauthTokenSecret) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OauthTokenSecret.
func (in *OauthTokenSecret) DeepCopy() *OauthTokenSecret {
	if in == nil {
		return nil
	}
	out := new(OauthTokenSecret)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProwJob) DeepCopyInto(out *ProwJob) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProwJob.
func (in *ProwJob) DeepCopy() *ProwJob {
	if in == nil {
		return nil
	}
	out := new(ProwJob)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ProwJob) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProwJobDefault) DeepCopyInto(out *ProwJobDefault) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProwJobDefault.
func (in *ProwJobDefault) DeepCopy() *ProwJobDefault {
	if in == nil {
		return nil
	}
	out := new(ProwJobDefault)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProwJobList) DeepCopyInto(out *ProwJobList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ProwJob, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProwJobList.
func (in *ProwJobList) DeepCopy() *ProwJobList {
	if in == nil {
		return nil
	}
	out := new(ProwJobList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ProwJobList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProwJobSpec) DeepCopyInto(out *ProwJobSpec) {
	*out = *in
	if in.Refs != nil {
		in, out := &in.Refs, &out.Refs
		*out = new(Refs)
		(*in).DeepCopyInto(*out)
	}
	if in.ExtraRefs != nil {
		in, out := &in.ExtraRefs, &out.ExtraRefs
		*out = make([]Refs, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.PodSpec != nil {
		in, out := &in.PodSpec, &out.PodSpec
		*out = new(corev1.PodSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.JenkinsSpec != nil {
		in, out := &in.JenkinsSpec, &out.JenkinsSpec
		*out = new(JenkinsSpec)
		**out = **in
	}
	if in.PipelineRunSpec != nil {
		in, out := &in.PipelineRunSpec, &out.PipelineRunSpec
		*out = new(v1beta1.PipelineRunSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.TektonPipelineRunSpec != nil {
		in, out := &in.TektonPipelineRunSpec, &out.TektonPipelineRunSpec
		*out = new(TektonPipelineRunSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.DecorationConfig != nil {
		in, out := &in.DecorationConfig, &out.DecorationConfig
		*out = new(DecorationConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.ReporterConfig != nil {
		in, out := &in.ReporterConfig, &out.ReporterConfig
		*out = new(ReporterConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.RerunAuthConfig != nil {
		in, out := &in.RerunAuthConfig, &out.RerunAuthConfig
		*out = new(RerunAuthConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.ProwJobDefault != nil {
		in, out := &in.ProwJobDefault, &out.ProwJobDefault
		*out = new(ProwJobDefault)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProwJobSpec.
func (in *ProwJobSpec) DeepCopy() *ProwJobSpec {
	if in == nil {
		return nil
	}
	out := new(ProwJobSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProwJobStatus) DeepCopyInto(out *ProwJobStatus) {
	*out = *in
	in.StartTime.DeepCopyInto(&out.StartTime)
	if in.PendingTime != nil {
		in, out := &in.PendingTime, &out.PendingTime
		*out = (*in).DeepCopy()
	}
	if in.CompletionTime != nil {
		in, out := &in.CompletionTime, &out.CompletionTime
		*out = (*in).DeepCopy()
	}
	if in.PrevReportStates != nil {
		in, out := &in.PrevReportStates, &out.PrevReportStates
		*out = make(map[string]ProwJobState, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProwJobStatus.
func (in *ProwJobStatus) DeepCopy() *ProwJobStatus {
	if in == nil {
		return nil
	}
	out := new(ProwJobStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProwPath) DeepCopyInto(out *ProwPath) {
	*out = *in
	if in.User != nil {
		in, out := &in.User, &out.User
		*out = new(url.Userinfo)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProwPath.
func (in *ProwPath) DeepCopy() *ProwPath {
	if in == nil {
		return nil
	}
	out := new(ProwPath)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Pull) DeepCopyInto(out *Pull) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Pull.
func (in *Pull) DeepCopy() *Pull {
	if in == nil {
		return nil
	}
	out := new(Pull)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Refs) DeepCopyInto(out *Refs) {
	*out = *in
	if in.Pulls != nil {
		in, out := &in.Pulls, &out.Pulls
		*out = make([]Pull, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Refs.
func (in *Refs) DeepCopy() *Refs {
	if in == nil {
		return nil
	}
	out := new(Refs)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReporterConfig) DeepCopyInto(out *ReporterConfig) {
	*out = *in
	if in.Slack != nil {
		in, out := &in.Slack, &out.Slack
		*out = new(SlackReporterConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.RocketChat != nil {
		in, out := &in.RocketChat, &out.RocketChat
		*out = new(RocketChatReporterConfig)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReporterConfig.
func (in *ReporterConfig) DeepCopy() *ReporterConfig {
	if in == nil {
		return nil
	}
	out := new(ReporterConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RerunAuthConfig) DeepCopyInto(out *RerunAuthConfig) {
	*out = *in
	if in.GitHubTeamIDs != nil {
		in, out := &in.GitHubTeamIDs, &out.GitHubTeamIDs
		*out = make([]int, len(*in))
		copy(*out, *in)
	}
	if in.GitHubTeamSlugs != nil {
		in, out := &in.GitHubTeamSlugs, &out.GitHubTeamSlugs
		*out = make([]GitHubTeamSlug, len(*in))
		copy(*out, *in)
	}
	if in.GitHubUsers != nil {
		in, out := &in.GitHubUsers, &out.GitHubUsers
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.GitHubOrgs != nil {
		in, out := &in.GitHubOrgs, &out.GitHubOrgs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RerunAuthConfig.
func (in *RerunAuthConfig) DeepCopy() *RerunAuthConfig {
	if in == nil {
		return nil
	}
	out := new(RerunAuthConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Resources) DeepCopyInto(out *Resources) {
	*out = *in
	if in.CloneRefs != nil {
		in, out := &in.CloneRefs, &out.CloneRefs
		*out = new(corev1.ResourceRequirements)
		(*in).DeepCopyInto(*out)
	}
	if in.InitUpload != nil {
		in, out := &in.InitUpload, &out.InitUpload
		*out = new(corev1.ResourceRequirements)
		(*in).DeepCopyInto(*out)
	}
	if in.PlaceEntrypoint != nil {
		in, out := &in.PlaceEntrypoint, &out.PlaceEntrypoint
		*out = new(corev1.ResourceRequirements)
		(*in).DeepCopyInto(*out)
	}
	if in.Sidecar != nil {
		in, out := &in.Sidecar, &out.Sidecar
		*out = new(corev1.ResourceRequirements)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Resources.
func (in *Resources) DeepCopy() *Resources {
	if in == nil {
		return nil
	}
	out := new(Resources)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RocketChatReporterConfig) DeepCopyInto(out *RocketChatReporterConfig) {
	*out = *in
	if in.JobStatesToReport != nil {
		in, out := &in.JobStatesToReport, &out.JobStatesToReport
		*out = make([]ProwJobState, len(*in))
		copy(*out, *in)
	}
	if in.Report != nil {
		in, out := &in.Report, &out.Report
		*out = new(bool)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RocketChatReporterConfig.
func (in *RocketChatReporterConfig) DeepCopy() *RocketChatReporterConfig {
	if in == nil {
		return nil
	}
	out := new(RocketChatReporterConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SlackReporterConfig) DeepCopyInto(out *SlackReporterConfig) {
	*out = *in
	if in.JobStatesToReport != nil {
		in, out := &in.JobStatesToReport, &out.JobStatesToReport
		*out = make([]ProwJobState, len(*in))
		copy(*out, *in)
	}
	if in.Report != nil {
		in, out := &in.Report, &out.Report
		*out = new(bool)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SlackReporterConfig.
func (in *SlackReporterConfig) DeepCopy() *SlackReporterConfig {
	if in == nil {
		return nil
	}
	out := new(SlackReporterConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TektonPipelineRunSpec) DeepCopyInto(out *TektonPipelineRunSpec) {
	*out = *in
	if in.V1Beta1 != nil {
		in, out := &in.V1Beta1, &out.V1Beta1
		*out = new(v1beta1.PipelineRunSpec)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TektonPipelineRunSpec.
func (in *TektonPipelineRunSpec) DeepCopy() *TektonPipelineRunSpec {
	if in == nil {
		return nil
	}
	out := new(TektonPipelineRunSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UtilityImages) DeepCopyInto(out *UtilityImages) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UtilityImages.
func (in *UtilityImages) DeepCopy() *UtilityImages {
	if in == nil {
		return nil
	}
	out := new(UtilityImages)
	in.DeepCopyInto(out)
	return out
}
