// Code generated by protoc-gen-go. DO NOT EDIT.
// source: message.proto

/*
Package msg is a generated protocol buffer package.

It is generated from these files:
	message.proto

It has these top-level messages:
	TransForm
	Position
	CallFuncInfo
	Error
	LoginInput
	RegistInput
	UserInfo
	RoomInfo
	Character
	Color
	Equipment
	Ability
*/
package msg

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/any"

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

type TransForm struct {
	X     float32 `protobuf:"fixed32,1,opt,name=X,json=x" json:"X,omitempty"`
	Y     float32 `protobuf:"fixed32,2,opt,name=Y,json=y" json:"Y,omitempty"`
	Z     float32 `protobuf:"fixed32,3,opt,name=Z,json=z" json:"Z,omitempty"`
	Yaw   float32 `protobuf:"fixed32,4,opt,name=Yaw,json=yaw" json:"Yaw,omitempty"`
	Pitch float32 `protobuf:"fixed32,5,opt,name=Pitch,json=pitch" json:"Pitch,omitempty"`
	Row   float32 `protobuf:"fixed32,6,opt,name=Row,json=row" json:"Row,omitempty"`
}

func (m *TransForm) Reset()                    { *m = TransForm{} }
func (m *TransForm) String() string            { return proto.CompactTextString(m) }
func (*TransForm) ProtoMessage()               {}
func (*TransForm) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *TransForm) GetX() float32 {
	if m != nil {
		return m.X
	}
	return 0
}

func (m *TransForm) GetY() float32 {
	if m != nil {
		return m.Y
	}
	return 0
}

func (m *TransForm) GetZ() float32 {
	if m != nil {
		return m.Z
	}
	return 0
}

func (m *TransForm) GetYaw() float32 {
	if m != nil {
		return m.Yaw
	}
	return 0
}

func (m *TransForm) GetPitch() float32 {
	if m != nil {
		return m.Pitch
	}
	return 0
}

func (m *TransForm) GetRow() float32 {
	if m != nil {
		return m.Row
	}
	return 0
}

