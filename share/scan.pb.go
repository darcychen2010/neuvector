// Code generated by protoc-gen-go. DO NOT EDIT.
// source: scan.proto

package share

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type ScanErrorCode int32

const (
	ScanErrorCode_ScanErrNone                ScanErrorCode = 0
	ScanErrorCode_ScanErrNetwork             ScanErrorCode = 1
	ScanErrorCode_ScanErrNotSupport          ScanErrorCode = 2
	ScanErrorCode_ScanErrSizeOverLimit       ScanErrorCode = 3
	ScanErrorCode_ScanErrPackage             ScanErrorCode = 4
	ScanErrorCode_ScanErrDatabase            ScanErrorCode = 5
	ScanErrorCode_ScanErrTimeout             ScanErrorCode = 6
	ScanErrorCode_ScanErrInProgress          ScanErrorCode = 7
	ScanErrorCode_ScanErrRegistryAPI         ScanErrorCode = 8
	ScanErrorCode_ScanErrFileSystem          ScanErrorCode = 9
	ScanErrorCode_ScanErrContainerAPI        ScanErrorCode = 10
	ScanErrorCode_ScanErrXrayAPI             ScanErrorCode = 11
	ScanErrorCode_ScanErrContainerExit       ScanErrorCode = 12
	ScanErrorCode_ScanErrAuthentication      ScanErrorCode = 13
	ScanErrorCode_ScanErrCertificate         ScanErrorCode = 14
	ScanErrorCode_ScanErrCanceled            ScanErrorCode = 15
	ScanErrorCode_ScanErrDriverAPINotSupport ScanErrorCode = 16
	ScanErrorCode_ScanErrImageNotFound       ScanErrorCode = 17
	ScanErrorCode_ScanErrAwsDownloadErr      ScanErrorCode = 18
	ScanErrorCode_ScanErrArgument            ScanErrorCode = 19
)

var ScanErrorCode_name = map[int32]string{
	0:  "ScanErrNone",
	1:  "ScanErrNetwork",
	2:  "ScanErrNotSupport",
	3:  "ScanErrSizeOverLimit",
	4:  "ScanErrPackage",
	5:  "ScanErrDatabase",
	6:  "ScanErrTimeout",
	7:  "ScanErrInProgress",
	8:  "ScanErrRegistryAPI",
	9:  "ScanErrFileSystem",
	10: "ScanErrContainerAPI",
	11: "ScanErrXrayAPI",
	12: "ScanErrContainerExit",
	13: "ScanErrAuthentication",
	14: "ScanErrCertificate",
	15: "ScanErrCanceled",
	16: "ScanErrDriverAPINotSupport",
	17: "ScanErrImageNotFound",
	18: "ScanErrAwsDownloadErr",
	19: "ScanErrArgument",
}
var ScanErrorCode_value = map[string]int32{
	"ScanErrNone":                0,
	"ScanErrNetwork":             1,
	"ScanErrNotSupport":          2,
	"ScanErrSizeOverLimit":       3,
	"ScanErrPackage":             4,
	"ScanErrDatabase":            5,
	"ScanErrTimeout":             6,
	"ScanErrInProgress":          7,
	"ScanErrRegistryAPI":         8,
	"ScanErrFileSystem":          9,
	"ScanErrContainerAPI":        10,
	"ScanErrXrayAPI":             11,
	"ScanErrContainerExit":       12,
	"ScanErrAuthentication":      13,
	"ScanErrCertificate":         14,
	"ScanErrCanceled":            15,
	"ScanErrDriverAPINotSupport": 16,
	"ScanErrImageNotFound":       17,
	"ScanErrAwsDownloadErr":      18,
	"ScanErrArgument":            19,
}

func (x ScanErrorCode) String() string {
	return proto.EnumName(ScanErrorCode_name, int32(x))
}
func (ScanErrorCode) EnumDescriptor() ([]byte, []int) { return fileDescriptor3, []int{0} }

type ScanObjectType int32

const (
	ScanObjectType_CONTAINER  ScanObjectType = 0
	ScanObjectType_HOST       ScanObjectType = 1
	ScanObjectType_IMAGE      ScanObjectType = 2
	ScanObjectType_PLATFORM   ScanObjectType = 3
	ScanObjectType_SERVERLESS ScanObjectType = 4
)

var ScanObjectType_name = map[int32]string{
	0: "CONTAINER",
	1: "HOST",
	2: "IMAGE",
	3: "PLATFORM",
	4: "SERVERLESS",
}
var ScanObjectType_value = map[string]int32{
	"CONTAINER":  0,
	"HOST":       1,
	"IMAGE":      2,
	"PLATFORM":   3,
	"SERVERLESS": 4,
}

func (x ScanObjectType) String() string {
	return proto.EnumName(ScanObjectType_name, int32(x))
}
func (ScanObjectType) EnumDescriptor() ([]byte, []int) { return fileDescriptor3, []int{1} }

type ScanProvider int32

const (
	ScanProvider_Neuvector ScanProvider = 0
	ScanProvider_JFrogXray ScanProvider = 1
)

var ScanProvider_name = map[int32]string{
	0: "Neuvector",
	1: "JFrogXray",
}
var ScanProvider_value = map[string]int32{
	"Neuvector": 0,
	"JFrogXray": 1,
}

func (x ScanProvider) String() string {
	return proto.EnumName(ScanProvider_name, int32(x))
}
func (ScanProvider) EnumDescriptor() ([]byte, []int) { return fileDescriptor3, []int{2} }

type ScanVulStatus int32

const (
	ScanVulStatus_Unpatched  ScanVulStatus = 0
	ScanVulStatus_FixExists  ScanVulStatus = 1
	ScanVulStatus_WillNotFix ScanVulStatus = 2
	ScanVulStatus_Unaffected ScanVulStatus = 3
)

var ScanVulStatus_name = map[int32]string{
	0: "Unpatched",
	1: "FixExists",
	2: "WillNotFix",
	3: "Unaffected",
}
var ScanVulStatus_value = map[string]int32{
	"Unpatched":  0,
	"FixExists":  1,
	"WillNotFix": 2,
	"Unaffected": 3,
}

func (x ScanVulStatus) String() string {
	return proto.EnumName(ScanVulStatus_name, int32(x))
}
func (ScanVulStatus) EnumDescriptor() ([]byte, []int) { return fileDescriptor3, []int{3} }

