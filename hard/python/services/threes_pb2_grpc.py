# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
import grpc

from messages import threes_pb2 as threes__pb2


class ThreesStub(object):
  # missing associated documentation comment in .proto file
  pass

  def __init__(self, channel):
    """Constructor.

    Args:
      channel: A grpc.Channel.
    """
    self.EchoThree = channel.unary_unary(
        '/Threes/EchoThree',
        request_serializer=threes__pb2.GetThreeRequest.SerializeToString,
        response_deserializer=threes__pb2.GetThreeResponse.FromString,
        )


class ThreesServicer(object):
  # missing associated documentation comment in .proto file
  pass

  def EchoThree(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')


def add_ThreesServicer_to_server(servicer, server):
  rpc_method_handlers = {
      'EchoThree': grpc.unary_unary_rpc_method_handler(
          servicer.EchoThree,
          request_deserializer=threes__pb2.GetThreeRequest.FromString,
          response_serializer=threes__pb2.GetThreeResponse.SerializeToString,
      ),
  }
  generic_handler = grpc.method_handlers_generic_handler(
      'Threes', rpc_method_handlers)
  server.add_generic_rpc_handlers((generic_handler,))