type Position struct {
	PosMap map[string]*TransForm `protobuf:"bytes,1,rep,name=PosMap,json=posMap" json:"PosMap,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *Position) Reset()                    { *m = Position{} }
func (m *Position) String() string            { return proto.CompactTextString(m) }
func (*Position) ProtoMessage()               {}
func (*Position) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Position) GetPosMap() map[string]*TransForm {
	if m != nil {
		return m.PosMap
	}
	return nil
}

type CallFuncInfo struct {
	RunnigNo   string                 `protobuf:"bytes,1,opt,name=RunnigNo,json=runnigNo" json:"RunnigNo,omitempty"`
	TargetId   string                 `protobuf:"bytes,2,opt,name=TargetId,json=targetId" json:"TargetId,omitempty"`
	TragetType []byte                 `protobuf:"bytes,3,opt,name=TragetType,json=tragetType,proto3" json:"TragetType,omitempty"`
	Cb         string                 `protobuf:"bytes,4,opt,name=Cb,json=cb" json:"Cb,omitempty"`
	Param      []*google_protobuf.Any `protobuf:"bytes,5,rep,name=Param,json=param" json:"Param,omitempty"`
	TimeStamp  int64                  `protobuf:"varint,6,opt,name=TimeStamp,json=timeStamp" json:"TimeStamp,omitempty"`
}

func (m *CallFuncInfo) Reset()                    { *m = CallFuncInfo{} }
func (m *CallFuncInfo) String() string            { return proto.CompactTextString(m) }
func (*CallFuncInfo) ProtoMessage()               {}
func (*CallFuncInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *CallFuncInfo) GetRunnigNo() string {
	if m != nil {
		return m.RunnigNo
	}
	return ""
}

func (m *CallFuncInfo) GetTargetId() string {
	if m != nil {
		return m.TargetId
	}
	return ""
}

func (m *CallFuncInfo) GetTragetType() []byte {
	if m != nil {
		return m.TragetType
	}
	return nil
}

func (m *CallFuncInfo) GetCb() string {
	if m != nil {
		return m.Cb
	}
	return ""
}

func (m *CallFuncInfo) GetParam() []*google_protobuf.Any {
	if m != nil {
		return m.Param
	}
	return nil
}

func (m *CallFuncInfo) GetTimeStamp() int64 {
	if m != nil {
		return m.TimeStamp
	}
	return 0
}

type Error struct {
	ErrType string `protobuf:"bytes,1,opt,name=ErrType,json=errType" json:"ErrType,omitempty"`
	ErrMsg  string `protobuf:"bytes,2,opt,name=ErrMsg,json=errMsg" json:"ErrMsg,omitempty"`
}

func (m *Error) Reset()                    { *m = Error{} }
func (m *Error) String() string            { return proto.CompactTextString(m) }
func (*Error) ProtoMessage()               {}
func (*Error) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Error) GetErrType() string {
	if m != nil {
		return m.ErrType
	}
	return ""
}

func (m *Error) GetErrMsg() string {
	if m != nil {
		return m.ErrMsg
	}
	return ""
}

type LoginInput struct {
	UserName string `protobuf:"bytes,1,opt,name=UserName,json=userName" json:"UserName,omitempty"`
	Pswd     string `protobuf:"bytes,2,opt,name=Pswd,json=pswd" json:"Pswd,omitempty"`
}

func (m *LoginInput) Reset()                    { *m = LoginInput{} }
func (m *LoginInput) String() string            { return proto.CompactTextString(m) }
func (*LoginInput) ProtoMessage()               {}
func (*LoginInput) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *LoginInput) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

func (m *LoginInput) GetPswd() string {
	if m != nil {
		return m.Pswd
	}
	return ""
}

type RegistInput struct {
	UserName string `protobuf:"bytes,1,opt,name=UserName,json=userName" json:"UserName,omitempty"`
	Pswd     string `protobuf:"bytes,2,opt,name=Pswd,json=pswd" json:"Pswd,omitempty"`
	Email    string `protobuf:"bytes,3,opt,name=Email,json=email" json:"Email,omitempty"`
}

func (m *RegistInput) Reset()                    { *m = RegistInput{} }
func (m *RegistInput) String() string            { return proto.CompactTextString(m) }
func (*RegistInput) ProtoMessage()               {}
func (*RegistInput) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *RegistInput) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

func (m *RegistInput) GetPswd() string {
	if m != nil {
		return m.Pswd
	}
	return ""
}

func (m *RegistInput) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

type UserInfo struct {
	UserName     string                `protobuf:"bytes,1,opt,name=UserName,json=userName" json:"UserName,omitempty"`
	Uuid         string                `protobuf:"bytes,2,opt,name=Uuid,json=uuid" json:"Uuid,omitempty"`
	OwnCharacter map[string]*Character `protobuf:"bytes,3,rep,name=OwnCharacter,json=ownCharacter" json:"OwnCharacter,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *UserInfo) Reset()                    { *m = UserInfo{} }
func (m *UserInfo) String() string            { return proto.CompactTextString(m) }
func (*UserInfo) ProtoMessage()               {}
func (*UserInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *UserInfo) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

func (m *UserInfo) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

func (m *UserInfo) GetOwnCharacter() map[string]*Character {
	if m != nil {
		return m.OwnCharacter
	}
	return nil
}

type RoomInfo struct {
	Uuid       string               `protobuf:"bytes,1,opt,name=Uuid,json=uuid" json:"Uuid,omitempty"`
	Name       string               `protobuf:"bytes,2,opt,name=Name,json=name" json:"Name,omitempty"`
	GameType   string               `protobuf:"bytes,3,opt,name=GameType,json=gameType" json:"GameType,omitempty"`
	OwnerUuid  string               `protobuf:"bytes,4,opt,name=OwnerUuid,json=ownerUuid" json:"OwnerUuid,omitempty"`
	UserInRoom map[string]*UserInfo `protobuf:"bytes,5,rep,name=UserInRoom,json=userInRoom" json:"UserInRoom,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *RoomInfo) Reset()                    { *m = RoomInfo{} }
func (m *RoomInfo) String() string            { return proto.CompactTextString(m) }
func (*RoomInfo) ProtoMessage()               {}
func (*RoomInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *RoomInfo) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

func (m *RoomInfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *RoomInfo) GetGameType() string {
	if m != nil {
		return m.GameType
	}
	return ""
}

func (m *RoomInfo) GetOwnerUuid() string {
	if m != nil {
		return m.OwnerUuid
	}
	return ""
}

func (m *RoomInfo) GetUserInRoom() map[string]*UserInfo {
	if m != nil {
		return m.UserInRoom
	}
	return nil
}

// character 即是沒有實體之腳色
// entity 則藉由character來初始化
type Character struct {
	Uuid          string                `protobuf:"bytes,1,opt,name=Uuid,json=uuid" json:"Uuid,omitempty"`
	CharacterType string                `protobuf:"bytes,2,opt,name=CharacterType,json=characterType" json:"CharacterType,omitempty"`
	Name          string                `protobuf:"bytes,3,opt,name=Name,json=name" json:"Name,omitempty"`
	Color         *Color                `protobuf:"bytes,4,opt,name=Color,json=color" json:"Color,omitempty"`
	Level         int32                 `protobuf:"varint,5,opt,name=Level,json=level" json:"Level,omitempty"`
	Exp           int32                 `protobuf:"varint,6,opt,name=Exp,json=exp" json:"Exp,omitempty"`
	Ability       *Ability              `protobuf:"bytes,7,opt,name=Ability,json=ability" json:"Ability,omitempty"`
	EquipmentMap  map[string]*Equipment `protobuf:"bytes,8,rep,name=EquipmentMap,json=equipmentMap" json:"EquipmentMap,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *Character) Reset()                    { *m = Character{} }
func (m *Character) String() string            { return proto.CompactTextString(m) }
func (*Character) ProtoMessage()               {}
func (*Character) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *Character) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

func (m *Character) GetCharacterType() string {
	if m != nil {
		return m.CharacterType
	}
	return ""
}

func (m *Character) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Character) GetColor() *Color {
	if m != nil {
		return m.Color
	}
	return nil
}

