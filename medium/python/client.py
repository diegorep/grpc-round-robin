import grpc
import time

# import the generated classes
from messages import ones_pb2
from services import ones_pb2_grpc

# open a gRPC channel
channel = grpc.insecure_channel('173.0.0.3:50051')

# create a stub (client)
stub = ones_pb2_grpc.OnesStub(channel)

# create a valid request message
for i in xrange(10):
    one = ones_pb2.One(value=1)
    request = ones_pb2.GetOneRequest(value=one)

    # make the call
    response = stub.EchoOne(request)

    # et voila
    print("Response type->", type(response))
    print("Response value->", response.value)
    print("Response's value's type->", type(response.value))
    print("Response's value's value->", response.value.value)
    time.sleep(1)
