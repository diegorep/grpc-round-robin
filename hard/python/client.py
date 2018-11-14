import grpc
import time

# import the generated classes for all three resources
from messages import ones_pb2, twos_pb2, threes_pb2
from services import ones_pb2_grpc, twos_pb2_grpc, threes_pb2_grpc

# open three gRPC channels, one per resource
ones_channel = grpc.insecure_channel('ones.digiapis.com:50051')
#twos_channel = grpc.insecure_channel('twos.digitapis.com:50050')
#threes_channel = grpc.insecure_channel('threes.digitapis.com:50049')

# create a stub (client) for each resource
ones_stub = ones_pb2_grpc.OnesStub(ones_channel)
#twos_stub = twos_pb2_grpc.TwosStub(twos_channel)
#threes_stub = threes_pb2_grpc.ThreesStub(threes_channel)

# create a valid request message
for i in xrange(10):
    one = ones_pb2.One(value=1)
    request = ones_pb2.GetOneRequest(value=one)

    # make the call
    response = ones_stub.EchoOne(request)

    # et voila
    print("Response's value's value->", response.value.value)
    ''' 
    # Repeat for Twos service
    two = twos_pb2.Two(value=2)
    request = twos_pb2.GetTwoRequest(value=two)
    response = twos_stub.EchoTwo(request)
    print("Response's value's value->", response.value.value)
    
    # Repeat for Threes service
    three = threes_pb2.Three(value=3)
    request = threes_pb2.GetThreeRequest(value=three)
    response = threes_stub.EchoThree(request)
    print("Response's value's value->", response.value.value)
    time.sleep(1)
    '''
