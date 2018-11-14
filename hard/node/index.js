var messages = require('./messages/twos_pb.js');
var services = require('./services/twos_grpc_pb.js');
const grpc = require('grpc');
const server = new grpc.Server();


function echoTwo(call, done) {
  requestValue = call.request.getValue();
  if (requestValue.getValue() === 2) {
    twoVal = 2.0;
  } else {
    twoVal = -2.0;
  }
  two = new messages.Two();
  two.setValue(twoVal);
  response = new messages.GetTwoResponse();
  response.setValue(two);
  done(null, response);
}

server.addService(services.TwosService, {
  echoTwo: echoTwo
});

server.bind('0.0.0.0:50050', grpc.ServerCredentials.createInsecure());

server.start();
console.log('grpc server running on port:', '0.0.0.0:50050');
