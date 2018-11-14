import grpc
from concurrent import futures
import time

# import the generated classes
from messages import ones_pb2
from services import ones_pb2_grpc

MAX_SERVER_WORKERS = 1

# Extend the auto-generated servicer class such that it
# actually implements the defined functions
class OnesServicer(ones_pb2_grpc.OnesServicer):
    def EchoOne(self, request, context):
        print "serving request..."
        if request.value.value == 1:
            one = ones_pb2.One(value=1)
        else:
            one = ones_pb2.One(value=-1)
        response = ones_pb2.GetOneResponse(value=one)
        return response

server = grpc.server(futures.ThreadPoolExecutor(max_workers=MAX_SERVER_WORKERS))

ones_pb2_grpc.add_OnesServicer_to_server(OnesServicer(), server)

# listen on port 50051
print('Starting python server. Listening on port 50051.')
server.add_insecure_port('[::]:50051')
server.start()


# since server.start() will not block,
# a sleep-loop is added to keep alive
try:
    while True:
        time.sleep(86400)
except KeyboardInterrupt:
    server.stop(0)