func (m *Character) GetLevel() int32 {
	if m != nil {
		return m.Level
	}
	return 0
}

func (m *Character) GetExp() int32 {
	if m != nil {
		return m.Exp
	}
	return 0
}

func (m *Character) GetAbility() *Ability {
	if m != nil {
		return m.Ability
	}
	return nil
}

func (m *Character) GetEquipmentMap() map[string]*Equipment {
	if m != nil {
		return m.EquipmentMap
	}
	return nil
}

type Color struct {
	R int32 `protobuf:"varint,1,opt,name=R,json=r" json:"R,omitempty"`
	G int32 `protobuf:"varint,2,opt,name=G,json=g" json:"G,omitempty"`
	B int32 `protobuf:"varint,3,opt,name=B,json=b" json:"B,omitempty"`
}

func (m *Color) Reset()                    { *m = Color{} }
func (m *Color) String() string            { return proto.CompactTextString(m) }
func (*Color) ProtoMessage()               {}
func (*Color) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *Color) GetR() int32 {
	if m != nil {
		return m.R
	}
	return 0
}

func (m *Color) GetG() int32 {
	if m != nil {
		return m.G
	}
	return 0
}

func (m *Color) GetB() int32 {
	if m != nil {
		return m.B
	}
	return 0
}

type Equipment struct {
	Name      string   `protobuf:"bytes,1,opt,name=Name,json=name" json:"Name,omitempty"`
	Type      string   `protobuf:"bytes,2,opt,name=Type,json=type" json:"Type,omitempty"`
	Uuid      string   `protobuf:"bytes,3,opt,name=Uuid,json=uuid" json:"Uuid,omitempty"`
	Color     *Color   `protobuf:"bytes,4,opt,name=Color,json=color" json:"Color,omitempty"`
	Ability   *Ability `protobuf:"bytes,5,opt,name=Ability,json=ability" json:"Ability,omitempty"`
	CD        int32    `protobuf:"varint,6,opt,name=CD,json=cD" json:"CD,omitempty"`
	Usable    int32    `protobuf:"varint,7,opt,name=Usable,json=usable" json:"Usable,omitempty"`
	Inventory int32    `protobuf:"varint,8,opt,name=Inventory,json=inventory" json:"Inventory,omitempty"`
}

func (m *Equipment) Reset()                    { *m = Equipment{} }
func (m *Equipment) String() string            { return proto.CompactTextString(m) }
func (*Equipment) ProtoMessage()               {}
func (*Equipment) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *Equipment) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Equipment) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Equipment) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

func (m *Equipment) GetColor() *Color {
	if m != nil {
		return m.Color
	}
	return nil
}

func (m *Equipment) GetAbility() *Ability {
	if m != nil {
		return m.Ability
	}
	return nil
}

func (m *Equipment) GetCD() int32 {
	if m != nil {
		return m.CD
	}
	return 0
}

func (m *Equipment) GetUsable() int32 {
	if m != nil {
		return m.Usable
	}
	return 0
}

func (m *Equipment) GetInventory() int32 {
	if m != nil {
		return m.Inventory
	}
	return 0
}

