// Code generated by protoc-gen-entgrpc. DO NOT EDIT.
package entpb

import (
	context "context"
	base64 "encoding/base64"
	entproto "entgo.io/contrib/entproto"
	ent "entgo.io/contrib/entproto/internal/todo/ent"
	nilexample "entgo.io/contrib/entproto/internal/todo/ent/nilexample"
	runtime "entgo.io/contrib/entproto/runtime"
	sqlgraph "entgo.io/ent/dialect/sql/sqlgraph"
	fmt "fmt"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
	strconv "strconv"
)

// NilExampleService implements NilExampleServiceServer
type NilExampleService struct {
	client *ent.Client
	UnimplementedNilExampleServiceServer
}

// NewNilExampleService returns a new NilExampleService
func NewNilExampleService(client *ent.Client) *NilExampleService {
	return &NilExampleService{
		client: client,
	}
}

// toProtoNilExample transforms the ent type to the pb type
func toProtoNilExample(e *ent.NilExample) (*NilExample, error) {
	v := &NilExample{}
	id := int32(e.ID)
	v.Id = id
	if e.StrNil != nil {
		strnil := wrapperspb.String(*e.StrNil)
		v.StrNil = strnil
	}
	if e.TimeNil != nil {
		timenil := timestamppb.New(*e.TimeNil)
		v.TimeNil = timenil
	}
	return v, nil
}

// Create implements NilExampleServiceServer.Create
func (svc *NilExampleService) Create(ctx context.Context, req *CreateNilExampleRequest) (*NilExample, error) {
	nilexample := req.GetNilExample()
	m := svc.client.NilExample.Create()
	if nilexample.GetStrNil() != nil {
		nilexampleStrNil := nilexample.GetStrNil().GetValue()
		m.SetStrNil(nilexampleStrNil)
	}
	if nilexample.GetTimeNil() != nil {
		nilexampleTimeNil := runtime.ExtractTime(nilexample.GetTimeNil())
		m.SetTimeNil(nilexampleTimeNil)
	}
	res, err := m.Save(ctx)
	switch {
	case err == nil:
		proto, err := toProtoNilExample(res)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "internal error: %s", err)
		}
		return proto, nil
	case sqlgraph.IsUniqueConstraintError(err):
		return nil, status.Errorf(codes.AlreadyExists, "already exists: %s", err)
	case ent.IsConstraintError(err):
		return nil, status.Errorf(codes.InvalidArgument, "invalid argument: %s", err)
	default:
		return nil, status.Errorf(codes.Internal, "internal error: %s", err)
	}

}

// Get implements NilExampleServiceServer.Get
func (svc *NilExampleService) Get(ctx context.Context, req *GetNilExampleRequest) (*NilExample, error) {
	var (
		err error
		get *ent.NilExample
	)
	id := int(req.GetId())
	switch req.GetView() {
	case GetNilExampleRequest_VIEW_UNSPECIFIED, GetNilExampleRequest_BASIC:
		get, err = svc.client.NilExample.Get(ctx, id)
	case GetNilExampleRequest_WITH_EDGE_IDS:
		get, err = svc.client.NilExample.Query().
			Where(nilexample.ID(id)).
			Only(ctx)
	default:
		return nil, status.Error(codes.InvalidArgument, "invalid argument: unknown view")
	}
	switch {
	case err == nil:
		return toProtoNilExample(get)
	case ent.IsNotFound(err):
		return nil, status.Errorf(codes.NotFound, "not found: %s", err)
	default:
		return nil, status.Errorf(codes.Internal, "internal error: %s", err)
	}
	return nil, nil

}

// Update implements NilExampleServiceServer.Update
func (svc *NilExampleService) Update(ctx context.Context, req *UpdateNilExampleRequest) (*NilExample, error) {
	nilexample := req.GetNilExample()
	nilexampleID := int(nilexample.GetId())
	m := svc.client.NilExample.UpdateOneID(nilexampleID)
	if nilexample.GetStrNil() != nil {
		nilexampleStrNil := nilexample.GetStrNil().GetValue()
		m.SetStrNil(nilexampleStrNil)
	}
	if nilexample.GetTimeNil() != nil {
		nilexampleTimeNil := runtime.ExtractTime(nilexample.GetTimeNil())
		m.SetTimeNil(nilexampleTimeNil)
	}
	res, err := m.Save(ctx)
	switch {
	case err == nil:
		proto, err := toProtoNilExample(res)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "internal error: %s", err)
		}
		return proto, nil
	case sqlgraph.IsUniqueConstraintError(err):
		return nil, status.Errorf(codes.AlreadyExists, "already exists: %s", err)
	case ent.IsConstraintError(err):
		return nil, status.Errorf(codes.InvalidArgument, "invalid argument: %s", err)
	default:
		return nil, status.Errorf(codes.Internal, "internal error: %s", err)
	}

}

// Delete implements NilExampleServiceServer.Delete
func (svc *NilExampleService) Delete(ctx context.Context, req *DeleteNilExampleRequest) (*emptypb.Empty, error) {
	var err error
	id := int(req.GetId())
	err = svc.client.NilExample.DeleteOneID(id).Exec(ctx)
	switch {
	case err == nil:
		return &emptypb.Empty{}, nil
	case ent.IsNotFound(err):
		return nil, status.Errorf(codes.NotFound, "not found: %s", err)
	default:
		return nil, status.Errorf(codes.Internal, "internal error: %s", err)
	}

}

// List implements NilExampleServiceServer.List
func (svc *NilExampleService) List(ctx context.Context, req *ListNilExampleRequest) (*ListNilExampleResponse, error) {
	var (
		err      error
		entList  []*ent.NilExample
		pageSize int
	)
	pageSize = int(req.GetPageSize())
	switch {
	case pageSize < 0:
		return nil, status.Errorf(codes.InvalidArgument, "page size cannot be less than zero")
	case pageSize == 0 || pageSize > entproto.MaxPageSize:
		pageSize = entproto.MaxPageSize
	}
	listQuery := svc.client.NilExample.Query().
		Order(ent.Desc(nilexample.FieldID)).
		Limit(pageSize + 1)
	if req.GetPageToken() != "" {
		bytes, err := base64.StdEncoding.DecodeString(req.PageToken)
		if err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "page token is invalid")
		}
		token, err := strconv.ParseInt(string(bytes), 10, 32)
		if err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "page token is invalid")
		}
		pageToken := int(token)
		listQuery = listQuery.
			Where(nilexample.IDLTE(pageToken))
	}
	switch req.GetView() {
	case ListNilExampleRequest_VIEW_UNSPECIFIED, ListNilExampleRequest_BASIC:
		entList, err = listQuery.All(ctx)
	case ListNilExampleRequest_WITH_EDGE_IDS:
		entList, err = listQuery.
			All(ctx)
	}
	switch {
	case err == nil:
		var nextPageToken string
		if len(entList) == pageSize+1 {
			nextPageToken = base64.StdEncoding.EncodeToString(
				[]byte(fmt.Sprintf("%v", entList[len(entList)-1].ID)))
			entList = entList[:len(entList)-1]
		}
		var pbList []*NilExample
		for _, entEntity := range entList {
			pbEntity, err := toProtoNilExample(entEntity)
			if err != nil {
				return nil, status.Errorf(codes.Internal, "internal error: %s", err)
			}
			pbList = append(pbList, pbEntity)
		}
		return &ListNilExampleResponse{
			NilExampleList: pbList,
			NextPageToken:  nextPageToken,
		}, nil
	default:
		return nil, status.Errorf(codes.Internal, "internal error: %s", err)
	}

}