type ScanVulnerability struct {
	Name                  string   `protobuf:"bytes,1,opt,name=Name" json:"Name,omitempty"`
	Score                 float32  `protobuf:"fixed32,2,opt,name=Score" json:"Score,omitempty"`
	Severity              string   `protobuf:"bytes,3,opt,name=Severity" json:"Severity,omitempty"`
	Description           string   `protobuf:"bytes,4,opt,name=Description" json:"Description,omitempty"`
	PackageNameDeprecated string   `protobuf:"bytes,5,opt,name=PackageNameDeprecated" json:"PackageNameDeprecated,omitempty"`
	PackageVersion        string   `protobuf:"bytes,6,opt,name=PackageVersion" json:"PackageVersion,omitempty"`
	FixedVersion          string   `protobuf:"bytes,7,opt,name=FixedVersion" json:"FixedVersion,omitempty"`
	Link                  string   `protobuf:"bytes,8,opt,name=Link" json:"Link,omitempty"`
	Vectors               string   `protobuf:"bytes,9,opt,name=Vectors" json:"Vectors,omitempty"`
	ScoreV3               float32  `protobuf:"fixed32,10,opt,name=ScoreV3" json:"ScoreV3,omitempty"`
	VectorsV3             string   `protobuf:"bytes,11,opt,name=VectorsV3" json:"VectorsV3,omitempty"`
	PublishedDate         string   `protobuf:"bytes,12,opt,name=PublishedDate" json:"PublishedDate,omitempty"`
	LastModifiedDate      string   `protobuf:"bytes,13,opt,name=LastModifiedDate" json:"LastModifiedDate,omitempty"`
	CPEs                  []string `protobuf:"bytes,14,rep,name=CPEs" json:"CPEs,omitempty"`
	CVEs                  []string `protobuf:"bytes,15,rep,name=CVEs" json:"CVEs,omitempty"`
	FeedRating            string   `protobuf:"bytes,16,opt,name=FeedRating" json:"FeedRating,omitempty"`
	InBase                bool     `protobuf:"varint,17,opt,name=InBase" json:"InBase,omitempty"`
	DBKey                 string   `protobuf:"bytes,18,opt,name=DBKey" json:"DBKey,omitempty"`
	FileName              string   `protobuf:"bytes,19,opt,name=FileName" json:"FileName,omitempty"`
	PackageName           string   `protobuf:"bytes,20,opt,name=PackageName" json:"PackageName,omitempty"`
}

func (m *ScanVulnerability) Reset()                    { *m = ScanVulnerability{} }
func (m *ScanVulnerability) String() string            { return proto.CompactTextString(m) }
func (*ScanVulnerability) ProtoMessage()               {}
func (*ScanVulnerability) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{0} }

func (m *ScanVulnerability) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ScanVulnerability) GetScore() float32 {
	if m != nil {
		return m.Score
	}
	return 0
}

func (m *ScanVulnerability) GetSeverity() string {
	if m != nil {
		return m.Severity
	}
	return ""
}

func (m *ScanVulnerability) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *ScanVulnerability) GetPackageNameDeprecated() string {
	if m != nil {
		return m.PackageNameDeprecated
	}
	return ""
}

func (m *ScanVulnerability) GetPackageVersion() string {
	if m != nil {
		return m.PackageVersion
	}
	return ""
}

func (m *ScanVulnerability) GetFixedVersion() string {
	if m != nil {
		return m.FixedVersion
	}
	return ""
}

func (m *ScanVulnerability) GetLink() string {
	if m != nil {
		return m.Link
	}
	return ""
}

func (m *ScanVulnerability) GetVectors() string {
	if m != nil {
		return m.Vectors
	}
	return ""
}

func (m *ScanVulnerability) GetScoreV3() float32 {
	if m != nil {
		return m.ScoreV3
	}
	return 0
}

func (m *ScanVulnerability) GetVectorsV3() string {
	if m != nil {
		return m.VectorsV3
	}
	return ""
}

func (m *ScanVulnerability) GetPublishedDate() string {
	if m != nil {
		return m.PublishedDate
	}
	return ""
}

func (m *ScanVulnerability) GetLastModifiedDate() string {
	if m != nil {
		return m.LastModifiedDate
	}
	return ""
}

func (m *ScanVulnerability) GetCPEs() []string {
	if m != nil {
		return m.CPEs
	}
	return nil
}

func (m *ScanVulnerability) GetCVEs() []string {
	if m != nil {
		return m.CVEs
	}
	return nil
}

func (m *ScanVulnerability) GetFeedRating() string {
	if m != nil {
		return m.FeedRating
	}
	return ""
}

func (m *ScanVulnerability) GetInBase() bool {
	if m != nil {
		return m.InBase
	}
	return false
}

func (m *ScanVulnerability) GetDBKey() string {
	if m != nil {
		return m.DBKey
	}
	return ""
}

func (m *ScanVulnerability) GetFileName() string {
	if m != nil {
		return m.FileName
	}
	return ""
}

func (m *ScanVulnerability) GetPackageName() string {
	if m != nil {
		return m.PackageName
	}
	return ""
}

type ScanLayerResult struct {
	Digest  string               `protobuf:"bytes,1,opt,name=Digest" json:"Digest,omitempty"`
	Vuls    []*ScanVulnerability `protobuf:"bytes,2,rep,name=Vuls" json:"Vuls,omitempty"`
	Cmds    string               `protobuf:"bytes,3,opt,name=Cmds" json:"Cmds,omitempty"`
	Size    int64                `protobuf:"varint,4,opt,name=Size" json:"Size,omitempty"`
	Secrets *ScanSecretResult    `protobuf:"bytes,5,opt,name=Secrets" json:"Secrets,omitempty"`
}

func (m *ScanLayerResult) Reset()                    { *m = ScanLayerResult{} }
func (m *ScanLayerResult) String() string            { return proto.CompactTextString(m) }
func (*ScanLayerResult) ProtoMessage()               {}
func (*ScanLayerResult) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{1} }

func (m *ScanLayerResult) GetDigest() string {
	if m != nil {
		return m.Digest
	}
	return ""
}

func (m *ScanLayerResult) GetVuls() []*ScanVulnerability {
	if m != nil {
		return m.Vuls
	}
	return nil
}

func (m *ScanLayerResult) GetCmds() string {
	if m != nil {
		return m.Cmds
	}
	return ""
}

func (m *ScanLayerResult) GetSize() int64 {
	if m != nil {
		return m.Size
	}
	return 0
}

func (m *ScanLayerResult) GetSecrets() *ScanSecretResult {
	if m != nil {
		return m.Secrets
	}
	return nil
}

type ScanModule struct {
	Name    string           `protobuf:"bytes,1,opt,name=Name" json:"Name,omitempty"`
	Version string           `protobuf:"bytes,2,opt,name=Version" json:"Version,omitempty"`
	Source  string           `protobuf:"bytes,3,opt,name=Source" json:"Source,omitempty"`
	Vuls    []*ScanModuleVul `protobuf:"bytes,4,rep,name=Vuls" json:"Vuls,omitempty"`
	CPEs    []string         `protobuf:"bytes,5,rep,name=CPEs" json:"CPEs,omitempty"`
}

func (m *ScanModule) Reset()                    { *m = ScanModule{} }
func (m *ScanModule) String() string            { return proto.CompactTextString(m) }
func (*ScanModule) ProtoMessage()               {}
func (*ScanModule) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{2} }

func (m *ScanModule) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ScanModule) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *ScanModule) GetSource() string {
	if m != nil {
		return m.Source
	}
	return ""
}

func (m *ScanModule) GetVuls() []*ScanModuleVul {
	if m != nil {
		return m.Vuls
	}
	return nil
}