type Ability struct {
	ATK  int32   `protobuf:"varint,1,opt,name=ATK,json=aTK" json:"ATK,omitempty"`
	DEF  int32   `protobuf:"varint,2,opt,name=DEF,json=dEF" json:"DEF,omitempty"`
	SPD  float32 `protobuf:"fixed32,3,opt,name=SPD,json=sPD" json:"SPD,omitempty"`
	MP   int32   `protobuf:"varint,4,opt,name=MP,json=mP" json:"MP,omitempty"`
	MAKT int32   `protobuf:"varint,5,opt,name=MAKT,json=mAKT" json:"MAKT,omitempty"`
}

func (m *Ability) Reset()                    { *m = Ability{} }
func (m *Ability) String() string            { return proto.CompactTextString(m) }
func (*Ability) ProtoMessage()               {}
func (*Ability) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *Ability) GetATK() int32 {
	if m != nil {
		return m.ATK
	}
	return 0
}

func (m *Ability) GetDEF() int32 {
	if m != nil {
		return m.DEF
	}
	return 0
}

func (m *Ability) GetSPD() float32 {
	if m != nil {
		return m.SPD
	}
	return 0
}

func (m *Ability) GetMP() int32 {
	if m != nil {
		return m.MP
	}
	return 0
}

func (m *Ability) GetMAKT() int32 {
	if m != nil {
		return m.MAKT
	}
	return 0
}

func init() {
	proto.RegisterType((*TransForm)(nil), "msg.TransForm")
	proto.RegisterType((*Position)(nil), "msg.Position")
	proto.RegisterType((*CallFuncInfo)(nil), "msg.CallFuncInfo")
	proto.RegisterType((*Error)(nil), "msg.Error")
	proto.RegisterType((*LoginInput)(nil), "msg.LoginInput")
	proto.RegisterType((*RegistInput)(nil), "msg.RegistInput")
	proto.RegisterType((*UserInfo)(nil), "msg.UserInfo")
	proto.RegisterType((*RoomInfo)(nil), "msg.RoomInfo")
	proto.RegisterType((*Character)(nil), "msg.Character")
	proto.RegisterType((*Color)(nil), "msg.Color")
	proto.RegisterType((*Equipment)(nil), "msg.Equipment")
	proto.RegisterType((*Ability)(nil), "msg.Ability")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Rpc service

type RpcClient interface {
	SyncPos(ctx context.Context, opts ...grpc.CallOption) (Rpc_SyncPosClient, error)
	CallMethod(ctx context.Context, opts ...grpc.CallOption) (Rpc_CallMethodClient, error)
	ErrorPipLine(ctx context.Context, opts ...grpc.CallOption) (Rpc_ErrorPipLineClient, error)
	Login(ctx context.Context, in *LoginInput, opts ...grpc.CallOption) (*UserInfo, error)
	CreateAccount(ctx context.Context, in *RegistInput, opts ...grpc.CallOption) (*Error, error)
}

type rpcClient struct {
	cc *grpc.ClientConn
}

func NewRpcClient(cc *grpc.ClientConn) RpcClient {
	return &rpcClient{cc}
}

func (c *rpcClient) SyncPos(ctx context.Context, opts ...grpc.CallOption) (Rpc_SyncPosClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Rpc_serviceDesc.Streams[0], c.cc, "/msg.Rpc/SyncPos", opts...)
	if err != nil {
		return nil, err
	}
	x := &rpcSyncPosClient{stream}
	return x, nil
}

type Rpc_SyncPosClient interface {
	Send(*Position) error
	Recv() (*Position, error)
	grpc.ClientStream
}

type rpcSyncPosClient struct {
	grpc.ClientStream
}

func (x *rpcSyncPosClient) Send(m *Position) error {
	return x.ClientStream.SendMsg(m)
}

func (x *rpcSyncPosClient) Recv() (*Position, error) {
	m := new(Position)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *rpcClient) CallMethod(ctx context.Context, opts ...grpc.CallOption) (Rpc_CallMethodClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Rpc_serviceDesc.Streams[1], c.cc, "/msg.Rpc/CallMethod", opts...)
	if err != nil {
		return nil, err
	}
	x := &rpcCallMethodClient{stream}
	return x, nil
}

type Rpc_CallMethodClient interface {
	Send(*CallFuncInfo) error
	Recv() (*CallFuncInfo, error)
	grpc.ClientStream
}

type rpcCallMethodClient struct {
	grpc.ClientStream
}

func (x *rpcCallMethodClient) Send(m *CallFuncInfo) error {
	return x.ClientStream.SendMsg(m)
}

