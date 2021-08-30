package policy

import (
	"github.com/paloaltonetworks/prisma-cloud-compute-go/collection"
)

type Rule struct {
	AntiMalware              AntiMalware             `json:"antiMalware,omitempty"`
	FileIntegrityRules       []FileIntegrityRule    `json:"fileIntegrityRules,omitempty"`
	Forensic                 Forensic                `json:"forensic,omitempty"`
	LogInspectionRules       []LogInspectionRules    `json:"logInspectionRules,omitempty"`
	Action                   []string                `json:"action,omitempty"`
	AlertThreshold           Threshold               `json:"alertThreshold,omitempty"`
	AllCompliance            bool                    `json:"allCompliance"`
	AuditAllowed             bool                    `json:"auditAllowed"`
	BlockMsg                 string                  `json:"blockMsg,omitempty"`
	BlockThreshold           Threshold               `json:"blockThreshold,omitempty"`
	Condition                Condition               `json:"condition,omitempty"`
	CveRules                 []CveRule               `json:"cveRules,omitempty"`
	Effect                   string                  `json:"effect,omitempty"`
	GraceDays                int                     `json:"graceDays,omitempty"`
	Group                    []string                `json:"group,omitempty"`
	License                  License                 `json:"license,omitempty"`
	OnlyFixed                bool                    `json:"onlyFixed"`
	Principal                []string                `json:"principal,omitempty"`
	Tags                     []Tag                   `json:"tags,omitempty"`
	Verbose                  bool                    `json:"verbose"`
	AdvancedProtection       bool                    `json:"advancedProtection"`
	CloudMetadataEnforcement bool                    `json:"cloudMetadataEnforcement"`
	Collections              []collection.Collection `json:"collections,omitempty"`
	CustomRules              []CustomRule            `json:"customRules,omitempty"`
	Disabled                 bool                    `json:"disabled"`
	Dns                      Dns                     `json:"dns,omitempty"`
	Filesystem               Filesystem              `json:"filesystem,omitempty"`
	KubernetesEnforcement    bool                    `json:"kubernetesEnforcement"`
	Modified                 string                  `json:"modified,omitempty"`
	Name                     string                  `json:"name,omitempty"`
	Network                  Network                 `json:"network,omitempty"`
	Notes                    string                  `json:"notes,omitempty"`
	Owner                    string                  `json:"owner,omitempty"`
	PreviousName             string                  `json:"previousName,omitempty"`
	Processes                Processes               `json:"processes,omitempty"`
	WildFireAnalysis         string                  `json:"wildFireAnalysis,omitempty"`
}

type FileIntegrityRule struct {
	Dir           bool     `json:"dir"`
	Exclusions    []string `json:"exclusions,omitempty"`
	Metadata      bool     `json:"metadata"`
	Path          string   `json:"path,omitempty"`
	ProcWhitelist []string `json:"procWhitelist,omitempty"`
	Read          bool     `json:"read"`
	Recursive     bool     `json:"recursive"`
	Write         bool     `json:"write"`
}

type Forensic struct {
	ActivitiesDisabled       bool `json:"activitiesDisabled"`
	DockerEnabled            bool `json:"dockerEnabled"`
	ReadonlyDockerEnabled    bool `json:"readonlyDockerEnabled"`
	ServiceActivitiesEnabled bool `json:"serviceActivitiesEnabled"`
	SshdEnabled              bool `json:"sshdEnabled"`
	SudoEnabled              bool `json:"sudoEnabled"`
}
type LogInspectionRules struct {
	Path  string   `json:"path,omitempty"`
	Regex []string `json:"regex,omitempty"`
}

type DeniedProcesses struct {
	Effect string   `json:"effect,omitempty"`
	Paths  []string `json:"paths,omitempty"`
}

type AntiMalware struct {
	DetectCompilerGeneratedBinary bool            `json:"detectCompilerGeneratedBinary"`
	EncryptedBinaries             string          `json:"encryptedBinaries,omitempty"`
	ExecutionFlowHijack           string          `json:"executionFlowHijack,omitempty"`
	IntelligenceFeed              string          `json:"intelligenceFeed,omitempty"`
	ReverseShell                  string          `json:"reverseShell,omitempty"`
	ServiceUnknownOriginBinary    string          `json:"serviceUnknownOriginBinary,omitempty"`
	SkipSSHTracking               bool            `json:"skipSSHTracking"`
	SuspiciousELFHeaders          string          `json:"suspiciousELFHeaders,omitempty"`
	TempFSProc                    string          `json:"tempFSProc,omitempty"`
	UserUnknownOriginBinary       string          `json:"userUnknownOriginBinary,omitempty"`
	WebShell                      string          `json:"webShell,omitempty"`
	WildFireAnalysis              string          `json:"wildFireAnalysis,omitempty"`
	AllowedProcesses              []string        `json:"allowedProcesses,omitempty"`
	CryptoMiner                   string          `json:"cryptoMiner,omitempty"`
	CustomFeed                    string          `json:"customFeed,omitempty"`
	DeniedProcesses               DeniedProcesses `json:"deniedProcesses,omitempty"`
	Paths                         []string        `json:"paths,omitempty"`
}

