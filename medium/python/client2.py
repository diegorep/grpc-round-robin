import grpc
import time

# import the generated classes for the resource
from messages import ones_pb2
from services import ones_pb2_grpc

# open three gRPC channels, one per resource
ones_channel = grpc.insecure_channel('ones.digitapis.com:50051')

# create a stub (client) for each resource
ones_stub = ones_pb2_grpc.OnesStub(ones_channel)

# wait for channels to be ready
grpc.channel_ready_future(ones_channel)

# create a valid request message
for i in xrange(10):
    one = ones_pb2.One(value=1)
    request = ones_pb2.GetOneRequest(value=one)

    # make the call
    response = ones_stub.EchoOne(request)

    # et voila
    print("Response's value's value->", response.value.value)