func (x *rpcCallMethodClient) Recv() (*CallFuncInfo, error) {
	m := new(CallFuncInfo)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *rpcClient) ErrorPipLine(ctx context.Context, opts ...grpc.CallOption) (Rpc_ErrorPipLineClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Rpc_serviceDesc.Streams[2], c.cc, "/msg.Rpc/ErrorPipLine", opts...)
	if err != nil {
		return nil, err
	}
	x := &rpcErrorPipLineClient{stream}
	return x, nil
}

type Rpc_ErrorPipLineClient interface {
	Send(*Error) error
	Recv() (*Error, error)
	grpc.ClientStream
}

type rpcErrorPipLineClient struct {
	grpc.ClientStream
}

func (x *rpcErrorPipLineClient) Send(m *Error) error {
	return x.ClientStream.SendMsg(m)
}

func (x *rpcErrorPipLineClient) Recv() (*Error, error) {
	m := new(Error)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *rpcClient) Login(ctx context.Context, in *LoginInput, opts ...grpc.CallOption) (*UserInfo, error) {
	out := new(UserInfo)
	err := grpc.Invoke(ctx, "/msg.Rpc/Login", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rpcClient) CreateAccount(ctx context.Context, in *RegistInput, opts ...grpc.CallOption) (*Error, error) {
	out := new(Error)
	err := grpc.Invoke(ctx, "/msg.Rpc/CreateAccount", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Rpc service

type RpcServer interface {
	SyncPos(Rpc_SyncPosServer) error
	CallMethod(Rpc_CallMethodServer) error
	ErrorPipLine(Rpc_ErrorPipLineServer) error
	Login(context.Context, *LoginInput) (*UserInfo, error)
	CreateAccount(context.Context, *RegistInput) (*Error, error)
}

func RegisterRpcServer(s *grpc.Server, srv RpcServer) {
	s.RegisterService(&_Rpc_serviceDesc, srv)
}

func _Rpc_SyncPos_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(RpcServer).SyncPos(&rpcSyncPosServer{stream})
}

type Rpc_SyncPosServer interface {
	Send(*Position) error
	Recv() (*Position, error)
	grpc.ServerStream
}

type rpcSyncPosServer struct {
	grpc.ServerStream
}

func (x *rpcSyncPosServer) Send(m *Position) error {
	return x.ServerStream.SendMsg(m)
}

func (x *rpcSyncPosServer) Recv() (*Position, error) {
	m := new(Position)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Rpc_CallMethod_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(RpcServer).CallMethod(&rpcCallMethodServer{stream})
}

type Rpc_CallMethodServer interface {
	Send(*CallFuncInfo) error
	Recv() (*CallFuncInfo, error)
	grpc.ServerStream
}

type rpcCallMethodServer struct {
	grpc.ServerStream
}

func (x *rpcCallMethodServer) Send(m *CallFuncInfo) error {
	return x.ServerStream.SendMsg(m)
}

func (x *rpcCallMethodServer) Recv() (*CallFuncInfo, error) {
	m := new(CallFuncInfo)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Rpc_ErrorPipLine_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(RpcServer).ErrorPipLine(&rpcErrorPipLineServer{stream})
}

type Rpc_ErrorPipLineServer interface {
	Send(*Error) error
	Recv() (*Error, error)
	grpc.ServerStream
}

type rpcErrorPipLineServer struct {
	grpc.ServerStream
}

func (x *rpcErrorPipLineServer) Send(m *Error) error {
	return x.ServerStream.SendMsg(m)
}

func (x *rpcErrorPipLineServer) Recv() (*Error, error) {
	m := new(Error)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Rpc_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RpcServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/msg.Rpc/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RpcServer).Login(ctx, req.(*LoginInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _Rpc_CreateAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegistInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RpcServer).CreateAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/msg.Rpc/CreateAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RpcServer).CreateAccount(ctx, req.(*RegistInput))
	}
	return interceptor(ctx, in, info, handler)
}

