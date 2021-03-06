// Code generated by protoc-gen-go.
// source: operation.proto
// DO NOT EDIT!

/*
Package v1beta1 is a generated protocol buffer package.

It is generated from these files:
	operation.proto

It has these top-level messages:
	Log
	DescribeRequest
	DescribeResponse
	LogDescribeRequest
	LogDescribeResponse
	Auth
	Metadata
	Operation
	DataBucketDeleteRequest
	NamespaceAdminTaskRequest
*/
package v1beta1

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis/google/api"
import _ "github.com/grpc-ecosystem/grpc-gateway/third_party/appscodeapis/appscode/api"
import appscode_dtypes "github.com/appscode/api/dtypes"
import appscode_kubernetes_v1beta1 "github.com/appscode/api/kubernetes/v1beta1"
import appscode_db_v1beta1 "github.com/appscode/api/db/v1beta1"
import appscode_db_v1beta11 "github.com/appscode/api/db/v1beta1"
import appscode_namespace_v1beta1 "github.com/appscode/api/namespace/v1beta1"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Next Id: 18
type OperationType int32

const (
	OperationType_UNKNOWN              OperationType = 0
	OperationType_CLUSTER_CREATE       OperationType = 1
	OperationType_CLUSTER_SCALE        OperationType = 2
	OperationType_CLUSTER_DELETE       OperationType = 3
	OperationType_CLUSTER_SET_VERSION  OperationType = 4
	OperationType_DATA_BUCKET_DELETE   OperationType = 9
	OperationType_BACKUP_SCHEDULE      OperationType = 10
	OperationType_NAMESPACE_CREATE     OperationType = 11
	OperationType_NAMESPACE_ADMIN_TASK OperationType = 12
	OperationType_DATABASE_CHECK       OperationType = 15
	OperationType_SNAPSHOT_CHECK       OperationType = 16
	OperationType_DATABASE_DELETE      OperationType = 17
)

var OperationType_name = map[int32]string{
	0:  "UNKNOWN",
	1:  "CLUSTER_CREATE",
	2:  "CLUSTER_SCALE",
	3:  "CLUSTER_DELETE",
	4:  "CLUSTER_SET_VERSION",
	9:  "DATA_BUCKET_DELETE",
	10: "BACKUP_SCHEDULE",
	11: "NAMESPACE_CREATE",
	12: "NAMESPACE_ADMIN_TASK",
	15: "DATABASE_CHECK",
	16: "SNAPSHOT_CHECK",
	17: "DATABASE_DELETE",
}
var OperationType_value = map[string]int32{
	"UNKNOWN":              0,
	"CLUSTER_CREATE":       1,
	"CLUSTER_SCALE":        2,
	"CLUSTER_DELETE":       3,
	"CLUSTER_SET_VERSION":  4,
	"DATA_BUCKET_DELETE":   9,
	"BACKUP_SCHEDULE":      10,
	"NAMESPACE_CREATE":     11,
	"NAMESPACE_ADMIN_TASK": 12,
	"DATABASE_CHECK":       15,
	"SNAPSHOT_CHECK":       16,
	"DATABASE_DELETE":      17,
}

func (x OperationType) String() string {
	return proto.EnumName(OperationType_name, int32(x))
}
func (OperationType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type Log struct {
	Id        string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Timestamp int64  `protobuf:"varint,2,opt,name=timestamp" json:"timestamp,omitempty"`
	Message   string `protobuf:"bytes,3,opt,name=message" json:"message,omitempty"`
}

func (m *Log) Reset()                    { *m = Log{} }
func (m *Log) String() string            { return proto.CompactTextString(m) }
func (*Log) ProtoMessage()               {}
func (*Log) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Log) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Log) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *Log) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type DescribeRequest struct {
	Phid string `protobuf:"bytes,1,opt,name=phid" json:"phid,omitempty"`
}

func (m *DescribeRequest) Reset()                    { *m = DescribeRequest{} }
func (m *DescribeRequest) String() string            { return proto.CompactTextString(m) }
func (*DescribeRequest) ProtoMessage()               {}
func (*DescribeRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *DescribeRequest) GetPhid() string {
	if m != nil {
		return m.Phid
	}
	return ""
}

type DescribeResponse struct {
	Status *appscode_dtypes.Status `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
	Op     *Operation              `protobuf:"bytes,2,opt,name=op" json:"op,omitempty"`
	Logs   []*Log                  `protobuf:"bytes,3,rep,name=logs" json:"logs,omitempty"`
}

func (m *DescribeResponse) Reset()                    { *m = DescribeResponse{} }
func (m *DescribeResponse) String() string            { return proto.CompactTextString(m) }
func (*DescribeResponse) ProtoMessage()               {}
func (*DescribeResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *DescribeResponse) GetStatus() *appscode_dtypes.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *DescribeResponse) GetOp() *Operation {
	if m != nil {
		return m.Op
	}
	return nil
}

func (m *DescribeResponse) GetLogs() []*Log {
	if m != nil {
		return m.Logs
	}
	return nil
}

type LogDescribeRequest struct {
	Phid  string `protobuf:"bytes,1,opt,name=phid" json:"phid,omitempty"`
	LogId string `protobuf:"bytes,2,opt,name=log_id,json=logId" json:"log_id,omitempty"`
}

func (m *LogDescribeRequest) Reset()                    { *m = LogDescribeRequest{} }
func (m *LogDescribeRequest) String() string            { return proto.CompactTextString(m) }
func (*LogDescribeRequest) ProtoMessage()               {}
func (*LogDescribeRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *LogDescribeRequest) GetPhid() string {
	if m != nil {
		return m.Phid
	}
	return ""
}

func (m *LogDescribeRequest) GetLogId() string {
	if m != nil {
		return m.LogId
	}
	return ""
}

type LogDescribeResponse struct {
	Status *appscode_dtypes.Status `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
	Log    *Log                    `protobuf:"bytes,2,opt,name=log" json:"log,omitempty"`
}