func (m *ScanModule) GetCPEs() []string {
	if m != nil {
		return m.CPEs
	}
	return nil
}

type ScanModuleVul struct {
	Name   string        `protobuf:"bytes,1,opt,name=Name" json:"Name,omitempty"`
	Status ScanVulStatus `protobuf:"varint,2,opt,name=Status,enum=share.ScanVulStatus" json:"Status,omitempty"`
}

func (m *ScanModuleVul) Reset()                    { *m = ScanModuleVul{} }
func (m *ScanModuleVul) String() string            { return proto.CompactTextString(m) }
func (*ScanModuleVul) ProtoMessage()               {}
func (*ScanModuleVul) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{3} }

func (m *ScanModuleVul) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ScanModuleVul) GetStatus() ScanVulStatus {
	if m != nil {
		return m.Status
	}
	return ScanVulStatus_Unpatched
}

type ScanSecretLog struct {
	Type       string `protobuf:"bytes,1,opt,name=Type" json:"Type,omitempty"`
	Text       string `protobuf:"bytes,2,opt,name=Text" json:"Text,omitempty"`
	File       string `protobuf:"bytes,3,opt,name=File" json:"File,omitempty"`
	RuleDesc   string `protobuf:"bytes,4,opt,name=RuleDesc" json:"RuleDesc,omitempty"`
	Suggestion string `protobuf:"bytes,5,opt,name=Suggestion" json:"Suggestion,omitempty"`
}

func (m *ScanSecretLog) Reset()                    { *m = ScanSecretLog{} }
func (m *ScanSecretLog) String() string            { return proto.CompactTextString(m) }
func (*ScanSecretLog) ProtoMessage()               {}
func (*ScanSecretLog) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{4} }

func (m *ScanSecretLog) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *ScanSecretLog) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

func (m *ScanSecretLog) GetFile() string {
	if m != nil {
		return m.File
	}
	return ""
}

func (m *ScanSecretLog) GetRuleDesc() string {
	if m != nil {
		return m.RuleDesc
	}
	return ""
}

func (m *ScanSecretLog) GetSuggestion() string {
	if m != nil {
		return m.Suggestion
	}
	return ""
}

type ScanSecretResult struct {
	Error ScanErrorCode    `protobuf:"varint,1,opt,name=Error,enum=share.ScanErrorCode" json:"Error,omitempty"`
	Logs  []*ScanSecretLog `protobuf:"bytes,2,rep,name=Logs" json:"Logs,omitempty"`
}

func (m *ScanSecretResult) Reset()                    { *m = ScanSecretResult{} }
func (m *ScanSecretResult) String() string            { return proto.CompactTextString(m) }
func (*ScanSecretResult) ProtoMessage()               {}
func (*ScanSecretResult) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{5} }

func (m *ScanSecretResult) GetError() ScanErrorCode {
	if m != nil {
		return m.Error
	}
	return ScanErrorCode_ScanErrNone
}

func (m *ScanSecretResult) GetLogs() []*ScanSecretLog {
	if m != nil {
		return m.Logs
	}
	return nil
}

type ScanSetIdPermLog struct {
	Type     string `protobuf:"bytes,1,opt,name=Type" json:"Type,omitempty"`
	File     string `protobuf:"bytes,2,opt,name=File" json:"File,omitempty"`
	Evidence string `protobuf:"bytes,3,opt,name=Evidence" json:"Evidence,omitempty"`
}

func (m *ScanSetIdPermLog) Reset()                    { *m = ScanSetIdPermLog{} }
func (m *ScanSetIdPermLog) String() string            { return proto.CompactTextString(m) }
func (*ScanSetIdPermLog) ProtoMessage()               {}
func (*ScanSetIdPermLog) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{6} }

func (m *ScanSetIdPermLog) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *ScanSetIdPermLog) GetFile() string {
	if m != nil {
		return m.File
	}
	return ""
}

func (m *ScanSetIdPermLog) GetEvidence() string {
	if m != nil {
		return m.Evidence
	}
	return ""
}

