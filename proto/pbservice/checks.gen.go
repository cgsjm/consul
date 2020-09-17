package pbservice

import structs "github.com/hashicorp/consul/agent/structs"

func CheckTypeToStructs(s CheckType) structs.CheckType {
	var t structs.CheckType
	t.CheckID = s.CheckID
	t.Name = s.Name
	t.Status = s.Status
	t.Notes = s.Notes
	t.ScriptArgs = s.ScriptArgs
	t.HTTP = s.HTTP
	t.Header = MapHeadersToStructs(s.Header)
	t.Method = s.Method
	t.Body = s.Body
	t.TCP = s.TCP
	t.Interval = s.Interval
	t.AliasNode = s.AliasNode
	t.AliasService = s.AliasService
	t.DockerContainerID = s.DockerContainerID
	t.Shell = s.Shell
	t.GRPC = s.GRPC
	t.GRPCUseTLS = s.GRPCUseTLS
	t.TLSSkipVerify = s.TLSSkipVerify
	t.Timeout = s.Timeout
	t.TTL = s.TTL
	t.SuccessBeforePassing = int(s.SuccessBeforePassing)
	t.FailuresBeforeCritical = int(s.FailuresBeforeCritical)
	t.ProxyHTTP = s.ProxyHTTP
	t.ProxyGRPC = s.ProxyGRPC
	t.DeregisterCriticalServiceAfter = s.DeregisterCriticalServiceAfter
	t.OutputMaxSize = int(s.OutputMaxSize)
	return t
}
func NewCheckTypeFromStructs(t structs.CheckType) CheckType {
	var s CheckType
	s.CheckID = t.CheckID
	s.Name = t.Name
	s.Status = t.Status
	s.Notes = t.Notes
	s.ScriptArgs = t.ScriptArgs
	s.HTTP = t.HTTP
	s.Header = NewMapHeadersFromStructs(t.Header)
	s.Method = t.Method
	s.Body = t.Body
	s.TCP = t.TCP
	s.Interval = t.Interval
	s.AliasNode = t.AliasNode
	s.AliasService = t.AliasService
	s.DockerContainerID = t.DockerContainerID
	s.Shell = t.Shell
	s.GRPC = t.GRPC
	s.GRPCUseTLS = t.GRPCUseTLS
	s.TLSSkipVerify = t.TLSSkipVerify
	s.Timeout = t.Timeout
	s.TTL = t.TTL
	s.SuccessBeforePassing = int32(t.SuccessBeforePassing)
	s.FailuresBeforeCritical = int32(t.FailuresBeforeCritical)
	s.ProxyHTTP = t.ProxyHTTP
	s.ProxyGRPC = t.ProxyGRPC
	s.DeregisterCriticalServiceAfter = t.DeregisterCriticalServiceAfter
	s.OutputMaxSize = int32(t.OutputMaxSize)
	return s
}
func HealthCheckToStructs(s HealthCheck) structs.HealthCheck {
	var t structs.HealthCheck
	t.Node = s.Node
	t.CheckID = s.CheckID
	t.Name = s.Name
	t.Status = s.Status
	t.Notes = s.Notes
	t.Output = s.Output
	t.ServiceID = s.ServiceID
	t.ServiceName = s.ServiceName
	t.ServiceTags = s.ServiceTags
	t.Type = s.Type
	t.Definition = HealthCheckDefinitionToStructs(s.Definition)
	t.EnterpriseMeta = EnterpriseMetaToStructs(s.EnterpriseMeta)
	t.RaftIndex = RaftIndexToStructs(s.RaftIndex)
	return t
}
func NewHealthCheckFromStructs(t structs.HealthCheck) HealthCheck {
	var s HealthCheck
	s.Node = t.Node
	s.CheckID = t.CheckID
	s.Name = t.Name
	s.Status = t.Status
	s.Notes = t.Notes
	s.Output = t.Output
	s.ServiceID = t.ServiceID
	s.ServiceName = t.ServiceName
	s.ServiceTags = t.ServiceTags
	s.Type = t.Type
	s.Definition = NewHealthCheckDefinitionFromStructs(t.Definition)
	s.EnterpriseMeta = NewEnterpriseMetaFromStructs(t.EnterpriseMeta)
	s.RaftIndex = NewRaftIndexFromStructs(t.RaftIndex)
	return s
}
func HealthCheckDefinitionToStructs(s HealthCheckDefinition) structs.HealthCheckDefinition {
	var t structs.HealthCheckDefinition
	t.HTTP = s.HTTP
	t.TLSSkipVerify = s.TLSSkipVerify
	t.Header = MapHeadersToStructs(s.Header)
	t.Method = s.Method
	t.Body = s.Body
	t.TCP = s.TCP
	t.Interval = s.Interval
	t.OutputMaxSize = uint(s.OutputMaxSize)
	t.Timeout = s.Timeout
	t.DeregisterCriticalServiceAfter = s.DeregisterCriticalServiceAfter
	t.ScriptArgs = s.ScriptArgs
	t.DockerContainerID = s.DockerContainerID
	t.Shell = s.Shell
	t.GRPC = s.GRPC
	t.GRPCUseTLS = s.GRPCUseTLS
	t.AliasNode = s.AliasNode
	t.AliasService = s.AliasService
	t.TTL = s.TTL
	return t
}
func NewHealthCheckDefinitionFromStructs(t structs.HealthCheckDefinition) HealthCheckDefinition {
	var s HealthCheckDefinition
	s.HTTP = t.HTTP
	s.TLSSkipVerify = t.TLSSkipVerify
	s.Header = NewMapHeadersFromStructs(t.Header)
	s.Method = t.Method
	s.Body = t.Body
	s.TCP = t.TCP
	s.Interval = t.Interval
	s.OutputMaxSize = uint32(t.OutputMaxSize)
	s.Timeout = t.Timeout
	s.DeregisterCriticalServiceAfter = t.DeregisterCriticalServiceAfter
	s.ScriptArgs = t.ScriptArgs
	s.DockerContainerID = t.DockerContainerID
	s.Shell = t.Shell
	s.GRPC = t.GRPC
	s.GRPCUseTLS = t.GRPCUseTLS
	s.AliasNode = t.AliasNode
	s.AliasService = t.AliasService
	s.TTL = t.TTL
	return s
}
