package grpc

import (
	"fmt"
	"log"
	"net"

	"github.com/Viquad/crud-audit-service/pkg/domain/audit"
	"google.golang.org/grpc"
)

type Server struct {
	audit audit.AuditServiceServer
	grpc  *grpc.Server
}

func NewServer(audit audit.AuditServiceServer) *Server {
	return &Server{
		audit: audit,
		grpc:  grpc.NewServer(),
	}
}

func (s *Server) ListenAndServe(port int) error {
	addr := fmt.Sprintf(":%d", port)

	lis, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}

	audit.RegisterAuditServiceServer(s.grpc, s.audit)

	log.Println("Audit service started")

	if err := s.grpc.Serve(lis); err != nil {
		return err
	}

	return nil
}
