// GENERATED CODE -- DO NOT EDIT!

'use strict';
var grpc = require('grpc');
var ones_pb = require('../messages/ones_pb.js');

function serialize_GetOneRequest(arg) {
  if (!(arg instanceof ones_pb.GetOneRequest)) {
    throw new Error('Expected argument of type GetOneRequest');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_GetOneRequest(buffer_arg) {
  return ones_pb.GetOneRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_GetOneResponse(arg) {
  if (!(arg instanceof ones_pb.GetOneResponse)) {
    throw new Error('Expected argument of type GetOneResponse');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_GetOneResponse(buffer_arg) {
  return ones_pb.GetOneResponse.deserializeBinary(new Uint8Array(buffer_arg));
}


// Interface exposed by the Ones Service
var OnesService = exports.OnesService = {
  // This servers's sole RPC: get a One, return a One
  echoOne: {
    path: '/Ones/EchoOne',
    requestStream: false,
    responseStream: false,
    requestType: ones_pb.GetOneRequest,
    responseType: ones_pb.GetOneResponse,
    requestSerialize: serialize_GetOneRequest,
    requestDeserialize: deserialize_GetOneRequest,
    responseSerialize: serialize_GetOneResponse,
    responseDeserialize: deserialize_GetOneResponse,
  },
};

exports.OnesClient = grpc.makeGenericClientConstructor(OnesService);