type Threshold struct {
	Enabled  bool `json:"enabled"`
	Disabled bool `json:"disabled"`
	Value    int  `json:"value,omitempty"`
}

type Condition struct {
	Device          string          `json:"device,omitempty"`
	Readonly        bool            `json:"readonly"`
	Vulnerabilities []Vulnerability `json:"vulnerabilities,omitempty"`
}

type Vulnerability struct {
	Block       bool `json:"block"`
	Id          int  `json:"id,omitempty"`
	MinSeverity int  `json:"minSeverity,omitempty"`
}

type CveRule struct {
	Description string     `json:"description,omitempty"`
	Effect      string     `json:"effect,omitempty"`
	Id          string     `json:"id,omitempty"`
	Expiration  Expiration `json:"expiration,omitempty"`
}

type Expiration struct {
	Date    string `json:"date,omitempty"`
	Enabled bool   `json:"enabled"`
}

type License struct {
	AlertThreshold Threshold `json:"alertThreshold,omitempty"`
	BlockThreshold Threshold `json:"blockThreshold,omitempty"`
	Critical       []string  `json:"critical,omitempty"`
	High           []string  `json:"high,omitempty"`
	Low            []string  `json:"low,omitempty"`
	Medium         []string  `json:"medium,omitempty"`
}

type Tag struct {
	Description string       `json:"description,omitempty"`
	Effect      string       `json:"effect,omitempty"`
	Expiration  []Expiration `json:"expiration,omitempty"`
	Name        string       `json:"name,omitempty"`
}

type CustomRule struct {
	Id     int      `json:"_id,omitempty"`
	Action []string `json:"action,omitempty"`
	Effect string   `json:"effect,omitempty"`
}

type Dns struct {
	Blacklist        []string `json:"blacklist,omitempty"`
	Effect           string   `json:"effect,omitempty"`
	Whitelist        []string `json:"whitelist,omitempty"`
	Allow            []string `json:"allow,omitempty"`
	Deny             []string `json:"deny,omitempty"`
	DenyListEffect   string   `json:"denyListEffect,omitempty"`
	IntelligenceFeed string   `json:"intelligenceFeed,omitempty"`
}

type Filesystem struct {
	BackdoorFiles         bool     `json:"backdoorFiles"`
	Blacklist             []string `json:"blacklist,omitempty"`
	CheckNewFiles         bool     `json:"checkNewFiles"`
	Effect                string   `json:"effect,omitempty"`
	SkipEncryptedBinaries bool     `json:"skipEncryptedBinaries"`
	SuspiciousELFHeaders  bool     `json:"suspiciousELFHeaders"`
	Whitelist             []string `json:"whitelist,omitempty"`
}

type Network struct {
	BlacklistIPs            []string   `json:"blacklistIPs,omitempty"`
	BlacklistListeningPorts []ListPort `json:"blacklistListeningPorts,omitempty"`
	BlacklistOutboundPorts  []ListPort `json:"blacklistOutboundPorts,omitempty"`
	DetectPortScan          bool       `json:"detectPortScan"`
	Effect                  string     `json:"effect,omitempty"`
	SkipModifiedProc        bool       `json:"skipModifiedProc"`
	SkipRawSockets          bool       `json:"skipRawSockets"`
	WhitelistIPs            []string   `json:"whitelistIPs,omitempty"`
	WhitelistListeningPorts []ListPort `json:"whitelistListeningPorts,omitempty"`
	WhitelistOutboundPorts  []ListPort `json:"whitelistOutboundPorts,omitempty"`
	AllowedOutboundIPs      []string   `json:"allowedOutboundIPs,omitempty"`
	CustomFeed              string     `json:"customFeed,omitempty"`
	DeniedListeningPorts    []ListPort `json:"deniedListeningPorts,omitempty"`
	DeniedOutboundIPs       []string   `json:"deniedOutboundIPs,omitempty"`
	DeniedOutboundPorts     []ListPort `json:"deniedOutboundPorts,omitempty"`
	DenyListEffect          string     `json:"denyListEffect,omitempty"`
	IntelligenceFeed        string     `json:"intelligenceFeed,omitempty"`
}

type Processes struct {
	Blacklist            []string `json:"blacklist,omitempty"`
	BlockAllBinaries     bool     `json:"blockAllBinaries"`
	CheckCryptoMiners    bool     `json:"checkCryptoMiners"`
	CheckLateralMovement bool     `json:"checkLateralMovement"`
	CheckNewBinaries     bool     `json:"checkNewBinaries"`
	CheckParentChild     bool     `json:"checkParentChild"`
	CheckSuidBinaries    bool     `json:"checkSuidBinaries"`
	Effect               string   `json:"effect,omitempty"`
	SkipModified         bool     `json:"skipModified"`
	SkipReverseShell     bool     `json:"skipReverseShell"`
	Whitelist            []string `json:"whitelist,omitempty"`
}

type ListPort struct {
	Deny  bool `json:"deny"`
	End   int  `json:"end,omitempty"`
	Start int  `json:"start,omitempty"`
}