func (m *LogDescribeResponse) Reset()                    { *m = LogDescribeResponse{} }
func (m *LogDescribeResponse) String() string            { return proto.CompactTextString(m) }
func (*LogDescribeResponse) ProtoMessage()               {}
func (*LogDescribeResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *LogDescribeResponse) GetStatus() *appscode_dtypes.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *LogDescribeResponse) GetLog() *Log {
	if m != nil {
		return m.Log
	}
	return nil
}

type Auth struct {
	Namespace string `protobuf:"bytes,1,opt,name=namespace" json:"namespace,omitempty"`
	Username  string `protobuf:"bytes,2,opt,name=username" json:"username,omitempty"`
	Secret    string `protobuf:"bytes,3,opt,name=secret" json:"secret,omitempty"`
	AuthType  string `protobuf:"bytes,4,opt,name=auth_type,json=authType" json:"auth_type,omitempty"`
}

func (m *Auth) Reset()                    { *m = Auth{} }
func (m *Auth) String() string            { return proto.CompactTextString(m) }
func (*Auth) ProtoMessage()               {}
func (*Auth) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *Auth) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *Auth) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *Auth) GetSecret() string {
	if m != nil {
		return m.Secret
	}
	return ""
}

func (m *Auth) GetAuthType() string {
	if m != nil {
		return m.AuthType
	}
	return ""
}

// Metadata holds other on request or operation specific data
// that could be used inside that operation.
// An welldefined message instead of a map is used
// so that the data fields can be explicitly defined with
// its own data type. Resolves data convertions.
type Metadata struct {
	// Contains PurchasePHID is this is a purchase request.
	PurchasePhids []string `protobuf:"bytes,1,rep,name=purchase_phids,json=purchasePhids" json:"purchase_phids,omitempty"`
	// PHID of the user who requested this operation.
	AuthorPhid string `protobuf:"bytes,2,opt,name=author_phid,json=authorPhid" json:"author_phid,omitempty"`
	AuthorName string `protobuf:"bytes,3,opt,name=author_name,json=authorName" json:"author_name,omitempty"`
}

func (m *Metadata) Reset()                    { *m = Metadata{} }
func (m *Metadata) String() string            { return proto.CompactTextString(m) }
func (*Metadata) ProtoMessage()               {}
func (*Metadata) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *Metadata) GetPurchasePhids() []string {
	if m != nil {
		return m.PurchasePhids
	}
	return nil
}

func (m *Metadata) GetAuthorPhid() string {
	if m != nil {
		return m.AuthorPhid
	}
	return ""
}

func (m *Metadata) GetAuthorName() string {
	if m != nil {
		return m.AuthorName
	}
	return ""
}

// Next Id: 22
type Operation struct {
	// Types that are valid to be assigned to Request:
	//	*Operation_ClusterCreateRequest
	//	*Operation_ClusterScaleRequest
	//	*Operation_ClusterDeleteRequest
	//	*Operation_ClusterSetVersionRequest
	//	*Operation_DataBucketDeleteRequest
	//	*Operation_BackupScheduleRequest
	//	*Operation_NamespaceCreateRequest
	//	*Operation_NamespaceAdminTaskRequest
	//	*Operation_DatabaseCheckRequest
	//	*Operation_SnapshotCheckRequest
	//	*Operation_DatabaseDeleteRequest
	Request  isOperation_Request `protobuf_oneof:"request"`
	Type     OperationType       `protobuf:"varint,14,opt,name=type,enum=appscode.operation.v1beta1.OperationType" json:"type,omitempty"`
	Phid     string              `protobuf:"bytes,15,opt,name=phid" json:"phid,omitempty"`
	Auth     *Auth               `protobuf:"bytes,16,opt,name=auth" json:"auth,omitempty"`
	Metadata *Metadata           `protobuf:"bytes,17,opt,name=metadata" json:"metadata,omitempty"`
}

func (m *Operation) Reset()                    { *m = Operation{} }
func (m *Operation) String() string            { return proto.CompactTextString(m) }
func (*Operation) ProtoMessage()               {}
func (*Operation) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

type isOperation_Request interface {
	isOperation_Request()
}