type ScanResult struct {
	Version         string               `protobuf:"bytes,1,opt,name=Version" json:"Version,omitempty"`
	Error           ScanErrorCode        `protobuf:"varint,2,opt,name=Error,enum=share.ScanErrorCode" json:"Error,omitempty"`
	Namespace       string               `protobuf:"bytes,3,opt,name=Namespace" json:"Namespace,omitempty"`
	Vuls            []*ScanVulnerability `protobuf:"bytes,4,rep,name=Vuls" json:"Vuls,omitempty"`
	ContainerID     string               `protobuf:"bytes,5,opt,name=ContainerID" json:"ContainerID,omitempty"`
	HostID          string               `protobuf:"bytes,6,opt,name=HostID" json:"HostID,omitempty"`
	Registry        string               `protobuf:"bytes,7,opt,name=Registry" json:"Registry,omitempty"`
	Repository      string               `protobuf:"bytes,8,opt,name=Repository" json:"Repository,omitempty"`
	Tag             string               `protobuf:"bytes,9,opt,name=Tag" json:"Tag,omitempty"`
	Digest          string               `protobuf:"bytes,10,opt,name=Digest" json:"Digest,omitempty"`
	ImageID         string               `protobuf:"bytes,11,opt,name=ImageID" json:"ImageID,omitempty"`
	Layers          []*ScanLayerResult   `protobuf:"bytes,12,rep,name=Layers" json:"Layers,omitempty"`
	Envs            []string             `protobuf:"bytes,13,rep,name=Envs" json:"Envs,omitempty"`
	Labels          map[string]string    `protobuf:"bytes,14,rep,name=Labels" json:"Labels,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Platform        string               `protobuf:"bytes,15,opt,name=Platform" json:"Platform,omitempty"`
	PlatformVersion string               `protobuf:"bytes,16,opt,name=PlatformVersion" json:"PlatformVersion,omitempty"`
	Author          string               `protobuf:"bytes,17,opt,name=Author" json:"Author,omitempty"`
	CVEDBCreateTime string               `protobuf:"bytes,18,opt,name=CVEDBCreateTime" json:"CVEDBCreateTime,omitempty"`
	Modules         []*ScanModule        `protobuf:"bytes,19,rep,name=Modules" json:"Modules,omitempty"`
	Secrets         *ScanSecretResult    `protobuf:"bytes,20,opt,name=Secrets" json:"Secrets,omitempty"`
	Cmds            []string             `protobuf:"bytes,21,rep,name=Cmds" json:"Cmds,omitempty"`
	SetIdPerms      []*ScanSetIdPermLog  `protobuf:"bytes,22,rep,name=SetIdPerms" json:"SetIdPerms,omitempty"`
	Provider        ScanProvider         `protobuf:"varint,23,opt,name=Provider,enum=share.ScanProvider" json:"Provider,omitempty"`
	Size            int64                `protobuf:"varint,24,opt,name=Size" json:"Size,omitempty"`
	SignatureInfo   *ScanSignatureInfo   `protobuf:"bytes,25,opt,name=SignatureInfo" json:"SignatureInfo,omitempty"`
	Created         string               `protobuf:"bytes,26,opt,name=Created" json:"Created,omitempty"`
}

func (m *ScanResult) Reset()                    { *m = ScanResult{} }
func (m *ScanResult) String() string            { return proto.CompactTextString(m) }
func (*ScanResult) ProtoMessage()               {}
func (*ScanResult) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{7} }

func (m *ScanResult) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *ScanResult) GetError() ScanErrorCode {
	if m != nil {
		return m.Error
	}
	return ScanErrorCode_ScanErrNone
}

func (m *ScanResult) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *ScanResult) GetVuls() []*ScanVulnerability {
	if m != nil {
		return m.Vuls
	}
	return nil
}

func (m *ScanResult) GetContainerID() string {
	if m != nil {
		return m.ContainerID
	}
	return ""
}

func (m *ScanResult) GetHostID() string {
	if m != nil {
		return m.HostID
	}
	return ""
}

func (m *ScanResult) GetRegistry() string {
	if m != nil {
		return m.Registry
	}
	return ""
}

func (m *ScanResult) GetRepository() string {
	if m != nil {
		return m.Repository
	}
	return ""
}

func (m *ScanResult) GetTag() string {
	if m != nil {
		return m.Tag
	}
	return ""
}

func (m *ScanResult) GetDigest() string {
	if m != nil {
		return m.Digest
	}
	return ""
}

func (m *ScanResult) GetImageID() string {
	if m != nil {
		return m.ImageID
	}
	return ""
}

func (m *ScanResult) GetLayers() []*ScanLayerResult {
	if m != nil {
		return m.Layers
	}
	return nil
}

func (m *ScanResult) GetEnvs() []string {
	if m != nil {
		return m.Envs
	}
	return nil
}

func (m *ScanResult) GetLabels() map[string]string {
	if m != nil {
		return m.Labels
	}
	return nil
}

func (m *ScanResult) GetPlatform() string {
	if m != nil {
		return m.Platform
	}
	return ""
}

func (m *ScanResult) GetPlatformVersion() string {
	if m != nil {
		return m.PlatformVersion
	}
	return ""
}

func (m *ScanResult) GetAuthor() string {
	if m != nil {
		return m.Author
	}
	return ""
}

func (m *ScanResult) GetCVEDBCreateTime() string {
	if m != nil {
		return m.CVEDBCreateTime
	}
	return ""
}

func (m *ScanResult) GetModules() []*ScanModule {
	if m != nil {
		return m.Modules
	}
	return nil
}

func (m *ScanResult) GetSecrets() *ScanSecretResult {
	if m != nil {
		return m.Secrets
	}
	return nil
}

func (m *ScanResult) GetCmds() []string {
	if m != nil {
		return m.Cmds
	}
	return nil
}

func (m *ScanResult) GetSetIdPerms() []*ScanSetIdPermLog {
	if m != nil {
		return m.SetIdPerms
	}
	return nil
}

func (m *ScanResult) GetProvider() ScanProvider {
	if m != nil {
		return m.Provider
	}
	return ScanProvider_Neuvector
}

func (m *ScanResult) GetSize() int64 {
	if m != nil {
		return m.Size
	}
	return 0
}

func (m *ScanResult) GetSignatureInfo() *ScanSignatureInfo {
	if m != nil {
		return m.SignatureInfo
	}
	return nil
}

func (m *ScanResult) GetCreated() string {
	if m != nil {
		return m.Created
	}
	return ""
}

type ScanSignatureInfo struct {
	Verifiers             []string      `protobuf:"bytes,1,rep,name=Verifiers" json:"Verifiers,omitempty"`
	VerificationTimestamp string        `protobuf:"bytes,2,opt,name=VerificationTimestamp" json:"VerificationTimestamp,omitempty"`
	VerificationError     ScanErrorCode `protobuf:"varint,3,opt,name=VerificationError,enum=share.ScanErrorCode" json:"VerificationError,omitempty"`
}

func (m *ScanSignatureInfo) Reset()                    { *m = ScanSignatureInfo{} }
func (m *ScanSignatureInfo) String() string            { return proto.CompactTextString(m) }
func (*ScanSignatureInfo) ProtoMessage()               {}
func (*ScanSignatureInfo) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{8} }

func (m *ScanSignatureInfo) GetVerifiers() []string {
	if m != nil {
		return m.Verifiers
	}
	return nil
}

func (m *ScanSignatureInfo) GetVerificationTimestamp() string {
	if m != nil {
		return m.VerificationTimestamp
	}
	return ""
}

func (m *ScanSignatureInfo) GetVerificationError() ScanErrorCode {
	if m != nil {
		return m.VerificationError
	}
	return ScanErrorCode_ScanErrNone
}

type ScanRunningRequest struct {
	Type             ScanObjectType `protobuf:"varint,1,opt,name=Type,enum=share.ScanObjectType" json:"Type,omitempty"`
	ID               string         `protobuf:"bytes,2,opt,name=ID" json:"ID,omitempty"`
	AgentID          string         `protobuf:"bytes,3,opt,name=AgentID" json:"AgentID,omitempty"`
	AgentRPCEndPoint string         `protobuf:"bytes,4,opt,name=AgentRPCEndPoint" json:"AgentRPCEndPoint,omitempty"`
}

func (m *ScanRunningRequest) Reset()                    { *m = ScanRunningRequest{} }
func (m *ScanRunningRequest) String() string            { return proto.CompactTextString(m) }
func (*ScanRunningRequest) ProtoMessage()               {}
func (*ScanRunningRequest) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{9} }

func (m *ScanRunningRequest) GetType() ScanObjectType {
	if m != nil {
		return m.Type
	}
	return ScanObjectType_CONTAINER
}

func (m *ScanRunningRequest) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *ScanRunningRequest) GetAgentID() string {
	if m != nil {
		return m.AgentID
	}
	return ""
}

func (m *ScanRunningRequest) GetAgentRPCEndPoint() string {
	if m != nil {
		return m.AgentRPCEndPoint
	}
	return ""
}

type ScanData struct {
	Error  ScanErrorCode `protobuf:"varint,1,opt,name=Error,enum=share.ScanErrorCode" json:"Error,omitempty"`
	Buffer []byte        `protobuf:"bytes,2,opt,name=Buffer,proto3" json:"Buffer,omitempty"`
}

func (m *ScanData) Reset()                    { *m = ScanData{} }
func (m *ScanData) String() string            { return proto.CompactTextString(m) }
func (*ScanData) ProtoMessage()               {}
func (*ScanData) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{10} }

func (m *ScanData) GetError() ScanErrorCode {
	if m != nil {
		return m.Error
	}
	return ScanErrorCode_ScanErrNone
}

func (m *ScanData) GetBuffer() []byte {
	if m != nil {
		return m.Buffer
	}
	return nil
}

type ScanAppPackage struct {
	AppName    string `protobuf:"bytes,1,opt,name=AppName" json:"AppName,omitempty"`
	ModuleName string `protobuf:"bytes,2,opt,name=ModuleName" json:"ModuleName,omitempty"`
	Version    string `protobuf:"bytes,3,opt,name=Version" json:"Version,omitempty"`
	FileName   string `protobuf:"bytes,4,opt,name=FileName" json:"FileName,omitempty"`
}

func (m *ScanAppPackage) Reset()                    { *m = ScanAppPackage{} }
func (m *ScanAppPackage) String() string            { return proto.CompactTextString(m) }
func (*ScanAppPackage) ProtoMessage()               {}
func (*ScanAppPackage) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{11} }

func (m *ScanAppPackage) GetAppName() string {
	if m != nil {
		return m.AppName
	}
	return ""
}

func (m *ScanAppPackage) GetModuleName() string {
	if m != nil {
		return m.ModuleName
	}
	return ""
}

func (m *ScanAppPackage) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *ScanAppPackage) GetFileName() string {
	if m != nil {
		return m.FileName
	}
	return ""
}

type ScanAppRequest struct {
	Packages []*ScanAppPackage `protobuf:"bytes,1,rep,name=Packages" json:"Packages,omitempty"`
}

func (m *ScanAppRequest) Reset()                    { *m = ScanAppRequest{} }
func (m *ScanAppRequest) String() string            { return proto.CompactTextString(m) }
func (*ScanAppRequest) ProtoMessage()               {}
func (*ScanAppRequest) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{12} }

func (m *ScanAppRequest) GetPackages() []*ScanAppPackage {
	if m != nil {
		return m.Packages
	}
	return nil
}

type ScanAwsLambdaRequest struct {
	ResType     string `protobuf:"bytes,1,opt,name=ResType" json:"ResType,omitempty"`
	FuncName    string `protobuf:"bytes,2,opt,name=FuncName" json:"FuncName,omitempty"`
	Region      string `protobuf:"bytes,3,opt,name=Region" json:"Region,omitempty"`
	FuncLink    string `protobuf:"bytes,4,opt,name=FuncLink" json:"FuncLink,omitempty"`
	ScanSecrets bool   `protobuf:"varint,5,opt,name=ScanSecrets" json:"ScanSecrets,omitempty"`
}

func (m *ScanAwsLambdaRequest) Reset()                    { *m = ScanAwsLambdaRequest{} }
func (m *ScanAwsLambdaRequest) String() string            { return proto.CompactTextString(m) }
func (*ScanAwsLambdaRequest) ProtoMessage()               {}
func (*ScanAwsLambdaRequest) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{13} }

func (m *ScanAwsLambdaRequest) GetResType() string {
	if m != nil {
		return m.ResType
	}
	return ""
}

func (m *ScanAwsLambdaRequest) GetFuncName() string {
	if m != nil {
		return m.FuncName
	}
	return ""
}

func (m *ScanAwsLambdaRequest) GetRegion() string {
	if m != nil {
		return m.Region
	}
	return ""
}

func (m *ScanAwsLambdaRequest) GetFuncLink() string {
	if m != nil {
		return m.FuncLink
	}
	return ""
}

func (m *ScanAwsLambdaRequest) GetScanSecrets() bool {
	if m != nil {
		return m.ScanSecrets
	}
	return false
}

func init() {
	proto.RegisterType((*ScanVulnerability)(nil), "share.ScanVulnerability")
	proto.RegisterType((*ScanLayerResult)(nil), "share.ScanLayerResult")
	proto.RegisterType((*ScanModule)(nil), "share.ScanModule")
	proto.RegisterType((*ScanModuleVul)(nil), "share.ScanModuleVul")
	proto.RegisterType((*ScanSecretLog)(nil), "share.ScanSecretLog")
	proto.RegisterType((*ScanSecretResult)(nil), "share.ScanSecretResult")
	proto.RegisterType((*ScanSetIdPermLog)(nil), "share.ScanSetIdPermLog")
	proto.RegisterType((*ScanResult)(nil), "share.ScanResult")
	proto.RegisterType((*ScanSignatureInfo)(nil), "share.ScanSignatureInfo")
	proto.RegisterType((*ScanRunningRequest)(nil), "share.ScanRunningRequest")
	proto.RegisterType((*ScanData)(nil), "share.ScanData")
	proto.RegisterType((*ScanAppPackage)(nil), "share.ScanAppPackage")
	proto.RegisterType((*ScanAppRequest)(nil), "share.ScanAppRequest")
	proto.RegisterType((*ScanAwsLambdaRequest)(nil), "share.ScanAwsLambdaRequest")
	proto.RegisterEnum("share.ScanErrorCode", ScanErrorCode_name, ScanErrorCode_value)
	proto.RegisterEnum("share.ScanObjectType", ScanObjectType_name, ScanObjectType_value)
	proto.RegisterEnum("share.ScanProvider", ScanProvider_name, ScanProvider_value)
	proto.RegisterEnum("share.ScanVulStatus", ScanVulStatus_name, ScanVulStatus_value)
}

func init() { proto.RegisterFile("scan.proto", fileDescriptor3) }

var fileDescriptor3 = []byte{
	// 1649 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x57, 0xcd, 0x73, 0x23, 0x39,
	0x15, 0x9f, 0xb6, 0x9d, 0xc4, 0x91, 0xf3, 0xa1, 0x28, 0x1f, 0xa3, 0x9d, 0x5a, 0xb6, 0x52, 0x2e,
	0x8a, 0x0a, 0x61, 0x2a, 0xd4, 0xce, 0x40, 0xf1, 0x71, 0xa0, 0x2a, 0xb1, 0x3b, 0xac, 0xc1, 0x49,
	0x4c, 0x3b, 0x6b, 0xb8, 0x2a, 0xee, 0x97, 0x4e, 0x33, 0xed, 0x6e, 0x23, 0xa9, 0x33, 0x31, 0x27,
	0x6e, 0x5c, 0xb9, 0xf0, 0x0f, 0x70, 0xe5, 0xc6, 0x7f, 0x04, 0x37, 0xfe, 0x0b, 0xea, 0x49, 0xea,
	0xb6, 0x3c, 0x49, 0x2d, 0xbb, 0x37, 0xbd, 0x8f, 0x7e, 0x7a, 0x3f, 0xbd, 0xf7, 0x7e, 0x52, 0x13,
	0xa2, 0xa6, 0x22, 0x3f, 0x9b, 0xcb, 0x42, 0x17, 0x6c, 0x4d, 0x3d, 0x08, 0x09, 0xdd, 0xff, 0xb6,
	0xc8, 0xde, 0x78, 0x2a, 0xf2, 0x49, 0x99, 0xe5, 0x20, 0xc5, 0x5d, 0x9a, 0xa5, 0x7a, 0xc1, 0x18,
	0x69, 0x5d, 0x8b, 0x19, 0xf0, 0xe0, 0x38, 0x38, 0xd9, 0x8c, 0xcc, 0x9a, 0x1d, 0x90, 0xb5, 0xf1,
	0xb4, 0x90, 0xc0, 0x1b, 0xc7, 0xc1, 0x49, 0x23, 0xb2, 0x02, 0x7b, 0x43, 0xda, 0x63, 0x78, 0x04,
	0x99, 0xea, 0x05, 0x6f, 0x1a, 0xef, 0x5a, 0x66, 0xc7, 0xa4, 0xd3, 0x07, 0x35, 0x95, 0xe9, 0x5c,
	0xa7, 0x45, 0xce, 0x5b, 0xc6, 0xec, 0xab, 0xd8, 0x4f, 0xc8, 0xe1, 0x48, 0x4c, 0x3f, 0x88, 0x04,
	0x70, 0x8b, 0x3e, 0xcc, 0x25, 0x4c, 0x85, 0x86, 0x98, 0xaf, 0x19, 0xdf, 0x97, 0x8d, 0xec, 0x07,
	0x64, 0xc7, 0x19, 0x26, 0x20, 0x15, 0x86, 0x5e, 0x37, 0xee, 0x9f, 0x68, 0x59, 0x97, 0x6c, 0x5d,
	0xa6, 0x4f, 0x10, 0x57, 0x5e, 0x1b, 0xc6, 0x6b, 0x45, 0x87, 0x48, 0x87, 0x69, 0xfe, 0x81, 0xb7,
	0x2d, 0x52, 0x5c, 0x33, 0x4e, 0x36, 0x26, 0x30, 0xd5, 0x85, 0x54, 0x7c, 0xd3, 0xa8, 0x2b, 0x11,
	0x2d, 0x06, 0xf6, 0xe4, 0x3d, 0x27, 0xe6, 0x14, 0x2a, 0x91, 0x7d, 0x4e, 0x36, 0x9d, 0xd3, 0xe4,
	0x3d, 0xef, 0x98, 0xaf, 0x96, 0x0a, 0xf6, 0x7d, 0xb2, 0x3d, 0x2a, 0xef, 0xb2, 0x54, 0x3d, 0x40,
	0xdc, 0x17, 0x1a, 0xf8, 0x96, 0xf1, 0x58, 0x55, 0xb2, 0x53, 0x42, 0x87, 0x42, 0xe9, 0xab, 0x22,
	0x4e, 0xef, 0x53, 0xe7, 0xb8, 0x6d, 0x1c, 0x9f, 0xe9, 0x31, 0xef, 0xde, 0x28, 0x54, 0x7c, 0xe7,
	0xb8, 0x89, 0x79, 0xe3, 0xda, 0xe8, 0x26, 0xa1, 0xe2, 0xbb, 0x4e, 0x37, 0x09, 0x15, 0xfb, 0x82,
	0x90, 0x4b, 0x80, 0x38, 0x12, 0x3a, 0xcd, 0x13, 0x4e, 0x4d, 0x34, 0x4f, 0xc3, 0x8e, 0xc8, 0xfa,
	0x20, 0xbf, 0x10, 0x0a, 0xf8, 0xde, 0x71, 0x70, 0xd2, 0x8e, 0x9c, 0x84, 0xd5, 0xee, 0x5f, 0xfc,
	0x16, 0x16, 0x9c, 0x99, 0x4f, 0xac, 0x80, 0xd5, 0xbe, 0x4c, 0x33, 0x53, 0x0f, 0xbe, 0x6f, 0xab,
	0x5d, 0xc9, 0x58, 0x6d, 0xaf, 0x5c, 0xfc, 0xc0, 0x56, 0xdb, 0x53, 0x75, 0xff, 0x15, 0x90, 0x5d,
	0xec, 0xb5, 0xa1, 0x58, 0x80, 0x8c, 0x40, 0x95, 0x99, 0xc6, 0xfd, 0xfb, 0x69, 0x02, 0x4a, 0xbb,
	0x5e, 0x73, 0x12, 0x7b, 0x4b, 0x5a, 0x93, 0x32, 0x53, 0xbc, 0x71, 0xdc, 0x3c, 0xe9, 0xbc, 0xe3,
	0x67, 0xa6, 0x5b, 0xcf, 0x9e, 0x75, 0x6a, 0x64, 0xbc, 0x0c, 0xf2, 0x59, 0xac, 0x5c, 0x07, 0x9a,
	0x35, 0xea, 0xc6, 0xe9, 0x9f, 0xc1, 0xb4, 0x5d, 0x33, 0x32, 0x6b, 0xf6, 0x25, 0xd9, 0x18, 0xc3,
	0x54, 0x82, 0x56, 0xa6, 0xc3, 0x3a, 0xef, 0x5e, 0x7b, 0x81, 0xad, 0xc5, 0xe6, 0x15, 0x55, 0x7e,
	0xdd, 0xbf, 0x05, 0x84, 0xa0, 0xf5, 0xaa, 0x88, 0xcb, 0x0c, 0x5e, 0x9c, 0x0c, 0xd3, 0x2f, 0xb6,
	0xc5, 0x1a, 0x55, 0xbf, 0xd8, 0xee, 0x3a, 0x22, 0xeb, 0xe3, 0xa2, 0x94, 0x53, 0x70, 0x99, 0x39,
	0x89, 0x9d, 0x38, 0x74, 0x2d, 0x83, 0xee, 0xc0, 0x4b, 0xc2, 0x6e, 0x33, 0x29, 0x33, 0x0f, 0x19,
	0xd6, 0x79, 0x6d, 0x59, 0xe7, 0xee, 0xef, 0xc8, 0xf6, 0x8a, 0xeb, 0x8b, 0x49, 0xbd, 0x25, 0xeb,
	0x63, 0x2d, 0x74, 0xa9, 0x4c, 0x4e, 0x3b, 0x2b, 0x9b, 0x4c, 0xca, 0xcc, 0xda, 0x22, 0xe7, 0xd3,
	0xfd, 0x6b, 0x60, 0x63, 0x5a, 0xd4, 0xc3, 0x22, 0xc1, 0x98, 0xb7, 0x8b, 0x79, 0x1d, 0x13, 0xd7,
	0x46, 0x07, 0x4f, 0xda, 0xa1, 0x34, 0x6b, 0xd4, 0x61, 0x0b, 0x54, 0x47, 0x8f, 0x6b, 0x6c, 0x93,
	0xa8, 0xcc, 0x00, 0x27, 0xdd, 0x4d, 0x7d, 0x2d, 0x63, 0x43, 0x8e, 0xcb, 0x04, 0x6b, 0x8c, 0xe7,
	0x65, 0xe7, 0xdc, 0xd3, 0x74, 0x1f, 0x08, 0xfd, 0xb4, 0x18, 0xec, 0x94, 0xac, 0x85, 0x52, 0x16,
	0xd2, 0x24, 0xb3, 0x0a, 0xc5, 0xe8, 0x7b, 0x45, 0x0c, 0x91, 0x75, 0xc1, 0xa3, 0x1d, 0x16, 0x49,
	0xd5, 0x38, 0x07, 0xcf, 0xea, 0x3b, 0x2c, 0x92, 0xc8, 0x78, 0x74, 0x27, 0xd5, 0x4e, 0x7a, 0x10,
	0x8f, 0x40, 0xce, 0xbe, 0x01, 0xb5, 0x41, 0xd8, 0x58, 0x45, 0x18, 0x3e, 0xa6, 0x31, 0xe4, 0x75,
	0x69, 0x6b, 0xb9, 0xfb, 0x9f, 0x0d, 0xdb, 0x31, 0x2e, 0x79, 0xaf, 0x3b, 0x82, 0xd5, 0xee, 0xa8,
	0x61, 0x35, 0xfe, 0x3f, 0xac, 0xcf, 0xc9, 0x26, 0x96, 0x55, 0xcd, 0x45, 0xbd, 0xe3, 0x52, 0x51,
	0x4f, 0x4b, 0xeb, 0x5b, 0x4d, 0xcb, 0x31, 0xe9, 0xf4, 0x8a, 0x5c, 0x8b, 0x34, 0x07, 0x39, 0xe8,
	0xbb, 0x1a, 0xf8, 0x2a, 0xec, 0xdb, 0xaf, 0x0a, 0xa5, 0x07, 0x7d, 0xc7, 0xac, 0x4e, 0x32, 0x85,
	0x85, 0x24, 0x55, 0x5a, 0x2e, 0x1c, 0x9b, 0xd6, 0x32, 0x16, 0x36, 0x82, 0x79, 0xa1, 0x52, 0x5d,
	0xc8, 0x85, 0xe3, 0x53, 0x4f, 0xc3, 0x28, 0x69, 0xde, 0x8a, 0xc4, 0x31, 0x2a, 0x2e, 0xbd, 0xd9,
	0x27, 0x2b, 0xb3, 0xcf, 0xc9, 0xc6, 0x60, 0x26, 0x12, 0x18, 0xf4, 0x1d, 0x93, 0x56, 0x22, 0x3b,
	0x23, 0xeb, 0x86, 0x3c, 0x14, 0xdf, 0x32, 0x48, 0x8f, 0x3c, 0xa4, 0x1e, 0xab, 0x44, 0xce, 0x0b,
	0x4b, 0x17, 0xe6, 0x8f, 0x8a, 0x6f, 0xdb, 0xe9, 0xc1, 0x35, 0xfb, 0x29, 0xc6, 0xb8, 0x83, 0xcc,
	0x72, 0x67, 0xe7, 0xdd, 0xf7, 0xbc, 0x18, 0xf6, 0xf3, 0x33, 0x6b, 0x0f, 0x73, 0x2d, 0x17, 0x91,
	0x73, 0x46, 0xe8, 0xa3, 0x4c, 0xe8, 0xfb, 0x42, 0xce, 0xf8, 0xae, 0x85, 0x5e, 0xc9, 0xec, 0x84,
	0xec, 0x56, 0xeb, 0xaa, 0xd4, 0x96, 0x69, 0x3f, 0x55, 0x23, 0xe4, 0xf3, 0x52, 0x3f, 0x14, 0xd2,
	0xd0, 0xed, 0x66, 0xe4, 0x24, 0x8c, 0xd0, 0x9b, 0x84, 0xfd, 0x8b, 0x9e, 0x04, 0xa1, 0xe1, 0x36,
	0x9d, 0x81, 0x23, 0xde, 0x4f, 0xd5, 0xec, 0x47, 0x64, 0xc3, 0x0e, 0xbe, 0xe2, 0xfb, 0x26, 0xff,
	0xbd, 0x67, 0xec, 0x11, 0x55, 0x1e, 0x3e, 0xdf, 0x1d, 0x7c, 0x3b, 0xbe, 0xab, 0xa9, 0xf4, 0xd0,
	0x11, 0x0e, 0x52, 0xe9, 0xcf, 0x08, 0xa9, 0xa7, 0x44, 0xf1, 0x23, 0xb3, 0xed, 0x6a, 0xa4, 0xe5,
	0x08, 0x45, 0x9e, 0x2b, 0xfb, 0x31, 0x69, 0x8f, 0x64, 0x81, 0x83, 0x21, 0xf9, 0x6b, 0xd3, 0xe4,
	0xfb, 0xde, 0x67, 0x95, 0x29, 0xaa, 0x9d, 0x6a, 0xd2, 0xe6, 0x1e, 0x69, 0xff, 0x8a, 0x6c, 0x8f,
	0xd3, 0x24, 0x17, 0xba, 0x94, 0x30, 0xc8, 0xef, 0x0b, 0xfe, 0x99, 0x81, 0xe2, 0x77, 0xf9, 0x8a,
	0x3d, 0x5a, 0x75, 0xc7, 0x76, 0xb2, 0xe7, 0x17, 0xf3, 0x37, 0xb6, 0x9d, 0x9c, 0xf8, 0xe6, 0x17,
	0xa4, 0xe3, 0x95, 0x1a, 0x3b, 0xf4, 0x03, 0x2c, 0xdc, 0x94, 0xe2, 0x12, 0x6f, 0xc1, 0x47, 0x91,
	0x95, 0xd5, 0xec, 0x5b, 0xe1, 0x97, 0x8d, 0x9f, 0x07, 0xdd, 0x7f, 0x06, 0xf6, 0xdd, 0xb4, 0xba,
	0x95, 0x79, 0x05, 0x48, 0xbc, 0xa5, 0xa5, 0xe2, 0x81, 0x39, 0xc1, 0xa5, 0x02, 0x5f, 0x3b, 0x56,
	0x98, 0x0a, 0xa4, 0x3a, 0x2c, 0xa7, 0xd2, 0x62, 0x36, 0x77, 0xd1, 0x5f, 0x36, 0xb2, 0x0b, 0xb2,
	0xe7, 0x1b, 0x2c, 0x63, 0x34, 0xbf, 0x81, 0x31, 0x9e, 0xbb, 0x77, 0xff, 0x1e, 0x10, 0x66, 0xfa,
	0xbb, 0xcc, 0xf3, 0x34, 0x4f, 0x22, 0xf8, 0x53, 0x89, 0x83, 0xf6, 0x43, 0x8f, 0xed, 0x76, 0xde,
	0x1d, 0x7a, 0xd1, 0x6e, 0xee, 0xfe, 0x08, 0x53, 0x8d, 0x46, 0x47, 0x82, 0x3b, 0xa4, 0x31, 0xe8,
	0xbb, 0x44, 0x1b, 0x83, 0x3e, 0x1e, 0xea, 0x79, 0x02, 0x39, 0x52, 0x84, 0x65, 0xa3, 0x4a, 0xc4,
	0x57, 0x8c, 0x59, 0x46, 0xa3, 0x5e, 0x98, 0xc7, 0xa3, 0x22, 0xcd, 0xb5, 0xbb, 0x04, 0x9e, 0xe9,
	0xbb, 0xd7, 0xa4, 0x8d, 0xbb, 0xf5, 0x85, 0x16, 0xdf, 0x89, 0xe4, 0x8f, 0xc8, 0xfa, 0x45, 0x79,
	0x7f, 0x0f, 0x96, 0x3a, 0xb7, 0x22, 0x27, 0x75, 0xff, 0x12, 0x90, 0x1d, 0xfc, 0xe0, 0x7c, 0x3e,
	0x77, 0x0f, 0x0f, 0x93, 0xe8, 0x7c, 0xee, 0x5d, 0x8f, 0x95, 0x88, 0x84, 0x65, 0xe7, 0xc4, 0x18,
	0x2d, 0x34, 0x4f, 0xe3, 0x13, 0x77, 0x73, 0x95, 0xb8, 0xfd, 0x67, 0x50, 0x6b, 0xf5, 0x19, 0xd4,
	0xed, 0xd5, 0x19, 0x54, 0xa7, 0xfc, 0x25, 0x69, 0xbb, 0x64, 0x6c, 0x4f, 0x74, 0x56, 0x4e, 0x7a,
	0x99, 0x6a, 0x54, 0xbb, 0x75, 0xff, 0x11, 0x90, 0x03, 0x63, 0xfc, 0xa8, 0x86, 0x62, 0x76, 0x17,
	0x8b, 0x2a, 0x16, 0x27, 0x1b, 0x11, 0x28, 0xef, 0x8a, 0xaa, 0x44, 0x93, 0x53, 0x99, 0x4f, 0x3d,
	0x2c, 0xb5, 0x8c, 0xc7, 0x85, 0x34, 0x5d, 0x03, 0x71, 0x52, 0xf5, 0x8d, 0x79, 0x00, 0xb7, 0x96,
	0xdf, 0x98, 0x47, 0xf0, 0x31, 0xe9, 0x2c, 0x49, 0xc2, 0x3e, 0x97, 0xda, 0x91, 0xaf, 0x3a, 0xfd,
	0x77, 0xd3, 0xbe, 0x19, 0xea, 0xea, 0xb0, 0x5d, 0xfb, 0x4d, 0x28, 0xe5, 0x75, 0x91, 0x03, 0x7d,
	0xc5, 0x98, 0x3d, 0x0c, 0x54, 0x80, 0xfe, 0x58, 0xc8, 0x0f, 0x34, 0x60, 0x87, 0x76, 0x70, 0x8c,
	0x93, 0x1e, 0x97, 0xf3, 0x79, 0x21, 0x35, 0x6d, 0x30, 0x6e, 0x11, 0x87, 0x52, 0xe2, 0xd0, 0xdf,
	0x3c, 0x82, 0x1c, 0xa6, 0xb3, 0x54, 0xd3, 0xa6, 0x17, 0xc4, 0x9d, 0x0f, 0x6d, 0xb1, 0x7d, 0xfb,
	0x92, 0x0c, 0xa5, 0xc4, 0xde, 0xb9, 0x13, 0x0a, 0xe8, 0x9a, 0xe7, 0x88, 0xd3, 0x53, 0x94, 0x9a,
	0xae, 0x7b, 0xbb, 0x0d, 0x90, 0x6f, 0x12, 0x09, 0x4a, 0xd1, 0x0d, 0x76, 0x64, 0xe7, 0x21, 0x94,
	0xb2, 0xba, 0xbf, 0xce, 0x47, 0x03, 0xda, 0xf6, 0xdc, 0xb1, 0xa0, 0xe3, 0x85, 0xd2, 0x30, 0xa3,
	0x9b, 0xec, 0x35, 0xd9, 0x77, 0xea, 0xfa, 0x96, 0x44, 0x7f, 0xe2, 0x6d, 0xf9, 0x07, 0x29, 0x4c,
	0x8c, 0x8e, 0x87, 0xa4, 0x76, 0x0e, 0x9f, 0x52, 0x4d, 0xb7, 0xd8, 0x67, 0xe4, 0xd0, 0x59, 0x90,
	0xf6, 0x21, 0xd7, 0x6e, 0x48, 0xe9, 0xb6, 0x97, 0x50, 0x0f, 0xa4, 0xb6, 0xe3, 0x0b, 0x74, 0xc7,
	0x03, 0xda, 0x13, 0xf9, 0x14, 0x32, 0x88, 0xe9, 0x2e, 0xfb, 0x82, 0xbc, 0xa9, 0xd0, 0xcb, 0xf4,
	0xd1, 0xe4, 0xe2, 0x9d, 0x25, 0xf5, 0x32, 0x30, 0x17, 0xe7, 0x75, 0xa1, 0x2f, 0x8b, 0x32, 0x8f,
	0xe9, 0x9e, 0x9f, 0xc1, 0x47, 0xd5, 0x2f, 0x3e, 0xe6, 0x59, 0x21, 0xe2, 0x50, 0x4a, 0xca, 0xbc,
	0x9d, 0xce, 0x65, 0x52, 0xce, 0x20, 0xd7, 0x74, 0xff, 0x34, 0xb2, 0xf8, 0x96, 0x74, 0xc0, 0xb6,
	0xc9, 0x66, 0xef, 0xe6, 0xfa, 0xf6, 0x7c, 0x70, 0x1d, 0x46, 0xf4, 0x15, 0x6b, 0x93, 0xd6, 0x57,
	0x37, 0xe3, 0x5b, 0x1a, 0xb0, 0x4d, 0xb2, 0x36, 0xb8, 0x3a, 0xff, 0x75, 0x48, 0x1b, 0x6c, 0x8b,
	0xb4, 0x47, 0xc3, 0xf3, 0xdb, 0xcb, 0x9b, 0xe8, 0x8a, 0x36, 0xd9, 0x0e, 0x21, 0xe3, 0x30, 0x9a,
	0x84, 0xd1, 0x30, 0x1c, 0x8f, 0x69, 0xeb, 0xf4, 0x2d, 0xd9, 0xf2, 0xd9, 0x1f, 0x23, 0x5e, 0x43,
	0xf9, 0x68, 0x7e, 0x96, 0xe8, 0x2b, 0x14, 0x7f, 0x73, 0x29, 0x8b, 0x04, 0x0f, 0x94, 0x06, 0xa7,
	0x57, 0xb6, 0xc9, 0xea, 0x27, 0x2b, 0xda, 0xbf, 0xce, 0xe7, 0x42, 0x4f, 0x1f, 0x20, 0xb6, 0xee,
	0x97, 0xe9, 0x53, 0xf8, 0x94, 0x2a, 0xad, 0x68, 0x80, 0x9b, 0xfd, 0x3e, 0xcd, 0x32, 0x84, 0x9c,
	0x3e, 0xd1, 0x06, 0xca, 0x5f, 0xe7, 0xe2, 0xfe, 0x1e, 0xa6, 0x1a, 0x62, 0xda, 0xbc, 0x5b, 0x37,
	0x7f, 0xbf, 0xef, 0xff, 0x17, 0x00, 0x00, 0xff, 0xff, 0xa0, 0x57, 0x79, 0x2e, 0x0b, 0x0f, 0x00,
	0x00,
}