var _Rpc_serviceDesc = grpc.ServiceDesc{
	ServiceName: "msg.Rpc",
	HandlerType: (*RpcServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Login",
			Handler:    _Rpc_Login_Handler,
		},
		{
			MethodName: "CreateAccount",
			Handler:    _Rpc_CreateAccount_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SyncPos",
			Handler:       _Rpc_SyncPos_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "CallMethod",
			Handler:       _Rpc_CallMethod_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "ErrorPipLine",
			Handler:       _Rpc_ErrorPipLine_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "message.proto",
}

func init() { proto.RegisterFile("message.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 979 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x55, 0x5f, 0x6f, 0xe3, 0x44,
	0x10, 0x3f, 0xc7, 0xd9, 0x24, 0x9e, 0x26, 0xf7, 0x67, 0x75, 0x42, 0x26, 0xe2, 0x4f, 0x65, 0x4e,
	0xa8, 0x20, 0x08, 0x5c, 0x79, 0x39, 0x10, 0x3c, 0x84, 0x24, 0x3d, 0x55, 0x6d, 0xaf, 0xd6, 0x36,
	0x95, 0x38, 0xde, 0x36, 0xce, 0x9e, 0x6b, 0x11, 0x7b, 0xcd, 0xda, 0x6e, 0x1a, 0x3e, 0x01, 0x1f,
	0x8b, 0x8f, 0xc0, 0x1b, 0xef, 0x7c, 0x0a, 0x78, 0x43, 0x3b, 0xbb, 0x71, 0x1d, 0x7a, 0x3a, 0x4e,
	0x3c, 0xc5, 0xbf, 0x99, 0xd9, 0xd9, 0xf9, 0xfd, 0x66, 0x32, 0x0b, 0x83, 0x54, 0x14, 0x05, 0x8f,
	0xc5, 0x28, 0x57, 0xb2, 0x94, 0xd4, 0x4d, 0x8b, 0x78, 0xe8, 0xf1, 0x6c, 0x63, 0x70, 0x20, 0xc1,
	0x9b, 0x2b, 0x9e, 0x15, 0x47, 0x52, 0xa5, 0xb4, 0x0f, 0xce, 0x0f, 0xbe, 0xb3, 0xef, 0x1c, 0xb4,
	0x98, 0x73, 0xa3, 0xd1, 0x4b, 0xbf, 0x65, 0xd0, 0x46, 0xa3, 0x1f, 0x7d, 0xd7, 0xa0, 0x5f, 0xe8,
	0x43, 0x70, 0x5f, 0xf2, 0xb5, 0xdf, 0x46, 0xec, 0x6e, 0xf8, 0x9a, 0x3e, 0x06, 0x12, 0x26, 0x65,
	0x74, 0xe5, 0x13, 0xb4, 0x91, 0x5c, 0x03, 0x1d, 0xc7, 0xe4, 0xda, 0xef, 0x98, 0x38, 0x25, 0xd7,
	0xc1, 0xaf, 0x0e, 0xf4, 0x42, 0x59, 0x24, 0x65, 0x22, 0x33, 0xfa, 0x14, 0x3a, 0xa1, 0x2c, 0xce,
	0x78, 0xee, 0x3b, 0xfb, 0xee, 0xc1, 0xde, 0xe1, 0xbb, 0xa3, 0xb4, 0x88, 0x47, 0x5b, 0xf7, 0xc8,
	0xf8, 0x66, 0x59, 0xa9, 0x36, 0xac, 0x93, 0x23, 0x18, 0x1e, 0xc3, 0x5e, 0xc3, 0xac, 0x2f, 0xf8,
	0x49, 0x6c, 0xb0, 0x68, 0x8f, 0xe9, 0x4f, 0xfa, 0x04, 0xc8, 0x35, 0x5f, 0x55, 0x02, 0x4b, 0xdf,
	0x3b, 0xbc, 0x8f, 0x29, 0x6b, 0x8e, 0xcc, 0x38, 0xbf, 0x69, 0x3d, 0x73, 0x82, 0xdf, 0x1c, 0xe8,
	0x4f, 0xf8, 0x6a, 0x75, 0x54, 0x65, 0xd1, 0x71, 0xf6, 0x4a, 0xd2, 0x21, 0xf4, 0x58, 0x95, 0x65,
	0x49, 0xfc, 0x42, 0xda, 0x8c, 0x3d, 0x65, 0xb1, 0xf6, 0xcd, 0xb9, 0x8a, 0x45, 0x79, 0xbc, 0xc4,
	0xcc, 0x1e, 0xeb, 0x95, 0x16, 0xd3, 0x0f, 0x00, 0xe6, 0x8a, 0xc7, 0xa2, 0x9c, 0x6f, 0x72, 0x81,
	0x22, 0xf5, 0x19, 0x94, 0xb5, 0x85, 0xde, 0x87, 0xd6, 0x64, 0x81, 0x62, 0x79, 0xac, 0x15, 0x2d,
	0xe8, 0xa7, 0x40, 0x42, 0xae, 0x78, 0xea, 0x13, 0x64, 0xfd, 0x78, 0x14, 0x4b, 0x19, 0xaf, 0x6c,
	0x8b, 0x16, 0xd5, 0xab, 0xd1, 0x38, 0xdb, 0x30, 0x92, 0xeb, 0x10, 0xfa, 0x1e, 0x78, 0xf3, 0x24,
	0x15, 0x17, 0x25, 0x4f, 0x73, 0xd4, 0xd1, 0x65, 0x5e, 0xb9, 0x35, 0x04, 0x5f, 0x03, 0x99, 0x29,
	0x25, 0x15, 0xf5, 0xa1, 0x3b, 0x53, 0x0a, 0xef, 0x37, 0x95, 0x77, 0x85, 0x81, 0xf4, 0x1d, 0xe8,
	0xcc, 0x94, 0x3a, 0x2b, 0x62, 0x5b, 0x76, 0x47, 0x20, 0x0a, 0xbe, 0x05, 0x38, 0x95, 0x71, 0x92,
	0x1d, 0x67, 0x79, 0x55, 0x6a, 0x7a, 0x97, 0x85, 0x50, 0x2f, 0x78, 0xba, 0x4d, 0xd0, 0xab, 0x2c,
	0xa6, 0x14, 0xda, 0x61, 0xb1, 0xde, 0xd2, 0x6e, 0xe7, 0xc5, 0x7a, 0x19, 0x5c, 0xc0, 0x1e, 0x13,
	0x71, 0x52, 0x94, 0xff, 0xeb, 0xb8, 0x9e, 0x96, 0x59, 0xca, 0x93, 0x15, 0x8a, 0xe5, 0x31, 0x22,
	0x34, 0x08, 0x7e, 0x77, 0x4c, 0x9a, 0x6d, 0x33, 0xde, 0x94, 0xf2, 0xb2, 0x4a, 0xea, 0x94, 0x55,
	0x95, 0x2c, 0xe9, 0x04, 0xfa, 0xe7, 0xeb, 0x6c, 0x72, 0xc5, 0x15, 0x8f, 0x4a, 0xa1, 0x7c, 0x17,
	0xb5, 0xfd, 0x10, 0xdb, 0xbf, 0x4d, 0x3a, 0x6a, 0x46, 0x98, 0xb9, 0xea, 0xcb, 0x86, 0x69, 0x78,
	0x0e, 0x8f, 0xee, 0x84, 0xbc, 0xed, 0x8c, 0xd5, 0xa7, 0x9a, 0x33, 0xf6, 0xb7, 0x03, 0x3d, 0x26,
	0x65, 0x8a, 0x94, 0xb6, 0x65, 0x3b, 0x8d, 0xb2, 0x29, 0xb4, 0x91, 0xa2, 0xa5, 0x92, 0x69, 0x7a,
	0x43, 0xe8, 0x3d, 0xe7, 0xa9, 0xa8, 0xa7, 0xc9, 0x63, 0xbd, 0xd8, 0x62, 0x3d, 0x0f, 0xe7, 0xeb,
	0x4c, 0x28, 0x4c, 0x64, 0x46, 0xca, 0x93, 0x5b, 0x03, 0xfd, 0x0e, 0xc0, 0x70, 0xd5, 0x77, 0xda,
	0xf1, 0x7a, 0x1f, 0xab, 0xdb, 0x16, 0x31, 0xba, 0xf5, 0x1b, 0x01, 0xa0, 0xaa, 0x0d, 0xc3, 0x53,
	0x78, 0xf0, 0x2f, 0xf7, 0x6b, 0xc8, 0x7f, 0xb4, 0x4b, 0x7e, 0xb0, 0xa3, 0x70, 0x93, 0xfb, 0x9f,
	0x2d, 0xf0, 0x6a, 0x51, 0x5e, 0x4b, 0xfe, 0x09, 0x0c, 0xea, 0x00, 0x64, 0x6b, 0x54, 0x18, 0x44,
	0x4d, 0x63, 0x2d, 0x91, 0xdb, 0x90, 0x68, 0x1f, 0xc8, 0x44, 0xae, 0xa4, 0x42, 0x09, 0xf6, 0x0e,
	0xc1, 0x74, 0x40, 0x5b, 0x18, 0x89, 0xf4, 0x8f, 0x1e, 0xb1, 0x53, 0x71, 0x2d, 0x56, 0xb8, 0x90,
	0x08, 0x23, 0x2b, 0x0d, 0x34, 0x9d, 0xd9, 0x8d, 0xf9, 0x23, 0x11, 0xe6, 0x8a, 0x9b, 0x9c, 0x7e,
	0x0c, 0xdd, 0xf1, 0x22, 0x59, 0x25, 0xe5, 0xc6, 0xef, 0x62, 0xae, 0x3e, 0xe6, 0xb2, 0x36, 0xd6,
	0xe5, 0xe6, 0x83, 0x4e, 0xa1, 0x3f, 0xfb, 0xb9, 0x4a, 0xf2, 0x54, 0x64, 0xa5, 0xde, 0x58, 0x3d,
	0x14, 0x77, 0x7f, 0xb7, 0xf5, 0xa3, 0x66, 0x88, 0x1d, 0x30, 0xd1, 0x30, 0xe9, 0x01, 0xbb, 0x13,
	0xf2, 0xb6, 0x03, 0x56, 0x1f, 0x6c, 0x8a, 0xfc, 0xd4, 0x0a, 0xa1, 0x17, 0x34, 0xc3, 0x14, 0x84,
	0x39, 0x88, 0x9e, 0xe3, 0x61, 0xc2, 0x9c, 0x58, 0xa3, 0xef, 0x51, 0x3e, 0xc2, 0x9c, 0x45, 0xf0,
	0x87, 0x03, 0x5e, 0x9d, 0xab, 0x56, 0xd7, 0x69, 0xa8, 0x4b, 0xa1, 0xdd, 0x68, 0x47, 0xbb, 0xb4,
	0x5d, 0xc0, 0xfe, 0xb9, 0x8d, 0xfe, 0xfd, 0x77, 0x17, 0x1a, 0xea, 0x92, 0x37, 0xa9, 0xab, 0x57,
	0xe4, 0xd4, 0xb6, 0xa5, 0x15, 0x4d, 0xf5, 0xd6, 0xba, 0x2c, 0xf8, 0x62, 0x25, 0xb0, 0x29, 0x84,
	0x75, 0x2a, 0x44, 0x7a, 0xfc, 0x8f, 0xb3, 0x6b, 0x91, 0x95, 0x52, 0x6d, 0xfc, 0x1e, 0xba, 0xbc,
	0x64, 0x6b, 0x08, 0xa2, 0xfa, 0x36, 0xad, 0xe9, 0x78, 0x7e, 0x62, 0x05, 0x71, 0xf9, 0xfc, 0x44,
	0x5b, 0xa6, 0xb3, 0x23, 0x2b, 0x8a, 0xbb, 0x9c, 0x1d, 0x69, 0xcb, 0x45, 0x38, 0xb5, 0xaf, 0x9a,
	0x5b, 0x84, 0x53, 0x5d, 0xc6, 0x59, 0x88, 0x6c, 0x08, 0x6b, 0xa5, 0xa1, 0x26, 0x7d, 0x36, 0x3e,
	0x99, 0xdb, 0x19, 0x6a, 0xa7, 0xe3, 0x93, 0xf9, 0xe1, 0x5f, 0x0e, 0xb8, 0x2c, 0x8f, 0xe8, 0xe7,
	0xd0, 0xbd, 0xd8, 0x64, 0x51, 0x28, 0x0b, 0x3a, 0xd8, 0x79, 0xb7, 0x86, 0xbb, 0x30, 0xb8, 0x77,
	0xe0, 0x7c, 0xe9, 0xd0, 0x67, 0x00, 0xfa, 0xb1, 0x39, 0x13, 0xe5, 0x95, 0x5c, 0xd2, 0x47, 0x46,
	0xaa, 0xc6, 0xeb, 0x33, 0xbc, 0x6b, 0xb2, 0x27, 0x3f, 0x83, 0x3e, 0x2e, 0xf9, 0x30, 0xc9, 0x4f,
	0x93, 0x4c, 0x50, 0x23, 0x33, 0x9a, 0x86, 0x8d, 0x6f, 0x1b, 0xfd, 0x09, 0x10, 0xdc, 0xeb, 0xf4,
	0x01, 0xba, 0x6e, 0x77, 0xfc, 0x70, 0xf7, 0x9f, 0x1a, 0xdc, 0xa3, 0x5f, 0xc0, 0x60, 0xa2, 0x04,
	0x2f, 0xc5, 0x38, 0x8a, 0x64, 0x95, 0x95, 0xf4, 0xa1, 0x59, 0x15, 0xb7, 0x8b, 0x7d, 0x37, 0xff,
	0xa2, 0x83, 0x2f, 0xd4, 0x57, 0xff, 0x04, 0x00, 0x00, 0xff, 0xff, 0xe8, 0x16, 0xbe, 0x48, 0x55,
	0x08, 0x00, 0x00,
}
