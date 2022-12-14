syntax = "proto3";

package proto;
option go_package="./{{.package}}/proto";

message Request {
  string ping = 1;
}

message Response {
  string pong = 1;
}

service {{.serviceName}} {
  rpc Ping(Request) returns(Response);
}
