package globals

import (
	pb "github.com/PretendoNetwork/grpc-go/friends"
	"github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/plogger-go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

var Logger = plogger.NewLogger()
var NEXServer *nex.Server
var GRPCFriendsClientConnection *grpc.ClientConn
var GRPCFriendsClient pb.FriendsClient
var GRPCFriendsCommonMetadata metadata.MD
