var messages = require('./messages/twos_pb.js');
var services = require('./services/twos_grpc_pb.js');
const grpc = require('grpc');
const server = new grpc.Server();


function echoTwo(call, done) {
  if (call.request.value.value === 2) {
    two = messages.twos_pb2.Two({value: 2});
  } else {
    two = messages.twos_pb2.Two({value: 2});
  }
  response = messages.twos_pb2.GetTwoResponse();
  response.setValue(two);
  done(null, response);
}

server.addService(services.TwosService, {
  echoTwo: echoTwo
});

server.bind('0.0.0.0:50050', grpc.ServerCredentials.createInsecure());

server.start();
console.log('grpc server running on port:', '0.0.0.0:50050');