type Operation_ClusterCreateRequest struct {
	ClusterCreateRequest *appscode_kubernetes_v1beta1.ClusterCreateRequest `protobuf:"bytes,1,opt,name=cluster_create_request,json=clusterCreateRequest,oneof"`
}
type Operation_ClusterScaleRequest struct {
	ClusterScaleRequest *appscode_kubernetes_v1beta1.ClusterScaleRequest `protobuf:"bytes,2,opt,name=cluster_scale_request,json=clusterScaleRequest,oneof"`
}
type Operation_ClusterDeleteRequest struct {
	ClusterDeleteRequest *appscode_kubernetes_v1beta1.ClusterDeleteRequest `protobuf:"bytes,3,opt,name=cluster_delete_request,json=clusterDeleteRequest,oneof"`
}
type Operation_ClusterSetVersionRequest struct {
	ClusterSetVersionRequest *appscode_kubernetes_v1beta1.ClusterSetVersionRequest `protobuf:"bytes,4,opt,name=cluster_set_version_request,json=clusterSetVersionRequest,oneof"`
}
type Operation_DataBucketDeleteRequest struct {
	DataBucketDeleteRequest *DataBucketDeleteRequest `protobuf:"bytes,9,opt,name=data_bucket_delete_request,json=dataBucketDeleteRequest,oneof"`
}
type Operation_BackupScheduleRequest struct {
	BackupScheduleRequest *appscode_db_v1beta11.BackupScheduleRequest `protobuf:"bytes,10,opt,name=backup_schedule_request,json=backupScheduleRequest,oneof"`
}
type Operation_NamespaceCreateRequest struct {
	NamespaceCreateRequest *appscode_namespace_v1beta1.CreateRequest `protobuf:"bytes,11,opt,name=namespace_create_request,json=namespaceCreateRequest,oneof"`
}
type Operation_NamespaceAdminTaskRequest struct {
	NamespaceAdminTaskRequest *NamespaceAdminTaskRequest `protobuf:"bytes,12,opt,name=namespace_admin_task_request,json=namespaceAdminTaskRequest,oneof"`
}
type Operation_DatabaseCheckRequest struct {
	DatabaseCheckRequest *appscode_db_v1beta1.DatabaseCheckRequest `protobuf:"bytes,19,opt,name=database_check_request,json=databaseCheckRequest,oneof"`
}
type Operation_SnapshotCheckRequest struct {
	SnapshotCheckRequest *appscode_db_v1beta11.SnapshotCheckRequest `protobuf:"bytes,20,opt,name=snapshot_check_request,json=snapshotCheckRequest,oneof"`
}
type Operation_DatabaseDeleteRequest struct {
	DatabaseDeleteRequest *appscode_db_v1beta1.DatabaseDeleteRequest `protobuf:"bytes,21,opt,name=database_delete_request,json=databaseDeleteRequest,oneof"`
}

func (*Operation_ClusterCreateRequest) isOperation_Request()      {}
func (*Operation_ClusterScaleRequest) isOperation_Request()       {}
func (*Operation_ClusterDeleteRequest) isOperation_Request()      {}
func (*Operation_ClusterSetVersionRequest) isOperation_Request()  {}
func (*Operation_DataBucketDeleteRequest) isOperation_Request()   {}
func (*Operation_BackupScheduleRequest) isOperation_Request()     {}
func (*Operation_NamespaceCreateRequest) isOperation_Request()    {}
func (*Operation_NamespaceAdminTaskRequest) isOperation_Request() {}
func (*Operation_DatabaseCheckRequest) isOperation_Request()      {}
func (*Operation_SnapshotCheckRequest) isOperation_Request()      {}
func (*Operation_DatabaseDeleteRequest) isOperation_Request()     {}

func (m *Operation) GetRequest() isOperation_Request {
	if m != nil {
		return m.Request
	}
	return nil
}

func (m *Operation) GetClusterCreateRequest() *appscode_kubernetes_v1beta1.ClusterCreateRequest {
	if x, ok := m.GetRequest().(*Operation_ClusterCreateRequest); ok {
		return x.ClusterCreateRequest
	}
	return nil
}

func (m *Operation) GetClusterScaleRequest() *appscode_kubernetes_v1beta1.ClusterScaleRequest {
	if x, ok := m.GetRequest().(*Operation_ClusterScaleRequest); ok {
		return x.ClusterScaleRequest
	}
	return nil
}

func (m *Operation) GetClusterDeleteRequest() *appscode_kubernetes_v1beta1.ClusterDeleteRequest {
	if x, ok := m.GetRequest().(*Operation_ClusterDeleteRequest); ok {
		return x.ClusterDeleteRequest
	}
	return nil
}

func (m *Operation) GetClusterSetVersionRequest() *appscode_kubernetes_v1beta1.ClusterSetVersionRequest {
	if x, ok := m.GetRequest().(*Operation_ClusterSetVersionRequest); ok {
		return x.ClusterSetVersionRequest
	}
	return nil
}

func (m *Operation) GetDataBucketDeleteRequest() *DataBucketDeleteRequest {
	if x, ok := m.GetRequest().(*Operation_DataBucketDeleteRequest); ok {
		return x.DataBucketDeleteRequest
	}
	return nil
}

func (m *Operation) GetBackupScheduleRequest() *appscode_db_v1beta11.BackupScheduleRequest {
	if x, ok := m.GetRequest().(*Operation_BackupScheduleRequest); ok {
		return x.BackupScheduleRequest
	}
	return nil
}

func (m *Operation) GetNamespaceCreateRequest() *appscode_namespace_v1beta1.CreateRequest {
	if x, ok := m.GetRequest().(*Operation_NamespaceCreateRequest); ok {
		return x.NamespaceCreateRequest
	}
	return nil
}

func (m *Operation) GetNamespaceAdminTaskRequest() *NamespaceAdminTaskRequest {
	if x, ok := m.GetRequest().(*Operation_NamespaceAdminTaskRequest); ok {
		return x.NamespaceAdminTaskRequest
	}
	return nil
}

func (m *Operation) GetDatabaseCheckRequest() *appscode_db_v1beta1.DatabaseCheckRequest {
	if x, ok := m.GetRequest().(*Operation_DatabaseCheckRequest); ok {
		return x.DatabaseCheckRequest
	}
	return nil
}

func (m *Operation) GetSnapshotCheckRequest() *appscode_db_v1beta11.SnapshotCheckRequest {
	if x, ok := m.GetRequest().(*Operation_SnapshotCheckRequest); ok {
		return x.SnapshotCheckRequest
	}
	return nil
}

func (m *Operation) GetDatabaseDeleteRequest() *appscode_db_v1beta1.DatabaseDeleteRequest {
	if x, ok := m.GetRequest().(*Operation_DatabaseDeleteRequest); ok {
		return x.DatabaseDeleteRequest
	}
	return nil
}

func (m *Operation) GetType() OperationType {
	if m != nil {
		return m.Type
	}
	return OperationType_UNKNOWN
}

func (m *Operation) GetPhid() string {
	if m != nil {
		return m.Phid
	}
	return ""
}

