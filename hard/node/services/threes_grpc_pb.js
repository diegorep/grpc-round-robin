// GENERATED CODE -- DO NOT EDIT!

'use strict';
var grpc = require('grpc');
var threes_pb = require('../messages/threes_pb.js');

function serialize_GetThreeRequest(arg) {
  if (!(arg instanceof threes_pb.GetThreeRequest)) {
    throw new Error('Expected argument of type GetThreeRequest');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_GetThreeRequest(buffer_arg) {
  return threes_pb.GetThreeRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_GetThreeResponse(arg) {
  if (!(arg instanceof threes_pb.GetThreeResponse)) {
    throw new Error('Expected argument of type GetThreeResponse');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_GetThreeResponse(buffer_arg) {
  return threes_pb.GetThreeResponse.deserializeBinary(new Uint8Array(buffer_arg));
}


var ThreesService = exports.ThreesService = {
  echoThree: {
    path: '/Threes/EchoThree',
    requestStream: false,
    responseStream: false,
    requestType: threes_pb.GetThreeRequest,
    responseType: threes_pb.GetThreeResponse,
    requestSerialize: serialize_GetThreeRequest,
    requestDeserialize: deserialize_GetThreeRequest,
    responseSerialize: serialize_GetThreeResponse,
    responseDeserialize: deserialize_GetThreeResponse,
  },
};

exports.ThreesClient = grpc.makeGenericClientConstructor(ThreesService);
