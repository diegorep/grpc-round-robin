// GENERATED CODE -- DO NOT EDIT!

'use strict';
var grpc = require('grpc');
var twos_pb = require('../messages/twos_pb.js');

function serialize_GetTwoRequest(arg) {
  if (!(arg instanceof twos_pb.GetTwoRequest)) {
    throw new Error('Expected argument of type GetTwoRequest');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_GetTwoRequest(buffer_arg) {
  return twos_pb.GetTwoRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_GetTwoResponse(arg) {
  if (!(arg instanceof twos_pb.GetTwoResponse)) {
    throw new Error('Expected argument of type GetTwoResponse');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_GetTwoResponse(buffer_arg) {
  return twos_pb.GetTwoResponse.deserializeBinary(new Uint8Array(buffer_arg));
}


var TwosService = exports.TwosService = {
  echoTwo: {
    path: '/Twos/EchoTwo',
    requestStream: false,
    responseStream: false,
    requestType: twos_pb.GetTwoRequest,
    responseType: twos_pb.GetTwoResponse,
    requestSerialize: serialize_GetTwoRequest,
    requestDeserialize: deserialize_GetTwoRequest,
    responseSerialize: serialize_GetTwoResponse,
    responseDeserialize: deserialize_GetTwoResponse,
  },
};

exports.TwosClient = grpc.makeGenericClientConstructor(TwosService);