func (m *Operation) GetAuth() *Auth {
	if m != nil {
		return m.Auth
	}
	return nil
}

func (m *Operation) GetMetadata() *Metadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Operation) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Operation_OneofMarshaler, _Operation_OneofUnmarshaler, _Operation_OneofSizer, []interface{}{
		(*Operation_ClusterCreateRequest)(nil),
		(*Operation_ClusterScaleRequest)(nil),
		(*Operation_ClusterDeleteRequest)(nil),
		(*Operation_ClusterSetVersionRequest)(nil),
		(*Operation_DataBucketDeleteRequest)(nil),
		(*Operation_BackupScheduleRequest)(nil),
		(*Operation_NamespaceCreateRequest)(nil),
		(*Operation_NamespaceAdminTaskRequest)(nil),
		(*Operation_DatabaseCheckRequest)(nil),
		(*Operation_SnapshotCheckRequest)(nil),
		(*Operation_DatabaseDeleteRequest)(nil),
	}
}

func _Operation_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Operation)
	// request
	switch x := m.Request.(type) {
	case *Operation_ClusterCreateRequest:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.ClusterCreateRequest); err != nil {
			return err
		}
	case *Operation_ClusterScaleRequest:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.ClusterScaleRequest); err != nil {
			return err
		}
	case *Operation_ClusterDeleteRequest:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.ClusterDeleteRequest); err != nil {
			return err
		}
	case *Operation_ClusterSetVersionRequest:
		b.EncodeVarint(4<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.ClusterSetVersionRequest); err != nil {
			return err
		}
	case *Operation_DataBucketDeleteRequest:
		b.EncodeVarint(9<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.DataBucketDeleteRequest); err != nil {
			return err
		}
	case *Operation_BackupScheduleRequest:
		b.EncodeVarint(10<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.BackupScheduleRequest); err != nil {
			return err
		}
	case *Operation_NamespaceCreateRequest:
		b.EncodeVarint(11<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.NamespaceCreateRequest); err != nil {
			return err
		}
	case *Operation_NamespaceAdminTaskRequest:
		b.EncodeVarint(12<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.NamespaceAdminTaskRequest); err != nil {
			return err
		}
	case *Operation_DatabaseCheckRequest:
		b.EncodeVarint(19<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.DatabaseCheckRequest); err != nil {
			return err
		}
	case *Operation_SnapshotCheckRequest:
		b.EncodeVarint(20<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.SnapshotCheckRequest); err != nil {
			return err
		}
	case *Operation_DatabaseDeleteRequest:
		b.EncodeVarint(21<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.DatabaseDeleteRequest); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("Operation.Request has unexpected type %T", x)
	}
	return nil
}

func _Operation_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Operation)
	switch tag {
	case 1: // request.cluster_create_request
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(appscode_kubernetes_v1beta1.ClusterCreateRequest)
		err := b.DecodeMessage(msg)
		m.Request = &Operation_ClusterCreateRequest{msg}
		return true, err
	case 2: // request.cluster_scale_request
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(appscode_kubernetes_v1beta1.ClusterScaleRequest)
		err := b.DecodeMessage(msg)
		m.Request = &Operation_ClusterScaleRequest{msg}
		return true, err
	case 3: // request.cluster_delete_request
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(appscode_kubernetes_v1beta1.ClusterDeleteRequest)
		err := b.DecodeMessage(msg)
		m.Request = &Operation_ClusterDeleteRequest{msg}
		return true, err
	case 4: // request.cluster_set_version_request
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(appscode_kubernetes_v1beta1.ClusterSetVersionRequest)
		err := b.DecodeMessage(msg)
		m.Request = &Operation_ClusterSetVersionRequest{msg}
		return true, err
	case 9: // request.data_bucket_delete_request
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(DataBucketDeleteRequest)
		err := b.DecodeMessage(msg)
		m.Request = &Operation_DataBucketDeleteRequest{msg}
		return true, err
	case 10: // request.backup_schedule_request
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(appscode_db_v1beta11.BackupScheduleRequest)
		err := b.DecodeMessage(msg)
		m.Request = &Operation_BackupScheduleRequest{msg}
		return true, err
	case 11: // request.namespace_create_request
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(appscode_namespace_v1beta1.CreateRequest)
		err := b.DecodeMessage(msg)
		m.Request = &Operation_NamespaceCreateRequest{msg}
		return true, err
	case 12: // request.namespace_admin_task_request
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(NamespaceAdminTaskRequest)
		err := b.DecodeMessage(msg)
		m.Request = &Operation_NamespaceAdminTaskRequest{msg}
		return true, err
	case 19: // request.database_check_request
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(appscode_db_v1beta1.DatabaseCheckRequest)
		err := b.DecodeMessage(msg)
		m.Request = &Operation_DatabaseCheckRequest{msg}
		return true, err
	case 20: // request.snapshot_check_request
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(appscode_db_v1beta11.SnapshotCheckRequest)
		err := b.DecodeMessage(msg)
		m.Request = &Operation_SnapshotCheckRequest{msg}
		return true, err
	case 21: // request.database_delete_request
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(appscode_db_v1beta1.DatabaseDeleteRequest)
		err := b.DecodeMessage(msg)
		m.Request = &Operation_DatabaseDeleteRequest{msg}
		return true, err
	default:
		return false, nil
	}
}

func _Operation_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Operation)
	// request
	switch x := m.Request.(type) {
	case *Operation_ClusterCreateRequest:
		s := proto.Size(x.ClusterCreateRequest)
		n += proto.SizeVarint(1<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Operation_ClusterScaleRequest:
		s := proto.Size(x.ClusterScaleRequest)
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Operation_ClusterDeleteRequest:
		s := proto.Size(x.ClusterDeleteRequest)
		n += proto.SizeVarint(3<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Operation_ClusterSetVersionRequest:
		s := proto.Size(x.ClusterSetVersionRequest)
		n += proto.SizeVarint(4<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Operation_DataBucketDeleteRequest:
		s := proto.Size(x.DataBucketDeleteRequest)
		n += proto.SizeVarint(9<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Operation_BackupScheduleRequest:
		s := proto.Size(x.BackupScheduleRequest)
		n += proto.SizeVarint(10<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Operation_NamespaceCreateRequest:
		s := proto.Size(x.NamespaceCreateRequest)
		n += proto.SizeVarint(11<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Operation_NamespaceAdminTaskRequest:
		s := proto.Size(x.NamespaceAdminTaskRequest)
		n += proto.SizeVarint(12<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Operation_DatabaseCheckRequest:
		s := proto.Size(x.DatabaseCheckRequest)
		n += proto.SizeVarint(19<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Operation_SnapshotCheckRequest:
		s := proto.Size(x.SnapshotCheckRequest)
		n += proto.SizeVarint(20<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Operation_DatabaseDeleteRequest:
		s := proto.Size(x.DatabaseDeleteRequest)
		n += proto.SizeVarint(21<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type DataBucketDeleteRequest struct {
	DataType  string `protobuf:"bytes,1,opt,name=data_type,json=dataType" json:"data_type,omitempty"`
	Namespace string `protobuf:"bytes,2,opt,name=namespace" json:"namespace,omitempty"`
	Prefix    string `protobuf:"bytes,3,opt,name=prefix" json:"prefix,omitempty"`
}

func (m *DataBucketDeleteRequest) Reset()                    { *m = DataBucketDeleteRequest{} }
func (m *DataBucketDeleteRequest) String() string            { return proto.CompactTextString(m) }
func (*DataBucketDeleteRequest) ProtoMessage()               {}
func (*DataBucketDeleteRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *DataBucketDeleteRequest) GetDataType() string {
	if m != nil {
		return m.DataType
	}
	return ""
}

func (m *DataBucketDeleteRequest) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *DataBucketDeleteRequest) GetPrefix() string {
	if m != nil {
		return m.Prefix
	}
	return ""
}

type NamespaceAdminTaskRequest struct {
	Namespace string `protobuf:"bytes,1,opt,name=namespace" json:"namespace,omitempty"`
}

func (m *NamespaceAdminTaskRequest) Reset()                    { *m = NamespaceAdminTaskRequest{} }
func (m *NamespaceAdminTaskRequest) String() string            { return proto.CompactTextString(m) }
func (*NamespaceAdminTaskRequest) ProtoMessage()               {}
func (*NamespaceAdminTaskRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *NamespaceAdminTaskRequest) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func init() {
	proto.RegisterType((*Log)(nil), "appscode.operation.v1beta1.Log")
	proto.RegisterType((*DescribeRequest)(nil), "appscode.operation.v1beta1.DescribeRequest")
	proto.RegisterType((*DescribeResponse)(nil), "appscode.operation.v1beta1.DescribeResponse")
	proto.RegisterType((*LogDescribeRequest)(nil), "appscode.operation.v1beta1.LogDescribeRequest")
	proto.RegisterType((*LogDescribeResponse)(nil), "appscode.operation.v1beta1.LogDescribeResponse")
	proto.RegisterType((*Auth)(nil), "appscode.operation.v1beta1.Auth")
	proto.RegisterType((*Metadata)(nil), "appscode.operation.v1beta1.Metadata")
	proto.RegisterType((*Operation)(nil), "appscode.operation.v1beta1.Operation")
	proto.RegisterType((*DataBucketDeleteRequest)(nil), "appscode.operation.v1beta1.DataBucketDeleteRequest")
	proto.RegisterType((*NamespaceAdminTaskRequest)(nil), "appscode.operation.v1beta1.NamespaceAdminTaskRequest")
	proto.RegisterEnum("appscode.operation.v1beta1.OperationType", OperationType_name, OperationType_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Operations service

type OperationsClient interface {
	Describe(ctx context.Context, in *DescribeRequest, opts ...grpc.CallOption) (*DescribeResponse, error)
	DescribeLog(ctx context.Context, in *LogDescribeRequest, opts ...grpc.CallOption) (*LogDescribeResponse, error)
}

type operationsClient struct {
	cc *grpc.ClientConn
}

func NewOperationsClient(cc *grpc.ClientConn) OperationsClient {
	return &operationsClient{cc}
}

func (c *operationsClient) Describe(ctx context.Context, in *DescribeRequest, opts ...grpc.CallOption) (*DescribeResponse, error) {
	out := new(DescribeResponse)
	err := grpc.Invoke(ctx, "/appscode.operation.v1beta1.Operations/Describe", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *operationsClient) DescribeLog(ctx context.Context, in *LogDescribeRequest, opts ...grpc.CallOption) (*LogDescribeResponse, error) {
	out := new(LogDescribeResponse)
	err := grpc.Invoke(ctx, "/appscode.operation.v1beta1.Operations/DescribeLog", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Operations service

type OperationsServer interface {
	Describe(context.Context, *DescribeRequest) (*DescribeResponse, error)
	DescribeLog(context.Context, *LogDescribeRequest) (*LogDescribeResponse, error)
}

func RegisterOperationsServer(s *grpc.Server, srv OperationsServer) {
	s.RegisterService(&_Operations_serviceDesc, srv)
}

func _Operations_Describe_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DescribeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OperationsServer).Describe(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appscode.operation.v1beta1.Operations/Describe",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OperationsServer).Describe(ctx, req.(*DescribeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Operations_DescribeLog_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LogDescribeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OperationsServer).DescribeLog(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appscode.operation.v1beta1.Operations/DescribeLog",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OperationsServer).DescribeLog(ctx, req.(*LogDescribeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Operations_serviceDesc = grpc.ServiceDesc{
	ServiceName: "appscode.operation.v1beta1.Operations",
	HandlerType: (*OperationsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Describe",
			Handler:    _Operations_Describe_Handler,
		},
		{
			MethodName: "DescribeLog",
			Handler:    _Operations_DescribeLog_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "operation.proto",
}

func init() { proto.RegisterFile("operation.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 1220 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x9c, 0x56, 0xcd, 0x6e, 0xdb, 0x46,
	0x10, 0x8e, 0x7e, 0x6a, 0x5b, 0xa3, 0xc4, 0x66, 0xd6, 0x7f, 0x8c, 0x13, 0x34, 0x06, 0xd1, 0x14,
	0x4e, 0xda, 0x4a, 0xb5, 0x93, 0x1c, 0x7a, 0x28, 0x12, 0x8a, 0x22, 0xe0, 0xc0, 0xb2, 0x6c, 0x90,
	0x72, 0x0a, 0xf4, 0x42, 0xac, 0xc8, 0x8d, 0x44, 0x58, 0x12, 0x59, 0xee, 0x32, 0x48, 0x10, 0xe4,
	0x92, 0x63, 0xaf, 0x7d, 0x88, 0x1e, 0x7b, 0xea, 0xb5, 0x2f, 0xd1, 0x57, 0xe8, 0x3b, 0xf4, 0x56,
	0x14, 0xbb, 0x24, 0x97, 0x14, 0x6d, 0xcb, 0x6e, 0x2f, 0x82, 0x38, 0xfb, 0x7d, 0xfb, 0xcd, 0xec,
	0xcc, 0xee, 0x0c, 0xac, 0x05, 0x21, 0x89, 0x30, 0xf3, 0x83, 0x59, 0x2b, 0x8c, 0x02, 0x16, 0xa0,
	0x1d, 0x1c, 0x86, 0xd4, 0x0d, 0x3c, 0xd2, 0xca, 0x57, 0xde, 0xee, 0x0f, 0x09, 0xc3, 0xfb, 0x3b,
	0x0f, 0x46, 0x41, 0x30, 0x9a, 0x90, 0x36, 0x0e, 0xfd, 0x36, 0x9e, 0xcd, 0x02, 0x26, 0x96, 0x69,
	0xc2, 0xdc, 0xf9, 0x3c, 0x63, 0x5e, 0xb1, 0xfe, 0x70, 0x6e, 0xdd, 0x63, 0xef, 0x43, 0x42, 0xdb,
	0xe2, 0x37, 0x05, 0x7c, 0x33, 0x07, 0x38, 0x8f, 0x87, 0x24, 0x9a, 0x11, 0x46, 0x68, 0x3b, 0xd5,
	0x6f, 0xbb, 0x93, 0x98, 0x32, 0x12, 0xa5, 0xf0, 0x2f, 0xe7, 0xf7, 0x1b, 0x4a, 0x98, 0x87, 0x19,
	0x1e, 0x62, 0x4a, 0xae, 0xc3, 0xd1, 0x19, 0x0e, 0xe9, 0x38, 0x60, 0x29, 0xee, 0xf1, 0x1c, 0x6e,
	0x86, 0xa7, 0x84, 0x86, 0xd8, 0x25, 0x12, 0xce, 0x08, 0x9e, 0x26, 0x50, 0xed, 0x18, 0x6a, 0xbd,
	0x60, 0x84, 0x56, 0xa1, 0xea, 0x7b, 0x6a, 0x65, 0xb7, 0xb2, 0xd7, 0xb0, 0xaa, 0xbe, 0x87, 0x1e,
	0x40, 0x83, 0xf9, 0x53, 0x42, 0x19, 0x9e, 0x86, 0x6a, 0x75, 0xb7, 0xb2, 0x57, 0xb3, 0x72, 0x03,
	0x52, 0x61, 0x79, 0x4a, 0x28, 0xc5, 0x23, 0xa2, 0xd6, 0x04, 0x25, 0xfb, 0xd4, 0x1e, 0xc1, 0x5a,
	0x97, 0x50, 0x37, 0xf2, 0x87, 0xc4, 0x22, 0x3f, 0xc5, 0x84, 0x32, 0x84, 0xa0, 0x1e, 0x8e, 0xe5,
	0xe6, 0xe2, 0xbf, 0xf6, 0x5b, 0x05, 0x94, 0x1c, 0x47, 0xc3, 0x60, 0x46, 0x09, 0x6a, 0xc3, 0x12,
	0x65, 0x98, 0xc5, 0x54, 0x40, 0x9b, 0x07, 0xdb, 0x2d, 0x99, 0xc0, 0xe4, 0x88, 0x5b, 0xb6, 0x58,
	0xb6, 0x52, 0x18, 0x7a, 0x0e, 0xd5, 0x20, 0xf1, 0xae, 0x79, 0xf0, 0xa8, 0x75, 0x75, 0xb6, 0x5b,
	0x27, 0x99, 0xc5, 0xaa, 0x06, 0x21, 0x7a, 0x0a, 0xf5, 0x49, 0x30, 0xa2, 0x6a, 0x6d, 0xb7, 0xb6,
	0xd7, 0x3c, 0x78, 0xb8, 0x88, 0xd8, 0x0b, 0x46, 0x96, 0x00, 0x6b, 0x2f, 0x00, 0xf5, 0x82, 0xd1,
	0x0d, 0x62, 0x43, 0x9b, 0xb0, 0x34, 0x09, 0x46, 0x8e, 0xef, 0x09, 0xcf, 0x1a, 0xd6, 0x67, 0x93,
	0x60, 0xf4, 0xca, 0xd3, 0xde, 0xc3, 0xfa, 0xdc, 0x06, 0xff, 0x37, 0xe8, 0x7d, 0xa8, 0x4d, 0x82,
	0x51, 0x1a, 0xf5, 0xb5, 0xce, 0x73, 0xac, 0x16, 0x43, 0x5d, 0x8f, 0xd9, 0x98, 0x27, 0x55, 0xd6,
	0x42, 0xea, 0x72, 0x6e, 0x40, 0x3b, 0xb0, 0x12, 0x53, 0x12, 0x71, 0x43, 0xea, 0xb9, 0xfc, 0x46,
	0x5b, 0xb0, 0x44, 0x89, 0x1b, 0x11, 0x96, 0xe6, 0x3b, 0xfd, 0x42, 0xf7, 0xa1, 0x81, 0x63, 0x36,
	0x76, 0xb8, 0xa7, 0x6a, 0x3d, 0x21, 0x71, 0xc3, 0xe0, 0x7d, 0x48, 0x34, 0x0a, 0x2b, 0xc7, 0x84,
	0x61, 0x5e, 0xc3, 0xe8, 0x11, 0xac, 0x86, 0x71, 0xe4, 0x8e, 0x31, 0x25, 0x0e, 0x3f, 0x25, 0x1e,
	0x6e, 0x6d, 0xaf, 0x61, 0xdd, 0xc9, 0xac, 0xa7, 0xdc, 0x88, 0x1e, 0x42, 0x93, 0xd3, 0x83, 0x48,
	0x80, 0x52, 0x37, 0x20, 0x31, 0x71, 0x44, 0x01, 0x20, 0xfc, 0xac, 0x15, 0x01, 0x7d, 0x3c, 0x25,
	0xda, 0x3f, 0x00, 0x0d, 0x99, 0x6e, 0xe4, 0xc3, 0x56, 0x7a, 0xd3, 0x1c, 0x37, 0x22, 0x98, 0x11,
	0x27, 0x4a, 0x32, 0x97, 0x9e, 0xf6, 0x7e, 0x7e, 0x7e, 0xf9, 0x25, 0x95, 0x07, 0x68, 0x24, 0x54,
	0x43, 0x30, 0xd3, 0x94, 0x1f, 0xde, 0xb2, 0x36, 0xdc, 0x4b, 0xec, 0xe8, 0x0d, 0x6c, 0x66, 0x52,
	0xd4, 0xc5, 0x93, 0x5c, 0x29, 0xc9, 0xd4, 0xb7, 0x37, 0x51, 0xb2, 0x39, 0x31, 0x17, 0x5a, 0x77,
	0x2f, 0x9a, 0x8b, 0x21, 0x79, 0x64, 0x42, 0x0a, 0x21, 0xd5, 0x6e, 0x1e, 0x52, 0x57, 0x30, 0x2f,
	0x86, 0x34, 0x67, 0x47, 0x6f, 0xe1, 0xbe, 0x0c, 0x89, 0x30, 0xe7, 0x2d, 0x89, 0xa8, 0x1f, 0xcc,
	0xa4, 0x5e, 0x5d, 0xe8, 0x3d, 0xbf, 0x51, 0x60, 0x84, 0xbd, 0x4e, 0xd8, 0xb9, 0xa6, 0xea, 0x5e,
	0xb1, 0x86, 0x22, 0xd8, 0xe1, 0x45, 0xe3, 0x0c, 0x63, 0xf7, 0x9c, 0xb0, 0x72, 0x98, 0x0d, 0x21,
	0xfb, 0x74, 0x51, 0xe5, 0x77, 0x31, 0xc3, 0x1d, 0x41, 0x2e, 0x07, 0xba, 0xed, 0x5d, 0xbe, 0x84,
	0x3c, 0xd8, 0x1e, 0x62, 0xf7, 0x3c, 0x0e, 0x1d, 0xea, 0x8e, 0x89, 0x17, 0x17, 0x12, 0x08, 0x42,
	0xf0, 0x49, 0xe1, 0x62, 0x0e, 0xa5, 0x52, 0x47, 0x70, 0xec, 0x94, 0x92, 0xeb, 0x6c, 0x0e, 0x2f,
	0x5b, 0x40, 0x04, 0x54, 0x79, 0xe1, 0xca, 0x15, 0xd9, 0x14, 0x32, 0x8f, 0x73, 0x19, 0x89, 0xcc,
	0x4f, 0xb3, 0x54, 0x89, 0x5b, 0x12, 0x32, 0x5f, 0x8b, 0xef, 0xe0, 0x41, 0x2e, 0x83, 0xbd, 0xa9,
	0x3f, 0x73, 0x18, 0xa6, 0xe7, 0x52, 0xea, 0x76, 0x39, 0x73, 0x17, 0x8f, 0xb0, 0x9f, 0xf1, 0x75,
	0x4e, 0x1f, 0x60, 0x7a, 0x9e, 0xcb, 0xde, 0x9b, 0x5d, 0xb5, 0x88, 0x30, 0x6c, 0x65, 0x3d, 0xcb,
	0x71, 0xc7, 0xc4, 0xcd, 0x35, 0xd7, 0xcb, 0xe1, 0x15, 0x4e, 0xb1, 0x9b, 0x52, 0x0c, 0xce, 0x28,
	0x54, 0xa5, 0x77, 0x89, 0x9d, 0x4b, 0x64, 0xed, 0xae, 0x24, 0xb1, 0xb1, 0x40, 0xc2, 0x4e, 0x29,
	0x65, 0x09, 0x7a, 0x89, 0x9d, 0x17, 0x83, 0x8c, 0xa2, 0x54, 0x7d, 0x9b, 0x0b, 0x8a, 0x21, 0x0b,
	0xa3, 0x5c, 0x74, 0x9b, 0xde, 0x65, 0x0b, 0xe8, 0x7b, 0xa8, 0x8b, 0x77, 0x73, 0x75, 0xb7, 0xb2,
	0xb7, 0x5a, 0x74, 0x7b, 0x41, 0x03, 0xe3, 0x0f, 0xab, 0x25, 0x68, 0xb2, 0xf7, 0xac, 0x15, 0x7a,
	0xcf, 0x33, 0xa8, 0xf3, 0xb7, 0x50, 0x55, 0x84, 0x97, 0xbb, 0x8b, 0xb6, 0xe4, 0x1d, 0xc1, 0x12,
	0x68, 0xf4, 0x12, 0x56, 0xa6, 0xe9, 0x43, 0xad, 0xde, 0x15, 0xcc, 0x2f, 0x16, 0x31, 0xb3, 0x47,
	0xdd, 0x92, 0xac, 0x4e, 0x03, 0x96, 0xd3, 0x03, 0xd2, 0x26, 0xb0, 0x7d, 0xc5, 0xf5, 0xe3, 0xdd,
	0x42, 0xdc, 0x6b, 0x11, 0x75, 0xd2, 0x7f, 0x56, 0xb8, 0x81, 0x07, 0x35, 0xdf, 0x9c, 0xaa, 0xe5,
	0xe6, 0xb4, 0x05, 0x4b, 0x61, 0x44, 0xde, 0xf8, 0xef, 0xb2, 0x06, 0x94, 0x7c, 0x69, 0xdf, 0xc1,
	0xbd, 0x2b, 0x2b, 0x75, 0x71, 0xbf, 0x7b, 0xf2, 0x73, 0x15, 0xee, 0xcc, 0x9d, 0x2b, 0x6a, 0xc2,
	0xf2, 0x59, 0xff, 0xa8, 0x7f, 0xf2, 0x43, 0x5f, 0xb9, 0x85, 0x10, 0xac, 0x1a, 0xbd, 0x33, 0x7b,
	0x60, 0x5a, 0x8e, 0x61, 0x99, 0xfa, 0xc0, 0x54, 0x2a, 0xe8, 0x2e, 0xdc, 0xc9, 0x6c, 0xb6, 0xa1,
	0xf7, 0x4c, 0xa5, 0x5a, 0x84, 0x75, 0xcd, 0x9e, 0x39, 0x30, 0x95, 0x1a, 0xda, 0x86, 0x75, 0x09,
	0x33, 0x07, 0xce, 0x6b, 0xd3, 0xb2, 0x5f, 0x9d, 0xf4, 0x95, 0x3a, 0xda, 0x02, 0xd4, 0xd5, 0x07,
	0xba, 0xd3, 0x39, 0x33, 0x8e, 0xcc, 0x41, 0x46, 0x68, 0xa0, 0x75, 0x58, 0xeb, 0xe8, 0xc6, 0xd1,
	0xd9, 0xa9, 0x63, 0x1b, 0x87, 0x66, 0xf7, 0xac, 0x67, 0x2a, 0x80, 0x36, 0x40, 0xe9, 0xeb, 0xc7,
	0xa6, 0x7d, 0xaa, 0x1b, 0x66, 0xe6, 0x42, 0x13, 0xa9, 0xb0, 0x91, 0x5b, 0xf5, 0xee, 0xf1, 0xab,
	0xbe, 0x33, 0xd0, 0xed, 0x23, 0xe5, 0x36, 0xf7, 0x84, 0x6f, 0xde, 0xd1, 0x6d, 0xd3, 0x31, 0x0e,
	0x4d, 0xe3, 0x48, 0x59, 0xe3, 0x36, 0xbb, 0xaf, 0x9f, 0xda, 0x87, 0x27, 0x83, 0xd4, 0xa6, 0x70,
	0x31, 0x89, 0x4b, 0x3d, 0xb8, 0x7b, 0xf0, 0x77, 0x15, 0x40, 0x1e, 0x06, 0x45, 0xbf, 0x56, 0x60,
	0x25, 0x1b, 0x55, 0xd0, 0x57, 0x0b, 0x9f, 0xda, 0xf9, 0x89, 0x68, 0xe7, 0xeb, 0x9b, 0x81, 0x93,
	0xe9, 0x47, 0x7b, 0xf9, 0xe9, 0x77, 0xb5, 0xba, 0x52, 0xf9, 0xf4, 0xe7, 0x5f, 0xbf, 0x54, 0x9f,
	0xa1, 0x83, 0xb6, 0x33, 0x37, 0xb8, 0xca, 0x0d, 0xe4, 0xe0, 0x2a, 0x2d, 0xb4, 0xfd, 0x81, 0x17,
	0xfc, 0x47, 0xf4, 0x47, 0x05, 0x9a, 0xd9, 0xb6, 0x7c, 0x90, 0x6d, 0x5d, 0x33, 0x11, 0x95, 0xfd,
	0x6d, 0xdf, 0x18, 0x9f, 0xba, 0xdc, 0x2f, 0xb8, 0xdc, 0x41, 0x2f, 0xff, 0xbb, 0xcb, 0x6d, 0x3e,
	0x4f, 0xb6, 0x3f, 0x24, 0x33, 0xe2, 0xc7, 0xce, 0x0b, 0xd0, 0xdc, 0x60, 0x9a, 0x7b, 0x81, 0x43,
	0xff, 0xa2, 0x27, 0x9d, 0x55, 0x99, 0x9c, 0x53, 0x3e, 0xb6, 0x9f, 0x56, 0x7e, 0x5c, 0x4e, 0x97,
	0x86, 0x4b, 0x62, 0x90, 0x7f, 0xfa, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x93, 0x17, 0x4c, 0x07,
	0x00, 0x0d, 0x00, 0x00,
}
